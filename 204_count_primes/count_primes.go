func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }

	var index_arr []bool = make([]bool, n)
	memsetRepeat(index_arr, true)

	var factor int = 2
	var exhaust bool = false

	for !exhaust && float64(factor) <= math.Sqrt(float64(n - 1)) {
		var multiple int = 2
		for factor*multiple <= n - 1 {
			index_arr[factor*multiple] = false
			multiple++
		}
		factor++
		var next_match int = indexTrue(index_arr[factor:])
		if next_match != -1 {
			factor += next_match
		} else {
			exhaust = true
		}
	}
	
	var count int = 0
    for _, p := range index_arr {
        if p {
            count += 1
        }
    }
    
    return count - 2
}

func indexTrueReverse(a []bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] {
			return i
		}
	}
	return -1
}

func indexTrue(a []bool) int {
	for i, e := range a {
		if e == true {
			return i
		}
	}
	return -1
}

func memsetRepeat(a []bool, v bool) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}