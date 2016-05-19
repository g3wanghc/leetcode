func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    var s string = strconv.Itoa(x)
    var l int = len(s)
    for i := 0; i < int(math.Floor(float64(l/2))); i++ {
        if s[i] != s[l - i - 1] {
            return false
        }
    }

    return true
}