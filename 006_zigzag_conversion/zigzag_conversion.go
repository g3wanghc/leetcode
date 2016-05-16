func convert(s string, numRows int) string {
    var converted_string string
 
    converted_string += filter_by_index(s, func(i int) bool {
    			return i % ((numRows - 1) * 2) == 0
    		})

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