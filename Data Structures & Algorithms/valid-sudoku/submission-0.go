func isValidSudoku(board [][]byte) bool {

    check := func(row []byte) bool {
        set := make(map[byte]struct{}, len(row))
        for _, b := range row {
            if string(b) == "." {
                continue
            }
            _, exists := set[b]
            if exists {
                return false
            }
            set[b] = struct{}{}
        }
        return true
    }

    for i := range board {
        if !check(board[i]) {
            return false
        }
    }

    for i := range board[0] {
        column := make([]byte, len(board), len(board))
        for j := 0; j < len(board); j++ {
            column[j] = board[j][i]
        }
        if !check(column) {
            return false
        }
    }

    for i := 0; i < 3; i++ { // horizontal
        for j := 0; j < 3; j++ { // vertical

            q := []byte{
                board[0 + i*3][0 + j*3], board[0 + i*3][1 + j*3], board[0 + i*3][2 + j*3],
                board[1 + i*3][0 + j*3], board[1 + i*3][1 + j*3], board[1 + i*3][2 + j*3],
                board[2 + i*3][0 + j*3], board[2 + i*3][1 + j*3], board[2 + i*3][2 + j*3],
            }
            if !check(q) {
                return false
            }

        }
    }

    return true
}
