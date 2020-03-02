package arrayTest

import "testing"

func TestArrayInit(t *testing.T) {
	var arr1 [3]int
	t.Log(arr1[0], arr1[1], arr1[2])
	arr2 := [3]int{1, 2, 3}
	t.Log(arr2[0], arr2[1], arr2[2])
	arr3 := [...]int{1,2,3,4,5}
	t.Log(len(arr3))
	t.Log(arr1, arr2, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1,2,3,4,5}
	//for i := 0; i < len(arr3); i++ {
	//	t.Log(arr3[i])
	//}
	//for idx, e := range arr3 {
	//	t.Log(idx, e)
	//}
	for _, value := range arr3 {
		t.Log(value)
	}
}

func TestArrayCut(t *testing.T) {
	arr := [...]int{1,2,3,4,5,6}
	arr2 := arr[1 : 3]
	arr3 := arr[1 : ]
	arr4 := arr[:4]
	arr5 := arr[1 : len(arr)]
	//arr6 := arr[-1 : 3] go语言不支持负数
	t.Log(arr2, arr3, arr4, arr5)
}

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))
	//len初始化的个数, cap代表容量
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	} //两倍增长
	//切片共享存储结构
}

func TestSliceShare(t *testing.T) {
	year := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	q1 := year[3 : 6]
	t.Log(q1, len(q1), cap(q1))
	q2 := year[5 : 8]
	t.Log(q2, len(q2), cap(q2))
	q2[0] = 13
	t.Log(q1, year)
}

//数组和切片的不同
// 1. 容量是否可伸缩
// 2. 是否可以比较