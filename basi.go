package main
import ("fmt"
		//"math"
		"bufio"
		"os"
		)




func main() {
	
	//multiplicar(3)
	//bizz_bozz()
	//bonacci()
	fmt.Println(add(42, 13))
	a, b := swap("hola", "mundo")
    fmt.Println(a, b)

    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)

    //arrays()
    read_write()
}

func multiplicar(num int) {
	var res int
	for i := 1; i <=10; i++ {
		res = num * i 
		fmt.Println(num, "*", i, "=", res)

	}
}

func bizz_bozz() {
	for i := 1; i <=100; i++ {
			if i % 3 == 0  && i % 5 == 0 {
				fmt.Println("BizzBozz")
			}else if i % 3 == 0{
				fmt.Println("Bizz")
			}else if i % 5 == 0 {
				fmt.Println("Bozz")
			}else{
				fmt.Println(i)
			}
		}	
}

func  fibonacci() {
	var x, anterior int  = 1, 0
	var temp int 
	for i := 1; i <= 20; i++ {
		fmt.Println(x)
		temp = x
		x = x + anterior
		anterior = temp
	}
}

func add(x int, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}



type Vertex struct {
    X int
    Y int
}

func arrays() {
	p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)
    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n",i, p[i])
    }
}

func read_write() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+ name)
	}
}

