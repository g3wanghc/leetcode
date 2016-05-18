func reverse(x int) int {
    var s string = strconv.Itoa(x)
    var negative bool = s[0] == byte('-')
    
    if negative {
        s = s[1:]
    }
    
    s = Reverse(s)
    
    i, err := strconv.ParseInt(s, 10, 32)
    
    if err != nil {
        return 0
    }
    
    if negative {
        i = -i
    }
    
    return int(i)
}

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}