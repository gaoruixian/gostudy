/*main表名这是一个可以独立执行的程序*/
package main

import (
	"fmt"
	// "io"
	"math/rand"
	// "os"
	"runtime"
	"time"
)

//变量类型在变量名的后面。当定义一个变量，默认赋值为其类型的null值
var a string = "i'm noob"
var number int = 1
var b string

type Vertex struct {
	X int
	Y int
}

/*
结构体
*/
type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

const (
	pi    = 3.14
	Big   = 1 << 10
	Small = Big >> 5
)

//定义接口
type Phone interface {
	//method_name [return_type]
	call()
}
//定义结构体
type IPhone struct {
	model string
}
//实现接口方法
/*func (struct_name_variable struct_name) method_name() [return_type]{

}*/
func (phone IPhone) call(){
	fmt.Printf("I'm %s,I can call you!\n",phone.model)
}

/*func (phone Phone) call(){
	fmt.Println("I'm IPhone,I can call you!")
}*/




/*print something*/
func main() {
	var phone Phone
	phone = IPhone{"IPhone"}
	phone.call()

	//errors.New("Error")


	fmt.Print("Hello Golang!\n")
	b = "golang"
	fmt.Println(a, number, b)
	test()
	x, y := swap("world", "Hello")
	fmt.Println(x, y)
	fmt.Println(pi)
	const pia = 3.1415
	fmt.Println(pia)
	fmt.Println(Big, Small)

	novartype := "没有var声明且没有类型的变量定义"
	fmt.Println(novartype)

	//循环语句，没有while;for循环没有括号
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("求和为：", sum)

	/*	suma := 0
		for suma < 10 {
			suma += suma

		}
		fmt.Println("for循环求和：", suma)*/
	//if语句
	if n := rand.Intn(100); n <= 2 {
		fmt.Println("[0,2]", n)
	} else {

		fmt.Println("[3,99]", n)

	}

	/**
		  switch语句,不用break语句跳出switch
	      可以像if一样带上一个短语
		  **/
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		fmt.Printf("%s.", os)

	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")

	case t.Hour() < 17:
		fmt.Println("Good afternoon.")

	default:
		fmt.Println("Good evening.")

	}

	for i := 0; i < 5; i++ {
		//deferred函数按照先进后出的顺序执行
		defer fmt.Print(i)

	}

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("struct", v)
	//Golang中存在指针，但是不支持算数运算
	p := Vertex{1, 2}
	q := &p
	q.X = 2 //直接访问域X
	fmt.Println(q)
	fmt.Println(p)

	//Name:为特定域设置值
	r := Vertex{X: 3}
	fmt.Println(r)

	//new函数:new(T)分配一个被初始化为0且类型为T的值，并且返回指向此值的指针
	vv := new(Vertex)
	fmt.Println(vv)
	vv.X, vv.Y = 11, 9
	fmt.Println(vv)

	//数组 [n]T 是一种类型：表示一个长度为n的数组，其中元素类型为T。定义完成后，数组长度是无法改变的

	var a [3]string
	a[0] = "gao"
	a[1] = "rui"
	a[2] = "xian"
	fmt.Println(a)

	//slice是一种数据结构，其指向一个数组某个连续的部分。slice用起来很像数组。[]T为slice类型。
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice ==", slice)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] == %d\n", i, slice[i])
	}
	//创建slice函数使用make函数（不是new函数，因为需要设置额外的参数来控制slice的创建）
	slice_a := make([]int, 5) //len(slice_a)为5
	fmt.Printf("slice_a:%d\n", slice_a)

	var maxOfTwo int = max(1, 2)
	fmt.Printf("比较大小函数返回:%d\n", maxOfTwo)

	fmt.Println("====================数组学习====================")
	/*
		var variable_name [size] variable_type
	*/
	var variable_varray [5]int = [5]int{1, 2, 3, 4, 5}
	variable_varray_onlimit := [...]int{1, 2, 3, 4, 5, 6}
	//variable_varray = int {0,2,1,3,5}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", variable_varray[i])
	}
	for _, y := range variable_varray_onlimit {
		fmt.Println(y)
	}
	//var varray_size  = len(variable_varray_a)
	//for i:=0;i< varray_size;i++  {
	//	fmt.Printf("%d\n",variable_varray_a[i])
	//}

	//创建新的结构体方式一
	fmt.Println(Books{"Go语言", "郜瑞仙", "语言编程", 90016})
	//创建新的结构体方式二:key-->value形式
	fmt.Println(Books{title: "Go语言", author: "郜瑞仙", subject: "语言编程", bookId: 90016})

	var book1 Books = Books{"Go语言", "郜瑞仙", "语言编程", 90016}
	bookTitle := book1.title
	fmt.Println(book1)
	fmt.Println(bookTitle)
	fmt.Println(&book1.subject)

	var book2 Books = Books{"java语言", "郜瑞仙", "编程语言", 90016}
	fmt.Println(book2)
	printBook(&book2)




	fmt.Println("Go语言切片Slice学习")
	/*
	Go语言切片是对数组的抽象。Go的数组长度不可变，其提供了一种灵活，功能强悍的内置类型切片（“动态数组”）。其优点是，长度不固定
	，可以追加元素。
	定义方法：
	var identifier []type 或者使用make()函数来创建切片
	var slice1 []type = make([]type,len),其简写为 slice1 := make([]type,len),其中len是切片的初始长度。
	切片初始化：
	slice1 := [] int{1,2,3}
	slice1 := arr[:];初始化切片s,是数组arr的引用
	slice1 := arr[startIndex:endIndex]
	slice1 := arr[startIndex:]
	slice1 := arr[:endIndex]
	slice1 := make([]int,len,cap)

	slice2 := slice1[startIndex,endIndex]
	 */

	sliceNumber := make([]int,3,5)
	printSlice(sliceNumber)

	var slice1 []int
	if(slice1 == nil){
		fmt.Println("slice1为空")
	}
	printSlice(slice1)


	slice2 := []int{1,2,3,4,5,6}
	printSlice(slice2)
	fmt.Printf("切片截取:%v\n",slice2[0:6])

	slice2 = append(slice2,7,8,9)
	printSlice(slice2)

	slice3 := make([]int,len(slice2),(cap(slice2))*2)
	copy(slice3,slice2)
	printSlice(slice3)
	//求slice3各元素之和
	sum1 := 0;
	for _,num:=range slice3{
		sum1 += num
	}
	fmt.Printf("slice3各元素之和%d\n",sum1 )

	fmt.Println("=============Go语言map学习===============")
	/*
	Map是一种无序的键值对的集合。Map最重要的一点是通过key来快速检索数据。由于map是无序的，所有无法决定他们返回顺序。
	map定义：
	1:var map_variable map[key_type]value_type,默认map是nil
	2:使用make
	map_variable := make(map[key_type]value_type)
	如果不初始化map,会创建一个nil map.nil map不能存放键值对
	 */
	var countryMap map[string]string
	countryMap = make(map[string]string)
	countryMap ["France"] = "Paris"
	countryMap ["China"] = "Beijing"
	countryMap ["Japan"] = "Tokoy"
	countryMap ["Italy"] = "Romon"
	delete(countryMap,"Tokoy")
	for country,capital:=range countryMap{
		if(capital == "Tokoy"){
			delete(countryMap,country)
		}
		fmt.Printf("国家%s的首都是%s\n",country,capital)


	}
	map1 := map[string]string{"a":"apple","b":"banana","c":"C++","d":"Go","e":"Java"}
	for key,value := range map1{

		fmt.Printf("map中的key:%s,value:%s\n",key,value )
	}
	mapv,ok := map1["a"];
	fmt.Println(mapv,ok)

	/*
	go语言类型转换:
	typename(expression)
	 */
	changeA := 17
	changeB  := 5
	var  mean = changeA/changeB
	fmt.Println(mean)
	var changMean = float32(changeA)/float32(changeB)
	fmt.Println(changMean)


	//字符串修改
	s := `Hello!
GaoRuiXian`
	s1 := "Hello!" +
		"Gao"
	fmt.Println(s1)
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

	var  com complex64  = 5+5i
	fmt.Println(com)

	fallThroughTest(2)


}
func fallThroughTest(num int){
	switch num {
	case 0:
		fmt.Println(num)
		fallthrough
	case 1:
		fmt.Println(num)
		fallthrough
	case 2:
		fmt.Println(num)
	default:
		fmt.Println(num)
	}



}

func printSlice(slice []int){

	fmt.Printf("len=%d,cap=%d slice=%v\n",len(slice),cap(slice),slice)
}


func printBook(book *Books) {
	fmt.Println(book.title)

}

/*
   函数定义
   func function_name([parameter list]) [return_types]{
      函数体

   }

*/

func max(num1 int, num2 int) int {

	var result int
	if num1 > num2 {
		result = num1

	} else {
		result = num2
	}

	return result

}

/*

   defer

*/

/*func CopyFile(dstName, srcName string) {
	src, err := os.Open(srcName)
	if err == nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err == nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}*/

func test() {
	var age int = 18
	fmt.Println("Hello Golang! This is my first Golang!\n", age)
}

func swap(x, y string) (string, string) {
	return y, x

}

//单行注释

/*
多行注释
*/
