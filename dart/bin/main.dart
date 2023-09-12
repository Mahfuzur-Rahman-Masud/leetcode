import 'package:dart/dart.dart' as dart;

void main(List<String> arguments) {
  print('Hello world: ${dart.calculate()}!');
}



 class ListNode {
     int val;
     ListNode? next;
     ListNode([this.val = 0, this.next]);
   }

class Solution {
  ListNode? addTwoNumbers(ListNode? l1, ListNode? l2) {
    var head = ListNode();
    var tail = head;
    int carry =0;



    while(l1 != null || l2!= null || carry > 0){
      int sum = (l1?.val??0) + (l2?.val??0) + carry;
      int digit = sum %10;
      carry = sum > 9 ? 1: 0;
      tail.next = ListNode(digit);
      tail = tail.next!;
      l1 = l1?.next;
      l2=l2?.next;
    }

    return head.next;
  }
}