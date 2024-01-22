package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
