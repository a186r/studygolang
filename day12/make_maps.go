package main

import "fmt"

func main() {
	// map是引用类型，所以可以使用如下声明
	// var map1 map[keytype]valuetype
	// var map2 map[string]int
	// 在声明的时候不需要知道map的长度，map是可以动态增长的
	//我未初始化的map值是nil
	// map传递给函数的代价很小：在32位机器上占4个字节，在64位机器上占8个字节，无论实际存储了多少数据
	// 通过key在map中找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢100倍
	// 所以，如果你很在乎性能的话，建议用切片来解决问题

	// map也可以用函数作为自己的值，这样就可以用来做分支结构，key用来选择要执行的函数
	// 常用的len(map1)方法可以获得map中的pair数目，这个数目是可以伸缩的，因为map-pairs在运行时
	// 可以动态添加和删除
	var mapLit map[string]int
	var mapAssigned map[string]int

	// map可以使用下面的方法来初始化，就像数组和结构体一样
	mapLit = map[string]int{"one": 1, "two": 2}
	// map是引用类型的，内存使用make方法来分配
	mapCreated := make(map[string]float32)
	// mapAssigned也是mapLit的引用，对mapAssigned的修改也会影响到mapLit的值
	// 不要使用new，永远使用make来构造map
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	mapsss()
}

// 体验一下map的值可以是任意类型的
func mapsss() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(mf)
}

// map的容量
// 和数组不同，map可以根据新增的key-value对动态的伸缩，map的大小会自动加一，但是你也可以选择表明map的初始容量
// map2 := make(map[string]float32,100)

// 当map的容量达到上限的时候，如果再增加新的键值对，map的大小会自动加一，所以处于性能考虑，对于大的map或者会快速扩张的map，即使只是大概知道大致容量，也最好先表明

// 以切片作为map的值
// 例如一个主线程作为key，多个子线程作为value，所以可以使用下面的方式定义map
mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)