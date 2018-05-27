package handlefile

import (
	"fmt"
	"testing"
)

func TestReadDir(t *testing.T) {
	path := "E:\\CodeSpace\\nsq"
	ana, _ := walkDirs(path)
	for _, v := range ana {
		fmt.Println(v)
	}
}
