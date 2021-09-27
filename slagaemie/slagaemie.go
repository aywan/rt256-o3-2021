package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var n, val int
	scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		scanf("%d\n", &val)
		if check11(val) {
			printf("YES\n")
		} else {
			printf("NO\n")
		}
	}
}

func check11(val int) bool {
	digits := digitCount(val)

	for i := digits; i > 1; i-- {
		d := create11(i)

		k := val / d
		val -= k * d

		if val == 0 {
			return true
		}

	}

	return false
}

func create11(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res = res*10 + 1
	}
	return res
}

func digitCount(val int) int {
	digits := 0
	for val > 0 {
		val = val / 10
		digits++
	}
	return digits
}
