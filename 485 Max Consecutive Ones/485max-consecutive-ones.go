func findMaxConsecutiveOnes(nums []int) int {
    ct := 0
    mxct := 0

    for _, val := range nums {
        if val == 1 {
            ct +=1
            mxct = max(mxct, ct)
        } else {
            ct = 0
        }
    }

    return mxct
}

