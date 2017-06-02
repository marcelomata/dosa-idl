// Code generated by thriftrw v1.2.0
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Dosa_MultiRemove_Args struct {
	Request *MultiRemoveRequest `json:"request,omitempty"`
}

func (v *Dosa_MultiRemove_Args) ToWire() (wire.Value, error) {
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

func _MultiRemoveRequest_Read(w wire.Value) (*MultiRemoveRequest, error) {
	var v MultiRemoveRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *Dosa_MultiRemove_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _MultiRemoveRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Dosa_MultiRemove_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("Dosa_MultiRemove_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_MultiRemove_Args) Equals(rhs *Dosa_MultiRemove_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}
	return true
}

func (v *Dosa_MultiRemove_Args) MethodName() string {
	return "multiRemove"
}

func (v *Dosa_MultiRemove_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Dosa_MultiRemove_Helper = struct {
	Args           func(request *MultiRemoveRequest) *Dosa_MultiRemove_Args
	IsException    func(error) bool
	WrapResponse   func(*MultiRemoveResponse, error) (*Dosa_MultiRemove_Result, error)
	UnwrapResponse func(*Dosa_MultiRemove_Result) (*MultiRemoveResponse, error)
}{}

func init() {
	Dosa_MultiRemove_Helper.Args = func(request *MultiRemoveRequest) *Dosa_MultiRemove_Args {
		return &Dosa_MultiRemove_Args{Request: request}
	}
	Dosa_MultiRemove_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}
	Dosa_MultiRemove_Helper.WrapResponse = func(success *MultiRemoveResponse, err error) (*Dosa_MultiRemove_Result, error) {
		if err == nil {
			return &Dosa_MultiRemove_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_MultiRemove_Result.ClientError")
			}
			return &Dosa_MultiRemove_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_MultiRemove_Result.ServerError")
			}
			return &Dosa_MultiRemove_Result{ServerError: e}, nil
		}
		return nil, err
	}
	Dosa_MultiRemove_Helper.UnwrapResponse = func(result *Dosa_MultiRemove_Result) (success *MultiRemoveResponse, err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Dosa_MultiRemove_Result struct {
	Success     *MultiRemoveResponse `json:"success,omitempty"`
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

func (v *Dosa_MultiRemove_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.ClientError != nil {
		w, err = v.ClientError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ServerError != nil {
		w, err = v.ServerError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Dosa_MultiRemove_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _MultiRemoveResponse_Read(w wire.Value) (*MultiRemoveResponse, error) {
	var v MultiRemoveResponse
	err := v.FromWire(w)
	return &v, err
}

func (v *Dosa_MultiRemove_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _MultiRemoveResponse_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.ClientError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ServerError, err = _InternalServerError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.ClientError != nil {
		count++
	}
	if v.ServerError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Dosa_MultiRemove_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Dosa_MultiRemove_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.ClientError != nil {
		fields[i] = fmt.Sprintf("ClientError: %v", v.ClientError)
		i++
	}
	if v.ServerError != nil {
		fields[i] = fmt.Sprintf("ServerError: %v", v.ServerError)
		i++
	}
	return fmt.Sprintf("Dosa_MultiRemove_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Dosa_MultiRemove_Result) Equals(rhs *Dosa_MultiRemove_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.ClientError == nil && rhs.ClientError == nil) || (v.ClientError != nil && rhs.ClientError != nil && v.ClientError.Equals(rhs.ClientError))) {
		return false
	}
	if !((v.ServerError == nil && rhs.ServerError == nil) || (v.ServerError != nil && rhs.ServerError != nil && v.ServerError.Equals(rhs.ServerError))) {
		return false
	}
	return true
}

func (v *Dosa_MultiRemove_Result) MethodName() string {
	return "multiRemove"
}

func (v *Dosa_MultiRemove_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
