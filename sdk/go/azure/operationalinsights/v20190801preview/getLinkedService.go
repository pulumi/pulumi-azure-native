


package v20190801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:operationalinsights/v20190801preview:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	LinkedServiceName string `pulumi:"linkedServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedServiceResult struct {
	Id                    string  `pulumi:"id"`
	Name                  string  `pulumi:"name"`
	ResourceId            *string `pulumi:"resourceId"`
	Type                  string  `pulumi:"type"`
	WriteAccessResourceId *string `pulumi:"writeAccessResourceId"`
}

func LookupLinkedServiceOutput(ctx *pulumi.Context, args LookupLinkedServiceOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedServiceResult, error) {
			args := v.(LookupLinkedServiceArgs)
			r, err := LookupLinkedService(ctx, &args, opts...)
			var s LookupLinkedServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedServiceResultOutput)
}

type LookupLinkedServiceOutputArgs struct {
	LinkedServiceName pulumi.StringInput `pulumi:"linkedServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupLinkedServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceArgs)(nil)).Elem()
}


type LookupLinkedServiceResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceResult)(nil)).Elem()
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutput() LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutputWithContext(ctx context.Context) LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupLinkedServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) WriteAccessResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) *string { return v.WriteAccessResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedServiceResultOutput{})
}
