package main

import "fmt"

func main(){
	i := 0
	for i < 3 {
		fmt.Println(i)
		i += 1
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//For inf
	j:=0 
	for {
		fmt.Println("Loop")
		j+=1
		if(j > 10){
			break
		}
	}

	for i := 0; i < 10; i+=1{
		if(i % 2 == 0){
			continue
		}
		fmt.Println("Odd", i)
	}
}