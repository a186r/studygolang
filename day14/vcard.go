package main

import (
	"time"
)

type Address struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirtData  time.Time
	Photo     string
	// 使用指针的方式存储，节省空间
	Address map[string]*Address
}

func main() {
	addr1 := &Address{"Elfenstraat", 12, "", "", "2600", "Mechelen", "Belgie"}
	addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "België"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2

	birthdt := time.Data(1956, 1, 17, 15, 4, 5, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"Ivo", "Balbaert", "", birthdt, photo, addrs}
	// fmt.Printf()
}
