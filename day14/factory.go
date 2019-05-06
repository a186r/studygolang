package main

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}

func main() {
	// 这样调用
	f := NewFile(10, "./test.txt")
}

// 在go语言中常常像上面一样在工厂方法里使用初始化来简便的实现构造函数
// 如果File是一个结构体类型，那么表达式new(File)和&File{}是等价的
