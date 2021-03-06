package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// inline static
// hid_t _go_hdf5_H5P_DEFAULT() { return H5P_DEFAULT; }
import "C"

import (
	"fmt"
	"runtime"
)

type PropType C.hid_t

type PropList struct {
	Identifier
}

var (
	P_DEFAULT *PropList = newPropList(C._go_hdf5_H5P_DEFAULT())
)

func newPropList(id C.hid_t) *PropList {
	p := &PropList{Identifier{id}}
	runtime.SetFinalizer(p, (*PropList).finalizer)
	return p
}

func (p *PropList) finalizer() {
	if err := p.Close(); err != nil {
		panic(fmt.Errorf("error closing PropList: %s", err))
	}
}

// NewPropList creates a new PropList as an instance of a property list class.
func NewPropList(cls_id PropType) (*PropList, error) {
	hid := C.H5Pcreate(C.hid_t(cls_id))
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPropList(hid), nil
}

// Close terminates access to a PropList.
func (p *PropList) Close() error {
	if p.id == 0 {
		return nil
	}
	err := h5err(C.H5Pclose(p.id))
	p.id = 0
	return err
}

// Copy copies an existing PropList to create a new PropList.
func (p *PropList) Copy() (*PropList, error) {
	hid := C.H5Pcopy(p.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPropList(hid), nil
}
