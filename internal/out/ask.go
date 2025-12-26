package out

import (
	"fmt"
)

func Ask(question string) string {
	fmt.Print(question + " ")
	var response string
	fmt.Scanln(&response)
	return response
}
