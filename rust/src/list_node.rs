// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }

    pub fn from(values: &[i32]) -> Option<Box<ListNode>> {
        let mut head = None;
        let mut tail = &mut head;

        for &val in values {
            *tail = Some(Box::new(ListNode::new(val)));
            tail = &mut tail.as_mut().unwrap().next
        }

        head
    }

    fn values(self) -> Vec<i32> {
        let mut array = Vec::new();
        let mut current = Some(Box::new(self));

        // while let Some(node) = current{
        //     array.push(node.val);
        //     current = node.next.as_mut().unwrap();
        // }

        loop {
            match current {
                Some(n) => {
                    array.push(n.val);
                    current = n.next
                }
                None => break
            }
        }

        array
    }
}


#[allow(dead_code)]
pub fn is_palindrome2(head: Option<Box<ListNode>>) -> bool {
    let (mut fast, mut slow) = (&head, &head);

    // Travel to the middle of the list
    while fast.is_some() && fast.as_ref().unwrap().next.is_some() {
        fast = &fast.as_ref().unwrap().next.as_ref().unwrap().next;
        slow = &slow.as_ref().unwrap().next;
    }


    // Break the list in the middle and reverse second half
    let (mut prev, mut curr) = (None, slow.clone());

    while let Some(mut node) = curr {
        curr = node.next;
        node.next = prev;
        prev = Some(node);
    }

    let (mut head, mut tail) = (head, prev);

    while head.is_some() && tail.is_some() {
        if head.as_ref().unwrap().val != tail.as_ref().unwrap().val {
            return false;
        }

        head = head.unwrap().next;
        tail = tail.unwrap().next;
    }

    true
}


#[allow(dead_code)]
pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {
    let mut c = vec![];
    let mut current = head;

    while let Some(node) = current {
        c.push(node.val);
        current = node.next;
    }

    let li = c.len() - 1;
    let m = c.len() / 2 + 1;

    for i in 0..m {
        if c[i] != c[li - i] {
            return false;
        }
    }

    return true;
}

#[allow(dead_code)]
pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let (mut prev, mut curr) = (None, head);

    while let Some(mut node) = curr {
        curr = node.next;
        node.next = prev;
        prev = Some(node);
    }

    prev
}


// pub fn reverse_list(head: Option<Box<ListNode>>)->Option<Box<ListNode>>{
//     let (mut prev, mut curr) = (None, head);
//
//     while let Some(mut node) = curr{
//         curr = node.next;
//         node.next = prev;
//         prev = Some(node)
//     }
//
//     prev
// }


// pub fn reverse_list( head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
//     if head.is_none(){
//         return None
//     }
//
//     let (mut prev, mut curr) = (None, head);
//
//     // 123
//     while let Some(mut node) = curr{
//         curr = node.next; //    node 1 / curr 2     node 2 / curr 3
//         node.next = prev; //    1->None             2->1
//         prev = Some(node); //   1                   2
//     }
//
//     prev
// }


#[allow(dead_code)]
pub fn remove_elements(mut head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
    let mut ptr = &mut head;
    loop {
        match ptr {
            None => break,
            Some(n) if n.val == val => *ptr = n.next.take(),
            Some(n) => ptr = &mut n.next,
        }
    }

    head
}

struct Solution();

impl Solution {
    pub fn remove_nth_from_end(mut head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }

        let mut count = 0;
        let mut ptr = &mut head;

        loop {
            match ptr {
                None => break,
                Some(n) => {
                    count += 1;
                    ptr = &mut n.next;
                }
            }
        }

        let nth = count - val;
        let mut count = 0;
        let mut ptr = &mut head;

        loop {
            match ptr {
                None => break,
                Some(n) if count == nth => {
                    *ptr = n.next.take();
                    break;
                }
                Some(n) => {
                    count += 1;
                    ptr = &mut n.next;
                }
            }
        }

        head
    }
}



#[cfg(test)]
mod test {
    use std::{assert_eq, println};

    use crate::list_node::{is_palindrome, is_palindrome2, ListNode, remove_elements, reverse_list, Solution};

    #[test]
    fn remove_nth_from_end_test() {
        let mut head = Solution::remove_nth_from_end(ListNode::from(&[1, 2, 3, 4, 5]), 2);
        println!("{:?}", if head.is_some() { head.unwrap().values() } else { Vec::new() });
        head = Solution::remove_nth_from_end(ListNode::from(&[1]),1 );
        println!("{:?}", if head.is_some() { head.unwrap().values() } else { Vec::new() });
        head = Solution::remove_nth_from_end(ListNode::from(&[1,2]),1 );
        println!("{:?}", if head.is_some() { head.unwrap().values() } else { Vec::new() });
    }

    #[test]
    fn reverse_list_test() {
        assert_eq!(reverse_list(ListNode::from(&[1, 2, 3, 4, 5])).unwrap().values(), &[5, 4, 3, 2, 1])
    }

    #[test]
    fn remove_elements_test() {
        let head = ListNode::from(&[1, 2, 6, 3, 4, 5, 6]);
        let h = head.unwrap();
        println!("{:?}", h.clone().values());
        let h2 = remove_elements(Some(h.clone()), 6);
        println!("{:?}", h2.clone().unwrap().values());
        assert_eq!(h2.clone().unwrap().values(), &[1, 2, 3, 4, 5])
    }


    #[test]
    fn is_palindrome_test() {
        assert!(is_palindrome(ListNode::from(&[1, 2, 2, 1])));
        assert!(is_palindrome(ListNode::from(&[1, 2, 1])));
        assert!(is_palindrome(ListNode::from(&[1, 1])));
        assert!(!is_palindrome(ListNode::from(&[1, 2])));
    }


    #[test]
    fn is_palindrome_test2() {
        assert!(is_palindrome2(ListNode::from(&[1, 2, 2, 1])));
        assert!(is_palindrome2(ListNode::from(&[1, 2, 1])));
        assert!(is_palindrome2(ListNode::from(&[1, 1])));
        assert!(!is_palindrome2(ListNode::from(&[1, 2])));
    }
}