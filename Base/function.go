package main
import "fmt"

func Function() {

	// defer: Trì hoãn việc thực thi hàm hoặc câu lệnh cho đến khi hàm bao bọc (Hàm gọi nó kết thúc)
    // Chỉ trì hoãn hàm chứ không trì hoãn đối số trong hàm (TH đối số được truyền vào là 1 hàm)
    // Nếu có nhiều lệnh defer trong hàm, chúng sẽ được thực thi theo thứ tự ngược lại với thứ tự chúng được gọi (như ngăn xếp hoặc LIFO),
    //  điều này có nghĩa là câu lệnh defer cuối cùng sẽ được thực thi đầu tiên và cứ như thế
	Function1()

    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d", i)
    }

    f()

}

func Function1() {
    // Ở đây thì sẽ chạy 2 cái printf trước sau đó mới gọi đến Function2()


    fmt.Printf("In Function1 at the top\n")
    defer Function2() // function deferred, will be executed after Function1 is done.
    fmt.Printf("In Function1 at the bottom!\n")
}

func Function2() {
    fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}



func f() {
    for i := 0; i < 4; i++{

        // g được gọi là hàm ẩn danh
        g := func(i int) {fmt.Printf("%d ", i)} // g assigned to the function literal
        g(i) // g used as a function
        fmt.Printf(" - g is of type %T and has value %v\n", g, g)
    }

    fplus := func(x, y int) int { return x + y }
    fplus(3, 4)

}