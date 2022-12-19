


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	LinkName          string `pulumi:"linkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedServiceResult struct {
	Id         string                     `pulumi:"id"`
	Identity   *IdentityResponse          `pulumi:"identity"`
	Location   *string                    `pulumi:"location"`
	Name       string                     `pulumi:"name"`
	Properties LinkedServicePropsResponse `pulumi:"properties"`
	Type       string                     `pulumi:"type"`
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
	LinkName          pulumi.StringInput `pulumi:"linkName"`
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

func (o LookupLinkedServiceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupLinkedServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLinkedServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) Properties() LinkedServicePropsResponseOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) LinkedServicePropsResponse { return v.Properties }).(LinkedServicePropsResponseOutput)
}

func (o LookupLinkedServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedServiceResultOutput{})
}
