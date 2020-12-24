// import "fmt"

func twoSum(nums []int, target int) []int {
    var seen = make(map[int]int)
    for i, num := range nums {
        if prevIndex, isFound := seen[target - num]; isFound  {
            // fmt.Print(prevIndex)
            return []int {prevIndex, i};
        }
        seen[num] = i
    }
    return nil
}
