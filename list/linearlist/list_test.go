package linearlist

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	l1 := New()

	if gotValue := l1.Len(); gotValue != 0 {
		t.Errorf("want %v but got %v", 0, gotValue)
	}
	if gotValue := l1.Empty(); gotValue != true {
		t.Errorf("want %v but got %v", true, gotValue)
	}
}

func TestList_Append(t *testing.T) {

	l1 := New()
	l1.Append("a", "b", "c", "d")
	if gotValue := l1.Len(); gotValue != 4 {
		t.Errorf("want %v but got %v", 4, gotValue)
	}
	if gotValue := l1.Empty(); gotValue != false {
		t.Errorf("want %v but got %v", false, gotValue)
	}

	if gotValue, ok := l1.Get(3); gotValue != "d" || !ok {
		t.Errorf("want %v but got %v", "d", gotValue)
	}
}

func TestList_Get(t *testing.T) {
	l1 := New("a", "b", "c", "d")
	if gotValue, ok := l1.Get(0); gotValue != "a" || !ok {
		t.Errorf("want %v but got %v", "a", gotValue)
	}
	if gotValue, ok := l1.Get(1); gotValue != "b" || !ok {
		t.Errorf("want %v but got %v", "b", gotValue)
	}
	if gotValue, ok := l1.Get(2); gotValue != "c" || !ok {
		t.Errorf("want %v but got %v", "c", gotValue)
	}
	if gotValue, ok := l1.Get(3); gotValue != "d" || !ok {
		t.Errorf("want %v but got %v", "d", gotValue)
	}
	l1.Remove("b")
	if gotValue, ok := l1.Get(3); gotValue != nil || ok {
		t.Errorf("want %v but got %v", nil, gotValue)
	}
	if gotValue, ok := l1.Get(2); gotValue != "d" || !ok {
		t.Errorf("want %v but got %v", "d", gotValue)
	}
}

func TestList_Indexof(t *testing.T) {
	l1 := New("a", "b", "c", "d")
	if gotValue := l1.Indexof("a"); gotValue != 0 {
		t.Errorf("want %v but got %v", 0, gotValue)
	}
	if gotValue := l1.Indexof("b"); gotValue != 1 {
		t.Errorf("want %v but got %v", 1, gotValue)
	}
	if gotValue := l1.Indexof("c"); gotValue != 2 {
		t.Errorf("want %v but got %v", 2, gotValue)
	}
	if gotValue := l1.Indexof("d"); gotValue != 3 {
		t.Errorf("want %v but got %v", 3, gotValue)
	}
	if gotValue := l1.Indexof("e"); gotValue != -1 {
		t.Errorf("want %v but got %v", -1, gotValue)
	}

}

func TestList_Insert(t *testing.T) {
	l1 := New()
	l1.Insert(0, "a", "c")
	l1.Insert(1, "b")
	l1.Insert(5, "f")
	if gotValue, ok := l1.Get(1); gotValue != "b" || !ok {
		t.Errorf("want %v but got %v", "b", gotValue)
	}
	if gotValue, ok := l1.Get(2); gotValue != "c" || !ok {
		t.Errorf("want %v but got %v", "c", gotValue)
	}
	if gotValue := l1.Len(); gotValue != 3 {
		t.Errorf("want %v but got %v", 3, gotValue)
	}
}

func TestList_Remove(t *testing.T) {
	l1 := New("a", "b", "c", "b")

	l1.Remove("b")
	if gotValue, ok := l1.Get(3); gotValue != nil || ok {
		t.Errorf("want %v but got %v", nil, gotValue)
	}
	if gotValue, ok := l1.Get(1); gotValue != "c" || !ok {
		t.Errorf("want %v but got %v", "c", gotValue)
	}
	if gotValue, ok := l1.Get(2); gotValue != "b" || !ok {
		t.Errorf("want %v but got %v", "b", gotValue)
	}
	l1.Remove("a")
	l1.Remove("b")
	l1.Remove("c")
	l1.Remove("d")
	if gotValue := l1.Len(); gotValue != 0 {
		t.Errorf("want %v but got %v", 0, gotValue)
	}
	if gotValue := l1.Empty(); gotValue != true {
		t.Errorf("want %v but got %v", true, gotValue)
	}

}
func TestList_Len(t *testing.T) {
	l1 := New("a", "b", "c", "d")
	if gotValue := l1.Len(); gotValue != 4 {
		t.Errorf("want %v but got %v", 4, gotValue)
	}
	l1.Append("f")
	if gotValue := l1.Len(); gotValue != 5 {
		t.Errorf("want %v but got %v", 5, gotValue)
	}
	l1.Remove("a")
	l1.Remove("b")
	l1.Remove("c")
	if gotValue := l1.Len(); gotValue != 2 {
		t.Errorf("want %v but got %v", 2, gotValue)
	}

}

func TestList_Empty(t *testing.T) {
	l1 := New("a", "b", "c", "d")
	if gotValue := l1.Empty(); gotValue != false {
		t.Errorf("want %v but got %v", false, gotValue)
	}
	if gotValue := l1.Len(); gotValue != 4 {
		t.Errorf("want %v but got %v", 4, gotValue)
	}
	l1.Remove("d")
	l1.Remove("c")
	l1.Remove("b")
	l1.Remove("a")
	if gotValue := l1.Empty(); gotValue != true {
		t.Errorf("want %v but got %v", true, gotValue)
	}
	if gotValue := l1.Len(); gotValue != 0 {
		t.Errorf("want %v but got %v", 0, gotValue)
	}
	l1.Append("a")
	if gotValue := l1.Empty(); gotValue != false {
		t.Errorf("want %v but got %v", false, gotValue)
	}
	if gotValue := l1.Len(); gotValue != 1 {
		t.Errorf("want %v but got %v", 1, gotValue)
	}
	l1.Clear()
	if gotValue := l1.Empty(); gotValue != true {
		t.Errorf("want %v but got %v", true, gotValue)
	}
	if gotValue := l1.Len(); gotValue != 0 {
		t.Errorf("want %v but got %v", 0, gotValue)
	}
}

func TestList_Clear(t *testing.T) {
	l1 := New("a", "b", "c", "d")
	if gotValue := l1.Empty(); gotValue != false {
		t.Errorf("want %v but got %v", false, gotValue)
	}
	if gotValue := l1.Len(); gotValue != 4 {
		t.Errorf("want %v but got %v", 4, gotValue)
	}
	l1.Clear()
	if gotValue := l1.Empty(); gotValue != true {
		t.Errorf("want %v but got %v", true, gotValue)
	}
	if gotValue := l1.Len(); gotValue != 0 {
		t.Errorf("want %v but got %v", 0, gotValue)
	}
}

func TestList_String(t *testing.T) {
	l1 := New("a", "b", "c", "d")
	if gotValue := l1.String(); !strings.HasPrefix(gotValue, "LinearList") {
		t.Errorf("Should start with %v", "a")
	}
}
