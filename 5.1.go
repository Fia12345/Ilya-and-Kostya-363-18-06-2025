package main 
import (
	"fmt"
	"time"
	"sync"

)

func worker(id int, done chan bool,wg*sync.WaitGroup){
	defer wg.Done()
	for{
		select {
		case<-done:
			fmt.Printf("Горутина %d завершает работу \n", id)
			return 
		default:
			fmt.Printf("Горутина %d работает\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main(){
	done:=make(chan bool)
	var wg sync.WaitGroup

	for i := 1; i<=3;i++{
		wg.Add(1)
		go worker(i, done, &wg )
	}
	time.Sleep(3*time.Second)
	close(done)
	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}