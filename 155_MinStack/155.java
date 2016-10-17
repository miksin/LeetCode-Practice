public class MinStack {

    ArrayList<Integer> stack;
    int minVal;

    /** initialize your data structure here. */
    public MinStack() {
        this.stack = new ArrayList<Integer>();
        this.minVal = Integer.MAX_VALUE;
    }
    
    public void push(int x) {
        this.stack.add(x);
        if (x < this.minVal){
            this.minVal = x;
        }
    }
    
    public void pop() {
        int popVal = this.top();
        this.stack.remove(this.stack.size() - 1);
        
        if (popVal <= this.minVal){
            if (this.stack.isEmpty()){
                this.minVal = Integer.MAX_VALUE;
            } else {
                this.minVal = Collections.min(this.stack);
            }
        }
    }
    
    public int top() {
        return this.stack.get(this.stack.size() - 1);
    }
    
    public int getMin() {
        return this.minVal;
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */