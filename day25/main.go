package main

import (
	"fmt"
)

func transform(subject, target int) int {
	return (target * subject) % 20201227
}

func loop(subject, loopsize int) int {
	target := 1
	for i := 0; i < loopsize; i++ {
		target = transform(subject, target)
	}
	return target
}

func find(subject int, public int) int {
	target := 1
	for i := 1; ; i++ {
		target = transform(subject, target)
		if target == public {
			return i
		}
	}
}

func day25a(subject, cardpub, doorpub int) {
	cardloop := find(subject, cardpub)
	doorloop := find(subject, doorpub)
	fmt.Println(cardloop, doorloop)
	fmt.Println(loop(cardpub, doorloop))
	fmt.Println(loop(doorpub, cardloop))
}

// This is a case where the compiled code really wins;
// on my machine Go ran over 50x faster than Python (.2 sec vs 10 sec)
func main() {
	day25a(7, 6929599, 2448427) // real
	// day25a(7, 5764801, 17807724) // sample
}
