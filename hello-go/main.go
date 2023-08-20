package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func soma(x int, y int) (int, bool) {
	res := x + y
	isBetter := x > y

	if x > 10 {
		return -1, false
	}

	return res, isBetter
}

type Course struct { // primeira letra no comeco maiusculo => publico, se minusculo privado
	Name        string `json:"name"`
	Description string `json:"description"` // isso sao tags para retornar min no json
	Price       int    `json:"price"`
}

func (c Course) GetFullInfo() string { //metodo da struct
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

// T1
func main() {

	channel := make(chan int)
	for i := 0; i < 1000000; i++ {
		go worker(i, channel) // cada thread custará 2k de memória
	}
	for i := 0; i < 10000000; i++ {
		channel <- i
	}

	// channel := make(chan string)

	// //T2
	// go func() {
	// 	channel <- "Hello World!" // publica no canal
	// }()
	// fmt.Println(<-channel) // lê de dentro do canal

	// var a string
	// a = "Hello, world!"
	// b := "Hello, world!" // := -> declara e seta a variável.
	// println(a, b)

	// //server em go
	// http.HandleFunc("/", home)
	// http.ListenAndServe(":8080", nil)

	// // go routines, o main tem 1 thread e os go tem +1 kd
	// go counter() // T2
	// go counter() // T3
	// counter()
}

func home(w http.ResponseWriter, r *http.Request) {

	//struct
	course := Course{
		Name:        "Full Cycle",
		Description: "Golang Course",
		Price:       100,
	}
	fmt.Println(course.Name, course.GetFullInfo())
	json.NewEncoder(w).Encode(course)
	w.Write([]byte("Hello, world!"))
}
