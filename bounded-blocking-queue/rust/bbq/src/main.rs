use std::{sync::Arc, thread};

use bbq::BBQ;

fn main() {
    let q: Arc<BBQ<i32>> = Arc::new(BBQ::new(10));

    let mut handles = vec![];

    for i in 1..=3 {
        let q_clone = Arc::clone(&q);
        let handle = thread::spawn(move || {
            for j in 1..=10 {
                q_clone.push(j);
                println!("Thread {} pushed: {}", i, j)
            }
        });
        handles.push(handle);
    }

    let q_clone = Arc::clone(&q);
    let consumer_handle = thread::spawn(move || {
        for _ in 1..=30 {
            let val = q_clone.pop();
            println!("Popped val: {val}");
        }
    });

    for handle in handles {
        handle.join().unwrap();
    }

    consumer_handle.join().unwrap();

    println!("Length of queue in the end: {}", q.len());
}
