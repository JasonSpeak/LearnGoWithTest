package main

import "fmt"

func majorityElement(nums []int) int {
	var currentMaxTimes = 0
	var currentMostValue = nums[0]

	kv := map[int]int{}

	for i := 0; i < len(nums); i++ {
		_,isExist := kv[nums[i]]
		if !isExist{
			kv[nums[i]] = 1
		}else{
			kv[nums[i]]++
		}
		if kv[nums[i]]>currentMaxTimes{
			currentMaxTimes = kv[nums[i]]
			currentMostValue = nums[i]
		}
	}

	return currentMostValue
}

func main(){
	test := []int{3,2,3}
	res := majorityElement(test)
	fmt.Println(res)
}