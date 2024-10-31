package slice

import (
	"fmt"
	"reflect"
)

// DedupeStrings returns a new slice with duplicates removed, maintains original order
func DedupeStrings(sli []string) []string {
	var result []string
	seen := map[string]bool{}
	for _, v := range sli {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

// DedupeInts returns a new slice with duplicates removed, maintains original order
func DedupeInts(sli []int) []int {
	var result []int
	seen := map[int]bool{}
	for _, v := range sli {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

// IntersectionStrings returns a slice of values in both argument slices, deduped
func IntersectionStrings(sli1, sli2 []string) []string {
	var result []string
	seen := map[string]bool{}
	for _, v := range sli1 {
		seen[v] = true
	}
	for _, v := range sli2 {
		if seen[v] {
			result = append(result, v)
			delete(seen, v)
		}
	}
	return result
}

// RemoveAllStrings returns a new slice with all instances of a given string removed
func RemoveAllStrings(sli []string, val string) []string {
	var result []string
	for _, v := range sli {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

// RemoveAllInts returns a new slice with all instances of a given int removed
func RemoveAllInts(sli []int, val int) []int {
	var result []int
	for _, v := range sli {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

// SpliceStrings removes a given number of elements starting at a given index
// if index + items >= len(sli) it does not throw an error
func SpliceStrings(sli []string, index int, items int) []string {
	if items < 0 {
		panic("cannot splice negative number of items")
	}
	if index+items >= len(sli) {
		return sli[:index]
	}
	copy(sli[index:], sli[index+items:])
	sli = sli[:len(sli)-items]
	return sli
}

// SpliceInts removes a given number of elements starting at a given index
// if index + items >= len(sli) it does not throw an error
func SpliceInts(sli []int, index int, items int) []int {
	if items < 0 {
		panic("cannot splice negative number of items")
	}
	if index+items >= len(sli) {
		return sli[:index]
	}
	copy(sli[index:], sli[index+items:])
	sli = sli[:len(sli)-items]
	return sli
}

type intTuple struct {
	A, B int
}

type intTupleSlice []intTuple

// NewZipTuple create intTupleSlice for use with .ZipTuple
func NewZipTuple() intTupleSlice {
	var t intTupleSlice
	t = make(intTupleSlice, 0)
	return t
}

// zip concatenate 2 int slices to a ZipTuple
func (t intTupleSlice) ZipTuple(a, b []int) error {

	if len(a) != len(b) {
		return fmt.Errorf("zip: arguments must be of same length")
	}

	t = make(intTupleSlice, len(a), len(b))

	for i, e := range a {
		t[i] = intTuple{e, b[i]}
	}

	return nil
}

// zipArr concatenate 2 int slices to one slice of arrays of two values
/*a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
b := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
z, err := slice.ZipArr(a, b)
if err != nil { panic(err) }
for _, v := range z { fmt.Printf("%d%d", v[0], v[1]) }
fmt.Println()
Output: 10293847566574839201
*/
func ZipArr(a, b []int) ([][2]int, error) {

	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([][2]int, len(a), len(a))

	for i, e := range a {
		r[i] = [2]int{e, b[i]}
	}

	return r, nil
}

// zipGen concatenate 2 slices of any type to one predefined slice
/*a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
b := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
c := [][2]int{}
e := zipGen(a, b, &c)
*/
func ZipGen(a, b, c interface{}) error {

	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)

	if ta.Kind() != reflect.Slice || tb.Kind() != reflect.Slice || ta != tb {
		return fmt.Errorf("zip: first two arguments must be slices of the same type")
	}

	if tc.Kind() != reflect.Ptr {
		return fmt.Errorf("zip: third argument must be pointer to slice")
	}

	for tc.Kind() == reflect.Ptr {
		tc = tc.Elem()
	}

	if tc.Kind() != reflect.Slice {
		return fmt.Errorf("zip: third argument must be pointer to slice")
	}

	eta, _, etc := ta.Elem(), tb.Elem(), tc.Elem()

	if etc.Kind() != reflect.Array || etc.Len() != 2 {
		return fmt.Errorf("zip: third argument's elements must be an array of length 2")
	}

	if etc.Elem() != eta {
		return fmt.Errorf("zip: third argument's elements must be an array of elements of the same type that the first two arguments are slices of")
	}

	va, vb, vc := reflect.ValueOf(a), reflect.ValueOf(b), reflect.ValueOf(c)

	for vc.Kind() == reflect.Ptr {
		vc = vc.Elem()
	}

	if va.Len() != vb.Len() {
		return fmt.Errorf("zip: first two arguments must have same length")
	}

	for i := 0; i < va.Len(); i++ {
		ea, eb := va.Index(i), vb.Index(i)
		tt := reflect.New(etc).Elem()
		tt.Index(0).Set(ea)
		tt.Index(1).Set(eb)
		vc.Set(reflect.Append(vc, tt))
	}

	return nil
}
