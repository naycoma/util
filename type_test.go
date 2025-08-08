package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/naycoma/util"
)

type Interface interface {
	Test()
}
type Interface2 interface {
	Interface
}

type Struct struct {
	Value string
}

type String string

func TestGetTypeName(t *testing.T) {
	a := assert.New(t)
	a.Equal("int", util.GetTypeName[int]())
	a.Equal("Interface", util.GetTypeName[Interface]())
	a.Equal("Struct", util.GetTypeName[Struct]())

	a.Equal("*Struct", util.GetTypeName[*Struct]())
	a.Equal("*Interface", util.GetTypeName[*Interface]())
}

var _ Interface = (*impl)(nil)

type impl struct {
}

// Test implements Interface.
func (i *impl) Test() {
	panic("unimplemented")
}

func TestGetTypeNameFromValue(t *testing.T) {
	a := assert.New(t)
	implInstance := &impl{}
	var i Interface = implInstance
	a.NotEqual("Interface", util.GetTypeNameFromValue(i))
	a.NotEqual("Interface", util.GetTypeNameFromValue(implInstance))
	a.Equal("*impl", util.GetTypeNameFromValue(implInstance))
	a.Equal("*impl", util.GetTypeNameFromValue(i))

	a.Equal("*Interface", util.GetTypeNameFromValue(&i))

	deep2 := &implInstance
	a.Equal("**impl", util.GetTypeNameFromValue(&implInstance))
	deep3 := &deep2
	a.Equal("***impl", util.GetTypeNameFromValue(deep3))
	deep4 := &deep3
	a.Equal("****impl", util.GetTypeNameFromValue(deep4))
	deep5 := &deep4
	a.Equal("*****impl", util.GetTypeNameFromValue(deep5))
	deep6 := &deep5
	a.Equal("******impl", util.GetTypeNameFromValue(deep6))

	// Additional test cases
	a.Equal("int", util.GetTypeNameFromValue(123))
	a.Equal("string", util.GetTypeNameFromValue("hello"))
	a.Equal("bool", util.GetTypeNameFromValue(true))
	a.Equal("Struct", util.GetTypeNameFromValue(Struct{}))
	a.Equal("String", util.GetTypeNameFromValue(String("test")))
	a.Equal("*Struct", util.GetTypeNameFromValue(&Struct{}))
	ptrToStruct := &Struct{}
	ptrToPtrToStruct := &ptrToStruct
	a.Equal("**Struct", util.GetTypeNameFromValue(ptrToPtrToStruct))
	var i2 Interface2 = implInstance
	a.Equal("*impl", util.GetTypeNameFromValue(i2))
	anonStruct := struct{ Name string }{Name: "test"}
	a.Equal("struct { Name string }", util.GetTypeNameFromValue(anonStruct))
	anonFunc := func() {}
	a.Equal("func()", util.GetTypeNameFromValue(anonFunc))
	var anyVal any = 123
	a.Equal("int", util.GetTypeNameFromValue(anyVal))
	var anyPtr any = &anyVal
	a.Equal("*int", util.GetTypeNameFromValue(anyPtr))

	assert.Panics(t, func() { util.GetTypeNameFromValue(nil) }, "GetTypeNameFromValue(nil) should panic")
}

func TestAnyPtr(t *testing.T) {
	a := assert.New(t)

	var typedInt int = 1
	a.Equal("int", util.GetTypeNameFromValue(typedInt))

	typedAutoPtr := &typedInt
	a.Equal("*int", util.GetTypeNameFromValue(typedAutoPtr))
	a.NotEqual("*interface {}", util.GetTypeNameFromValue(typedAutoPtr))

	var anyInt any = typedInt
	a.Equal("int", util.GetTypeNameFromValue(anyInt))

	var anyPtr *any = &anyInt
	a.Equal("*int", util.GetTypeNameFromValue(anyPtr))

	autoPtr := &anyInt
	a.Equal("*int", util.GetTypeNameFromValue(autoPtr))

	autoFromAnyPtr := *anyPtr
	a.Equal("int", util.GetTypeNameFromValue(autoFromAnyPtr))
	autoFromAutoPtr := *autoPtr
	a.Equal("int", util.GetTypeNameFromValue(autoFromAutoPtr))
	var fromAnyPtr any = *anyPtr
	a.Equal("int", util.GetTypeNameFromValue(fromAnyPtr))
	var fromAutoPtr any = *autoPtr
	a.Equal("int", util.GetTypeNameFromValue(fromAutoPtr))
}

func TestGetTypeString(t *testing.T) {
	a := assert.New(t)
	a.Equal("int", util.GetTypeString[int]())
	a.Equal("util_test.Interface", util.GetTypeString[Interface]())
	a.Equal("util_test.Struct", util.GetTypeString[Struct]())

	a.Equal("*util_test.Struct", util.GetTypeString[*Struct]())
	a.Equal("*util_test.Interface", util.GetTypeString[*Interface]())
}

func TestGetTypeStringFromValue(t *testing.T) {
	a := assert.New(t)

	// New test cases
	a.Equal("int", util.GetTypeStringFromValue(123))
	a.Equal("string", util.GetTypeStringFromValue("hello"))
	a.Equal("bool", util.GetTypeStringFromValue(true))
	a.Equal("util_test.Struct", util.GetTypeStringFromValue(Struct{}))
	a.Equal("util_test.String", util.GetTypeStringFromValue(String("test")))
	a.Equal("*util_test.Struct", util.GetTypeStringFromValue(&Struct{}))
	ptrToStruct := &Struct{}
	ptrToPtrToStruct := &ptrToStruct
	a.Equal("**util_test.Struct", util.GetTypeStringFromValue(ptrToPtrToStruct))
	implInstance := &impl{}
	var i Interface = implInstance
	a.Equal("*util_test.impl", util.GetTypeStringFromValue(i))
	var i2 Interface2 = implInstance
	a.Equal("*util_test.impl", util.GetTypeStringFromValue(i2))
	anonStruct := struct{ Name string }{Name: "test"}
	a.Equal("struct { Name string }", util.GetTypeStringFromValue(anonStruct))
	anonFunc := func() {}
	a.Equal("func()", util.GetTypeStringFromValue(anonFunc))
	
	var anyVal any = 123
	a.Equal("int", util.GetTypeStringFromValue(anyVal))
	var anyPtr any = &anyVal
	a.Equal("*interface {}", util.GetTypeStringFromValue(anyPtr))
	
	var anyVal2 any = int(123)
	a.Equal("int", util.GetTypeStringFromValue(anyVal2))
	// TODO
	var anyPtr2 any = &anyVal2
	a.Equal("*interface {}", util.GetTypeStringFromValue(anyPtr2))

	assert.Panics(t, func() { util.GetTypeStringFromValue(nil) }, "GetTypeStringFromValue(nil) should panic")
}