//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, ch chan *Tweet) {
	defer close(ch)

	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}

		ch <- tweet
	}
}

func consumer(ch chan *Tweet) {
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	ch := make(chan *Tweet)

	// Consumer
	go consumer(ch)

	// Producer
	producer(stream, ch)

	fmt.Printf("Process took %s\n", time.Since(start))
}
