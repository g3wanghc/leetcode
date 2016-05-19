func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    r, _ := regexp.Compile("((-|\\+){0,1}[0-9]+)")

    var ind []int = r.FindStringIndex(s)
    if len(ind) == 0 || ind[0] != 0 {
        return 0
    }    
    s = r.FindString(s)

    i, err := strconv.ParseInt(s, 10, 32)
	
    if err != nil {
	    return int(i)
    }
	
    return int(i)
}