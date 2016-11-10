import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Solution {
    public int myAtoi(String str) {
        
        Pattern p = Pattern.compile("^ *([\\+-]?[0-9]+).*$");
        Matcher m = p.matcher(str);
        
        if (!m.matches()){
            return 0;
        }
        
        int result = 0;
        
        try {
            result = Integer.parseInt(m.group(1));
        } catch (Exception e) {
            if (m.group(1).contains("-"))
                result = Integer.MIN_VALUE;
            else 
                result = Integer.MAX_VALUE;
        }
        
        return result;
    }
}
