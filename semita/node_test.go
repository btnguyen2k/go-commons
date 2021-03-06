package semita

import (
	"reflect"
	"strings"
	"testing"
)

func Test_createEmptyGenericSlice(t *testing.T) {
	v := createEmptyGenericSlice()
	if v.Kind() != reflect.Slice || v.Len() != 0 || v.Type().Elem().Kind() != reflect.Interface {
		t.Errorf("Test_createEmptyGenericSlice failed, expected empty generic slice, but received {%#v}", v)
	}
}

func Test_createEmptyGenericMap(t *testing.T) {
	v := createEmptyGenericMap()
	if v.Kind() != reflect.Map || v.Len() != 0 || v.Type().Key().Kind() != reflect.String || v.Type().Elem().Kind() != reflect.Interface {
		t.Errorf("Test_createEmptyGenericMap failed, expected empty generic map, but received {%#v}", v)
	}
}

func TestNode_unwrap(t *testing.T) {
	name := "TestNode_unwrap"
	{
		v := true
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(bool) != v {
			t.Errorf("%s failed, expected {%#v}, but received {%#v}", name, v, uv)
		}
	}
	{
		v := 1
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(int) != v {
			t.Errorf("%s failed, expected {%#v}, but received {%#v}", name, v, uv)
		}
	}
	{
		v := 1.2
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(float64) != v {
			t.Errorf("%s failed, expected {%#v}, but received {%#v}", name, v, uv)
		}
	}
	{
		v := "123"
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		uv := n.unwrap()
		if uv.(string) != v {
			t.Errorf("%s failed, expected {%#v}, but received {%#v}", name, v, uv)
		}
	}
}

func TestNode_nextInvalid(t *testing.T) {
	name := "TestNode_nextInvalid"
	{
		v := false
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("path")
		if err == nil { // primitive type cannot be "next"
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		v := 1
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("[0]")
		if err == nil { // primitive type cannot be "next"
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		v := 1.2
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("path")
		if err == nil { // primitive type cannot be "next"
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		v := "123"
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("[0]")
		if err == nil { // primitive type cannot be "next"
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		v := map[int]string{1: "one"}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("1")
		if err == nil { // map key must be string
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		v := []int{0, 1, 2, 3}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.next("[a]")
		if err == nil { // index value must be a number
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
}

func TestNode_nextEmptyOrNonExist(t *testing.T) {
	name := "TestNode_nextEmptyOrNonExist"
	{
		v := []int{0, 1, 2, 3}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		next, err := n.next("[]")
		if err != nil || next != nil { // index out-of-bound: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
		next, err = n.next("[-1]")
		if err != nil || next != nil { // index out-of-bound: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
		next, err = n.next("[99]")
		if err != nil || next != nil { // index out-of-bound: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		v := [4]int{0, 1, 2, 3}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		next, err := n.next("[]")
		if err != nil || next != nil { // index out-of-bound: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
		next, err = n.next("[-1]")
		if err != nil || next != nil { // index out-of-bound: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
		next, err = n.next("[99]")
		if err != nil || next != nil { // index out-of-bound: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		v := map[string]int{"a": 1, "b": 2, "c": 3}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		next, err := n.next("e")
		if err != nil || next != nil { // index not exist: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
	{
		type MyStruct struct {
			FieldMap   map[string]interface{}
			FieldSlice []int
		}
		v := MyStruct{}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		next, err := n.next("FieldNotExist")
		if err != nil || next != nil { // field not exist: "next" should be successful, but node should be nil
			t.Errorf("%s failed for value {%#v}", name, v)
		}
		next, err = n.next("FieldMap")
		if err != nil || next == nil || !next.value.IsNil() { // next should be successful, but node's value should be 'empty' for uninitialized field
			t.Errorf("%s failed for value {%#v}", name, v)
		}
		next, err = n.next("FieldSlice")
		if err != nil || next == nil || !next.value.IsNil() { // next should be successful, but node's value should be 'empty' for uninitialized field
			t.Errorf("%s failed for value {%#v}", name, v)
		}
	}
}

type Inner struct {
	b interface{}
	f interface{}
	i interface{}
	s interface{}
}

type Outter struct {
	A       interface{}
	B       interface{}
	M       interface{}
	S       interface{}
	private interface{}
}

func genDataOuter() Outter {
	return Outter{
		A: []int{0, 1, 2, 3, 4, 5},
		B: [3]string{"a", "b", "c"},
		M: map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		S:       Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
		private: 1.2,
	}
}

func genDataMap() map[string]interface{} {
	return map[string]interface{}{
		"a": []int{0, 1, 2, 3, 4, 5},
		"b": [3]string{"a", "b", "c"},
		"m": map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		"s": genDataOuter(),
	}
}

func TestNode_nextMap(t *testing.T) {
	v := genDataMap()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}", v, p)
	}

	p = "not exist"
	node, err = root.next(p)
	if node != nil || err != nil {
		// non-exists entry
		t.Errorf("TestNode_nextMap failed with data %#v at index {%#v}, error: %e", v, p, err)
	}

	for _, path := range []string{"a.[0]", "b.[1]", "m.z", "s.A.[0]", "s.B.[1]", "s.M.z", "s.S.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextMap failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func TestNode_nextStruct(t *testing.T) {
	v := genDataOuter()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "not exist"
	node, err = root.next(p)
	if node != nil || err != nil {
		// non-exists entry
		t.Errorf("TestNode_nextStruct failed with data %#v at index {%#v}", v, p)
	}

	for _, path := range []string{"A.[0]", "B.[1]", "M.z", "S.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextStruct failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func genDataSlice() []interface{} {
	return []interface{}{
		"a string",
		103,
		19.81,
		true,
		[]int{0, 1, 2, 3, 4, 5},
		[3]string{"a", "b", "c"},
		map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		genDataOuter(),
	}
}

func TestNode_nextSlice(t *testing.T) {
	v := genDataSlice()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "abc"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextSlice failed with data %#v at index {%#v}", v, p)
	}

	for _, path := range []string{"[4].[0]", "[5].[1]", "[6].z", "[7].A.[0]", "[7].B.[1]", "[7].M.z", "[7].S.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextSlice failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func genDataArray() [8]interface{} {
	return [8]interface{}{
		"a string",
		103,
		19.81,
		true,
		[]int{0, 1, 2, 3, 4, 5},
		[3]string{"a", "b", "c"},
		map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		genDataOuter(),
	}
}

func TestNode_nextArray(t *testing.T) {
	v := genDataArray()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var err error
	var p string
	var node *node

	p = "abc"
	node, err = root.next(p)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_nextArray failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextArray failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextArray failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.next(p)
	if node != nil || err != nil {
		// index out-of-bound: silent nil should be return
		t.Errorf("TestNode_nextArray failed with data %#v at index {%#v}", v, p)
	}

	for _, path := range []string{"[4].[0]", "[5].[1]", "[6].z", "[7].A.[0]", "[7].B.[1]", "[7].M.z", "[7].S.s"} {
		node = root
		for _, p = range strings.Split(path, ".") {
			node, err = node.next(p)
			if node == nil || err != nil {
				t.Errorf("TestNode_nextArray failed with data %#v at path {%#v}", v, path)
			}
		}
	}
}

func TestNode_setValueInvalid(t *testing.T) {
	{
		v := false
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("path", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
		}
	}
	{
		v := 1
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("[0]", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
		}
	}
	{
		v := 1.2
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("path", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
		}
	}
	{
		v := "123"
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("[0]", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
		}
	}
	{
		v := map[int]string{1: "one"}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("1", reflect.ValueOf("value"))
		if err == nil {
			t.Errorf("TestNode_setValueInvalid failed for value {%#v}", v)
		}
	}
}

func TestNode_setValueMapInvalidType(t *testing.T) {
	{
		v := map[int]string{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		node, err := root.setValue("1", reflect.ValueOf("string"))
		if node != nil || err == nil {
			// invalid key type
			t.Errorf("TestNode_setValueMapInvalidType failed with data %#v", v)
		}
	}
	{
		v := map[string]int{"one": 1}
		n := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		_, err := n.setValue("two", reflect.ValueOf("2"))
		if err == nil {
			// invalid element type
			t.Errorf("TestNode_setValueMapInvalidType failed for value {%#v}", v)
		}
	}
}

func TestNode_setValueMap(t *testing.T) {
	v := genDataMap()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}

	p = "notExist"
	node, err = root.setValue(p, data)
	if node == nil || err != nil || node.unwrap() != data.Interface() {
		// non-exists entry
		t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"a", "b", "m", "s"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueMap failed with data %#v at index {%#v}", v, p)
		}
	}
}

func TestNode_setValueStructInvalidType(t *testing.T) {
	type MyStruct struct {
		S string
		I int
	}

	v := MyStruct{S: "string", I: 123}
	n := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	_, err := n.setValue("S", reflect.ValueOf(1981))
	if err == nil {
		// invalid element type
		t.Errorf("TestNode_setValueStructInvalidType failed for value {%#v}", v)
	}
}

func TestNode_setValueStructUnaddressable(t *testing.T) {
	type MyStruct struct {
		A       interface{}
		B       interface{}
		M       interface{}
		S       interface{}
		private interface{}
	}
	v := MyStruct{
		A: []int{0, 1, 2, 3, 4, 5},
		B: [3]string{"a", "b", "c"},
		M: map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		S:       Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
		private: 1.2,
	}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v), // for struct: only addressable struct is settable
	}

	p := "A"
	node, err := root.setValue(p, reflect.ValueOf("data"))
	if node != nil || err == nil {
		t.Errorf("TestNode_setValueStructUnaddressable failed with data %#v at index {%#v}", v, p)
	}
}

func TestNode_setValueStruct(t *testing.T) {
	type MyStruct struct {
		A       interface{}
		B       interface{}
		M       interface{}
		S       interface{}
		private interface{}
	}
	v := MyStruct{
		A: []int{0, 1, 2, 3, 4, 5},
		B: [3]string{"a", "b", "c"},
		M: map[string]interface{}{
			"x": "x",
			"y": 19.81,
			"z": true,
		},
		S:       Inner{b: true, f: 1.03, i: 1981, s: "btnguyen2k"},
		private: 1.2,
	}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(&v), // for struct: only addressable struct is settable
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "notExist"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// non-exists entry
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}

	p = "private"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// non-exists entry
		t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"A", "B", "M", "S"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueStruct failed with data %#v at index {%#v}", v, p)
		}
	}
}

func TestNode_setValueSliceInvalidType(t *testing.T) {
	v := []string{"a", "b", "c"}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	{
		node, err := root.setValue("[1]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueSliceInvalidType failed with data %#v", v)
		}
	}
	{
		node, err := root.setValue("[a]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueSliceInvalidType failed with data %#v", v)
		}
	}
}

func TestNode_setValueSliceAppend(t *testing.T) {
	name := "TestNode_setValueSliceAppend"
	v := map[string]interface{}{
		"s": []interface{}{0, "1", false, true},
	}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	_p := "s"
	_node, _err := root.next("s")
	if _node == nil || _err != nil {
		t.Errorf("%s failed with data %#v at index {%#v}", name, v, _p)
	}

	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	l := len(_node.unwrap().([]interface{}))
	p = "[]"
	node, err = _node.setValue(p, data)
	_node, _err = root.next("s")
	if _node == nil || _err != nil {
		t.Errorf("%s failed with data %#v at index {%#v}", name, v, _p)
	}
	if node == nil || err != nil || node.unwrap() != data.Interface() || len(_node.unwrap().([]interface{})) != l+1 {
		// non-exists entry
		t.Errorf("%s failed with data %#v at index {%#v}", name, v, p)
	}

	for _, p = range []string{"[0]", "[1]", "[2]", "[]"} {
		node, err = _node.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, p)
		}
	}
}

func TestNode_setValueSlice(t *testing.T) {
	v := genDataSlice()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "xyz"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}

	l := root.value.Len()
	p = "[]"
	node, err = root.setValue(p, data)
	if node == nil || err != nil || node.unwrap() != data.Interface() || root.value.Len() != l+1 {
		// non-exists entry
		t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"[0]", "[1]", "[2]", "[]"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueSlice failed with data %#v at index {%#v}", v, p)
		}
	}
}

func TestNode_setValueArrayInvalidType(t *testing.T) {
	v := [3]string{"a", "b", "c"}
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v),
	}
	{
		node, err := root.setValue("[1]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueArrayInvalidType failed with data %#v", v)
		}
	}
	{
		node, err := root.setValue("[a]", reflect.ValueOf(1))
		if node != nil || err == nil {
			// invalid type
			t.Errorf("TestNode_setValueArrayInvalidType failed with data %#v", v)
		}
	}
}

func TestNode_setValueArrayUnaddressable(t *testing.T) {
	v := genDataArray()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(v), // for array: only addressable array is settable
	}
	p := "[0]"
	node, err := root.setValue(p, reflect.ValueOf("data"))
	if node != nil || err == nil {
		t.Errorf("TestNode_setValueArrayUnaddressable failed with data %#v at index {%#v}", v, p)
	}
}

func TestNode_setValueArray(t *testing.T) {
	v := genDataArray()
	root := &node{
		prev:     nil,
		prevType: nil,
		key:      "",
		value:    reflect.ValueOf(&v), // for array: only addressable array is settable
	}
	var data = reflect.ValueOf("data")
	var err error
	var p string
	var node *node

	p = "xyz"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// invalid type
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}

	p = "[-1]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}
	p = "[999]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}
	p = "[]"
	node, err = root.setValue(p, data)
	if node != nil || err == nil {
		// index out-of-bound
		t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
	}

	for _, p = range []string{"[0]", "[1]", "[2]", "[3]"} {
		node, err = root.setValue(p, data)
		if node == nil || err != nil || node.unwrap() != data.Interface() {
			t.Errorf("TestNode_setValueArray failed with data %#v at index {%#v}", v, p)
		}
	}
}

/*----------------------------------------------------------------------*/
func TestNode_createChild_Invalid(t *testing.T) {
	name := "TestNode_createChild_Invalid"
	{
		v := "a string"
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "key"
		_, err := root.createChild(index, "")
		if err == nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
	{
		v := 103
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v),
		}
		index := "key"
		_, err := root.createChild(index, "")
		if err == nil {
			t.Errorf("%s failed with data %#v", name, v)
		}
	}
}

func TestNode_createChildOfMap(t *testing.T) {
	name := "TestNode_createChildOfMap"
	{
		v := map[string]int{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "key"
		nextIndex := "" // map has specific element type, hence 'nextIndex' is not used
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.unwrap() != 0 {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := map[string]interface{}{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "key"
		nextIndex := "[]" // child node should be a slice
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Slice {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := map[string]interface{}{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "key"
		nextIndex := "nestedKey" // child node should be a map
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Map {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
}

func TestNode_createChildOfStruct(t *testing.T) {
	name := "TestNode_createChildOfStruct"
	type MyStruct struct {
		FieldInt     int
		FieldPtr     *string
		FieldGeneric interface{}
		fieldPrivate int
	}
	{
		v := MyStruct{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: wrapped object must be addressable
		}
		index := "NotExist"
		_, err := root.createChild(index, "")
		if err == nil { // 'createChild' should fail because field does not exist
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, "")
		}
	}
	{
		v := MyStruct{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: wrapped object must be addressable
		}
		index := "fieldPrivate"
		_, err := root.createChild(index, "")
		if err == nil { // 'createChild' should fail because field is not exported
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, "")
		}
	}
	{
		v := MyStruct{FieldInt: 2}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: wrapped object must be addressable
		}
		index := "FieldInt"
		nextIndex := "" // struct field has specific element type, hence 'nextIndex' is not used
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.unwrap() != 0 {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		str := "a string"
		v := MyStruct{FieldPtr: &str}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: wrapped object must be addressable
		}
		index := "FieldPtr"
		nextIndex := "" // struct field has specific element type, hence 'nextIndex' is not used
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.unwrap() != nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := MyStruct{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: wrapped object must be addressable
		}
		index := "FieldGeneric"
		nextIndex := "[]" // child node should be a slice
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Slice {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := MyStruct{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: wrapped object must be addressable
		}
		index := "FieldGeneric"
		nextIndex := "nested" // child node should be a map
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Map {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
}

func TestNode_createChildOfSlice(t *testing.T) {
	name := "TestNode_createChildOfSlice"
	{
		v := []int{0, 1, 2}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "[-1]"
		nextIndex := "" // slice has specific element type, hence 'nextIndex' is not used
		_, err := root.createChild(index, nextIndex)
		if err == nil { // should failed because index is out-of-bound
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		index = "[99]"
		nextIndex = "" // slice has specific element type, hence 'nextIndex' is not used
		_, err = root.createChild(index, nextIndex)
		if err == nil { // should failed because index is out-of-bound
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		index = "[a]"
		nextIndex = "" // slice has specific element type, hence 'nextIndex' is not used
		_, err = root.createChild(index, nextIndex)
		if err == nil { // should failed because index is invalid
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := []int{}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "[]"
		nextIndex := "" // slice has specific element type, hence 'nextIndex' is not used
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil || root.value.Len() == 0 {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next("[0]")
		if err != nil || next == nil || next.unwrap() != 0 {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := []interface{}{1, 2, 3}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "[1]"
		nextIndex := "[]" // child node should be a slice
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Slice {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := []interface{}{"a", "b", "c"}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		index := "[2]"
		nextIndex := "nestedKey" // child node should be a map
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Map {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
}

func TestNode_createChildOfArray(t *testing.T) {
	name := "TestNode_createChildOfArray"
	{
		v := [3]int{0, 1, 2}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: wrapped object must be addressable
		}
		index := "[-1]"
		nextIndex := "" // array has specific element type, hence 'nextIndex' is not used
		_, err := root.createChild(index, nextIndex)
		if err == nil { // should failed because index is out-of-bound
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		index = "[99]"
		nextIndex = "" // array has specific element type, hence 'nextIndex' is not used
		_, err = root.createChild(index, nextIndex)
		if err == nil { // should failed because index is out-of-bound
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		index = "[a]"
		nextIndex = "" // array has specific element type, hence 'nextIndex' is not used
		_, err = root.createChild(index, nextIndex)
		if err == nil { // should failed because index is invalid
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		index = "[]"
		nextIndex = "" // array has specific element type, hence 'nextIndex' is not used
		_, err = root.createChild(index, nextIndex)
		if err == nil { // should failed because index is out-of-bound
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := [3]int{1, 2, 3}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: wrapped object must be addressable
		}
		index := "[1]"
		nextIndex := "" // array has specific element type, hence 'nextIndex' is not used
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.unwrap() != 0 {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := [3]interface{}{false, 1, "two"}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: wrapped object must be addressable
		}
		index := "[1]"
		nextIndex := "[]" // child node should be a slice
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Slice {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
	{
		v := [3]interface{}{true, "1", 2.0}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: wrapped object must be addressable
		}
		index := "[2]"
		nextIndex := "nestedKey" // child node should be a map
		child, err := root.createChild(index, nextIndex)
		if err != nil || child == nil {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
		next, err := root.next(index)
		if err != nil || next == nil || next.value.Elem().Kind() != reflect.Map {
			t.Errorf("%s failed with data %#v at index {%v.%v}", name, v, index, nextIndex)
		}
	}
}

/*--------------------------------------------------------------------------------*/

func TestNode_removeValue_Invalid(t *testing.T) {
	name := "TestNode_removeValue_Invalid"
	{
		v := []interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "[a]" // invalid index: 'a' is not a number
		err := root.removeValue(path)
		if err == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		v := []interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "[10]" // index out-of-bound
		err := root.removeValue(path)
		if err == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		v := [3]interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "[1]" // unaddressable array, value cannot be set
		err := root.removeValue(path)
		if err == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		v := map[string]interface{}{"a": 1, "b": "2", "c": true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "[1]" // invalid type: expecting array or slice
		err := root.removeValue(path)
		if err == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		v := map[int]interface{}{1: 1, 2: "2", 3: true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "a" // invalid map: map's key must be string
		err := root.removeValue(path)
		if err == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		type MyStruct struct {
			A interface{}
			B interface{}
			C interface{}
		}
		v := MyStruct{A: 1, B: "2", C: true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "A" // unaddressable struct, value cannot be set
		err := root.removeValue(path)
		if err == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		v := []interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "a" // invalid type: expecting map or struct
		err := root.removeValue(path)
		if err == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
}

func TestNode_removeValue_Map(t *testing.T) {
	name := "TestNode_removeKey_Map"
	{
		v := map[string]interface{}{"a": 1, "b": "2", "c": true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		if root.value.Kind() != reflect.Map {
			t.Errorf("%s failed: expecting data to be a map, but received {%T}", name, root.unwrap())
		}
		path := "a"
		node, _ := root.next(path)
		if node == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		err := root.removeValue(path)
		node, _ = root.next(path)
		if err != nil || node != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		v := map[string]interface{}{"a": 1, "b": "2", "c": true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v),
		}
		if root.value.Elem().Kind() != reflect.Map {
			t.Errorf("%s failed: expecting data to be a map, but received {%T}", name, root.unwrap())
		}
		path := "a"
		node, _ := root.next(path)
		if node == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		err := root.removeValue(path)
		node, _ = root.next(path)
		if err != nil || node != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}

	{
		v := map[string]interface{}{"a": 1, "b": "2", "c": true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "a"
		_, err := root.setValue(path, reflect.ValueOf(nil)) // set 'nil' should be equivalent to 'remove'
		node, _ := root.next(path)
		if err != nil || node != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
}

func TestNode_removeValue_Struct(t *testing.T) {
	name := "TestNode_removeValue_Struct"
	type MyStruct struct {
		fieldPrivate   int
		FieldInt       int
		FieldString    string
		FieldPointer   *int
		FieldInterface interface{}
	}

	{
		i := 123
		v := MyStruct{fieldPrivate: 1, FieldInt: 2, FieldString: "3", FieldPointer: &i, FieldInterface: true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: only addressable struct is settable
		}
		var path string
		var next *node
		var err error

		if root.value.Elem().Kind() != reflect.Struct {
			t.Errorf("%s failed: expecting data to be a struct, but received {%T}", name, root.unwrap())
		}

		path = "fieldPrivate"
		err = root.removeValue(path)
		if err == nil { // private field should not be writable
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}

		path = "FieldInt"
		err = root.removeValue(path)
		if err != nil { // removal should be successful
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		next, err = root.next(path)
		if err != nil || next == nil || next.unwrap() != 0 { // after removal, field's value should be reset to 'zero'
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}

		path = "FieldString"
		err = root.removeValue(path)
		if err != nil { // removal should be successful
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		next, err = root.next(path)
		if err != nil || next == nil || next.unwrap() != "" { // after removal, field's value should be reset to 'zero'
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}

		path = "FieldPointer"
		err = root.removeValue(path)
		if err != nil { // removal should be successful
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		next, err = root.next(path)
		if err != nil || next == nil || next.unwrap() != nil { // after removal, field's value should be reset to 'zero'
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}

		path = "FieldInterface"
		err = root.removeValue(path)
		if err != nil { // removal should be successful
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		next, err = root.next(path)
		if err != nil || next == nil || next.unwrap() != nil { // after removal, field's value should be reset to 'zero'
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}

	{
		i := 123
		v := MyStruct{fieldPrivate: 1, FieldInt: 2, FieldString: "3", FieldPointer: &i, FieldInterface: true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for struct: only addressable struct is settable
		}

		path := "FieldInterface"
		_, err := root.setValue(path, reflect.ValueOf(nil)) // set 'nil' should be equivalent to 'remove'
		node, _ := root.next(path)
		if err != nil || node.unwrap() != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
}

func TestNode_removeValue_Slice(t *testing.T) {
	name := "TestNode_removeValue_Slice"
	{
		v := []interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		if root.value.Kind() != reflect.Slice {
			t.Errorf("%s failed: expecting data to be a slice, but received {%T}", name, root.unwrap())
		}
		path := "[1]"
		node, _ := root.next(path)
		if node == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		l1 := len(root.unwrap().([]interface{}))
		err := root.removeValue(path)
		l2 := len(root.unwrap().([]interface{}))
		node, _ = root.next(path)
		if err != nil || l1 != l2 || node.unwrap() != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
	{
		v := []interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v),
		}
		if root.value.Elem().Kind() != reflect.Slice {
			t.Errorf("%s failed: expecting data to be a slice, but received {%T}", name, root.unwrap())
		}
		path := "[1]"
		node, _ := root.next(path)
		if node == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		l1 := len(*root.unwrap().(*[]interface{}))
		err := root.removeValue(path)
		l2 := len(*root.unwrap().(*[]interface{}))
		node, _ = root.next(path)
		if err != nil || l1 != l2 || node.unwrap() != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}

	{
		v := []interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(v),
		}
		path := "[1]"
		l1 := len(root.unwrap().([]interface{}))
		_, err := root.setValue(path, reflect.ValueOf(nil)) // set 'nil' should be equivalent to 'remove'
		l2 := len(root.unwrap().([]interface{}))
		node, _ := root.next(path)
		if err != nil || l1 != l2 || node.unwrap() != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
}

func TestNode_removeValue_Array(t *testing.T) {
	name := "TestNode_removeValue_Array"
	{
		v := [3]interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: only addressable array is settable
		}
		if root.value.Elem().Kind() != reflect.Array {
			t.Errorf("%s failed: expecting data to be an array, but received {%T}", name, root.unwrap())
		}
		path := "[1]"
		node, _ := root.next(path)
		if node == nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
		l1 := len(root.unwrap().(*[3]interface{}))
		err := root.removeValue(path)
		l2 := len(root.unwrap().(*[3]interface{}))
		node, _ = root.next(path)
		if err != nil || l1 != l2 || node.unwrap() != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}

	{
		v := [3]interface{}{1, "2", true}
		root := &node{
			prev:     nil,
			prevType: nil,
			key:      "",
			value:    reflect.ValueOf(&v), // for array: only addressable array is settable
		}
		path := "[1]"
		l1 := len(root.unwrap().(*[3]interface{}))
		_, err := root.setValue(path, reflect.ValueOf(nil)) // set 'nil' should be equivalent to 'remove'
		l2 := len(root.unwrap().(*[3]interface{}))
		node, _ := root.next(path)
		if err != nil || l1 != l2 || node.unwrap() != nil {
			t.Errorf("%s failed with data %#v at index {%#v}", name, v, path)
		}
	}
}
