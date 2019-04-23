package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	// fmt.Println(t.Format("02 Jan 2006 15:04"))
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	// 计算时间
	week = 60 * 60 * 24 * 7 * 1e9 //必须是纳秒

	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)

	// 格式化时间
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))

	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}

// 如果需要在应用程序经过一定时间或者周期执行某项任务，则可以使用time.After或者time.Ticker.
// time.Sleep(Duration d)可以实现对某个进程时长为d的暂停
