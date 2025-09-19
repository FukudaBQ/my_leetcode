import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.Stack;

public class leetcode150 {
    public int evalRPN(String[] tokens) {
        Set<String> operator = Set.of("+", "-", "*", "/");
        Stack<String> stack = new Stack<>();
        for (int i = 0; i < tokens.length; i++) {
            String c = tokens[i];
            if (operator.contains(c)) {
                int rightOp = Integer.parseInt(stack.pop());
                int leftOp = Integer.parseInt(stack.pop());
                if (c.equals("+")) {
                    stack.push(String.valueOf(rightOp + leftOp));
                } else if(c.equals("-")) {
                    stack.push(String.valueOf(leftOp - rightOp));
                } else if (c.equals("*")) {
                    stack.push(String.valueOf(leftOp * rightOp));
                } else if(c.equals("/")) {
                    stack.push(String.valueOf(leftOp / rightOp));
                }
            } else {
                stack.push(c);
            }
        }
        return Integer.parseInt(stack.pop());
    }

    public static void main(String[] args) {
        leetcode150 leetcode150 = new leetcode150();
        int res = leetcode150.evalRPN(new String[]{"2","1","+","3","*"});
        System.out.println(res);
    }
}
