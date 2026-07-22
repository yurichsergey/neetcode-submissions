
import (
    "unicode/utf8"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var res strings.Builder
    for _, s := range strs {
        l := utf8.RuneCountInString(s)
        res.WriteString(strconv.Itoa(l) + "#" + s)
    }
    fmt.Println(res.String())
    return res.String()
}

func (s *Solution) Decode(encoded string) []string {
    readSize := true
    var sizeStr strings.Builder
    size := 0
    count := 0

    var word strings.Builder

    res := []string{}
    for _, r := range encoded {
        if (readSize && r != '#') {
            sizeStr.WriteRune(r)
            continue            
        }

        if (readSize) {
            num, err := strconv.Atoi(sizeStr.String())
            if err != nil {
                fmt.Printf("Error strconv.Atoi(sizeStr)")
                return []string{}
            }
            size = num
            readSize = false
            sizeStr.Reset()

            if size == 0 {
                res = append(res, "")
                readSize = true
            }

            continue
        }

        count ++
        word.WriteRune(r)
        if (count >= size) {
            readSize = true
            res = append(res, word.String())
            word.Reset()
            count = 0
        }
    }

    return res
}
