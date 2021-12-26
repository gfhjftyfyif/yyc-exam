package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func a (n int){
	for num:=(n-1)*250000+1;num<n*250000;num++{
		var b =true
		for i:=2;i<num;i++{
			if num%i==0{
               b=false
			   break
			}
		}
		if b{
			fmt.Println(num)
		}
	}
	wg.Done()
}
func main(){
for i:=1;i<=4;i++{
	wg.Add(1)
	go a(i)
	}
	wg.Wait()
}
//参考b站教学视频