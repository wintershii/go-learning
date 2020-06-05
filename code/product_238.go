package main
import "fmt"

func main()  {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{0}
	}
	res := make([]int, length)
	num := 1
	for i := 0; i < len(nums); i++ {
		res[i] = num
		num *= nums[i]
	}
	num = 1
	for i := length-1 ; i >= 0; i-- {
		res[i] *= num
		num *= nums[i]
	}
	return res
}
