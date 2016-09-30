/* Notes about finalizers.
 *
 * SetFinalizer will not work on an empty struct.
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

type empty struct {
}

type nonEmpty struct {
	x int
}

var counts struct {
	empty    int
	nonEmpty int
}

func do() {
	e := &empty{}
	ne := &nonEmpty{}

	runtime.SetFinalizer(e, func(x *empty) {
		// fmt.Printf("finalizer called for empty\n")
		counts.empty++
	})
	runtime.SetFinalizer(ne, func(x *nonEmpty) {
		// fmt.Printf("finalizer called for nonEmpty\n")
		counts.nonEmpty++
	})
}

func main() {
	last := time.Now().Unix()

	for {
		do()
		runtime.GC()

		now := time.Now().Unix()
		if now-last > 1 {
			fmt.Printf("%+v\n", counts)
			last = now
		}
	}
}
