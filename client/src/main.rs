use client::{generate_all_keys, publish_all_keys, kdf};
use serde_json::json;


// #[tokio::main]
fn main() {
    // generate_all_keys(false);
    // let email = "asd@yahoo.com";
    // let username = "asdasf";
    // let password = "bobobobo";
    // let js = json!({
    //     "username": username,
    //     "email": email,
    //     "password": password,
    // });
    // println!("{}", publish_all_keys(&js.to_string(),"http://10.169.129.170:8080/register").await);
    let b: [u8; 32] = [12; 32];
    let h = kdf(&b);
    println!("{:?}", h);
}