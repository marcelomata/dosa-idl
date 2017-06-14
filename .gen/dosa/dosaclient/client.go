// Code generated by thriftrw-plugin-yarpc
// @generated

package dosaclient

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc"
	"github.com/uber/dosa-idl/.gen/dosa"
)

// Interface is a client for the Dosa service.
type Interface interface {
	CheckSchema(
		ctx context.Context,
		Request *dosa.CheckSchemaRequest,
		opts ...yarpc.CallOption,
	) (*dosa.CheckSchemaResponse, error)

	CheckSchemaStatus(
		ctx context.Context,
		Request *dosa.CheckSchemaStatusRequest,
		opts ...yarpc.CallOption,
	) (*dosa.CheckSchemaStatusResponse, error)

	CreateIfNotExists(
		ctx context.Context,
		Request *dosa.CreateRequest,
		opts ...yarpc.CallOption,
	) error

	CreateScope(
		ctx context.Context,
		Request *dosa.CreateScopeRequest,
		opts ...yarpc.CallOption,
	) error

	DropScope(
		ctx context.Context,
		Request *dosa.DropScopeRequest,
		opts ...yarpc.CallOption,
	) error

	MultiRead(
		ctx context.Context,
		Request *dosa.MultiReadRequest,
		opts ...yarpc.CallOption,
	) (*dosa.MultiReadResponse, error)

	MultiRemove(
		ctx context.Context,
		Request *dosa.MultiRemoveRequest,
		opts ...yarpc.CallOption,
	) (*dosa.MultiRemoveResponse, error)

	MultiUpsert(
		ctx context.Context,
		Request *dosa.MultiUpsertRequest,
		opts ...yarpc.CallOption,
	) (*dosa.MultiUpsertResponse, error)

	Range(
		ctx context.Context,
		Request *dosa.RangeRequest,
		opts ...yarpc.CallOption,
	) (*dosa.RangeResponse, error)

	Read(
		ctx context.Context,
		Request *dosa.ReadRequest,
		opts ...yarpc.CallOption,
	) (*dosa.ReadResponse, error)

	Remove(
		ctx context.Context,
		Request *dosa.RemoveRequest,
		opts ...yarpc.CallOption,
	) error

	RemoveRange(
		ctx context.Context,
		Request *dosa.RemoveRangeRequest,
		opts ...yarpc.CallOption,
	) error

	Scan(
		ctx context.Context,
		Request *dosa.ScanRequest,
		opts ...yarpc.CallOption,
	) (*dosa.ScanResponse, error)

	ScopeExists(
		ctx context.Context,
		Request *dosa.ScopeExistsRequest,
		opts ...yarpc.CallOption,
	) (*dosa.ScopeExistsResponse, error)

	Search(
		ctx context.Context,
		Request *dosa.SearchRequest,
		opts ...yarpc.CallOption,
	) (*dosa.SearchResponse, error)

	TruncateScope(
		ctx context.Context,
		Request *dosa.TruncateScopeRequest,
		opts ...yarpc.CallOption,
	) error

	Upsert(
		ctx context.Context,
		Request *dosa.UpsertRequest,
		opts ...yarpc.CallOption,
	) error

	UpsertSchema(
		ctx context.Context,
		Request *dosa.UpsertSchemaRequest,
		opts ...yarpc.CallOption,
	) (*dosa.UpsertSchemaResponse, error)
}

// New builds a new client for the Dosa service.
//
// 	client := dosaclient.New(dispatcher.ClientConfig("dosa"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "Dosa",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(func(c transport.ClientConfig) Interface {
		return New(c)
	})
}

type client struct {
	c thrift.Client
}

func (c client) CheckSchema(
	ctx context.Context,
	_Request *dosa.CheckSchemaRequest,
	opts ...yarpc.CallOption,
) (success *dosa.CheckSchemaResponse, err error) {

	args := dosa.Dosa_CheckSchema_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_CheckSchema_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_CheckSchema_Helper.UnwrapResponse(&result)
	return
}

func (c client) CheckSchemaStatus(
	ctx context.Context,
	_Request *dosa.CheckSchemaStatusRequest,
	opts ...yarpc.CallOption,
) (success *dosa.CheckSchemaStatusResponse, err error) {

	args := dosa.Dosa_CheckSchemaStatus_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_CheckSchemaStatus_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_CheckSchemaStatus_Helper.UnwrapResponse(&result)
	return
}

func (c client) CreateIfNotExists(
	ctx context.Context,
	_Request *dosa.CreateRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := dosa.Dosa_CreateIfNotExists_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_CreateIfNotExists_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = dosa.Dosa_CreateIfNotExists_Helper.UnwrapResponse(&result)
	return
}

func (c client) CreateScope(
	ctx context.Context,
	_Request *dosa.CreateScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := dosa.Dosa_CreateScope_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_CreateScope_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = dosa.Dosa_CreateScope_Helper.UnwrapResponse(&result)
	return
}

func (c client) DropScope(
	ctx context.Context,
	_Request *dosa.DropScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := dosa.Dosa_DropScope_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_DropScope_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = dosa.Dosa_DropScope_Helper.UnwrapResponse(&result)
	return
}

func (c client) MultiRead(
	ctx context.Context,
	_Request *dosa.MultiReadRequest,
	opts ...yarpc.CallOption,
) (success *dosa.MultiReadResponse, err error) {

	args := dosa.Dosa_MultiRead_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_MultiRead_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_MultiRead_Helper.UnwrapResponse(&result)
	return
}

func (c client) MultiRemove(
	ctx context.Context,
	_Request *dosa.MultiRemoveRequest,
	opts ...yarpc.CallOption,
) (success *dosa.MultiRemoveResponse, err error) {

	args := dosa.Dosa_MultiRemove_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_MultiRemove_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_MultiRemove_Helper.UnwrapResponse(&result)
	return
}

func (c client) MultiUpsert(
	ctx context.Context,
	_Request *dosa.MultiUpsertRequest,
	opts ...yarpc.CallOption,
) (success *dosa.MultiUpsertResponse, err error) {

	args := dosa.Dosa_MultiUpsert_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_MultiUpsert_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_MultiUpsert_Helper.UnwrapResponse(&result)
	return
}

func (c client) Range(
	ctx context.Context,
	_Request *dosa.RangeRequest,
	opts ...yarpc.CallOption,
) (success *dosa.RangeResponse, err error) {

	args := dosa.Dosa_Range_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_Range_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_Range_Helper.UnwrapResponse(&result)
	return
}

func (c client) Read(
	ctx context.Context,
	_Request *dosa.ReadRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ReadResponse, err error) {

	args := dosa.Dosa_Read_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_Read_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_Read_Helper.UnwrapResponse(&result)
	return
}

func (c client) Remove(
	ctx context.Context,
	_Request *dosa.RemoveRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := dosa.Dosa_Remove_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_Remove_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = dosa.Dosa_Remove_Helper.UnwrapResponse(&result)
	return
}

func (c client) RemoveRange(
	ctx context.Context,
	_Request *dosa.RemoveRangeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := dosa.Dosa_RemoveRange_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_RemoveRange_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = dosa.Dosa_RemoveRange_Helper.UnwrapResponse(&result)
	return
}

func (c client) Scan(
	ctx context.Context,
	_Request *dosa.ScanRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ScanResponse, err error) {

	args := dosa.Dosa_Scan_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_Scan_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_Scan_Helper.UnwrapResponse(&result)
	return
}

func (c client) ScopeExists(
	ctx context.Context,
	_Request *dosa.ScopeExistsRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ScopeExistsResponse, err error) {

	args := dosa.Dosa_ScopeExists_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_ScopeExists_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_ScopeExists_Helper.UnwrapResponse(&result)
	return
}

func (c client) Search(
	ctx context.Context,
	_Request *dosa.SearchRequest,
	opts ...yarpc.CallOption,
) (success *dosa.SearchResponse, err error) {

	args := dosa.Dosa_Search_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_Search_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_Search_Helper.UnwrapResponse(&result)
	return
}

func (c client) TruncateScope(
	ctx context.Context,
	_Request *dosa.TruncateScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := dosa.Dosa_TruncateScope_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_TruncateScope_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = dosa.Dosa_TruncateScope_Helper.UnwrapResponse(&result)
	return
}

func (c client) Upsert(
	ctx context.Context,
	_Request *dosa.UpsertRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := dosa.Dosa_Upsert_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_Upsert_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = dosa.Dosa_Upsert_Helper.UnwrapResponse(&result)
	return
}

func (c client) UpsertSchema(
	ctx context.Context,
	_Request *dosa.UpsertSchemaRequest,
	opts ...yarpc.CallOption,
) (success *dosa.UpsertSchemaResponse, err error) {

	args := dosa.Dosa_UpsertSchema_Helper.Args(_Request)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result dosa.Dosa_UpsertSchema_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = dosa.Dosa_UpsertSchema_Helper.UnwrapResponse(&result)
	return
}
