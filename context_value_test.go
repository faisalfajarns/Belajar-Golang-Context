package belajargolangcontex

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContextValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("g"))

	contextA.Done()
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
				select{
				case <- ctx.Done():

					return 
				default :
				destination <- counter
				counter++
				time.Sleep(time.Second)
				}

		}
	}()

	return destination
}


func TestContextCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	 
	parent:= context.Background()
	// ctx with cancel
	ctx, cancel := context.WithCancel(parent) 
	group := sync.WaitGroup{}
	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n:= range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	
	}
	
	time.Sleep(5* time.Second)
	group.Wait()
	cancel() // mengirim sinyal cancel ke context

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextCancelWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	 
	parent:= context.Background()
	// ctx cancel with timeout
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second) 
	// group := sync.WaitGroup{}

	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n:= range destination {
		fmt.Println("Counter", n)
	
	
	}
	
	time.Sleep(2* time.Second)
	// group.Wait()
	cancel() // mengirim sinyal cancel ke context

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextCancelWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	 
	parent:= context.Background()
	// ctx cancel with timeout
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second)) 
	// group := sync.WaitGroup{}

	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n:= range destination {
		fmt.Println("Counter", n)
	
	
	}
	
	time.Sleep(2* time.Second)
	// group.Wait()
	cancel() // mengirim sinyal cancel ke context

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}