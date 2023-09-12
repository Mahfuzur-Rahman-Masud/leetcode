use std::collections::VecDeque;

struct MyStack {
    q: VecDeque<i32>,
}


/**
 * `&self` means the method takes an immutable reference
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyStack {
    fn new() -> Self {
        Self {
            q: VecDeque::new()
        }
    }

    fn push(&mut self, x: i32) {
        self.q.push_back(x);
    }

    fn pop(&mut self) -> i32 {
        match self.q.pop_back() {
            None => 0,
            Some(v) => v
        }
    }

    fn top(&self) -> i32 {
        match self.q.back(){
            None=>0,
            Some(v)=>*v
        }
    }

    fn empty(&self) -> bool {
        self.q.is_empty()
    }
}


#[cfg(test)]
mod test {
    use crate::stack::MyStack;

    #[test]
    fn stack_test() {
        let mut s = MyStack::new();
        assert_eq!(s.pop(), 0);
        assert_eq!(s.top(), 0);
        s.push(2);
        assert_eq!(s.top(), 2);
        assert_eq!(s.pop(), 2);
        assert_eq!(s.pop(), 0);
        assert_eq!(s.empty(), true);
    }
}


