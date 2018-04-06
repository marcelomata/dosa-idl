// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package dosa

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Dosa_Scan_Args represents the arguments for the Dosa.scan function.
//
// The arguments for scan are sent and received over the wire as this struct.
type Dosa_Scan_Args struct {
	Request *ScanRequest `json:"request,omitempty"`
}

// ToWire translates a Dosa_Scan_Args struct into a Thrift-level intermediate
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
func (v *Dosa_Scan_Args) ToWire() (wire.Value, error) {
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

func _ScanRequest_Read(w wire.Value) (*ScanRequest, error) {
	var v ScanRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_Scan_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_Scan_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_Scan_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_Scan_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _ScanRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Dosa_Scan_Args
// struct.
func (v *Dosa_Scan_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Dosa_Scan_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_Scan_Args match the
// provided Dosa_Scan_Args.
//
// This function performs a deep comparison.
func (v *Dosa_Scan_Args) Equals(rhs *Dosa_Scan_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "scan" for this struct.
func (v *Dosa_Scan_Args) MethodName() string {
	return "scan"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Dosa_Scan_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Dosa_Scan_Helper provides functions that aid in handling the
// parameters and return values of the Dosa.scan
// function.
var Dosa_Scan_Helper = struct {
	// Args accepts the parameters of scan in-order and returns
	// the arguments struct for the function.
	Args func(
		request *ScanRequest,
	) *Dosa_Scan_Args

	// IsException returns true if the given error can be thrown
	// by scan.
	//
	// An error can be thrown by scan only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for scan
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// scan into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by scan
	//
	//   value, err := scan(args)
	//   result, err := Dosa_Scan_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from scan: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*ScanResponse, error) (*Dosa_Scan_Result, error)

	// UnwrapResponse takes the result struct for scan
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if scan threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Dosa_Scan_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Dosa_Scan_Result) (*ScanResponse, error)
}{}

func init() {
	Dosa_Scan_Helper.Args = func(
		request *ScanRequest,
	) *Dosa_Scan_Args {
		return &Dosa_Scan_Args{
			Request: request,
		}
	}

	Dosa_Scan_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BadRequestError:
			return true
		case *InternalServerError:
			return true
		default:
			return false
		}
	}

	Dosa_Scan_Helper.WrapResponse = func(success *ScanResponse, err error) (*Dosa_Scan_Result, error) {
		if err == nil {
			return &Dosa_Scan_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Scan_Result.ClientError")
			}
			return &Dosa_Scan_Result{ClientError: e}, nil
		case *InternalServerError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Dosa_Scan_Result.ServerError")
			}
			return &Dosa_Scan_Result{ServerError: e}, nil
		}

		return nil, err
	}
	Dosa_Scan_Helper.UnwrapResponse = func(result *Dosa_Scan_Result) (success *ScanResponse, err error) {
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

// Dosa_Scan_Result represents the result of a Dosa.scan function call.
//
// The result of a scan execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Dosa_Scan_Result struct {
	// Value returned by scan after a successful execution.
	Success     *ScanResponse        `json:"success,omitempty"`
	ClientError *BadRequestError     `json:"clientError,omitempty"`
	ServerError *InternalServerError `json:"serverError,omitempty"`
}

// ToWire translates a Dosa_Scan_Result struct into a Thrift-level intermediate
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
func (v *Dosa_Scan_Result) ToWire() (wire.Value, error) {
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
		return wire.Value{}, fmt.Errorf("Dosa_Scan_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ScanResponse_Read(w wire.Value) (*ScanResponse, error) {
	var v ScanResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Dosa_Scan_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Dosa_Scan_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Dosa_Scan_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Dosa_Scan_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ScanResponse_Read(field.Value)
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
		return fmt.Errorf("Dosa_Scan_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Dosa_Scan_Result
// struct.
func (v *Dosa_Scan_Result) String() string {
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

	return fmt.Sprintf("Dosa_Scan_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Dosa_Scan_Result match the
// provided Dosa_Scan_Result.
//
// This function performs a deep comparison.
func (v *Dosa_Scan_Result) Equals(rhs *Dosa_Scan_Result) bool {
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

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "scan" for this struct.
func (v *Dosa_Scan_Result) MethodName() string {
	return "scan"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Dosa_Scan_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
