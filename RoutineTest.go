package main

import "fmt"

func appendint(s string, i int) map[string]int {
	ret := make(map[string]int, 1)
	ret[s] = i
	return ret
}

func TestRoutine() {

	ret := make([]map[string]int, 0)
	s1 := "hello"
	success := make(chan map[string]int)
	for i := 0; i < 4; i++ {

		go func(iig int) {
			tmp := appendint(s1, iig)
			success <- tmp
		}(i)
	}
	t := time.After(time.Second * 10)
A:
	for len(ret) < 4 {
		select {
		case <-t:
			fmt.Println("break ")
			break A
		case map1 := <-success:
			//fmt.Println("success", <-success)
			fmt.Println("success")
			ret = append(ret, map1)
		}
	}

	fmt.Println("final ret:", ret)
	return
}

func main() {
	TestRoutine()
}
