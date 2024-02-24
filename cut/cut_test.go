package cut_test

import (
	"cut/cut"
	"fmt"
	"reflect"
	"testing"
)

func TestGetFieldsWithSpace(t *testing.T) {
	want := []int{1, 2, 3}
	got, err := cut.GetFields("1 2 3")
	if err != nil {
		t.Errorf("Error whiles parsing fields")
	} else if !reflect.DeepEqual(want, got) {
		t.Error("Go back and look at the code you just wrote")
	}
}

func TestGetFieldsWithComma(t *testing.T) {
	want := []int{1, 2, 3}
	got, err := cut.GetFields("1,2,3")
	if err != nil {
		t.Errorf("Error whiles parsing fields")
	} else if !reflect.DeepEqual(want, got) {
		t.Error("Go back and look at the code you just wrote")
	}
}

func TestGetFieldsWithInvalidFormat(t *testing.T) {
	_, err := cut.GetFields("dis int des | ig in uplink, 2")
	if err == nil {
		t.Errorf("Error whiles parsing fields")
	}
}

func TestCutSpecifiedField(t *testing.T) {
	text := "Julius Nicholas \n Smith Pepperwood \n"
	want := []string{"Nicholas", "Pepperwood"}
	got := cut.CutSpecificField([]byte(text), []int{2})
	if !reflect.DeepEqual(want, got) {
		for _, v := range got {
			fmt.Println(v)
		}
		t.Error("Go back and look at the code you just wrote")
	}
}

func TestCutSpecifiedMultipleFields(t *testing.T) {
	text := "Julius Nicholas \n Smith Pepperwood \n"
	want := []string{"Julius	Nicholas", "Smith	Pepperwood"}
	got := cut.CutSpecificField([]byte(text), []int{1, 2})
	if !reflect.DeepEqual(want, got) {
		for _, v := range got {
			fmt.Println(v)
		}
		t.Error("Go back and look at the code you just wrote")
	}
}

func TestCutSpecifiedFieldWithDelimiter(t *testing.T) {
	text := "Julius,Nicholas \nSmith,Pepperwood \n"
	want := []string{"Julius", "Smith"}
	got := cut.CutSpecificFieldWithDelimiter([]byte(text), []int{1}, ",")
	if !reflect.DeepEqual(want, got) {
		for _, v := range got {
			fmt.Println(v)
		}
		t.Error("Go back and look at the code you just wrote")
	}
}
