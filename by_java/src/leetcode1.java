import java.util.Arrays;

public class leetcode1 {
    public int[] twoSum(int[] nums, int target) {
        int a[]=new int[2];
        for (int i=0;i<nums.length;i++){
            int b=nums[i];
            for(int k=i+1;k<nums.length;k++){
                int c=nums[k];
                if(b+c==target){
                    a[0]=i;
                    a[1]=k;
                }
            }
        }
        return a;
    }

    public static void main(String[] args) {
        leetcode1 leetcode1 = new leetcode1();
        int[] nums=new int[]{2,7,11,15};
        int target=9;
        int[] res = leetcode1.twoSum(nums,target);
        System.out.println(Arrays.toString(res));
    }
}
