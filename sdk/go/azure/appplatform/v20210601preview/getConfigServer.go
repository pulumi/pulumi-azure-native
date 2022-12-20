


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigServer(ctx *pulumi.Context, args *LookupConfigServerArgs, opts ...pulumi.InvokeOption) (*LookupConfigServerResult, error) {
	var rv LookupConfigServerResult
	err := ctx.Invoke("azure-native:appplatform/v20210601preview:getConfigServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupConfigServerResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties ConfigServerPropertiesResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}

func LookupConfigServerOutput(ctx *pulumi.Context, args LookupConfigServerOutputArgs, opts ...pulumi.InvokeOption) LookupConfigServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigServerResult, error) {
			args := v.(LookupConfigServerArgs)
			r, err := LookupConfigServer(ctx, &args, opts...)
			var s LookupConfigServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigServerResultOutput)
}

type LookupConfigServerOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupConfigServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigServerArgs)(nil)).Elem()
}


type LookupConfigServerResultOutput struct{ *pulumi.OutputState }

func (LookupConfigServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigServerResult)(nil)).Elem()
}

func (o LookupConfigServerResultOutput) ToLookupConfigServerResultOutput() LookupConfigServerResultOutput {
	return o
}

func (o LookupConfigServerResultOutput) ToLookupConfigServerResultOutputWithContext(ctx context.Context) LookupConfigServerResultOutput {
	return o
}

func (o LookupConfigServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConfigServerResultOutput) Properties() ConfigServerPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigServerResult) ConfigServerPropertiesResponse { return v.Properties }).(ConfigServerPropertiesResponseOutput)
}

func (o LookupConfigServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigServerResultOutput{})
}
