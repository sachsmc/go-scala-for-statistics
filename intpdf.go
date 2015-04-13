package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Dnorm(x float64) float64 {

	return math.Exp(-x*x/2) / math.Sqrt(2*math.Pi)

}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	N := 0
	if len(os.Args) == 1 {

		N = 5000

	} else {

		Na, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}

		N = Na

	}

	res := float64(0.0)
	for i := 0; i < N; i++ {

		func(i int) {
			candx := rand.Float64()*10 - 5
			candy := rand.Float64() * .5

			if candy <= Dnorm(candx) {
				res = res + 1
			}
		}(i)
	}

	// the rectangle is 10 by 0.5 = 5.0

	fmt.Fprintf(os.Stdout, "%f8\n", 5.0*res/float64(N))

}
