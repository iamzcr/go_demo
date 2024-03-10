package main

import "fmt"

// 数组,同一类型数据的集合,长度不可变,是值数据类型非引用数据类型
// 数组的元素个数称为数组长度,长度是类型的一部分[5]int,[10]int是不同的类型
// 数组只是声明,不赋值,元素是对应类型的初始值,如int为0
// 数组赋值给新变量或者传递给函数，都会产生一个数组副本
// 数组也是一种值，函数通过值传递来接收参数，所以数组作为函数的参数是非常低效的
// 数组越界：赋值如果超出数组长度，编译器在检测到对越界元素的访问时会报错
// 数组越界：如果在编译时没有发现越界错误，程序会在运行时出现panic
func main() {
	//声明
	var arr [3]int
	fmt.Printf("arr:值%v;类型%T,内存地址%p\n", arr, arr, &arr)
	//数组赋值
	arr[0] = 1
	fmt.Println(cap(arr))
	fmt.Printf("arr:值%v;类型%T,内存地址%p\n", arr, arr, &arr)

	var arr1 = [3]int{1, 2, 3}
	fmt.Printf("arr1:值%v;类型%T,内存地址%p\n", arr1, arr1, &arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Printf("arr2:值%v;类型%T,内存地址%p\n", arr2, arr2, &arr2)

	//另外的声明赋值方式,让编译器根据初始值的个数自行推断数组的长度
	arr3 := [...]int{1, 2, 100, 5, 3, 4, 8}
	fmt.Printf("arr3:值%v;类型%T,内存地址%p\n", arr3, arr3, &arr3)

	//指定下标初始化,因为数组是连续的,从0开始,如果跳过下标,对应的下标没赋值,会给元素类型的默认值,如果0的下标不赋值,那会自动补上,值为元素类型的默认值
	arr4 := [...]int{0: 1, 2: 10, 3: 100}
	fmt.Printf("arr4:值%v;类型%T,内存地址%p\n", arr4, arr4, &arr4)
	//数组遍历
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("for:下标值%v;下标类型%T,元素值%v;元素类型%T,元素内存地址%p\n", i, i, arr4[i], arr4[i], &arr4[i])
	}
	//当使用range迭代数组时，每次迭代都会返回数组中的一个元素的副本，而不是直接引用原始数组的内存地址,这是为了安全性和一致性考虑。
	for index, item := range arr4 {
		fmt.Printf("range:下标值%v;下标类型%T,元素值%v;元素类型%T,元素内存地址%p\n", index, index, item, item, &item)
	}

	//多为数组初始化
	var arrMuti = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arrMuti1 = [3][3]int{}
	var arrMuti2 = [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Printf("arrMuti:值%v;类型%T,内存地址%p\n", arrMuti, arrMuti, &arrMuti)
	fmt.Printf("arrMuti1:值%v;类型%T,内存地址%p\n", arrMuti1, arrMuti1, &arrMuti1)
	fmt.Printf("arrMuti2:值%v;类型%T,内存地址%p\n", arrMuti2, arrMuti2, &arrMuti2)

	//testArr := [...]int{1, 2, 100, 5, 3, 4, 8}
	//
	////求出数组testArr元素的最大值
	//max, index := testArr[0], 0 // 假设第一个元素为最大值,下标为0
	//for i := 0; i < len(testArr); i++ {
	//	if testArr[i] > max {
	//		max = testArr[i]
	//		index = i
	//	}
	//}
	//fmt.Printf("testArr最大值为%v:下标值%v\n", max, index)

	testArr := [...]int{1, 2, 100, 5, 3, 4, 8}
	//求出数组testArr两个元素相加等于7的元素
	//方法1:嵌套循环
	targetSum := 7
	for i := 0; i < len(testArr)-1; i++ {
		for j := i + 1; j < len(testArr); j++ {
			if testArr[i]+testArr[j] == targetSum {
				fmt.Printf("找到和为%d的两个数：%d 和 %d\n", targetSum, testArr[i], testArr[j])
			}
		}
	}
	//方法2哈希表
	valueMap := make(map[int]bool)
	for _, num := range testArr {
		complement := targetSum - num
		if valueMap[complement] {
			fmt.Printf("找到和为%d的两个数：%d 和 %d\n", targetSum, num, complement)
		}
		valueMap[num] = true
	}
}
