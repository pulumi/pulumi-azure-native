


package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobInventoryPolicy(ctx *pulumi.Context, args *LookupBlobInventoryPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBlobInventoryPolicyResult, error) {
	var rv LookupBlobInventoryPolicyResult
	err := ctx.Invoke("azure-native:storage:getBlobInventoryPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobInventoryPolicyArgs struct {
	AccountName             string `pulumi:"accountName"`
	BlobInventoryPolicyName string `pulumi:"blobInventoryPolicyName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupBlobInventoryPolicyResult struct {
	Id               string                            `pulumi:"id"`
	LastModifiedTime string                            `pulumi:"lastModifiedTime"`
	Name             string                            `pulumi:"name"`
	Policy           BlobInventoryPolicySchemaResponse `pulumi:"policy"`
	SystemData       SystemDataResponse                `pulumi:"systemData"`
	Type             string                            `pulumi:"type"`
}

func LookupBlobInventoryPolicyOutput(ctx *pulumi.Context, args LookupBlobInventoryPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBlobInventoryPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobInventoryPolicyResult, error) {
			args := v.(LookupBlobInventoryPolicyArgs)
			r, err := LookupBlobInventoryPolicy(ctx, &args, opts...)
			var s LookupBlobInventoryPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobInventoryPolicyResultOutput)
}

type LookupBlobInventoryPolicyOutputArgs struct {
	AccountName             pulumi.StringInput `pulumi:"accountName"`
	BlobInventoryPolicyName pulumi.StringInput `pulumi:"blobInventoryPolicyName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobInventoryPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobInventoryPolicyArgs)(nil)).Elem()
}


type LookupBlobInventoryPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBlobInventoryPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobInventoryPolicyResult)(nil)).Elem()
}

func (o LookupBlobInventoryPolicyResultOutput) ToLookupBlobInventoryPolicyResultOutput() LookupBlobInventoryPolicyResultOutput {
	return o
}

func (o LookupBlobInventoryPolicyResultOutput) ToLookupBlobInventoryPolicyResultOutputWithContext(ctx context.Context) LookupBlobInventoryPolicyResultOutput {
	return o
}

func (o LookupBlobInventoryPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobInventoryPolicyResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupBlobInventoryPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobInventoryPolicyResultOutput) Policy() BlobInventoryPolicySchemaResponseOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) BlobInventoryPolicySchemaResponse { return v.Policy }).(BlobInventoryPolicySchemaResponseOutput)
}

func (o LookupBlobInventoryPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBlobInventoryPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobInventoryPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobInventoryPolicyResultOutput{})
}
