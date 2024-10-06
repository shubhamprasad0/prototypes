use std::{
    collections::VecDeque,
    sync::{Condvar, Mutex},
};

pub struct BBQ<T> {
    queue: Mutex<VecDeque<T>>,
    not_empty: Condvar,
    not_full: Condvar,
    cap: usize,
}

impl<T> BBQ<T> {
    pub fn new(cap: usize) -> Self {
        return BBQ {
            queue: Mutex::new(VecDeque::with_capacity(cap)),
            not_empty: Condvar::new(),
            not_full: Condvar::new(),
            cap,
        };
    }

    pub fn len(&self) -> usize {
        let queue = self.queue.lock().unwrap();
        queue.len()
    }

    pub fn push(&self, val: T) {
        let mut queue = self.queue.lock().unwrap();

        while queue.len() == self.cap {
            queue = self.not_full.wait(queue).unwrap()
        }

        queue.push_back(val);

        self.not_empty.notify_one();
    }

    pub fn pop(&self) -> T {
        let mut queue = self.queue.lock().unwrap();

        while queue.len() == 0 {
            queue = self.not_empty.wait(queue).unwrap();
        }

        let val = queue.pop_front().unwrap();

        self.not_full.notify_one();

        val
    }
}

#[cfg(test)]
mod tests {
    use std::{sync::Arc, thread};

    use super::*;

    #[test]
    fn create_new_bbq() {
        let q: BBQ<i64> = BBQ::new(2);
        assert_eq!(q.cap, 2)
    }

    #[test]
    fn test_push_pull() {
        let q: BBQ<i64> = BBQ::new(2);
        q.push(2);
        q.push(3);

        let val1 = q.pop();
        assert_eq!(val1, 2);

        let val2 = q.pop();
        assert_eq!(val2, 3);
    }

    #[test]
    fn test_blocking_on_push() {
        let q = Arc::new(BBQ::new(2));
        q.push(2);
        q.push(3);

        let q_clone = Arc::clone(&q);
        let handle = thread::spawn(move || {
            q_clone.push(4);
        });
        assert_eq!(q.len(), 2);
        let val = q.pop();
        handle.join().unwrap();
        assert_eq!(val, 2);
        assert_eq!(q.len(), 2);

        let val = q.pop();
        assert_eq!(val, 3);

        let val = q.pop();
        assert_eq!(val, 4);
    }

    #[test]
    fn test_blocking_on_pop() {
        let q = Arc::new(BBQ::new(2));

        let q_clone = Arc::clone(&q);
        let handle = thread::spawn(move || {
            let val = q_clone.pop();
            assert_eq!(val, 1);
        });

        q.push(1);
        handle.join().unwrap();
    }
}
