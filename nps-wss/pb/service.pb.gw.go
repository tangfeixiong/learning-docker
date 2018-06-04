// Code generated by protoc-gen-grpc-gateway
// source: pb/service.proto
// DO NOT EDIT!

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_NpsWssService_LinkStatsCreation_0(ctx context.Context, marshaler runtime.Marshaler, client NpsWssServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq LinkReqResp
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.LinkStatsCreation(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_NpsWssService_NetFlowStats_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_NpsWssService_NetFlowStats_0(ctx context.Context, marshaler runtime.Marshaler, client NpsWssServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StatsReqResp
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_NpsWssService_NetFlowStats_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.NetFlowStats(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterNpsWssServiceHandlerFromEndpoint is same as RegisterNpsWssServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNpsWssServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterNpsWssServiceHandler(ctx, mux, conn)
}

// RegisterNpsWssServiceHandler registers the http handlers for service NpsWssService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNpsWssServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewNpsWssServiceClient(conn)

	mux.Handle("POST", pattern_NpsWssService_LinkStatsCreation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_NpsWssService_LinkStatsCreation_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_NpsWssService_LinkStatsCreation_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_NpsWssService_NetFlowStats_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_NpsWssService_NetFlowStats_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_NpsWssService_NetFlowStats_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_NpsWssService_LinkStatsCreation_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "links"}, ""))

	pattern_NpsWssService_NetFlowStats_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "stats"}, ""))
)

var (
	forward_NpsWssService_LinkStatsCreation_0 = runtime.ForwardResponseMessage

	forward_NpsWssService_NetFlowStats_0 = runtime.ForwardResponseMessage
)
