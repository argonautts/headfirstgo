// pass_fail сообщает, сдал ли пользователь экзамен.
package chrpt4

import (
	"fmt"
	"headFirst/chrpt4/keyboard"
	"log"
)

func Pass_fail() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
