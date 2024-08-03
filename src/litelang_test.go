package litelang

import (
	"fmt"
	"os"
	"testing"
)

func TestCode(t *testing.T) {
	file, _ := os.ReadFile("testfunc.txt")
	code := string(file)
	variables := make(map[string]interface{})
	_, logs := ExecuteCode(code, variables)
	fmt.Println(logs)
}
