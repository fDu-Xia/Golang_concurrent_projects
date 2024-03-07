package out

import (
	"fmt"
)

func Out(OutputBuffer chan string) {
	for i := range OutputBuffer {
		fmt.Println(i)
	}
}
