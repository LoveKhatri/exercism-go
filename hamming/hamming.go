package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("lengths are not equal")
	}

	count := 0

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
