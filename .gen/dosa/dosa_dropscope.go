// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Dosa_DropScope_Args represents the arguments for the Dosa.dropScope function.
//
// The arguments for dropScope are sent and received over the wire as this struct.
type Dosa_DropScope_Args struct {
	Request *DropScopeRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_DropScope_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Dosa_DropScope_Args) ToWire() (wire.Value, error) {
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

func _DropScopeRequest_Read(w wire.Value) (*DropScopeRequest, error) {
	var v DropScopeRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_DropScope_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_DropScope_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_DropScope_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_DropScope_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _DropScopeRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_DropScope_Args
// struct.
func (v *Dosa_DropScope_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_DropScope_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_DropScope_Args match the
// provided Dosa_DropScope_Args.
//
// This function performs a deep comparison.
func (v *Dosa_DropScope_Args) Equals(rhs *Dosa_DropScope_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "dropScope" for this struct.
func (v *Dosa_DropScope_Args) MethodName() string {
	return "dropScope"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_DropScope_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_DropScope_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.dropScope
// function.
var Dosa_DropScope_Helper = struct {
	// Args accepts the parameters of dropScope in-order and returns
	// the arguments struct for the function.
	Args func(
		request *DropScopeRequest,
	) *Dosa_DropScope_Args

	// IsException returns true if the given error can be thrown
	// by dropScope.
	//
	// An error can be thrown by dropScope only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for dropScope
	// given the error returned by it. The provided error may
	// be nil if dropScope did not fail.
	//
	// This allows mapping errors returned by dropScope into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// dropScope
	//
	//   err := dropScope(args)
	//   result, err := Dosa_DropScope_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from dropScope: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Dosa_DropScope_Result, error)

	// UnwrapResponse takes the result struct for dropScope
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if dropScope threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Dosa_DropScope_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_DropScope_Result) error
}{}

func init() {
	Dosa_DropScope_Helper.Args = func(
		request *DropScopeRequest,
	) *Dosa_DropScope_Args {
		return &Dosa_DropScope_Args{
			Request: request,
		}
	}

	Dosa_DropScope_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}

	Dosa_DropScope_Helper.WrapResponse = func(err error) (*Dosa_DropScope_Result, error) {
		if err == nil {
			return &Dosa_DropScope_Result{}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_DropScope_Result.ClientError")
			}
			return &Dosa_DropScope_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_DropScope_Result.ServerError")
			}
			return &Dosa_DropScope_Result{ServerError: e}, nil
		}

		return nil, err
	}
	Dosa_DropScope_Helper.UnwrapResponse = func(result *Dosa_DropScope_Result) (err error) {
		if result.ClientError != nil {
			err = result.ClientError
			return
		}
		if result.ServerError != nil {
			err = result.ServerError
			return
		}
		return
	}

}

// Dosa_DropScope_Result represents the result of a Dosa.dropScope function call.
//
// The result of a dropScope execution is sent and received over the wire as this struct.
type Dosa_DropScope_Result struct {
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

// ToWire translates a Dosa_DropScope_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Dosa_DropScope_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

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

	if i > 1 {
		return wire.Value{}, fmt.Errorf("Dosa_DropScope_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Dosa_DropScope_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_DropScope_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_DropScope_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_DropScope_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
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
	if v.ClientError != nil {
		count++
	}
	if v.ServerError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Dosa_DropScope_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_DropScope_Result
// struct.
func (v *Dosa_DropScope_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.ClientError != nil {
		fields[i] = fmt.Sprintf("ClientError: %v", v.ClientError)
		i++
	}
	if v.ServerError != nil {
		fields[i] = fmt.Sprintf("ServerError: %v", v.ServerError)
		i++
	}

	return fmt.Sprintf("Dosa_DropScope_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_DropScope_Result match the
// provided Dosa_DropScope_Result.
//
// This function performs a deep comparison.
func (v *Dosa_DropScope_Result) Equals(rhs *Dosa_DropScope_Result) bool {
	if !((v.ClientError == nil && rhs.ClientError == nil) || (v.ClientError != nil && rhs.ClientError != nil && v.ClientError.Equals(rhs.ClientError))) {
		return false
	}
	if !((v.ServerError == nil && rhs.ServerError == nil) || (v.ServerError != nil && rhs.ServerError != nil && v.ServerError.Equals(rhs.ServerError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "dropScope" for this struct.
func (v *Dosa_DropScope_Result) MethodName() string {
	return "dropScope"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_DropScope_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
