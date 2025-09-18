package main
import(
	"fmt"
						"time"
)

func main(){
	ch:= make(chan int )
	select{
	case <- ch:	
	fmt.Println("Данные получены")
	default:
		fmt.Println("Данных нет")
	}
	go func(){
		time.Sleep(100*time.Millisecond)
		ch<- 42
	}()
	time.Sleep(200*time.Millisecond)
	select{
	case<-ch:
		fmt.Println("Данные получены")
	default:
		fmt.Println("Данных нет")
	}
}