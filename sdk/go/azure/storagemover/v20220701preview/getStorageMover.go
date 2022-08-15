


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageMover(ctx *pulumi.Context, args *LookupStorageMoverArgs, opts ...pulumi.InvokeOption) (*LookupStorageMoverResult, error) {
	var rv LookupStorageMoverResult
	err := ctx.Invoke("azure-native:storagemover/v20220701preview:getStorageMover", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageMoverArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageMoverName  string `pulumi:"storageMoverName"`
}


type LookupStorageMoverResult struct {
	Description       *string            `pulumi:"description"`
	Id                string             `pulumi:"id"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupStorageMoverOutput(ctx *pulumi.Context, args LookupStorageMoverOutputArgs, opts ...pulumi.InvokeOption) LookupStorageMoverResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageMoverResult, error) {
			args := v.(LookupStorageMoverArgs)
			r, err := LookupStorageMover(ctx, &args, opts...)
			var s LookupStorageMoverResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageMoverResultOutput)
}

type LookupStorageMoverOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageMoverName  pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupStorageMoverOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageMoverArgs)(nil)).Elem()
}


type LookupStorageMoverResultOutput struct{ *pulumi.OutputState }

func (LookupStorageMoverResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageMoverResult)(nil)).Elem()
}

func (o LookupStorageMoverResultOutput) ToLookupStorageMoverResultOutput() LookupStorageMoverResultOutput {
	return o
}

func (o LookupStorageMoverResultOutput) ToLookupStorageMoverResultOutputWithContext(ctx context.Context) LookupStorageMoverResultOutput {
	return o
}

func (o LookupStorageMoverResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupStorageMoverResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageMoverResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStorageMoverResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageMoverResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStorageMoverResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStorageMoverResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStorageMoverResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageMoverResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageMoverResultOutput{})
}
