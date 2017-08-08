package Zigzag_6

import "fmt"

type arr struct {
	Row []rune
}

func Convert(s string, numRows int) string {

	// we need to create a slice of fixed arrays

	//ar:=arr{Row:make([]rune,numRows)}

	obj := []arr{}

	for i, r := range s {
		//fmt.Printf("i %d r %r\n", i, r)
		if i%numRows == 0 {
			ar := arr{Row: make([]rune, numRows)}
			obj = append(obj, ar)
		}
		obj[len(obj)-1].Row[i%numRows] = r
	}

	fmt.Println(obj)
	fmt.Println(len(obj))

	for v := 0; v < numRows; v++ {
		for i := 0; i < len(obj); i++ {
			//fmt.Println(obj[i])
			fmt.Print(string(obj[i].Row[v]))
		}
	}
	fmt.Println(obj)
	return "a"
}
