package main

//数组的操作，遍历
func traverse(arr []int) {
	//arr:=[]int{10,20,30,40,50}
	/*for(let i=0;i< len(arr);i++){

	}*/

	//就发现这种遍历方式
	for _, v := range arr {
		println(v)
	}

	//index:=0
	//l:=len(arr)
	//居然没由 while语法
	//while(index<l){
	//}
}

// 截取
func slice(arr []int) {
	//panic: runtime error: slice bounds out of range [:8] with capacity 5
	//如果超过了长度,就会报错
	arr = arr[3:8]
	traverse(arr)
}
