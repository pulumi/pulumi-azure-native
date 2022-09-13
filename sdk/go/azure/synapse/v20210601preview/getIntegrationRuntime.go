


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationRuntime(ctx *pulumi.Context, args *LookupIntegrationRuntimeArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationRuntimeResult, error) {
	var rv LookupIntegrationRuntimeResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getIntegrationRuntime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationRuntimeArgs struct {
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type LookupIntegrationRuntimeResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupIntegrationRuntimeOutput(ctx *pulumi.Context, args LookupIntegrationRuntimeOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationRuntimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationRuntimeResult, error) {
			args := v.(LookupIntegrationRuntimeArgs)
			r, err := LookupIntegrationRuntime(ctx, &args, opts...)
			var s LookupIntegrationRuntimeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationRuntimeResultOutput)
}

type LookupIntegrationRuntimeOutputArgs struct {
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName          pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIntegrationRuntimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationRuntimeArgs)(nil)).Elem()
}


type LookupIntegrationRuntimeResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationRuntimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationRuntimeResult)(nil)).Elem()
}

func (o LookupIntegrationRuntimeResultOutput) ToLookupIntegrationRuntimeResultOutput() LookupIntegrationRuntimeResultOutput {
	return o
}

func (o LookupIntegrationRuntimeResultOutput) ToLookupIntegrationRuntimeResultOutputWithContext(ctx context.Context) LookupIntegrationRuntimeResultOutput {
	return o
}

func (o LookupIntegrationRuntimeResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIntegrationRuntimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationRuntimeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationRuntimeResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupIntegrationRuntimeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationRuntimeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationRuntimeResultOutput{})
}
