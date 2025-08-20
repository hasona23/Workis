package helpers

import "fmt"

func LogError(err error) {
	if err != nil {
		fmt.Println("[ERROR] ", err)
	}
}

func LogInfo(str string) {
	fmt.Println("[INFO] ", str)
}
