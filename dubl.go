package main

import "fmt"

func main() {
	a:=[10]int{}
	for i:=0; i<10; i++ {
	fmt.Print("Enter a ", i, " number ")
	fmt.Scan(&a[i])	
	}
	for i:=0; i<10; i++{
		sum:=0
		for j:=i+1; j<10; j++{
        if a[i]==a[j]{
			sum++
		}
	} 
	if sum==1 || sum==3 || sum==5{
	fmt.Print(a[i], " ")
	}
	}
	fmt.Println()
	

}