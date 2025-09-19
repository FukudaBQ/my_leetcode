import java.util.Arrays;
import java.util.Map;
import java.util.Stack;

public class leetcode20 {
    public boolean isValid(String s) {
        Map<Character, Character> map = Map.of(')', '(', '}', '{', ']', '[');
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (map.containsKey(c) && stack.size() > 0) {
                char comp = stack.peek();
                if (map.get(c).equals(comp)) {
                    stack.pop();
                } else {
                    stack.push(c);
                }
            } else {
                stack.push(c);
            }
        }
        return stack.isEmpty();
    }

    public static void main(String[] args) {
        leetcode20 leetcode20 = new leetcode20();
        boolean res = leetcode20.isValid("()[]{}");
        System.out.println(res);
    }
}
