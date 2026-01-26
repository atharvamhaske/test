func findMin(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    } 

    i := len(nums) / 2
    left := findMin(nums[0:i])
    right := findMin(nums[i:len(nums)])

    return min(left, right)
}