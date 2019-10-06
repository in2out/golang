package main

import "fmt"

func main() {
	s := "abc"

	// break 사용 안해도 범위 정해져 있다.
	switch s {
	case "abc":
		fmt.Println("s is abc")
	case "cde":
		fmt.Println("s is cde")
	default:
		fmt.Println("unkown")
	}

	// 위에서 부터 순차적으로 실행
	switch {
	case s == "abc":
		fmt.Println("s is abc")
	case s > "cde":
		fmt.Println("s is cde")
	default:
		fmt.Println("unkown")
	}

	// break == fallthrough
	switch {
	case s == "abc":
		fmt.Println("s is abc")
		fallthrough
	case s > "cde":
		fmt.Println("s is cde")
	default:
		fmt.Println("unkown")
	}

}
