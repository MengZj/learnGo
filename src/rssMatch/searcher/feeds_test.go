package searcher

import (
	"fmt"
	"testing"
)

func TestRetrieve(t *testing.T) {
	filename := "data.json"
	fs,err := Retrieve(filename)
	if err != nil {
		t.Errorf("retrieve error %v\n",err)
	}
	for _, value := range fs {
		fmt.Println(value)
	}
}
