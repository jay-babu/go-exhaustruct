// Package blank_identifier tests blank identifier assignment behavior.
// Empty struct literals assigned to _ via var declarations are always allowed,
// as this is a common Go idiom for interface compliance verification.
package blank_identifier

// TestStruct is a simple struct.
type TestStruct struct {
	A string
	B int
}

// SomeInterface is a test interface.
type SomeInterface interface {
	DoSomething()
}

func (TestStruct) DoSomething() {}

// Package-level interface compliance checks (the primary use case).
var _ SomeInterface = TestStruct{}
var _ SomeInterface = &TestStruct{}

func shouldPassVarBlankIdentifier() {
	var _ = TestStruct{}
}

func shouldPassVarBlankIdentifierWithType() {
	var _ SomeInterface = TestStruct{}
}

func shouldPassVarBlankIdentifierPointer() {
	var _ SomeInterface = &TestStruct{}
}

func shouldFailRegularDeclaration() {
	var test = TestStruct{} // want "blank_identifier.TestStruct is missing fields A, B"
	_ = test
}

func shouldFailAssignment() {
	_ = TestStruct{} // want "blank_identifier.TestStruct is missing fields A, B"
}

func shouldFailInReturn() TestStruct {
	return TestStruct{} // want "blank_identifier.TestStruct is missing fields A, B"
}
