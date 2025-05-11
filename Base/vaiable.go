package main

import (
	fm "fmt"
	"math/rand"
)

func Variable() {

	// Khai báo biến cơ bản
	var number int = 5	// 4 byte trên máy 32bit và 64 byte trên máy 64bit
	var checked bool = true

	// unit: số nguyên không dấu, chỉ được lưu trữ giá trị không âm
	var xx uint = 500
	fm.Printf("Type: %T, value: %v \n", xx, xx)


	// float32 đảm bảo chính xác đến khoảng 7 chữ số thập phân và float64 cho khoảng 15 chữ số thập phân.
	// Hãy luôn sử dụng float64 khi có thể vì tất cả các hàm của gói math mong đợi kiểu dữ liệu này.
	var so_thuc_1 float32 = 3.15
	var so_thuc_2 float64 = 3.19999

	// Số phức (complex64 (mỗi phần thực và phần ảo là 32 bit), complex128 (mỗi phần thực và phần ảo là 64 bit))
	var c1 complex64 = 5 + 10i

	// Ưu tiên sử dụng. Nhưng chỉ được sử dụng trong hàm, không phải trong phạm vi package
	other_num := 10
	other_checked := false

	// Hằng số được coi như không có kiểu dữ liệu
	const HANGSO = 999

	fm.Println(number, checked, other_num, other_checked, so_thuc_1, so_thuc_2, c1)
	fm.Println(HANGSO)


	var a, b = 6, "Hello"
	c, d := 7, "World!"
	fm.Println(a)
	fm.Println(b)
	fm.Println(c)
	fm.Println(d)

	var (
		x int
		y int = 1
		z string = "hello"
	  )
   
	fm.Println(x)
	fm.Println(y)
	fm.Println(z)

	random_x := rand.Int()
	fm.Printf("a is: %d\n", random_x)

	// Byte. Kiểu byte là một bí danh cho uint8 và điều này phù hợp với việc mã hóa ký tự sử dụng bảng ASCII truyền thống (1 byte)
	var ch1 byte = 'A'
	var ch2 byte = 65
	var ch3 byte = '\x41'

	fm.Println(ch1, ch2, ch3)

	b3 := 10 > 5 && 7 < 15     // AND operator
    fm.Println(b3)
    b3 = 10 < 5 || 2 > 7       // OR operator
    fm.Println(b3)
    b3 = !b3                   // NOT operator
    fm.Println(b3)

	/*
		Các chuỗi và file văn bản Go chiếm ít bộ nhớ/dung lượng đĩa hơn
			(do có các ký tự có độ rộng biến thiên).
		Vì sử dụng chuẩn UTF-8 nên Go không cần phải mã hóa và giải mã chuỗi như những ngôn ngữ khác.

		Chuỗi là kiểu giá trị và không thể thay đổi,
			điều đó có nghĩa là sau khi tạo chuỗi,
			không thể thay đổi nội dung của chuỗi.
			Nói cách khác, chuỗi là một mảng byte bất biến.

		Có 2 loại chuỗi: Chuỗi tự giải thích và chuỗi nguyên thuỷ
	*/

	chuoi_tu_giai_thich := "Tự hiểu được xuống dòng \n" // represents a newline
	chuoi_nguyen_thuy := `In nguyên ra cả \n thay vì xuống dòng`	// Có thể trải rộng được trên nhiều dòng
	fm.Println(chuoi_tu_giai_thich)
	fm.Println(chuoi_nguyen_thuy)
	fm.Println(len(chuoi_nguyen_thuy))

}