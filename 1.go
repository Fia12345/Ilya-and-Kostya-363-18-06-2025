package main 
import(
   "fmt"
   "time"
)

func main(){
	timeOut := 4 * time.Second // вместо цифры можно поставить другое число

	resultChan:= make(chan string)
	go longRequest(resultChan)
	select{
	case result := <-resultChan :
		fmt.Println("ответ наверное", result)
	case<- time.After(timeOut):
		fmt.Println("Ошибка")
	}

}

func longRequest(resultChan chan <-string) {
	time.Sleep(3*time.Second)
	resultChan<-"Данные получены"
}




