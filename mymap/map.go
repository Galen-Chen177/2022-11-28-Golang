package mymap

import "sort"

// leetcode 242
// 字符串有效字母是否相等
// abcd 与 dbca 返回true
func IsAnagram1(s, t string) bool {
	mapS := map[rune]int{}
	for _, v := range s {
		mapS[v]++
	}
	mapT := map[rune]int{}
	for _, v := range t {
		mapT[v]++
	}
	if len(mapS) != len(mapT) {
		return false
	}
	for k, v := range mapS {
		if mapT[k] != v {
			return false
		}
	}
	return true
}

func IsAnagram2(s, t string) bool {
	intS := []rune{}
	intT := []rune{}
	for _, v := range s {
		intS = append(intS, v)
	}
	for _, v := range t {
		intT = append(intT, v)
	}
	sort.Slice(intS, func(i, j int) bool {
		return intS[i] > intS[j]
	})
	sort.Slice(intT, func(i, j int) bool {
		return intT[i] > intT[j]
	})
	if len(intS) != len(intT) {
		return false
	}
	for i, v := range intS {
		if intT[i] != v {
			return false
		}
	}
	return true
}

// 给定数组，找出所有三个数相加为0的情况
func ThreeNums(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return [][]int{
			nums,
		}
	}

	sort.Ints(nums)
	res := [][]int{}
	for i, v := range nums[:len(nums)-2] {
		tmpM := map[int]struct{}{}
		for _, v2 := range nums[i+1:] {
			if _, ok := tmpM[v2]; ok {
				res = append(res, []int{v, target - v - v2, v2})
				break
			} else {
				tmpM[target-v-v2] = struct{}{}
			}
		}
	}

	return res
}

func TwoNums(nums []int, target int) []int {
	tmpM := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := tmpM[v]; ok {
			return []int{target - v, v}
		} else {
			tmpM[target-v] = struct{}{}
		}
	}
	return []int{}
}
