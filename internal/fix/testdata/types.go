package testdata

type Nested struct {
	Value int
}

type TestStruct struct {
	stringField    string
	intField       int
	boolField      bool
	float64Field   float64
	pointerField   *int
	sliceField     []string
	mapField       map[string]int
	chanField      chan int
	interfaceField interface{}
	funcField      func()
	structField    Nested
}
