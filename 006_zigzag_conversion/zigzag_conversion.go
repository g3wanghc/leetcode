func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    var converted_string string

    for r := 0; r < numRows; r++ {
        converted_string += filter_by_index(s, func(i int) bool {
            if i % ((numRows - 1) * 2) == r {
                return true
            }
            if (i + r) % ((numRows - 1) * 2) == 0 {
                return true
            }
            return false
        })
    }

    return converted_string
}

    
func filter_by_index(s string, f func(int) bool) string {
    var sf string
    for i, c := range s {
        if f(i) {
            sf += string(c)
        }
    }
    return sf
}