/*
 * @Descripttion: 
 * @version: 
 * @Author: renjie.zhang
 * @Date: 2019-10-24 17:30:27
 * @LastEditTime: 2019-10-24 17:43:21
 */
class Solution{
    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(3);
        ListNode node4 = new ListNode(4);
        ListNode node5 = new ListNode(5);
        node1.next = node2;
        node2.next = node3;
        node3.next = node4;
        node1.next = node5;
        ListNode node =  removeNthFromEnd(node1, 2);
        System.out.print(node.val);
    }

    public static ListNode removeNthFromEnd(ListNode head, int n) {
        int length = 1;
        ListNode temp = head;
        while (temp.next != null){
            length++;
            temp = temp.next;
        }
        if(n<0||n>length){
            return null;
        }
        if(n == length){
            head = head.next;
            return head;
        }
        
        ListNode cureent = head;
        for (int i= length-n-1;i>0;i--){
            cureent = cureent.next;
        }
        
        cureent.next = cureent.next.next;
        return head;
    }
}

class ListNode{
    int val;
    ListNode next;
    ListNode(int x){
        this.val = x;
    }
}