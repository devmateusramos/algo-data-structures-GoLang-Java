package algo.com;

public class Main {
    public static void main(String[] args) {
        int[] numbers = {3, 5, -4, 8, 11, 1, -1, 6};
        int[] target = Program.TwoNumberSum(numbers, 10);
        System.out.println("Values " + target[0] + " E " + target[1]);
    }
}