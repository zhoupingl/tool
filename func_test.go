package tool

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {

	arr := []string{"3", "2", "1"}
	InsertSort(arr)
	fmt.Println(arr)

}