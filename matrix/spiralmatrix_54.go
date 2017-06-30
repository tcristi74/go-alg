package matrix

import "fmt"

//Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
//Given the following matrix:
//[
//[ 1, 2, 3 ],
//[ 4, 5, 6 ],
//[ 7, 8, 9 ]
//]
//You should return [1,2,3,6,9,8,7,4,5].
func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	if len(matrix) ==0 {
		return ret
	}
	readsquare(&matrix, 0, 0, len(matrix[0]), len(matrix), &ret)
	fmt.Println(ret)
	return ret
}

func readsquare(matrix *[][]int, xstart int, ystart int, xend int, yend int, ret *[]int) {

	fmt.Printf("xstart=%d  xend=%d  ystart=%d yend=%d \n",xstart,xend,ystart,yend)
	// get first row
	row1 :=(*matrix)[ystart]
	*ret = append(*ret,row1[xstart:xend]...)
	//fmt.Println(*ret)
	for i:=ystart+1;i<=yend-2;i++ {
		fmt.Printf("i1=%d\n",i)
		*ret = append(*ret,(*matrix)[i][xend-1])
		//fmt.Println(*ret)
	}
	if ystart+1<yend {
		// bottom row
		for i := xend - 1; i >= xstart; i-- {
			fmt.Printf("i2=%d\n", i)
			//fmt.Printf("i=%d\n",i)
			*ret = append(*ret, (*matrix)[yend-1][i])

		}
	}
	fmt.Println(*ret)
	if xstart+1<xend {
		for i := yend - 2; i > ystart; i-- {
			fmt.Printf("i3=%d\n", i)
			*ret = append(*ret, (*matrix)[i][xstart])
		}
	}

	xstart++
	xend--
	ystart++
	yend--

	if xstart+1<=xend  && ystart+1<=yend{
		readsquare(matrix , xstart , ystart , xend , yend , ret )
	} else {
		return
	}
}
