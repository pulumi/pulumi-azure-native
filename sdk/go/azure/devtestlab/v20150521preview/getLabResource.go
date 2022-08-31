


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupLabResource(ctx *pulumi.Context, args *LookupLabResourceArgs, opts ...pulumi.InvokeOption) (*LookupLabResourceResult, error) {
	var rv LookupLabResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getLabResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabResourceArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLabResourceResult struct {
	ArtifactsStorageAccount *string           `pulumi:"artifactsStorageAccount"`
	CreatedDate             *string           `pulumi:"createdDate"`
	DefaultStorageAccount   *string           `pulumi:"defaultStorageAccount"`
	DefaultVirtualNetworkId *string           `pulumi:"defaultVirtualNetworkId"`
	Id                      *string           `pulumi:"id"`
	LabStorageType          *string           `pulumi:"labStorageType"`
	Location                *string           `pulumi:"location"`
	Name                    *string           `pulumi:"name"`
	ProvisioningState       *string           `pulumi:"provisioningState"`
	StorageAccounts         []string          `pulumi:"storageAccounts"`
	Tags                    map[string]string `pulumi:"tags"`
	Type                    *string           `pulumi:"type"`
	VaultName               *string           `pulumi:"vaultName"`
}

func LookupLabResourceOutput(ctx *pulumi.Context, args LookupLabResourceOutputArgs, opts ...pulumi.InvokeOption) LookupLabResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResourceResult, error) {
			args := v.(LookupLabResourceArgs)
			r, err := LookupLabResource(ctx, &args, opts...)
			var s LookupLabResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResourceResultOutput)
}

type LookupLabResourceOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLabResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResourceArgs)(nil)).Elem()
}


type LookupLabResourceResultOutput struct{ *pulumi.OutputState }

func (LookupLabResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResourceResult)(nil)).Elem()
}

func (o LookupLabResourceResultOutput) ToLookupLabResourceResultOutput() LookupLabResourceResultOutput {
	return o
}

func (o LookupLabResourceResultOutput) ToLookupLabResourceResultOutputWithContext(ctx context.Context) LookupLabResourceResultOutput {
	return o
}

func (o LookupLabResourceResultOutput) ArtifactsStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.ArtifactsStorageAccount }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) DefaultStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.DefaultStorageAccount }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) DefaultVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.DefaultVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) LabStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.LabStorageType }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) StorageAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLabResourceResult) []string { return v.StorageAccounts }).(pulumi.StringArrayOutput)
}

func (o LookupLabResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLabResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupLabResourceResultOutput) VaultName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResourceResult) *string { return v.VaultName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResourceResultOutput{})
}
