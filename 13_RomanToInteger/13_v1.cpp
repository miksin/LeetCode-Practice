class Solution {
public:
    int romanToInt(string s) {
        map<char, int> alpha;
        alpha['I'] = 1;
        alpha['V'] = 5;
        alpha['X'] = 10;
        alpha['L'] = 50;
        alpha['C'] = 100;
        alpha['D'] = 500;
        alpha['M'] = 1000;
        
        int buf = 0;
        int sum = 0;
        
        for (char c: s){
            int num = alpha[c];
            
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
};
