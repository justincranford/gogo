package timesutil

import (
	"fmt"
	"time"
)

func TimeTests() {
	start := time.Now()
	fmt.Println("Cranford", "Justin", 1+1, "test", true && false, true || false, 0x11001100^0x10101010)
	t := time.Now()
	fmt.Println("Time now", t)

	time.Sleep(10 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println("Elapsed", elapsed)
}
