class Solution {
public:
    int fourSumCount(vector<int>& A, vector<int>& B, vector<int>& C, vector<int>& D) {
        unordered_map<int, int> hashTable;
        int count = 0;

        for (const int& a : A) {
            for (const int& b : B) {
                int value = a + b;
                auto iter = hashTable.find(value);
                if (iter == hashTable.end()) {
                    hashTable[value] = 0;
                }
                hashTable[value]++;
            }
        }


        for (const int& c : C) {
            for (const int& d : D) {
                int value = c + d;
                auto iter = hashTable.find(-value);
                if (iter != hashTable.end()) {
                    count += iter->second;
                }
            }
        }

        return count;
    }
};
