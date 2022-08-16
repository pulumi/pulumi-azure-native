


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBatchEndpoint(ctx *pulumi.Context, args *LookupBatchEndpointArgs, opts ...pulumi.InvokeOption) (*LookupBatchEndpointResult, error) {
	var rv LookupBatchEndpointResult
	err := ctx.Invoke("azure-native:machinelearningservices:getBatchEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type LookupBatchEndpointResult struct {
	Id         string                    `pulumi:"id"`
	Identity   *ResourceIdentityResponse `pulumi:"identity"`
	Kind       *string                   `pulumi:"kind"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties BatchEndpointResponse     `pulumi:"properties"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}

func LookupBatchEndpointOutput(ctx *pulumi.Context, args LookupBatchEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupBatchEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchEndpointResult, error) {
			args := v.(LookupBatchEndpointArgs)
			r, err := LookupBatchEndpoint(ctx, &args, opts...)
			var s LookupBatchEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBatchEndpointResultOutput)
}

type LookupBatchEndpointOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBatchEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchEndpointArgs)(nil)).Elem()
}

type LookupBatchEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupBatchEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchEndpointResult)(nil)).Elem()
}

func (o LookupBatchEndpointResultOutput) ToLookupBatchEndpointResultOutput() LookupBatchEndpointResultOutput {
	return o
}

func (o LookupBatchEndpointResultOutput) ToLookupBatchEndpointResultOutputWithContext(ctx context.Context) LookupBatchEndpointResultOutput {
	return o
}

func (o LookupBatchEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBatchEndpointResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o LookupBatchEndpointResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupBatchEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBatchEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBatchEndpointResultOutput) Properties() BatchEndpointResponseOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) BatchEndpointResponse { return v.Properties }).(BatchEndpointResponseOutput)
}

func (o LookupBatchEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBatchEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBatchEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchEndpointResultOutput{})
}
