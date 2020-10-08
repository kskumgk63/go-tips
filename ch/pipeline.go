package ch

type doneCh <-chan struct{}

type intStream <-chan int

func newIntStream(done doneCh, integers ...int) intStream {
	intStream := make(chan int, len(integers))
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

func multiply(done doneCh, intStream <-chan int, multiplier int) intStream {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream
}

func pipelineMultiply(intStream intStream, multipliers [3]int) []int {
	done := make(chan struct{})
	pipeline := multiply(done, multiply(done, multiply(done, intStream, multipliers[0]), multipliers[1]), multipliers[2])

	var result []int
	for i := range pipeline {
		result = append(result, i)
	}

	return result
}
