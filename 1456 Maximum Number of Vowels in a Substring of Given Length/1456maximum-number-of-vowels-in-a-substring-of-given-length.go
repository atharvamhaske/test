func maxVowels(s string, k int) int {
    ct := 0
    res := 0

    vowels := map[byte]bool {
        'a' : true,
        'e' : true,
        'i' : true,
        'o' : true,
        'u' : true,
    }

    for i := 0; i < k; i++ {
        if vowels[s[i]] {
            ct++
        }
    }
    res = ct
    for i := k; i < len(s); i++ {
        
        if vowels[s[i-k]] {
            ct--
        } 
        if vowels[s[i]] {
            ct++
        }
        if ct > res {
            res = ct
        }
        
    }
    return res

}