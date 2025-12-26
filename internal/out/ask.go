package out

import (
	"fmt"

	"go.mattglei.ch/timber"
)

func Ask(question string) string {
	fmt.Print(question + " ")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		timber.Fatal(err, "failed to ask question")
	}
	return response
}
