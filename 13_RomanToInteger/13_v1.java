public class Solution {
    public int romanToInt(String s) {
        HashMap<Character, Integer> alpha = new HashMap<Character, Integer>();
        alpha.put('I', 1);
        alpha.put('V', 5);
        alpha.put('X', 10);
        alpha.put('L', 50);
        alpha.put('C', 100);
        alpha.put('D', 500);
        alpha.put('M', 1000);
        
        int buf = 0;
        int sum = 0;
        
        for (Character c: s.toCharArray()){
            int num = alpha.get(c);
            
            if (buf == 0){
                buf = num;
            } else if (buf < num){
                sum = sum - buf + num;
                buf = 0;
            } else {
                sum += buf;
                buf = num;
            }
        }
        
        sum += buf;
        return sum;
    }
}
