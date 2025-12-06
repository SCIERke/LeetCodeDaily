package main

import "fmt"

func reverseString(s []byte)  {
	tmp := make([]byte, len(s))
	copy(tmp, s);

	for idx , value := range tmp {
		s[len(s) - idx - 1] = value;
	}

}


func main() {
	s := []byte{'h','e','l','l','o'};
	reverseString(s);

	fmt.Println(string(s));
}