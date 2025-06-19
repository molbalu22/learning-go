package concurrentprimes

import (
	"math"
	"sort"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func GeneratePrimeChunk(chunkStart, chunkEnd int, pr chan int, done chan bool) {
	defer func() {
		done <- true
	}()

	if chunkStart%2 == 0 {
		chunkStart++
	}

	for i := chunkStart; i < chunkEnd; i += 2 {
		isPrime := true
		for j := 3; j <= int(math.Sqrt(float64(i))); j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			pr <- i
		}
	}
}

func GeneratePrimes(n int) []int {
	if n < 2 {
		return []int{}
	}

	const numWorkers = 4

	chunkSize := n / numWorkers
	res := make([]int, 1, int(math.Sqrt(float64(n))))
	res[0] = 2

	pr := make(chan int)
	running := numWorkers
	done := make(chan bool)
	for i := range numWorkers {
		chunkStart := i * chunkSize
		chunkEnd := (i + 1) * chunkSize

		if i == 0 {
			chunkStart = 3
		}

		if i == numWorkers-1 {
			chunkEnd = n + 1
		}

		go GeneratePrimeChunk(chunkStart, chunkEnd, pr, done)
	}

	for {
		select {
		case d := <-done:
			if d {
				running--
			}
		case p := <-pr:
			res = append(res, p)
		}
		if running == 0 {
			break
		}
	}

	sort.Ints(res)

	return res
}
