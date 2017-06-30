package tests

//import (
//	"fmt"
//	"testing"
//)

//LoadArray1
func LoadArray1() []int {

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 7}
	return arr

}

//LoadArray2
func LoadArray2() []int {

	arr := []int{-2, -1}
	return arr

}

//LoadDictionary
func LoadDictionary() map[string]bool {

	dic := make(map[string]bool)

	dic["de"] = true
	dic["dec"] = true
	dic["deca"] = true
	dic["decal"] = true
	dic["do"] = true
	dic["dor"] = true
	dic["di"] = true
	dic["dir"] = true
	dic["arm"] = true
	dic["med"] = true
	dic["ree"] = true
	dic["re"] = true

	return dic

}

//LoadMatrix
func LoadMatrix(matrix *[3][4]string) {
	matrix[0][0] = "d"
	matrix[0][1] = "d"
	matrix[0][2] = "e"
	matrix[0][3] = "m"

	matrix[1][0] = "e"
	matrix[1][1] = "d"
	matrix[1][2] = "e"
	matrix[1][3] = "d"

	matrix[2][0] = "c"
	matrix[2][1] = "m"
	matrix[2][2] = "r"
	matrix[2][3] = "a"
}

//TestpointExists l
//func TestpointExists(t *testing.T) {
//
//
//
//
//	var matrix [3][4]string
//
//	LoadMatrix(matrix)
//	result := pointExists(matrix, [2]int {2,3})
//	fmt.Println("pointExists", result)
//
//	//fmt.Printf("Time elapsed QuickSort Concurrent: %d\n", timeQuickSort)
//	//fmt.Printf("Time elapsed GenericSort: %d\n", timeGenericSorted)
//
//	fmt.Println("________end________")
//}

//TestgetNewPosition
//func TestgetNewPosition(t *testing.T) {
//	fmt.Println("pointExists")
//
//
//	//fmt.Printf("Time elapsed QuickSort Concurrent: %d\n", timeQuickSort)
//	//fmt.Printf("Time elapsed GenericSort: %d\n", timeGenericSorted)
//
//	fmt.Println("________end________")
//}
