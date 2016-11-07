class Solution {
public:
    int reverse(int x) {
        stringstream ss;
        string s;
        int result;
        
        ss << abs(x);
        ss >> s;
        
        std::reverse(s.begin(), s.end());
        
        try {
            result = stoi(s);
        } catch (...) {
            result = 0;
        }
        
        result = (x<0)? -result: result;
        
        return result;
    }
};
