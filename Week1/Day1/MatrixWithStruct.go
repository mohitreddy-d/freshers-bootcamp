package main

import (
	"encoding/json"
	"fmt"
)

//Create a struct called 'Matrix'. The Matrix struct has the following information:
//number of rows of matrix
//number of columns of matrix
//elements of matrix (You can use 2D vector)
//
//The Matrix struct has methods for each of the following:
//get the number of rows
//get the number of columns
//set the elements of the matrix at a given position (i,j)
//adding two matrices. (2nd matrix can be taken as input to the method)
//print matrix structure as json
//You can assume that the dimensions are correct for the multiplication and addition.
//
//Expectation: Use structs, slices, methods

type Matrix struct {
	Rows     int     `json:"rows"`
	Cols     int     `json:"cols"`
	Elements [][]int `json:"elements"`
}

func (This Matrix) Add(Other Matrix) Matrix {
	res := make([][]int, This.Rows)
	for i := 0; i < This.Rows; i++ {
		res[i] = make([]int, This.Cols)
	}

	for idx1 := range This.Elements {
		for idx2, val2 := range This.Elements[idx1] {
			res[idx1][idx2] = val2 + Other.Elements[idx1][idx2]
		}
	}
	return Matrix{
		Rows:     This.Rows,
		Cols:     This.Cols,
		Elements: res,
	}
}

func main() {

	fmt.Println("")
	TwoDArr1 := [][]int{{1, 1}, {2, 2}}
	TwoDArr2 := [][]int{{10, 10}, {20, 20}}
	matrix1 := Matrix{
		Rows:     len(TwoDArr1),
		Cols:     len(TwoDArr1[0]),
		Elements: TwoDArr1,
	}

	matrix2 := Matrix{
		Rows:     len(TwoDArr1),
		Cols:     len(TwoDArr1[0]),
		Elements: TwoDArr2,
	}

	res := matrix1.Add(matrix2)
	fmt.Println(res)

	jsonData, err := json.Marshal(res)
	// can also use MarshallIndent(3 params)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println(string(jsonData))

}
