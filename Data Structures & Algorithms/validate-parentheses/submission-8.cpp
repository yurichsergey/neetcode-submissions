class Solution {
public:
    bool isValid(string s) {
        if (s.length() % 2 != 0) {
            return false;
        }

        std::stack<char> st;
        char open;
        for (char c : s) {
            open = (c == ')' ? '(' :
                ( c == '}' ? '{' :
                    ( c == ']' ? '[' :
                       '0'
                    )
                )
            );

            if (open == '0') {
                st.push(c);
            } else if (st.empty() || open != st.top()) {
                return false;
            } else {
                st.pop();
            }
        }
        return st.empty();
    }
};
