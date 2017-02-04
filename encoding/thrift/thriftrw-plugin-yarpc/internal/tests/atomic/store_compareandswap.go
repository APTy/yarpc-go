// Code generated by thriftrw v1.0.0
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Store_CompareAndSwap_Args struct {
	Request *CompareAndSwap `json:"request,omitempty"`
}

func (v *Store_CompareAndSwap_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _CompareAndSwap_Read(w wire.Value) (*CompareAndSwap, error) {
	var v CompareAndSwap
	err := v.FromWire(w)
	return &v, err
}

func (v *Store_CompareAndSwap_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _CompareAndSwap_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Store_CompareAndSwap_Args) String() string {
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("Store_CompareAndSwap_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Store_CompareAndSwap_Args) MethodName() string {
	return "compareAndSwap"
}

func (v *Store_CompareAndSwap_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Store_CompareAndSwap_Helper = struct {
	Args           func(request *CompareAndSwap) *Store_CompareAndSwap_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*Store_CompareAndSwap_Result, error)
	UnwrapResponse func(*Store_CompareAndSwap_Result) error
}{}

func init() {
	Store_CompareAndSwap_Helper.Args = func(request *CompareAndSwap) *Store_CompareAndSwap_Args {
		return &Store_CompareAndSwap_Args{Request: request}
	}
	Store_CompareAndSwap_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *IntegerMismatchError:
			return true
		default:
			return false
		}
	}
	Store_CompareAndSwap_Helper.WrapResponse = func(err error) (*Store_CompareAndSwap_Result, error) {
		if err == nil {
			return &Store_CompareAndSwap_Result{}, nil
		}
		switch e := err.(type) {
		case *IntegerMismatchError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Store_CompareAndSwap_Result.Mismatch")
			}
			return &Store_CompareAndSwap_Result{Mismatch: e}, nil
		}
		return nil, err
	}
	Store_CompareAndSwap_Helper.UnwrapResponse = func(result *Store_CompareAndSwap_Result) (err error) {
		if result.Mismatch != nil {
			err = result.Mismatch
			return
		}
		return
	}
}

type Store_CompareAndSwap_Result struct {
	Mismatch *IntegerMismatchError `json:"mismatch,omitempty"`
}

func (v *Store_CompareAndSwap_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Mismatch != nil {
		w, err = v.Mismatch.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("Store_CompareAndSwap_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _IntegerMismatchError_Read(w wire.Value) (*IntegerMismatchError, error) {
	var v IntegerMismatchError
	err := v.FromWire(w)
	return &v, err
}

func (v *Store_CompareAndSwap_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Mismatch, err = _IntegerMismatchError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Mismatch != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Store_CompareAndSwap_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *Store_CompareAndSwap_Result) String() string {
	var fields [1]string
	i := 0
	if v.Mismatch != nil {
		fields[i] = fmt.Sprintf("Mismatch: %v", v.Mismatch)
		i++
	}
	return fmt.Sprintf("Store_CompareAndSwap_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Store_CompareAndSwap_Result) MethodName() string {
	return "compareAndSwap"
}

func (v *Store_CompareAndSwap_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
