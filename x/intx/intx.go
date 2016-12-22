package intx

func Contains(lst []int, v int) bool {
	for _, it := range lst {
		if it == v {
			return true
		}
	}
	return false
}

func Intersect(lst1, lst2 []int) (lst []int) {
	for _, it := range lst1 {
		if Contains(lst2, it) {
			lst = append(lst, it)
		}
	}
	return
}
