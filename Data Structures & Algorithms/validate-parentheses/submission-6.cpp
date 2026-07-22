class Solution {
public:
    bool isValid(string s) {
        unordered_map<char, char> p = {{')', '('}, {']', '['}, {'}', '{'}};
        std::stack<char> st;
        for (char c : s) {
            if (p.count(c)) {
                if (st.empty() || p[c] != st.top()) {
                    return false;
                }
                st.pop();
            } else {
                st.push(c);
            }
        }
        return st.empty();
    }
};
