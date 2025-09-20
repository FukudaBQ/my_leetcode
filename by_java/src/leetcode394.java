import java.util.ArrayDeque;
import java.util.Deque;

public class leetcode394 {
    public String decodeString(String s) {
        Deque <String> stack = new ArrayDeque<>();
        String res = "";
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c != ']') {
                String lastChar = stack.peek();
                if (lastChar != null && lastChar.matches("\\d+") && Character.isDigit(c)) {
                    String digit = stack.pop();
                    digit += c;
                    stack.push(digit);
                } else if (lastChar != null && lastChar.matches("(?i)[a-z]+") && Character.isLetter(c)) {
                    String letter = stack.pop();
                    letter += c;
                    stack.push(letter);
                } else {
                    stack.push(String.valueOf(c));
                }
            } else {
                String buildString = "";
                String lastBuild = "";
                for (String a = stack.pop(); !a.equals("["); a = stack.pop()) {
                    if (a.matches("\\d+")) {
                        int num = Integer.parseInt(a);
                        for (int j = 0; j < num - 1; j++) {
                            buildString = lastBuild + buildString;
                        }
                    } else if (!a.isEmpty()){
                        lastBuild = a;
                        buildString = a + buildString;
                    }
                }
                stack.push(buildString);
                stack.push("");
            }
        }
        String lastString = "";
        while (!stack.isEmpty()) {
            String str = stack.pop();
            if (str.matches("\\d+")) {
                int num = Integer.parseInt(str);
                for (int j = 0; j < num - 1; j++) {
                    res = lastString + res;
                }
            } else if (!str.isEmpty()){
                lastString = str;
                res = str + res;
            }
        }
        return res;
    }

    public static void main(String[] args) {
        leetcode394 leetcode394 = new leetcode394();
        String res = leetcode394.decodeString("3[4[a5[nc]6[d]]]");
        System.out.println(res);
    }
}
