package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 20; i += 1 {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
}
