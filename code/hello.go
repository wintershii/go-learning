package main

import "fmt"

type Books struct {
	title string
	author string
	id int
}

func main()  {
	fmt.Print("hello world")
	arr1 := [3]int{1,2,3}
	fmt.Print(arr1)
	slice := make([]int,0,10)
	fmt.Print(slice)
	//ptr, len, cap := slice
	slice = append(slice, 1,2,3)
	fmt.Println(slice,cap(slice),len(slice))
	var ip *int
	a := 20
	ip = &a
	fmt.Println(*ip)
	b := Books{"111","1111",1}
	fmt.Println(b)
	var struct_poi *Books
	struct_poi = &b
	fmt.Println(struct_poi.author)

	for i, num := range slice  {
		fmt.Println(i , num)
	}

	mapp := make(map[string]string)
	mapp["a"] = "A"
	mapp["b"] = "B"
	for k, v := range mapp {
		fmt.Println(k, v)
	}
	fmt.Println(mapp["b"])
	delete(mapp,"b")
	fmt.Println(mapp["b"])
}
