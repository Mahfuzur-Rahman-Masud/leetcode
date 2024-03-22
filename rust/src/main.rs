// use actix_web::{App, HttpServer, Responder, web};
// use crate::list_node::{is_palindrome, is_palindrome2, ListNode};

use std::time::SystemTime;

use crate::list_node::{is_palindrome2, ListNode};

mod bitwise;
mod numbers;
mod practice;
mod list_node;
mod strings;
mod arrays;
mod b_tree;
mod stack;
mod queue;
mod api;
mod combinations;
mod mix;
mod matrix;


// async fn hello() -> impl Responder {
//     "Hello, World!"
// }
//
// #[tokio::main]
// async fn main() -> std::io::Result<()> {
//     HttpServer::new(|| {
//         App::new()
//             .route("/", web::get().to(hello))
//     })
//         .bind("127.0.0.1:5555")?
//         .run()
//         .await
// }
#[allow(dead_code)]
fn main() {
    println!("Warming up...");
    for _ in 0..1000 {
        is_palindrome2(ListNode::from(&[1, 2, 2, 1]));
    }


    println!("Firing up...");
    let start = SystemTime::now();
    // Your code...
    for _ in 0..5000_000 {
        is_palindrome2(ListNode::from(&[1, 2, 2, 1]));
    }
    let end = SystemTime::now();

    if let Ok(duration) = end.duration_since(start) {
        // duration represents the elapsed time
        println!("Elapsed time: {:?}", duration);
    } else {
        println!("Time went backwards!");
    }
}