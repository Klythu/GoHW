package main

import (
	"fmt"
	"os"
)

type Chash interface {
	record()
	look()
	erase()
}
type Mem struct {
	mem map[int]int
}

func NewMem() Mem {
	p := new(Mem)
	p.mem = make(map[int]int)
	return (*p)
}

func record(mem Mem, id int, inf int) {
	mem.mem[id] = inf
}
func erase(mem Mem, id int) {
	delete(mem.mem, id)
}
func look(mem Mem) {
	fmt.Println(mem.mem)
}
func main() {
	mem := NewMem()
	var info int
	var key int
	var rejim string
	for {
		fmt.Println("Команды: [a]dd , [d]elite ,[l]ook")
		fmt.Fscan(os.Stdin, &rejim)
		switch rejim {
		case "a":
			fmt.Println("Key")
			fmt.Fscan(os.Stdin, &key)
			fmt.Println("DATA")
			fmt.Fscan(os.Stdin, &info)
			record(mem, int(key), int(info))
		case "d":
			fmt.Println("Key")
			fmt.Fscan(os.Stdin, &key)
			erase(mem, int(key))
		case "l":
			look(mem)
		}

	}
}
