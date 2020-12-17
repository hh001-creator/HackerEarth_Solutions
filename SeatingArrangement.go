package main

import "fmt"

func main(){
	var t int;
	seats := [6] string{"WS" , "MS" , "AS", "AS" , "MS" , "WS"};
	fmt.Scanf("%d",&t);
	for i := 0;i < t;i++ {
		var n int;
		fmt.Scanf("%d",&n);
		if(n % 12 == 0){
			fmt.Println(n - 11 , "WS");
		}else if(n % 12 >= 1 && n % 12 <= 6){
			fmt.Println(n + (13 - (2 * (n % 12))) , seats[n % 12 - 1]);
		}else if(n % 12 > 6){
			fmt.Println(n - (((n % 6) * 2) - 1) , seats[n % 12 - 7]);
		}
	}
}