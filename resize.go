package terminal

import (
	"fmt"
)

const (
	ScreenSizeSeq = "%d;%dt"
)

func Resize(x,y int) {
	fmt.Printf(CSI+ScreenSizeSeq, y, x)
}
