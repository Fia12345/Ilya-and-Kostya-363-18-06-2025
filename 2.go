package main
import(
	"fmt"
	"time"
)

func worker(done chan bool){
	for {
		select{
		case<-done:
			fmt.Println("Завершение работы ")
			return
		default:
			fmt.Println("Работаем")
			time.Sleep(time.Second)
		}
	}
}

func main(){
	done := make(chan bool, 1)

	go worker(done)
	time.Sleep(3*time.Second)
	
	fmt.Println("\nПолучен сигнал завершения")
	done<-true

	time.Sleep(100*time.Millisecond)
}