package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	fmt.Println("Hash abcde :", dataStruct.Hash("abcde"))
	fmt.Println("Hash abcde :", dataStruct.Hash("abcde"))
	fmt.Println("Hash abcdf :", dataStruct.Hash("abcdf"))
	fmt.Println("Hash qwer :", dataStruct.Hash("qwer"))

	m := dataStruct.CreateMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0108888888")
	m.Add("ZXCVBNM", "011111111")
	m.Add("CCC", "0123323123")

	fmt.Println("AAA =", m.Get("AAA"))
	fmt.Println("BBB =", m.Get("BBB"))
	fmt.Println("CCC =", m.Get("CCC"))
	fmt.Println("DDD =", m.Get("DDD"))
	fmt.Println("ZXCVBNM =", m.Get("ZXCVBNM"))

	/*******************************************/
	mp := make(map[string]string)

	mp["bcd"] = "ccc"
	fmt.Println("bcd", mp["bcd"])

	mp1 := make(map[int]string)
	mp1[53] = "ddd"
	fmt.Println("53", mp1[53])

	mp1[10] = "test"
	mp1[20] = "list"
	v, ok := mp1[10]
	fmt.Println("key 10 v:", v, "ok:", ok)
	v, ok = mp1[20]
	fmt.Println("key 20 v:", v, "ok:", ok)
	v, ok = mp1[30]
	fmt.Println("key 30 v:", v, "ok:", ok)

	delete(mp1, 20)
	v, ok = mp1[20]
	fmt.Println("key 20 v:", v, "ok:", ok)

	for key, value := range mp1 {
		fmt.Println(key, ":", value)
	}
}
