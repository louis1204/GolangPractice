func kLengthApart(nums []int, k int) bool {
    curr := k
    i := 0;
    for i < len(nums) && nums[i] != 1 {
        i++
    }
    for ; i < len(nums); i++ {
        if nums[i] == 0 {
            curr++
        } else {
            if curr < k {
                return false
            }
            curr = 0
        }
    }
    return true
}
