package main

func merge(s[] int, mid int)  {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft, helperRight := 0, mid
	current := 0
	high := len(s)

	for helperLeft < mid && helperRight < high{
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft < mid {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func mergesort(s[] int)  {
	if len(s) > 1 {
		mid := len(s)/2
		mergesort(s[:mid])
		mergesort(s[mid:])
		merge(s, mid)
	}
}

//func main() {
//	a := []int{5,5,6,1,2,3,9,7,5,0}
//	mergesort(a)
//	fmt.Println(a)
//}
