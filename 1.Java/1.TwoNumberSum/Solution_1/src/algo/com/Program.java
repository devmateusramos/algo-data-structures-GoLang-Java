package algo.com;

import java.util.HashSet;
import java.util.Set;

public class Program {
    //Solution Time O(n) | Space O(n)
    public static int[] TwoNumberSum(int[] array, int targetSum) {
        Set<Integer> visitedSet = new HashSet<>();
        for (int num: array) {
            if (visitedSet.contains(targetSum- num)) {
                return new int[]{targetSum - num, num};
            } else {
                visitedSet.add(num);
            }
        }
        return new int[0];
    }
}
