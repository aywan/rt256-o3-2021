package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	boxes := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &boxes[i])
	}
	sort.Ints(boxes)

	c := 0
	for i := 0; i < n; i = i + 2 {
		c += boxes[i+1] - boxes[i]
	}

	printf("%d\n", c)
}
