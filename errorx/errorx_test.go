package errorx

import (
	"fmt"
	"testing"
)

func TestErrorX(t *testing.T) {
	// 測試Format
	err := New("test")
	err.Format(nil, 'v')
	// 測試Wrap
	err = Wrap(err)
	err.Format(nil, 'v')
	// 測試Cause
	cause := err.Cause()
	fmt.Println(cause)
	err2 := C(1001, "test2")
	// 測試Code
	fmt.Println(err2.Code())

}
