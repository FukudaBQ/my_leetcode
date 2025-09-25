import java.util.*;

public class leetcode373 {
    public List<List<Integer>> kSmallestPairs(int[] nums1, int[] nums2, int k) {
        PriorityQueue<List<Integer>> pq = new PriorityQueue<>(Comparator.comparingInt(o -> nums1[o.get(0)] + nums2[o.get(1)]));
        List<List<Integer>> ans = new LinkedList<>();
        HashSet<List<Integer>> visited = new HashSet<>();
        pq.offer(Arrays.asList(0, 0));
        while (k > 0) {
            List<Integer> smallest = pq.poll();
            int i = smallest.get(0),  j = smallest.get(1);
            if (j + 1 < nums2.length) {
                if (!visited.contains(Arrays.asList(i, j + 1))){
                    pq.add(Arrays.asList(i, j + 1));
                    visited.add(Arrays.asList(i, j + 1));
                }
            }
            if (i + 1 < nums1.length) {
                if (!visited.contains(Arrays.asList(i + 1, j))){
                    pq.add(Arrays.asList(i + 1, j));
                    visited.add(Arrays.asList(i + 1, j));
                }
            }
            ans.add(List.of(nums1[smallest.get(0)], nums2[smallest.get(1)]));
            k --;
        }
        return ans;
    }

    public static void main(String[] args) {
        leetcode373 leetcode = new leetcode373();

        List<List<java.lang.Integer>> res = leetcode.kSmallestPairs(new int[]{1, 7, 11}, new int[]{2, 4, 6}, 3);
        System.out.println(res.toString());
    }
}
