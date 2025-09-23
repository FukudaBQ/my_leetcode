import java.util.ArrayDeque;
import java.util.Deque;
import java.util.LinkedList;
import java.util.Stack;

public class leetcode224 {

    public int calculate_failed(String s) {
        Deque<Object> stack = new ArrayDeque<>();
        int res = 0;
        char curSign = '+';
        int curNum = 0;
        int curTotal = 0;
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (Character.isDigit(c)) {
                curNum = curNum * 10 + c - '0';
            } else if (c == '+' || c =='-') {
                if (curSign == '+') {
                    curTotal += curNum;
                } else {
                    curTotal -= curNum;
                }
                curSign = c;
                curNum = 0;
            } else if (c == '(') {
                stack.push(curTotal);
                stack.push(curSign);
                curNum = 0;
                curTotal = 0;
                curSign = '+';
            } else if (c == ')') {
                if (curSign == '+') {
                    curTotal += curNum;
                } else {
                    curTotal -= curNum;
                }
                char sign = (char) stack.pop();
                int num = (int) stack.pop();
                if (sign == '+') {
                    curTotal += num;
                } else {
                    curTotal -= num;
                }
                curNum = 0;
                curSign = '+';
            }
        }
        if (curSign == '+') {
            curTotal += curNum;
        } else {
            curTotal -= curNum;
        }
        return curTotal;
    }

    public int calculate(String s) {
        Deque<Object> stack = new ArrayDeque<>();
        int curValue = 0;
        char curSign = '+';
        int curNum = 0;
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (Character.isDigit(c)) {
                curNum = curNum * 10 + c - '0';
            } else if (c == '+' || c =='-') {
                if (curSign == '+') {
                    curValue += curNum;
                    curNum = 0;
                    curSign = c;
                } else {
                    curValue -= curNum;
                    curNum = 0;
                    curSign = c;
                }
            } else if (c == '(') {
                stack.push(curValue);
                stack.push(curSign);
                curNum = 0;
                curSign = '+';
                curValue = 0;
            } else if (c == ')') {
                if (curSign == '+') {
                    curValue += curNum;
                } else {
                    curValue -= curNum;
                }
                char lastSign = (char) stack.pop();
                int lastValue = (int) stack.pop();
                if (lastSign == '+') {
                    lastValue += curValue;
                } else  {
                    lastValue -= curValue;
                }
                curValue = lastValue;
                curSign = '+';
                curNum = 0;
            }
        }
        if (curSign == '+') {
            curValue += curNum;
        } else if (curSign == '-') {
            curValue -= curNum;
        }
        return curValue;
    }

    public static void main(String[] args) {
        leetcode224 leetcode = new leetcode224();
        int res = leetcode.calculate("1 - ( - 2)");
        System.out.println(res);
    }
}
