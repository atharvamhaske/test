func nextPermutation(nums []int)  {
    n := len(nums)

    peakIdx := n-1
    largeIdx := n-1

    //find the final idx of increasing order
    for; peakIdx > 0; peakIdx-- {
        if nums[peakIdx-1] < nums[peakIdx] {
            break
        }
    }

    //find the largest element from right to peak, when compared
    //to the element immediate left of the peakIdx

    if peakIdx != 0 {
        for; largeIdx >= peakIdx; largeIdx-- {
            if nums[peakIdx-1] < nums[largeIdx] {
                //swap the elements
                nums[largeIdx], nums[peakIdx-1] = nums[peakIdx-1], nums[largeIdx]
                break
            }
        }
    }

    // Reverse the order of elements starting from the end of the array, 
    //uptill the peak
    for i:=n-1; i>peakIdx; i, peakIdx = i-1, peakIdx+1 {
        nums[i], nums[peakIdx] = nums[peakIdx], nums[i]
    }
}