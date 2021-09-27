package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var parties int
	scanf("%d\n", &parties)

	for i := 0; i < parties; i++ {
		calculate()
	}
}

func calculate() {
	var n, val int64
	scanf("%d\n", &n)

	eventsCnt := int64(1) << (n - 1)
	k := eventsCnt
	sum := big.NewInt(0)

	for i := int64(0); i < n; i++ {

		for j := int64(0); j <= i; j++ {

			scanf("%d", &val)
			bigVal := big.NewInt(val)

			if j == 0 || j == i {
				sum.Add(sum, big.NewInt(0).Mul(big.NewInt(k), bigVal))
			} else {
				binK := big.NewInt(0).Binomial(i, j)
				valK := new(big.Int).Mul(binK, big.NewInt(int64(1)<<(n-i-1)))
				sum.Add(sum, new(big.Int).Mul(valK, bigVal))
			}
		}
		k = k >> 1
		scanf("\n")
	}

	if sum.Cmp(big.NewInt(0)) == 0 {
		printf("0 1\n")
	} else {
		events := big.NewInt(eventsCnt)
		d := new(big.Int).GCD(nil, nil, sum, events)

		printf("%s %s\n", new(big.Int).Div(sum, d).String(), new(big.Int).Div(events, d).String())
	}
}
