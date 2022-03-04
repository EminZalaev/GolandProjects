//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func main() {
	go SendingData()
	time.Sleep(2 * time.Second)

}

func SendingData() {
	ch := make(chan int)
	//создаем и закрываем канал, что бы горутина чтения завершилась. так же закидываем в канал данные
	defer close(ch)
	//вызываем горутину чтения из канала
	go dataReception(ch)
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second / 10)
	}

}

func dataReception(ch chan int) {
	for {
		//читаем данные из канала
		for val := range ch {
			fmt.Println(val)
		}
	}
}
