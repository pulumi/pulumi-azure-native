


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConsoleWithLocation(ctx *pulumi.Context, args *LookupConsoleWithLocationArgs, opts ...pulumi.InvokeOption) (*LookupConsoleWithLocationResult, error) {
	var rv LookupConsoleWithLocationResult
	err := ctx.Invoke("azure-native:portal/v20181001:getConsoleWithLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsoleWithLocationArgs struct {
	ConsoleName string `pulumi:"consoleName"`
	Location    string `pulumi:"location"`
}


type LookupConsoleWithLocationResult struct {
	Properties ConsolePropertiesResponse `pulumi:"properties"`
}

func LookupConsoleWithLocationOutput(ctx *pulumi.Context, args LookupConsoleWithLocationOutputArgs, opts ...pulumi.InvokeOption) LookupConsoleWithLocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConsoleWithLocationResult, error) {
			args := v.(LookupConsoleWithLocationArgs)
			r, err := LookupConsoleWithLocation(ctx, &args, opts...)
			var s LookupConsoleWithLocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConsoleWithLocationResultOutput)
}

type LookupConsoleWithLocationOutputArgs struct {
	ConsoleName pulumi.StringInput `pulumi:"consoleName"`
	Location    pulumi.StringInput `pulumi:"location"`
}

func (LookupConsoleWithLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleWithLocationArgs)(nil)).Elem()
}


type LookupConsoleWithLocationResultOutput struct{ *pulumi.OutputState }

func (LookupConsoleWithLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleWithLocationResult)(nil)).Elem()
}

func (o LookupConsoleWithLocationResultOutput) ToLookupConsoleWithLocationResultOutput() LookupConsoleWithLocationResultOutput {
	return o
}

func (o LookupConsoleWithLocationResultOutput) ToLookupConsoleWithLocationResultOutputWithContext(ctx context.Context) LookupConsoleWithLocationResultOutput {
	return o
}

func (o LookupConsoleWithLocationResultOutput) Properties() ConsolePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConsoleWithLocationResult) ConsolePropertiesResponse { return v.Properties }).(ConsolePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConsoleWithLocationResultOutput{})
}
