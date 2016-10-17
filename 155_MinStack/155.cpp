class MinStack {
public:
    /** initialize your data structure here. */
    MinStack(): min(INT_MAX) {
        
    }
    
    void push(int x) {
        if (x <= min){
            // Push second min value for pop()
            stack.push(min);
            min = x;
        }
        
        stack.push(x);
    }
    
    void pop() {
        // If pop value is min, update min to second min
        int pop_value = stack.top();
        stack.pop();

        if (pop_value == min){
            min = stack.top();
            stack.pop();
        }
    }
    
    int top() {
        return stack.top();
    }
    
    int getMin() {
        return min;
    }
    
private:
    std::stack<int> stack;
    int min;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */