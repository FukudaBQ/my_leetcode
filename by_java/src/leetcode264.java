import java.util.PriorityQueue;

public class leetcode264 {
    public int nthUglyNumber(int n) {
        PriorityQueue<Long> heap = new PriorityQueue<>();
        Long[] dp = new Long[n];
        if (n == 1) {
            return 1;
        }
        heap.add(1L);
        for (int i = 0; i < n; i ++) {
            Long last = heap.poll();
            if (i > 1 && dp[i - 1].equals(last)) {
                i -= 1;
                continue;
            }
            dp[i]  = last;
            heap.add(last * 2);
            heap.add(last * 3);
            heap.add(last * 5);
        }
        return dp[n-1].intValue();
    }

    public static void main(String[] args) {
        leetcode264 leetcode = new leetcode264();
        int res = leetcode.nthUglyNumber(1407);
        System.out.println(res);
    }
}
