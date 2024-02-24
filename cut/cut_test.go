package cut_test

import (
	"cut/cut"
	"fmt"
	"reflect"
	"testing"
)

func TestCutSpecifiedField(t *testing.T) {
	text := "Julius Nicholas \n Smith Pepperwood \n"
	want := []string{"Nicholas", "Pepperwood"}
	got := cut.CutSpecificField([]byte(text), 1)
	if !reflect.DeepEqual(want, got) {
		for _, v := range got {
			fmt.Println(v)
		}
		t.Error("Go back and look at the code you just wrote")
	}
}
