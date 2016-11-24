public class Solution {
    public boolean isValid(String s) {
        HashMap mapping = new HashMap<Character, Character>();
        mapping.put(']', '[');
        mapping.put('}', '{');
        mapping.put(')', '(');
        
        Stack check = new Stack<Character>();
        
        for (int i = 0; i < s.length(); i++){
            if (mapping.get(s.charAt(i)) != null){
                if (check.isEmpty() || mapping.get(s.charAt(i)) != check.peek())
                    return false;
                check.pop();
            } else {
                check.push(s.charAt(i));
            }
        }
        
        return check.isEmpty();
    }
}
