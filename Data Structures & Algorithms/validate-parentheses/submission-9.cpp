class Solution {
public:
    bool isValid(string s) {
        if (s.length() % 2 != 0) {
            return false;
        }

        string st;
        st.reserve(s.length());
        for (char c : s) {
            if (c == '(') st.push_back(')');
            else if (c == '{') st.push_back('}');
            else if (c == '[') st.push_back(']');
            else {
                // other symbol - i expect closed symbol
                if (st.empty() || c != st.back()) {
                    return false;
                }
                st.pop_back();
            }
        }
        return st.empty();
    }
};
