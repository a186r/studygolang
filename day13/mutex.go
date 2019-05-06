// 在一些复杂的程序中，通常通过不同线程执行不同应用来实现程序的并发，当不同线程使用同一个变量时，
// 经常会出现一个问题：无法预知变量被不同线程的程序修改的顺序（这通常被称为资源竞争）

// 经典的做法是一次只能让一个线程对共享变量进行操作。当变量被一个线程改变时，我们为它上锁，直到这个线程
// 执行完成并解锁后，其他线程才能访问它

// 特别是我们之前学习的map类型是不存在这种锁的机制来实现这种效果，所以map类型是非线程安全的。当并行
// 访问一个共享的map类型的数据，map数据将会出错

// 在go语言中这种锁的机制是通过sync包中的Mutex来实现的，sync是synchronized,意味着线程将有序的对同一变量
// 进行访问

// type Info struct {
// 	mu sync.Mutex
// }

// 如果一个函数想要改变一个变量可以这样写
// func Update(info *Info) {
// 	info.mu.Lock()
// 	info.Str = ""
// 	info.mu.Unlock()
// }

// 通过Mutex来实现一个可以上锁的共享缓冲器
// type SyncedBuffer struct {
// 	lock   sync.Mutex
// 	buffer bytes.Buffe
// }