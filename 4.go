package main
import(
	"fmt"
	"time"
)
func main(){
	highPriority :=make(chan string, 100)
	lowPriority :=make(chan string,100)
	go func(){
		for{
			select{
			case task:= <-highPriority:
				fmt.Printf("Срочная задача: %s \n",task)
			default: 
			select{
			case task:= <-highPriority:
				fmt.Printf("Срочная задача: %s \n",task)
			case task:= <-lowPriority:
				fmt.Printf("Обычная задача: %s \n",task)
			}
		}
		time.Sleep(100*time.Millisecond)
	}
}()
lowPriority<-"отчёт за неделю"
highPriority<-"Срочной платёж"
lowPriority<-"Обновление БД"
highPriority<-"Критический баг "  

time.Sleep(1*time.Second)
}