class Solution {
public:
    bool isValid(string s) {
        map<char, char> mapping;
        mapping[')'] = '(';
        mapping['}'] = '{';
        mapping[']'] = '[';
        
        stack<char> check;
        
        for (char c: s){
            map<char, char>::iterator it(mapping.find(c));
            if (it != mapping.end()){
                if (check.empty() || it->second != check.top())
                    return false;
                check.pop();
            } else {
                check.push(c);
            }
        }
        
        return check.empty();
    }
};
