// Auto-generated by avdl-compiler v1.3.24 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/pprof.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ProcessorProfileArg struct {
	SessionID              int         `codec:"sessionID" json:"sessionID"`
	ProfileFile            string      `codec:"profileFile" json:"profileFile"`
	ProfileDurationSeconds DurationSec `codec:"profileDurationSeconds" json:"profileDurationSeconds"`
}

type LogProcessorProfileArg struct {
	SessionID              int         `codec:"sessionID" json:"sessionID"`
	LogDirForMobile        string      `codec:"logDirForMobile" json:"logDirForMobile"`
	ProfileDurationSeconds DurationSec `codec:"profileDurationSeconds" json:"profileDurationSeconds"`
}

type TraceArg struct {
	SessionID            int         `codec:"sessionID" json:"sessionID"`
	TraceFile            string      `codec:"traceFile" json:"traceFile"`
	TraceDurationSeconds DurationSec `codec:"traceDurationSeconds" json:"traceDurationSeconds"`
}

type LogTraceArg struct {
	SessionID            int         `codec:"sessionID" json:"sessionID"`
	LogDirForMobile      string      `codec:"logDirForMobile" json:"logDirForMobile"`
	TraceDurationSeconds DurationSec `codec:"traceDurationSeconds" json:"traceDurationSeconds"`
}

type PprofInterface interface {
	ProcessorProfile(context.Context, ProcessorProfileArg) error
	LogProcessorProfile(context.Context, LogProcessorProfileArg) error
	Trace(context.Context, TraceArg) error
	LogTrace(context.Context, LogTraceArg) error
}

func PprofProtocol(i PprofInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.pprof",
		Methods: map[string]rpc.ServeHandlerDescription{
			"processorProfile": {
				MakeArg: func() interface{} {
					ret := make([]ProcessorProfileArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ProcessorProfileArg)
					if !ok {
						err = rpc.NewTypeError((*[]ProcessorProfileArg)(nil), args)
						return
					}
					err = i.ProcessorProfile(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"logProcessorProfile": {
				MakeArg: func() interface{} {
					ret := make([]LogProcessorProfileArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LogProcessorProfileArg)
					if !ok {
						err = rpc.NewTypeError((*[]LogProcessorProfileArg)(nil), args)
						return
					}
					err = i.LogProcessorProfile(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"trace": {
				MakeArg: func() interface{} {
					ret := make([]TraceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TraceArg)
					if !ok {
						err = rpc.NewTypeError((*[]TraceArg)(nil), args)
						return
					}
					err = i.Trace(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"logTrace": {
				MakeArg: func() interface{} {
					ret := make([]LogTraceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LogTraceArg)
					if !ok {
						err = rpc.NewTypeError((*[]LogTraceArg)(nil), args)
						return
					}
					err = i.LogTrace(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type PprofClient struct {
	Cli rpc.GenericClient
}

func (c PprofClient) ProcessorProfile(ctx context.Context, __arg ProcessorProfileArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pprof.processorProfile", []interface{}{__arg}, nil)
	return
}

func (c PprofClient) LogProcessorProfile(ctx context.Context, __arg LogProcessorProfileArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pprof.logProcessorProfile", []interface{}{__arg}, nil)
	return
}

func (c PprofClient) Trace(ctx context.Context, __arg TraceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pprof.trace", []interface{}{__arg}, nil)
	return
}

func (c PprofClient) LogTrace(ctx context.Context, __arg LogTraceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.pprof.logTrace", []interface{}{__arg}, nil)
	return
}
