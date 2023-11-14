use std::{fs::{self, File}, io::Write};

use aes_gcm::{Aes256Gcm, Key, KeyInit, AeadCore, aead::{Payload, Aead}};
use hkdf::Hkdf;
use pqc_kyber::*;
use rand::{rngs::StdRng, SeedableRng};
use serde_json::{json, Value};
use sha2::Sha512;
use x25519_dalek::{StaticSecret, PublicKey, x25519};

use crate::api::xeddsa;

use super::BASE_URL;

const HKDF_INFO: &[u8; 53] = b"LATIFAProtocol_CURVE25519_SHA-512_CRYSTALS-KYBER-1024";
const HKDF_SALT: [u8; 64] = [0; 64];
const HKDF_F: [u8; 32] = [0xFF; 32];

fn kdf(km: Vec<u8>) -> [u8; 32] {
    let ikm = [&HKDF_F, km.as_slice()].concat();
    let hk = Hkdf::<Sha512>::new(Some(&HKDF_SALT), &ikm);
    let mut okm: [u8; 32] = [0; 32];
    hk.expand(HKDF_INFO, &mut okm)
        .expect("valid");
    okm
}

pub fn pending_requests() -> bool {
    let token = fs::read_to_string("auth").unwrap();
    let client = reqwest::blocking::Client::new();
    let res = match client.get(BASE_URL.to_owned()+"/requests/pending")
        .header("Authorization", "Bearer ".to_owned() + &token)
        .send() {
        Ok(e) => e,
        Err(_) => return false,
    };
    match res.json::<Value>() {
        Ok(r) => {
            println!("{:?}",r);
        },
        Err(_) => return false,
    }
    true
}

pub fn approved_requests() {
    let token = fs::read_to_string("auth").unwrap();
    let client = reqwest::blocking::Client::new();
    let res = match client.get(BASE_URL.to_owned()+"/requests/approved").send() {
        Ok(r) => r,
        Err(err) => return,
    };

    let binding = res.json::<Value>().unwrap();
    let arr = binding.as_array().unwrap();
    for req in arr {
        fetch_keys_handshake(req["id"].to_string());
    }
}

/**
 * Requests a conncetion to the client with the email
 * Returns true if request was made successfully.
 * This DOES NOT mean that the request was approved,
 * only that it was posted onto the server.
 */
pub fn request_connection(email: String) -> bool {
    let token = fs::read_to_string("auth").unwrap();
    let client = reqwest::blocking::Client::new();
    let res = match client.post(BASE_URL.to_owned()+"/requests/email/"+&email)
        .header("Authorization", "Bearer ".to_owned() + &token)
        .send() {
        Ok(e) => e,
        Err(_) => return false,
    };
    if res.status().as_u16() == 204 {
        true
    } else {
        false
    }
}

/**
 * Upon approval of key fetching (aka the person allowed
 * you to initiate contact), fetch all keys needed and
 * initiate the handshake protocol.
 */
pub fn fetch_keys_handshake(req_id: String) -> Vec<u8> {
    let x = xeddsa::XEdDSA::new();
    // FETCH from the server and get a JSON response with all the keys needed
    let client = reqwest::blocking::Client::new();
    let body = json!({
        "request_id": req_id
    });
    let res = match client.get(BASE_URL.to_owned() + "/handshakes")
        .json(&body)
        .send() {
            Ok(r) => r,
            Err(_) => return vec![],
        };


    let keys: Value = res.json().unwrap();
    let mut ik_b: [u8; 32] = [0; 32];
    let mut spk_b: [u8; 32] = [0; 32];
    let mut spk_b_sig: [u8; 64] = [0; 64];
    let mut pqpk_b: [u8; KYBER_PUBLICKEYBYTES] = [0; KYBER_PUBLICKEYBYTES];
    let mut pqpk_b_sig: [u8; 64] = [0; 64];
    let mut opk_b: [u8; 32] = [0; 32];
    hex::decode_to_slice(keys["ik"].as_str().unwrap(), &mut ik_b).unwrap();
    hex::decode_to_slice(keys["spk"].as_str().unwrap(), &mut spk_b).unwrap();
    hex::decode_to_slice(keys["spk_sig"].as_str().unwrap(), &mut spk_b_sig).unwrap();
    hex::decode_to_slice(keys["pqpk"].as_str().unwrap(), &mut pqpk_b).unwrap();
    hex::decode_to_slice(keys["pqpk_sig"].as_str().unwrap(), &mut pqpk_b_sig).unwrap();
    // hex::decode_to_slice(keys["opk"].as_str().unwrap(), &mut opk_b).unwrap();

    // Verify signatures
    let b1 = x.verify(ik_b.clone(), &spk_b, spk_b_sig);
    let b2 = x.verify(ik_b.clone(), &pqpk_b, pqpk_b_sig);
    if b1 || b2 {
        // Abort
        return vec![]
    }

    // Just ephemeral secret
    let ek_sec: StaticSecret = StaticSecret::random_from_rng(StdRng::from_entropy());
    let ek_pub: PublicKey = PublicKey::from(&ek_sec);

    // PQKEM stuff
    let (ct, ss) = encapsulate(&pqpk_b, &mut StdRng::from_entropy()).unwrap();

    // Read in the identity key
    let f = fs::read_to_string("keys.json").unwrap();
    let keys: Value = serde_json::from_str(&f).unwrap();
    let mut ik_pub: [u8; 32] = [0; 32]; 
    hex::decode_to_slice(keys["ik_pub"].as_str().unwrap(), &mut ik_pub).unwrap();
    let mut ik_sec: [u8; 32] = [0; 32]; 
    hex::decode_to_slice(keys["ik_sec"].as_str().unwrap(), &mut ik_sec).unwrap();

    // Compute triple diffie hellman
    // Needs to account for dh4 with OPK
    let dh1 = x25519(ik_sec.clone(), spk_b.clone());
    let dh2 = x25519(ek_sec.to_bytes(), ik_b.clone());
    let dh3 = x25519(ek_sec.to_bytes(), spk_b.clone());

    let km = [dh1, dh2, dh3, ss].concat();
    let sk = kdf(km);
    let mut file = File::create("sk").unwrap();
    file.write_all(&sk).unwrap();

    // Compute associated data
    let ad = [ik_pub, ik_b].concat();

    // Let the first message of this protocol be the 
    let aes_key = Key::<Aes256Gcm>::from_slice(&sk);
    let cipher = Aes256Gcm::new(&aes_key); 
    let nonce = Aes256Gcm::generate_nonce(&mut StdRng::from_entropy());
    let payload = Payload {
        msg: &ik_pub,
        aad: &ad,
    };
    let handshake = cipher.encrypt(&nonce, payload).unwrap();

    // Convert all needed info into hex strings
    let s_ek_pub = hex::encode(ek_pub.as_bytes());
    let s_ct = hex::encode(ct);
    let s_handshake = hex::encode(&handshake);
    let s_pqpk_b = hex::encode(pqpk_b);
    let s_opk_b = hex::encode(opk_b);

    // Construct the JSON
    let body = json!({
        "ik": keys["ik_pub"].as_str().unwrap(),
        "ek": s_ek_pub,
        "ct": s_ct,
        "handshake": s_handshake,
        "pqpk_used": s_pqpk_b,
        "opk_used": s_opk_b,
    });
    let handshake = body.to_string();
    let body = json!({
        "request_id": req_id,
        "handshake": handshake
    });

    let client = reqwest::blocking::Client::new();
    let res = match client.post(BASE_URL.to_owned()+"/handshake")
        .json(&body)
        .send() {
            Ok(re) => re,
            Err(_) => return vec![],
        };
    vec![]
}

// pub fn complete_handshake(handshake: String) {
//     let hs: Value = serde_json::from_str(&handshake).unwrap();
//     let s_ik_a = hs["ik"].to_string();
//     let s_ek_a = hs["ek"].to_string();

//     let mut ik_a: [u8; 32] = [0; 32];
//     hex::decode_to_slice(s_ik_a, &mut ik_a).unwrap();
//     let mut ek_a: [u8; 32] = [0; 32];
//     hex::decode_to_slice(s_ek_a, &mut ek_a).unwrap();

//     let f = fs::read_to_string("keys.json").unwrap();
//     let keys: Value = serde_json::from_str(&f).unwrap();

//     let s_spk_pub = keys["spk_pub"].to_string();

// }