/*
A few small generic channel utilities.
*/
package channelgenerics

// Join joins the two input channel in the one resulting channel.
// Note: Spawns a goroutine, than terminates when both input channels are
// closed.
func Join[T any](a, b <-chan T) <-chan T {
	c := make(chan T)
	go func() {
		adone := false
		bdone := false
		for {
			select {
			case elm, ok := <-a:
				if !ok {
					adone = true
				} else {
					c <- elm
				}
			case elm, ok := <-b:
				if !ok {
					bdone = true
				} else {
					c <- elm
				}
			}
			if adone && bdone {
				close(c)
				return
			}
		}
	}()
	return c
}

type Fn[T any] func(t T) bool

// Filter filters the input channel using fn, if fn returns true for an element
// it is pushed into the resulting channel.
// Note: Spawns a goroutine, than terminates when the input channel is closed.
func Filter[T any](in <-chan T, fn Fn[T]) <-chan T {
	out := make(chan T)
	go func() {
		for {
			select {
			case elm, ok := <-in:
				if !ok {
					close(out)
					return
				}
				if fn(elm) {
					out <- elm
				}
			}
		}
	}()
	return out
}
