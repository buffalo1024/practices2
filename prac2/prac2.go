package main

func Func(a, b []int) []int {
	newArray := []int{}

	lengthA := len(a)
	lengthB := len(b)
	i := 0
	j := 0
	for {
		if i >= lengthA {
			break
		}
		if j >= lengthB {
			newArray = append(newArray, a[i:]...)
			break
		}
		if a[i] < b[j] {
			newArray = append(newArray, a[i])
			i = i + 1
			continue
		}
		if a[i] == b[j] {
			i = i + 1
			continue
		}
		if a[i] > b[j] {
			j = j + 1
			continue
		}
	}

	return newArray
}
