import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.PriorityQueue;

public class leetcode23 {
    public ListNode mergeKLists(ListNode[] lists) {
        PriorityQueue<ListNode> heap = new PriorityQueue<>(Comparator.comparing(x -> x.val));
        for (int i = 0; i < lists.length; i++) {
            if (lists[i] != null) {
                heap.add(lists[i]);
                lists[i] = lists[i].next;
            }
        }
        ListNode head = new ListNode(0);
        ListNode cur = head;
        while (!heap.isEmpty()) {
            ListNode node = heap.poll();
            cur.next = node;
            cur = cur.next;
            if (node.next != null) {
                heap.add(node.next);
            }
        }
        return head.next;
    }

    public static void main(String[] args) {
        leetcode23 leetcode = new leetcode23();

        ListNode[] lists = new ListNode[] {
            ListNode.buildListNode(1, 4, 5),
            ListNode.buildListNode(1, 3, 4),
            ListNode.buildListNode(2, 6)
        };

        ListNode res = leetcode.mergeKLists(lists);
        System.out.println(ListNode.printListNode(res));
    }
}


