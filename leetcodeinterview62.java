import java.util.ArrayList;

public class leetcodeinterview62 {
    public static int lastRemaining(int n, int m) {
        ArrayList<Integer> list = new ArrayList<>(n);
        for (int i = 0; i < n; i++) {
            list.add(i);
        }
        int idx = 0;
        while (n > 1) {
            idx = (idx + m - 1) % n;
            list.remove(idx);
            n--;
        }
        return list.get(0);
    }

    public static void main(String[] args){
        int res = lastRemaining(15,3);
        System.out.println(res);
    }
}
