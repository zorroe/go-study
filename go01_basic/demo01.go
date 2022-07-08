package main

import (
	"fmt"
	"strings"
)

var name string = "jack"
var age int32 = 19
var (
	gender = "male"
	email  = "1318753541@qq.com"
)

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5}
var strs = [5]string{1: "Hello", 3: "Jack"}

var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

const pi = 3.1415926
const e = 2.74

func calArea(r float32) float32 {
	return r * r * pi
}

func stringMethod() {
	var s = "bcdedit"
	s += "hijab"

	split := strings.Split(s, "i")
	println(len(s))
	for i := range split {
		println(split[i])
	}

	println(strings.Contains(s, "ijk"))

	println(strings.HasPrefix(s, "abc"))
	println(strings.HasSuffix(s, "jkl"))

	println(strings.Index(s, "def"))
	println(strings.LastIndex(s, "def"))
}

func stringMethod2() {
	s := "Hello World 太阳"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

	for _, r := range s {
		fmt.Printf("%c ", r)
	}
}

func changeString() {
	s1 := "hello"
	byteS1 := []byte(s1)

	byteS1[0] = 'H'

	s2 := string(byteS1)
	fmt.Println(s2)

	s3 := "阳光"
	runeS3 := []rune(s3)

	runeS3[0] = '月'

	s4 := string(runeS3)
	fmt.Println(s4)
}

func arrayMethod() {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3, 4, 5}
	c := [5]int{1: 100, 3: 300}
	d := [...]struct {
		name string
		age  int
	}{
		{"jack", 17},
		{"Lucy", 20},
	}
	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	f := [...][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(strs)
	fmt.Println(arr3)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func test(x [2]int) {
	fmt.Printf("x: %p\n\n", &x)
	x[1] = 1000
}

func traverse(x [3][3]int) {
	for k1, v1 := range x {
		for k2, v2 := range v1 {
			fmt.Printf("[%d,%d]:%d  ", k1, k2, v2)
		}
		fmt.Println()
	}
}

func printArr(arr *[5]int) {
	println(arr)
	arr[0] = 10
	println(arr)
}

func getSum(arr [10]int) int {
	sum := 0
	for _, v1 := range arr {
		sum += v1
	}
	return sum
}

func getTarget(arr [5]int, total int) {
	for i := 0; i < len(arr); i++ {
		target := total - arr[i]
		for k := i + 1; k < len(arr); k++ {
			if arr[k] == target {
				fmt.Printf("[%d,%d]:%d,%d \n", i, k, arr[i], arr[k])
			}
		}
	}
}

func sliceMethod() {
	// 1.
	var numberList []int
	fmt.Println(numberList)

	// 2.
	var numberListEmpty = []int{}
	fmt.Println(numberListEmpty)

	// 3.
	numList := make([]int, 2, 5)
	fmt.Println(numList)
	fmt.Println(len(numList))
	fmt.Println(cap(numList))
	// 指针：是指向第一个切片元素对应的底层数组元素的地址。（切片的第一个元素不一定是数组中的第一个元素）
	// 长度：切片的元素个数
	// 容量：从切片的开始位置到底层数组的结尾位置

	arr := [5]string{"Hello", "World", "Jack", "How", "Are"}
	var s1 = arr[1:4]
	fmt.Println(s1)
	fmt.Println(arr)
}

func emptySlice() {
	var numberList []int
	fmt.Println(numberList == nil)
	fmt.Println(len(numberList) == 0)
}

func modifySlice() {
	var arr = [...]string{"Go", "Java", "PHP", "Python"}
	s := arr[:]

	fmt.Println(arr)
	fmt.Println(s)
	s[0] = "Ruby"
	fmt.Println(arr)
	fmt.Println(s)
}

func appendSlice() {
	s := []string{"Go"}
	fmt.Println(cap(s))
	//追加一个元素
	s = append(s, "Java")
	fmt.Println(cap(s))
	//追加两个元素
	s = append(s, "C", "C++")
	fmt.Println(cap(s))
	//追加一个切片
	s = append(s, []string{"Rust", "PHP"}...)
	fmt.Println(cap(s))
	fmt.Println(s)
}

func mSlice() {
	var s = [][]string{{"1", "Java"}, {"2", "Python"}, {"3", "Go"}}
	fmt.Println(s)
}

func createMap() {
	// make(map[keyType]valueType)
	steps1 := make(map[string]string)
	fmt.Println(steps1)

	steps2 := make(map[string]int)
	fmt.Println(steps2)

	var step3 map[string]string = map[string]string{
		"name":   "Jack",
		"gender": "female",
		"score":  "80",
	}
	fmt.Println(step3)

	step4 := make(map[string]string)
	fmt.Println(step4)
}

func mapMethod() {
	var steps map[string]string = map[string]string{
		"name":   "Jack",
		"gender": "female",
		"score":  "80",
	}
	fmt.Println(steps)
	steps["email"] = "131@qq.com"
	fmt.Println(steps)
	steps["email"] = "222@153.com"
	fmt.Println(steps["email"])
	delete(steps, "email")

	k1, v1 := steps["email"]
	fmt.Println(k1, v1)
	k2, v2 := steps["name"]
	fmt.Println(k2, v2)

	for key, value := range steps {
		fmt.Println(key, value)
	}
}

func mapByReference() {
	var steps map[string]string = map[string]string{
		"name":   "Jack",
		"gender": "female",
		"score":  "80",
	}

	steps2 := steps

	steps2["name"] = "Lucy"
	steps2["gender"] = "male"
	steps2["score"] = "99"

	fmt.Println(steps)
	fmt.Println(steps2)
}

func main() {
	mapByReference()
}
