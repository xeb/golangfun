package channelsTest

import (
	"fmt"
	"github.com/xeb/golangfun/channels"
	"testing"
)

func TestSimpleChanValue(t *testing.T) {

	value := channels.SimpleChanValue()

	if value == 1 {
		t.Fatalf("\n\tBad test case\n")
	}

	fmt.Printf("\n\tDone! value %d\n", value)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}
}
