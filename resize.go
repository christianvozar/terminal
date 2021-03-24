package terminal

import (
	"fmt"
)

// printf '\e[8;23;79t' 

func Resize(x,y int) {
	fmt.Printf("printf '\\e[8;%d;%d'\n", x, y)
}