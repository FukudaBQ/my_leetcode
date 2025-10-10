public class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    /**
     * 把 int 数组快速建成链表，返回头结点。
     * 例：buildListNode(1,2,3) -> 1->2->3
     */
    public static ListNode buildListNode(int... nums) {
        if (nums == null || nums.length == 0) return null;
        ListNode dummy = new ListNode();
        ListNode cur = dummy;
        for (int v : nums) {
            cur.next = new ListNode(v);
            cur = cur.next;
        }
        return dummy.next;
    }

    /**
     * 把链表打印成一行字符串：1->2->3
     * 空链表返回 "null"
     */
    public static String printListNode(ListNode head) {
        if (head == null) return "null";
        StringBuilder sb = new StringBuilder();
        while (head != null) {
            sb.append(head.val);
            if (head.next != null) sb.append("->");
            head = head.next;
        }
        return sb.toString();
    }
}
