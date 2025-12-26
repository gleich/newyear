package out

import (
	"fmt"

	"github.com/gleich/lumber/v2"
)

func Ask(question string) string {
	fmt.Print(question + " ")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		lumber.Fatal(err, "failed to ask question")
	}
	return response
}
