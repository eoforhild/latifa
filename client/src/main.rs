use client::{generate_keys_and_dump, register_and_publish, login, request_connection};
use rand::{rngs::StdRng, SeedableRng};
use serde_json::json;
use x25519_dalek::{StaticSecret, PublicKey};


#[tokio::main]
async fn main() {
    generate_keys_and_dump();
    let email = "asdfd@yahoo.com";
    let username = "asddgasf";
    let password = "dgsdfwse";

    let email2 = "asdafvv@lasd.com";
    let user2 = "ppppp";
    let pass2 = "sdfvse";
    let js = json!({
        "username": user2,
        "email": email2,
        "password": pass2,
    });

    let stat = request_connection(email2.to_string()).await;
    println!("{}", stat)
    // let b: [u8; 32] = [12; 32];
    // let h = kdf(&b);
    // println!("{:?}", h);
}