package main
import "fmt"

func ForExample() {
	for i := 0; i < 5; i++ {}

	// for { }

	str := "Go is a beautiful language!"
	for pos, char := range str {
        fmt.Printf("Character on position %d is: %c \n", pos, char)
    }

	/*
	labelName:
	for condition {
		// ...
		break labelName
		continue labelName
	}


	Sử dụng label để điều khiển luồng chương trình chính xác hơn
	*/


}