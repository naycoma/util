package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/naycoma/util"
)

type MyInt int
type MyBool bool

func TestCompare(t *testing.T) {
	a := assert.New(t)
	a.True(util.Compare(1, 2) < 0)
	a.True(util.Compare(2, 1) > 0)
	a.True(util.Compare(1, 1) == 0)

	a.True(util.Compare(MyInt(1), MyInt(2)) < 0)
	a.True(util.Compare(MyInt(2), MyInt(1)) > 0)
	a.True(util.Compare(MyInt(1), MyInt(1)) == 0)

	a.True(util.Compare(MyInt(1), 2) < 0)
	a.True(util.Compare(2, MyInt(1)) > 0)
	a.True(util.Compare(1, MyInt(1)) == 0)

	a.True(util.Compare(MyBool(true), MyBool(false)) > 0)
	a.True(util.Compare(MyBool(false), MyBool(true)) < 0)
	a.True(util.Compare(MyBool(true), MyBool(true)) == 0)
}

type Person struct {
	Name string
	Age  int
	City string
}

func TestSorted(t *testing.T) {
	a := assert.New(t)

	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "London"},
		{"Charlie", 35, "New York"},
		{"Alice", 20, "London"},
	}

	// Sort by Name
	sortedByName := util.Sorted(people, func(p Person) string { return p.Name })
	expectedByName := []Person{
		{"Alice", 30, "New York"},
		{"Alice", 20, "London"},
		{"Bob", 25, "London"},
		{"Charlie", 35, "New York"},
	}
	a.Equal(expectedByName, sortedByName)

	// Sort by Age
	sortedByAge := util.Sorted(people, func(p Person) int { return p.Age })
	expectedByAge := []Person{
		{"Alice", 20, "London"},
		{"Bob", 25, "London"},
		{"Alice", 30, "New York"},
		{"Charlie", 35, "New York"},
	}
	a.Equal(expectedByAge, sortedByAge)
}

func TestSorted2(t *testing.T) {
	a := assert.New(t)

	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "London"},
		{"Charlie", 35, "New York"},
		{"Alice", 20, "London"},
	}

	// Sort by Name then Age
	sortedByNameAge := util.Sorted2(people, func(p Person) (string, int) { return p.Name, p.Age })
	expectedByNameAge := []Person{
		{"Alice", 20, "London"},
		{"Alice", 30, "New York"},
		{"Bob", 25, "London"},
		{"Charlie", 35, "New York"},
	}
	a.Equal(expectedByNameAge, sortedByNameAge)
}

func TestSorted3(t *testing.T) {
	a := assert.New(t)

	people := []Person{
		{"Alice", 30, "New York"},
		{"Bob", 25, "London"},
		{"Charlie", 35, "New York"},
		{"Alice", 20, "London"},
	}

	// Sort by City, Name then Age
	sortedByCityNameAge := util.Sorted3(people, func(p Person) (string, string, int) { return p.City, p.Name, p.Age })
	expectedByCityNameAge := []Person{
		{"Alice", 20, "London"},
		{"Bob", 25, "London"},
		{"Alice", 30, "New York"},
		{"Charlie", 35, "New York"},
	}
	a.Equal(expectedByCityNameAge, sortedByCityNameAge)
}

func TestSorted4(t *testing.T) {
	a := assert.New(t)

	// Example with 4 keys (just extending the pattern)
	type Item struct{ A, B, C, D int }
	items := []Item{{1, 2, 3, 4}, {1, 2, 3, 3}, {1, 2, 4, 1}}
	sortedItems := util.Sorted4(items, func(i Item) (int, int, int, int) { return i.A, i.B, i.C, i.D })
	expectedItems := []Item{{1, 2, 3, 3}, {1, 2, 3, 4}, {1, 2, 4, 1}}
	a.Equal(expectedItems, sortedItems)
}

func TestSorted5(t *testing.T) {
	a := assert.New(t)

	// Example with 5 keys
	type Item struct{ A, B, C, D, E int }
	items := []Item{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 4}, {1, 2, 3, 5, 1}}
	sortedItems := util.Sorted5(items, func(i Item) (int, int, int, int, int) { return i.A, i.B, i.C, i.D, i.E })
	expectedItems := []Item{{1, 2, 3, 4, 4}, {1, 2, 3, 4, 5}, {1, 2, 3, 5, 1}}
	a.Equal(expectedItems, sortedItems)
}

func TestSorted6(t *testing.T) {
	a := assert.New(t)

	// Example with 6 keys
	type Item struct{ A, B, C, D, E, F int }
	items := []Item{{1, 2, 3, 4, 5, 6}, {1, 2, 3, 4, 5, 5}, {1, 2, 3, 4, 6, 1}}
	sortedItems := util.Sorted6(items, func(i Item) (int, int, int, int, int, int) { return i.A, i.B, i.C, i.D, i.E, i.F })
	expectedItems := []Item{{1, 2, 3, 4, 5, 5}, {1, 2, 3, 4, 5, 6}, {1, 2, 3, 4, 6, 1}}
	a.Equal(expectedItems, sortedItems)
}
