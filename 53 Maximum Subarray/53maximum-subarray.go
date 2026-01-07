func maxSubArray(nums []int) int {
    curSum := nums[0]
    maxSum := nums[0]

    for _, val := range nums[1:] {
        if curSum+val > val {
            curSum = curSum + val
        } else {
            curSum = val
        }

        if curSum > maxSum {
            maxSum = curSum
        }
    }
    return maxSum
}
