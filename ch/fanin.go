package ch

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeat(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	fmt.Println("INVOKE repeat()")
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
				fmt.Println("case valueStream <- fn(): repeat()")
			}
		}
	}()
	fmt.Println("RETURN repeat()")
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	fmt.Println("INVOKE take()")
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
				fmt.Println("case takeStream <- <-valueStream: take()")
			}
		}
	}()
	fmt.Println("RETURN repeat()")
	return takeStream
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	fmt.Println("INVOKE toInt()")
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
				fmt.Println("case intStream <- v.(int): toInt()")
			}
		}
	}()
	fmt.Println("RETURN toInt()")
	return intStream
}

func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	fmt.Println("INVOKE primeFinder()")
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer--
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
					fmt.Println("case primeStream <- integer: primeFinder()")
				}
			}
		}
	}()
	fmt.Println("RETURN primeFinder()")
	return primeStream
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	fmt.Println("INVOKE fanIn()")
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
				fmt.Println("case multiplexedStream <- i: fanIn()")
			}
		}
	}

	// Select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	fmt.Println("RETURN fanIn()")
	return multiplexedStream
}

func fanInMain() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	rand := func() interface{} { return rand.Intn(50000000) }

	randIntStream := toInt(done, repeat(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	takenCount := 10
	for prime := range take(done, fanIn(done, finders...), takenCount) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))
}
