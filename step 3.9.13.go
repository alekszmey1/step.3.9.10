package main

import "fmt"

func main()  {
first := make(chan int)
second := make(chan int)
third := make(chan struct{})
first <- 3
second <- 5
calculator(first, second, third)
}
/*Вам необходимо написать функцию calculator следующего вида:
Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит всего одно значение в один из каналов - получили значение
, обработали его, завершили работу.
После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите предельное время работы.*/
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int{
	t := make(chan int)
	var s int
	select {
	case s = <-firstChan:
		fmt.Println(s)
		s = s*s
		fmt.Println(s)
		t <-s
	case <- stopChan:
		break
	case s = <-secondChan:
		fmt.Println(s)
		s = s*3
		fmt.Println(s)
		t <-s
	}
	return t
}