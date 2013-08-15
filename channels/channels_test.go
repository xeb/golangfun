package channels

import (
	"fmt"
	"testing"
)

// func TestChanWorks(t *testing.T) {

// 	value := ChanTest()

// 	if value == 1 {
// 		t.Fatalf("\n\tBad test case\n")
// 	}

// 	fmt.Printf("\n\tDone! value %d\n", value)
// }

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// fmt.Sprintf("Hello")
	}
}