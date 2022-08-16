


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConsole(ctx *pulumi.Context, args *LookupConsoleArgs, opts ...pulumi.InvokeOption) (*LookupConsoleResult, error) {
	var rv LookupConsoleResult
	err := ctx.Invoke("azure-native:portal/v20181001:getConsole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsoleArgs struct {
	ConsoleName string `pulumi:"consoleName"`
}


type LookupConsoleResult struct {
	Properties ConsolePropertiesResponse `pulumi:"properties"`
}

func LookupConsoleOutput(ctx *pulumi.Context, args LookupConsoleOutputArgs, opts ...pulumi.InvokeOption) LookupConsoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConsoleResult, error) {
			args := v.(LookupConsoleArgs)
			r, err := LookupConsole(ctx, &args, opts...)
			var s LookupConsoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConsoleResultOutput)
}

type LookupConsoleOutputArgs struct {
	ConsoleName pulumi.StringInput `pulumi:"consoleName"`
}

func (LookupConsoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleArgs)(nil)).Elem()
}


type LookupConsoleResultOutput struct{ *pulumi.OutputState }

func (LookupConsoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleResult)(nil)).Elem()
}

func (o LookupConsoleResultOutput) ToLookupConsoleResultOutput() LookupConsoleResultOutput {
	return o
}

func (o LookupConsoleResultOutput) ToLookupConsoleResultOutputWithContext(ctx context.Context) LookupConsoleResultOutput {
	return o
}

func (o LookupConsoleResultOutput) Properties() ConsolePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConsoleResult) ConsolePropertiesResponse { return v.Properties }).(ConsolePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConsoleResultOutput{})
}
