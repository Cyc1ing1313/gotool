package errors

import (
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
)

func Ck(err *error) bool {
	if *err != nil {
		typName := reflect.ValueOf(*err).Elem().Type().Name()
		if typName != "withStack" && typName != "fundamental" {
			*err = errors.WithStack(*err)
			errData := unsafe.Pointer(*((*uintptr)(unsafe.Pointer((uintptr)(unsafe.Pointer(err)) + uintptr(8)))))
			stack := (*[]uintptr)(unsafe.Pointer(*((*uintptr)(unsafe.Pointer((uintptr)(unsafe.Pointer(errData)) + uintptr(16))))))
			(*stack)[0] = ((*stack)[1] - uintptr(16))
		}
	}
	return *err != nil
}
