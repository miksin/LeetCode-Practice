class Solution {
public:
    int myAtoi(string str) {
        if (!std::regex_match(str, std::regex("^( )*(\\+|\\-)?([0-9])+.*$", std::regex::ECMAScript)))
            return 0;
            
        str.erase(0, str.find_first_of("0123456789-"));
        str.erase(str.begin()+str.find_last_of("0123456789")+1, str.end());

        int result = 0;
        
        try {
            result = stoi(str);
        } catch (...) {
            if (str.find("-") == string::npos)
                result = INT_MAX;
            else
                result = INT_MIN;
        }

        return result;
    }
};
