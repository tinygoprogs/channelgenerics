package main

import (
	"context"
	"fmt"
	cg "github.com/tinygoprogs/channelgenerics"
)

func main() {
	// generic channel utility usage
	a := make(chan string)
	b := make(chan string)
	c := cg.Filter(cg.Join(a, b), func(s string) bool {
		if s == "b1" {
			return false
		}
		return true
	})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case elm, ok := <-c:
				if !ok {
					cancel()
					return
				}
				fmt.Println(elm)
			}
		}
	}()
	a <- "a1"
	b <- "b1"
	a <- "a2"
	b <- "b2"
	close(a)
	close(b)
	select {
	case <-ctx.Done():
		fmt.Println("done")
	}
}
