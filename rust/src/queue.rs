struct MyQueue {
    s: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyQueue {
    fn new() -> Self {
        Self {
            s: vec![]
        }
    }

    fn push(&mut self, x: i32) {
      self.s.insert(0, x);
    }

    fn pop(&mut self) -> i32 {
        if self.s.is_empty(){
            return  0
        }

        self.s.pop().unwrap()
    }

    fn peek(&self) -> i32 {
        if self.s.is_empty(){
            return 0
        }

        self.s.get(self.s.len() -1).unwrap().clone()
    }

    fn empty(&self) -> bool {
        self.s.is_empty()
    }
}

#[cfg(test)]
mod test{
    use crate::queue::MyQueue;

    #[test]
    fn queue_test(){
        let mut  q = MyQueue::new();
        q.push(5);
        assert!(!q.empty());
        assert_eq!(q.peek(), 5);
        assert_eq!(q.pop(), 5);
        assert_eq!(q.peek(), 0);
        assert_eq!(q.pop(), 0);
        assert!(q.empty());

        q.push(5);
        q.push(7);
        assert_eq!(q.peek(), 5);
        assert_eq!(q.pop(), 5);
        assert_eq!(q.pop(), 7);
    }
}
