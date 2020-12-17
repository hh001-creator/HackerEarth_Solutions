package main

import "fmt"


func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func make_anagrams(a string , b string) int {
	var arr[256] int;
	var arr1[256] int;
	for i := 0 ; i < len(a) ;i++ {
		arr[(int)(a[i])]++;
	}
	for i := 0 ; i < len(b) ;i++ {
		arr1[(int)(b[i])]++;
	}
	count := 0;
	for i := 0;i < 256;i++ {
		if(arr[i] != arr1[i]){
			count += Abs(arr[i] - arr1[i]);
		}
	}
	return count;
}

func main(){
	var t int;
	fmt.Scanf("%d",&t);
	for i := 0;i < t;i++ {
		var a , b string;
		fmt.Scanf("%s",&a);
		fmt.Scanf("%s",&b);
		fmt.Println(make_anagrams(a , b));
	}
}