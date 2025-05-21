package main

import "fmt"

func main() {
	ar := []int{2, 3, 6, 7}
	stcount := 2

	sum := 0
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	
	st, end, ans := 0, sum-1, -1

	for st <= end {
		mid := st + (end - st)/2

		if IsValid(ar, stcount,mid) {
			ans=mid
			end = mid - 1

		} else {
			st = mid + 1

		}

	}

	fmt.Println(ans)

}

func IsValid(ar []int, stcount, maxAllowed int) bool {
	st, page := 1, 0

	for i := 0; i < len(ar); i++ {
		if ar[i] > maxAllowed {
			return false
		}
		if page+ar[i] <= maxAllowed {
			page += ar[i]
		} else {
			st++
			page = ar[i]
		}
	}

	return st == stcount
}
