package wordcount

import "strings"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func wordCountsInString(str chan string, count chan map[string]int) {
	for s := range str {
		c := make(map[string]int)
		for word := range strings.FieldsSeq(s) {
			c[word] += 1
		}
		count <- c
	}
}

func CountWords(strings []string) map[string]int {
	res := make(map[string]int)

	const numWorkers = 4

	str := make(chan string)
	count := make(chan map[string]int)

	go func() {
		defer close(str)
		for _, s := range strings {
			str <- s
		}
	}()

	for range numWorkers {
		go wordCountsInString(str, count)
	}

	for range len(strings) {
		c := <-count
		for k, v := range c {
			res[k] += v
		}
	}

	return res
}
