package main

import (
	"fmt"
)

type  Virtmach struct {
	ip string
	hostname string
	diskgb int
	ram int
}

func ( v *Virtmach)ipset(ip string) {
	v.ip =ip
}


func (v *Virtmach)diskexpand(gb int){
	v.diskgb = v.diskgb+gb
}

func (v Virtmach) display() {
	fmt.Println("Primary IP address:", v.ip)
	fmt.Println("Hostname:", v.hostname)
	fmt.Println("Disk GB:", v.diskgb)
	fmt.Println("RAM:", v.ram)
}

func main() {

	vm1 := Virtmach{"10.20.0.1", "Zebra",20,8}

	vm1.display()
	vm1.diskexpand(3)
	vm1.ipset("192.168.77.33")

	vm1.display()
}

