package misc

/*
	I intend this for semi-useful and likely redundant(wheel re-inventing) functions and structures;
*/
/**/
func IsNumInIntArray(num int, intArray []int) bool {
	answer := false
	for _, a := range intArray {
		if a == num {
			answer = true
		}
	}
	return answer
}

/**/
func IsNumInIntArrayAndWhat_singleint(voidValue, num int, intArray []int) int {
	answer := voidValue
	for _, a := range intArray {
		if a == num {
			answer = a
		}
	}
	return answer
}

/**/
func IsNumInIntArrayAndWhat_boolInt(num int, intArray []int) (bool, int) {
	answer0 := false
	answer1 := -7
	for _, a := range intArray {
		if a == num {
			answer0 = true
			answer1 = a
		}
	}
	return answer0, answer1
}
