package terminal

import (
	"fmt"
)

const (
	ScreenSizeSeq = "%d;%dt"
)

// printf '\e[8;23;79t' 

func Resize(x,y int) {
	fmt.Printf("printf '\\e[8;%d;%dt'\n", x, y)
}