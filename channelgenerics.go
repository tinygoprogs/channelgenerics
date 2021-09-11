package channelgenerics

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
