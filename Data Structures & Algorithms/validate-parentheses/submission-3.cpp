class Solution {
public:
    bool isValid(string s) {
        if (s.length() % 2 != 0) {
            return false;
        }

        unordered_map<char, char> p = {{')', '('}, {']', '['}, {'}', '{'}};
        std::stack<char> st;
        char last;
        for (char c : s) {
            switch (c) {
                case ')':
                    last = '(';
                    break;
                case '}':
                    last = '{';
                    break;
                case ']':
                    last = '[';
                    break;
                default:
                    last = '0';
            }

            if (last == '0') {
                st.push(c);
            } else if (st.empty() || last != st.top()) {
                return false;
            } else {
                st.pop();
            }
        }
        return st.empty();
    }
};
