package kvmap

import (
	"fmt"
	"testing"
)

func TestPlainData(t *testing.T) {
	key := "hello"
	data := "world"
	m := ToKV(key, data)
	v, ok := m[key]
	if !ok {
		t.Errorf("%s not found", data)
	}
	if v != data {
		t.Errorf("expected %s, found %s", data, v)
	}
}

func TestSliceData(t *testing.T) {
	key := "slices"
	data := []string{"zero", "one", "two"}
	m := ToKV(key, data)
	for i, d := range data {
		kk := fmt.Sprintf("%s/%d", key, i)
		vv := m[kk]
		if vv != d {
			t.Errorf("expected %s, found %s", d, vv)
		}
	}
}

func TestMapData(t *testing.T) {
	key := "maps"
	data := map[string]int{"one": 1, "two": 2}
	m := ToKV(key, data)
	for k, v := range data {
		kk := fmt.Sprintf("%s/%v", key, k)
		vv := m[kk]
		if vv != v {
			t.Errorf("expected %v, found %v", v, vv)
		}
	}
}

type TestDataAlpha struct {
	One string `kv:"sais"`
}

type TestData struct {
	One   string        `kv:"uno"`
	Two   int           `kv:"dos"`
	Three []string      `kv:"tres"`
	Four  TestDataAlpha `kv:"quatro"`
	Five  map[int]int   `kv:"singko"`
}

func TestStructData(t *testing.T) {
	key := "struct"
	data := &TestData{
		One:   "1",
		Two:   2,
		Three: []string{"3", "3"},
		Four: TestDataAlpha{
			One: "111",
		},
		Five: map[int]int{6: 6},
	}
	m := ToKV(key, data)

	if uno := m[key+"/uno"]; uno != data.One {
		t.Errorf("expected %v, found: %v", data.One)
	}

	if dos := m[key+"/dos"]; dos != data.Two {
		t.Errorf("expected %v, found: %v", data.Two)
	}

	if tresuno := m[key+"/tres/0"]; tresuno != data.Three[0] {
		t.Errorf("expected %v, found %v", data.Three[0], tresuno)
	}

	if tresdos := m[key+"/tres/1"]; tresdos != data.Three[1] {
		t.Errorf("expected %v, found %v", data.Three[1], tresdos)
	}

	if quatro := m[key+"/quatro/sais"]; quatro != data.Four.One {
		t.Errorf("expected %v, found %v", data.Four.One, quatro)
	}

	if singko := m[key+"/singko/6"]; singko != data.Five[6] {
		t.Errorf("expected %v, found %v", data.Five[6], singko)
	}
}
