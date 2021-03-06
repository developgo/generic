package generic

import "testing"

func TestErrInvalidGenericValueNull(t *testing.T) {
	expected := "invalid value: (nil)"
	err := ErrInvalidGenericValue{Value: nil}
	if err.Error() != expected {
		t.Errorf("actual:%s, expected:%s", err.Error(), expected)
	}
}

func TestErrInvalidGenericValueString(t *testing.T) {
	expected := "invalid value: (string)"
	testVal := "aaaaaaa"
	err := ErrInvalidGenericValue{Value: testVal}
	if err.Error() != expected {
		t.Errorf("actual:%s, expected:%s", err.Error(), expected)
	}
}

func TestValidFlagReset(t *testing.T) {
	var v ValidFlag = true
	v.Reset()
	if v {
		t.Error("actual:true, expected:false")
	}
}
