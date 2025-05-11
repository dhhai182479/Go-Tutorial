package main
import fm "fmt"	// Alias

// Hàm main là điểm bắt đầu của GO. Khi chạy, hàm main sẽ được gọi đầu tiên 
func Hello() {
	fm.Println("Hello, world!")

	text := "Hehehaha"

	fm.Printf("Orther Print: %s\n", text)	// Thường sử dụng format string
	/* 
		func Printf(format string, list of variables to be printed)

		%d - định dạng các giá trị số nguyên
		%s - định dạng các giá trị chuỗi.
		%v - định dạng mặc định chung
	*/

	// Print(): In ra định dạng nguyên bản của đối số
	var i,j string = "Hellooooooo","Worldddddd"
	fm.Print(i)
	fm.Print(j, "\n")


	// Để nhập giá trị từ bàn phím
	// Scan(a ...any) (n int, err error)

	var age int
    var name string
    fm.Print("Enter your age and name:")
    fm.Scan(&age, &name)
    fm.Println("Name:", name)
    fm.Println("Age:", age)




}	