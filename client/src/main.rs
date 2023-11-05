use client::{generate_keys_and_dump, register_and_publish, login, request_connection, pending_requests};
use rand::{rngs::StdRng, SeedableRng};
use serde_json::json;
use x25519_dalek::{StaticSecret, PublicKey};


#[tokio::main]
async fn main() {
    // let email = "asdf@yahoo.com";
    // let username = "mmmmm";
    // let password = "dgsdfwse";
    // generate_keys_and_dump();
    // let reg1 = json!({
    //     "username": username,
    //     "email": email,
    //     "password": password,
    // });
    // let reg1stat = register_and_publish(&reg1.to_string()).await;
    // println!("Register 1 {}", reg1stat);

    let email2 = "zxcv@lasd.com";
    let user2 = "ppppp";
    let pass2 = "sdfvse";
    let reg2 = json!({
        "username": user2,
        "email": email2,
        "password": pass2,
    });
    // let reg2stat = register_and_publish(&reg2.to_string()).await;
    // println!("Register 2 {}", reg2stat);

    // let log2 = json!({
    //     "login": email2,
    //     "password": pass2
    // });
    // let log2stat = login(&log2.to_string()).await;
    // println!("Login 2 {}", log2stat);

    // let stat = request_connection(email2.to_string()).await;
    // println!("{}", stat)
    // let b: [u8; 32] = [12; 32];
    // let h = kdf(&b);
    // println!("{:?}", h);
    let stat = pending_requests().await;
    println!("{}", stat)
}