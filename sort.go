package util

import (
	"cmp"
	"reflect"
	"sort"
	"unsafe"

	"github.com/samber/lo"
)

// Ordered is a constraint that permits any type that is ordered by the Go language.
// This includes all integer, float, and string types, as well as bool.
type Ordered interface {
	cmp.Ordered | ~bool
}

// Compare compares two Ordered values x and y.
// It returns -1 if x is less than y, 0 if x equals y, and +1 if x is greater than y.
// It handles various ordered types including bool.
func Compare[T Ordered](x, y T) int {
	size := unsafe.Sizeof(x)
	if size != unsafe.Sizeof(y) {
		panic("x and y must be of the same type")
	}
	xPtr := unsafe.Pointer(&x)
	yPtr := unsafe.Pointer(&y)
	anyY := any(y)
	switch xT := any(x).(type) {
	case bool:
		return cmp.Compare(*(*uint8)(xPtr), *(*uint8)(yPtr))
	case int:
		return cmp.Compare(xT, anyY.(int))
	case int8:
		return cmp.Compare(xT, anyY.(int8))
	case int16:
		return cmp.Compare(xT, anyY.(int16))
	case int32:
		return cmp.Compare(xT, anyY.(int32))
	case int64:
		return cmp.Compare(xT, anyY.(int64))
	case uint:
		return cmp.Compare(xT, anyY.(uint))
	case uint8:
		return cmp.Compare(xT, anyY.(uint8))
	case uint16:
		return cmp.Compare(xT, anyY.(uint16))
	case uint32:
		return cmp.Compare(xT, anyY.(uint32))
	case uint64:
		return cmp.Compare(xT, anyY.(uint64))
	case float32:
		return cmp.Compare(xT, anyY.(float32))
	case float64:
		return cmp.Compare(xT, anyY.(float64))
	case string:
		return cmp.Compare(xT, anyY.(string))
	}
	valX := reflect.ValueOf(x)
	valY := reflect.ValueOf(y)
	switch typ := valX.Kind(); typ {
	case reflect.Bool:
		return cmp.Compare(*(*uint8)(xPtr), *(*uint8)(yPtr))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cmp.Compare(valX.Int(), valY.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return cmp.Compare(valX.Uint(), valY.Uint())
	case reflect.Float32, reflect.Float64:
		return cmp.Compare(valX.Float(), valY.Float())
	case reflect.String:
		return cmp.Compare(valX.String(), valY.String())
	}
	panic("x or y is not Ordered")
}

// pNType is a helper struct used for multi-key sorting.
type pNType[P1, P2, P3, P4, P5, P6 Ordered] struct {
	N  uint8
	P1 P1
	P2 P2
	P3 P3
	P4 P4
	P5 P5
	P6 P6
}

// pN compares the N-th key of p with the N-th key of other.
func (p pNType[P1, P2, P3, P4, P5, P6]) pN(i uint8, other pNType[P1, P2, P3, P4, P5, P6]) int {
	switch i {
	case 1:
		return Compare(p.P1, other.P1)
	case 2:
		return Compare(p.P2, other.P2)
	case 3:
		return Compare(p.P3, other.P3)
	case 4:
		return Compare(p.P4, other.P4)
	case 5:
		return Compare(p.P5, other.P5)
	case 6:
		return Compare(p.P6, other.P6)
	}
	panic("invalid index")
}

// Compare compares two pNType instances based on their keys up to N.
func (p pNType[P1, P2, P3, P4, P5, P6]) Compare(other pNType[P1, P2, P3, P4, P5, P6]) int {
	if p.N == 0 {
		p.N = 6
	}
	for i := range p.N {
		if result := p.pN(i+1, other); result != 0 {
			return result
		}
	}
	return 0
}

// nComparable is an interface for types that can be compared.
type nComparable[T any] interface {
	Compare(other T) int
}

// sortedBy sorts a slice by transforming its elements into a comparable type C.
func sortedBy[T any, C nComparable[C]](slice []T, replace func(item T) C) (sorted []T) {
	zip := lo.Zip2(lo.Map(slice, func(item T, _ int) C {
		return replace(item)
	}), slice)
	sort.SliceStable(zip, func(i, j int) bool {
		return zip[i].A.Compare(zip[j].A) < 0
	})
	_, sorted = lo.Unzip2(zip)
	return
}

// Sorted sorts a slice by a single key.
func Sorted[S ~[]T, T any, P1 Ordered](
	slice S, keys func(item T) P1,
) (sorted []T) {
	return sortedBy(slice, func(item T) pNType[P1, int, int, int, int, int] {
		return pNType[P1, int, int, int, int, int]{
			N:  1,
			P1: keys(item),
		}
	})
}

// Sorted2 sorts a slice by two keys.
func Sorted2[S ~[]T, T any, P1, P2 Ordered](
	slice S, keys func(item T) (P1, P2),
) (sorted []T) {
	return sortedBy(slice, func(item T) pNType[P1, P2, int, int, int, int] {
		p1, p2 := keys(item)
		return pNType[P1, P2, int, int, int, int]{
			N:  2,
			P1: p1,
			P2: p2,
		}
	})
}

// Sorted3 sorts a slice by three keys.
func Sorted3[S ~[]T, T any, P1, P2, P3 Ordered](
	slice S, keys func(item T) (P1, P2, P3),
) (sorted []T) {
	return sortedBy(slice, func(item T) pNType[P1, P2, P3, int, int, int] {
		p1, p2, p3 := keys(item)
		return pNType[P1, P2, P3, int, int, int]{
			N:  3,
			P1: p1,
			P2: p2,
			P3: p3,
		}
	})
}

// Sorted4 sorts a slice by four keys.
func Sorted4[S ~[]T, T any, P1, P2, P3, P4 Ordered](
	slice S, keys func(item T) (P1, P2, P3, P4),
) (sorted []T) {
	return sortedBy(slice, func(item T) pNType[P1, P2, P3, P4, int, int] {
		p1, p2, p3, p4 := keys(item)
		return pNType[P1, P2, P3, P4, int, int]{
			N:  4,
			P1: p1,
			P2: p2,
			P3: p3,
			P4: p4,
		}
	})
}

// Sorted5 sorts a slice by five keys.
func Sorted5[S ~[]T, T any, P1, P2, P3, P4, P5 Ordered](
	slice S, keys func(item T) (P1, P2, P3, P4, P5),
) (sorted []T) {
	return sortedBy(slice, func(item T) pNType[P1, P2, P3, P4, P5, int] {
		p1, p2, p3, p4, p5 := keys(item)
		return pNType[P1, P2, P3, P4, P5, int]{
			N:  5,
			P1: p1,
			P2: p2,
			P3: p3,
			P4: p4,
			P5: p5,
		}
	})
}

// Sorted6 sorts a slice by six keys.
func Sorted6[S ~[]T, T any, P1, P2, P3, P4, P5, P6 Ordered](
	slice S, keys func(item T) (P1, P2, P3, P4, P5, P6),
) (sorted []T) {
	return sortedBy(slice, func(item T) pNType[P1, P2, P3, P4, P5, P6] {
		p1, p2, p3, p4, p5, p6 := keys(item)
		return pNType[P1, P2, P3, P4, P5, P6]{
			N:  6,
			P1: p1,
			P2: p2,
			P3: p3,
			P4: p4,
			P5: p5,
			P6: p6,
		}
	})
}