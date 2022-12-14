


package v20180501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSerialPort(ctx *pulumi.Context, args *LookupSerialPortArgs, opts ...pulumi.InvokeOption) (*LookupSerialPortResult, error) {
	var rv LookupSerialPortResult
	err := ctx.Invoke("azure-native:serialconsole/v20180501:getSerialPort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSerialPortArgs struct {
	ParentResource            string `pulumi:"parentResource"`
	ParentResourceType        string `pulumi:"parentResourceType"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	SerialPort                string `pulumi:"serialPort"`
}


type LookupSerialPortResult struct {
	Id    string  `pulumi:"id"`
	Name  string  `pulumi:"name"`
	State *string `pulumi:"state"`
	Type  string  `pulumi:"type"`
}

func LookupSerialPortOutput(ctx *pulumi.Context, args LookupSerialPortOutputArgs, opts ...pulumi.InvokeOption) LookupSerialPortResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSerialPortResult, error) {
			args := v.(LookupSerialPortArgs)
			r, err := LookupSerialPort(ctx, &args, opts...)
			var s LookupSerialPortResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSerialPortResultOutput)
}

type LookupSerialPortOutputArgs struct {
	ParentResource            pulumi.StringInput `pulumi:"parentResource"`
	ParentResourceType        pulumi.StringInput `pulumi:"parentResourceType"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceProviderNamespace pulumi.StringInput `pulumi:"resourceProviderNamespace"`
	SerialPort                pulumi.StringInput `pulumi:"serialPort"`
}

func (LookupSerialPortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSerialPortArgs)(nil)).Elem()
}


type LookupSerialPortResultOutput struct{ *pulumi.OutputState }

func (LookupSerialPortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSerialPortResult)(nil)).Elem()
}

func (o LookupSerialPortResultOutput) ToLookupSerialPortResultOutput() LookupSerialPortResultOutput {
	return o
}

func (o LookupSerialPortResultOutput) ToLookupSerialPortResultOutputWithContext(ctx context.Context) LookupSerialPortResultOutput {
	return o
}

func (o LookupSerialPortResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSerialPortResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSerialPortResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSerialPortResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSerialPortResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSerialPortResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupSerialPortResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSerialPortResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSerialPortResultOutput{})
}
