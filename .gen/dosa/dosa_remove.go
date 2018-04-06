// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Dosa_Remove_Args represents the arguments for the Dosa.remove function.
//
// The arguments for remove are sent and received over the wire as this struct.
type Dosa_Remove_Args struct {
	Request *RemoveRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_Remove_Args struct into a Thrift-level intermediate
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
func (v *Dosa_Remove_Args) ToWire() (wire.Value, error) {
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

func _RemoveRequest_Read(w wire.Value) (*RemoveRequest, error) {
	var v RemoveRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_Remove_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_Remove_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_Remove_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_Remove_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _RemoveRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_Remove_Args
// struct.
func (v *Dosa_Remove_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_Remove_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_Remove_Args match the
// provided Dosa_Remove_Args.
//
// This function performs a deep comparison.
func (v *Dosa_Remove_Args) Equals(rhs *Dosa_Remove_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "remove" for this struct.
func (v *Dosa_Remove_Args) MethodName() string {
	return "remove"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_Remove_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_Remove_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.remove
// function.
var Dosa_Remove_Helper = struct {
	// Args accepts the parameters of remove in-order and returns
	// the arguments struct for the function.
	Args func(
		request *RemoveRequest,
	) *Dosa_Remove_Args

	// IsException returns true if the given error can be thrown
	// by remove.
	//
	// An error can be thrown by remove only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for remove
	// given the error returned by it. The provided error may
	// be nil if remove did not fail.
	//
	// This allows mapping errors returned by remove into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// remove
	//
	//   err := remove(args)
	//   result, err := Dosa_Remove_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from remove: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Dosa_Remove_Result, error)

	// UnwrapResponse takes the result struct for remove
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if remove threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Dosa_Remove_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_Remove_Result) error
}{}

func init() {
	Dosa_Remove_Helper.Args = func(
		request *RemoveRequest,
	) *Dosa_Remove_Args {
		return &Dosa_Remove_Args{
			Request: request,
		}
	}

	Dosa_Remove_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}

	Dosa_Remove_Helper.WrapResponse = func(err error) (*Dosa_Remove_Result, error) {
		if err == nil {
			return &Dosa_Remove_Result{}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Remove_Result.ClientError")
			}
			return &Dosa_Remove_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Remove_Result.ServerError")
			}
			return &Dosa_Remove_Result{ServerError: e}, nil
		}

		return nil, err
	}
	Dosa_Remove_Helper.UnwrapResponse = func(result *Dosa_Remove_Result) (err error) {
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

// Dosa_Remove_Result represents the result of a Dosa.remove function call.
//
// The result of a remove execution is sent and received over the wire as this struct.
type Dosa_Remove_Result struct {
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

// ToWire translates a Dosa_Remove_Result struct into a Thrift-level intermediate
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
func (v *Dosa_Remove_Result) ToWire() (wire.Value, error) {
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
		return wire.Value{}, fmt.Errorf("Dosa_Remove_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Dosa_Remove_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_Remove_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_Remove_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_Remove_Result) FromWire(w wire.Value) error {
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
		return fmt.Errorf("Dosa_Remove_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_Remove_Result
// struct.
func (v *Dosa_Remove_Result) String() string {
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

	return fmt.Sprintf("Dosa_Remove_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_Remove_Result match the
// provided Dosa_Remove_Result.
//
// This function performs a deep comparison.
func (v *Dosa_Remove_Result) Equals(rhs *Dosa_Remove_Result) bool {
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
// This will always be "remove" for this struct.
func (v *Dosa_Remove_Result) MethodName() string {
	return "remove"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_Remove_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
