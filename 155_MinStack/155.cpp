class MinStack {
public:
    /** initialize your data structure here. */
    MinStack(): minVal(INT_MAX) {
        
    }
    
    void push(int x) {
        if (x < minVal){
            minVal = x;
        }
        
        stack.push_back(x);
    }
    
    void pop() {
        int currentMin = minVal;
        stack.pop_back();
        if (currentMin <= minVal){
            minVal = INT_MAX;
            for (int &val : stack){
                if (val < minVal){
                    minVal = val;
                }
            }
        }
    }
    
    int top() {
        return stack.back();
    }
    
    int getMin() {
        return minVal;
    }
    
private:
    vector<int> stack;
    int minVal;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */