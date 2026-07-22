import "slices"

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    uniq := map[int]struct{}{} 
    for _, i := range nums {
        uniq[i] = struct{}{}
    }

    uniqNums := make([]int, 0, len(uniq))
    for num := range uniq {
        uniqNums = append(uniqNums, num)
    }
    slices.Sort(uniqNums)

    fmt.Println(uniqNums)
    maxLen := 1
    curLen := 1

    currInd := 0
    for i := 1; i < len(uniqNums); i++ {
        fmt.Println(curLen, currInd, i, uniqNums[currInd] + i, uniqNums[i])
        if uniqNums[currInd] + i - currInd != uniqNums[i] {
            currInd = i
            curLen = 1
            continue;
        }

        curLen ++

        if curLen > maxLen {
            maxLen = curLen
        }
    }
    fmt.Println(maxLen)
    return maxLen
}
