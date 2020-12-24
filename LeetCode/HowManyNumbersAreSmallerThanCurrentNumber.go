// O(n2) approach
func smallerNumbersThanCurrent(nums []int) []int {
    res := make([]int, len(nums))
    for i := range nums {
        res[i] = countSmaller(nums, i)
    }
    return res
}

func countSmaller(nums []int, index int) int {
    count := 0
    for _, num := range nums {
        if (num < nums[index]) {
            count++
        }
    }
    return count
}

