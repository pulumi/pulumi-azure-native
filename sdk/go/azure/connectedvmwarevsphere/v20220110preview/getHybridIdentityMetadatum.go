


package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridIdentityMetadatum(ctx *pulumi.Context, args *LookupHybridIdentityMetadatumArgs, opts ...pulumi.InvokeOption) (*LookupHybridIdentityMetadatumResult, error) {
	var rv LookupHybridIdentityMetadatumResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220110preview:getHybridIdentityMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridIdentityMetadatumArgs struct {
	MetadataName       string `pulumi:"metadataName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualMachineName string `pulumi:"virtualMachineName"`
}


type LookupHybridIdentityMetadatumResult struct {
	Id                string             `pulumi:"id"`
	Identity          IdentityResponse   `pulumi:"identity"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	PublicKey         *string            `pulumi:"publicKey"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
	VmId              *string            `pulumi:"vmId"`
}

func LookupHybridIdentityMetadatumOutput(ctx *pulumi.Context, args LookupHybridIdentityMetadatumOutputArgs, opts ...pulumi.InvokeOption) LookupHybridIdentityMetadatumResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridIdentityMetadatumResult, error) {
			args := v.(LookupHybridIdentityMetadatumArgs)
			r, err := LookupHybridIdentityMetadatum(ctx, &args, opts...)
			var s LookupHybridIdentityMetadatumResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridIdentityMetadatumResultOutput)
}

type LookupHybridIdentityMetadatumOutputArgs struct {
	MetadataName       pulumi.StringInput `pulumi:"metadataName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupHybridIdentityMetadatumOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridIdentityMetadatumArgs)(nil)).Elem()
}


type LookupHybridIdentityMetadatumResultOutput struct{ *pulumi.OutputState }

func (LookupHybridIdentityMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridIdentityMetadatumResult)(nil)).Elem()
}

func (o LookupHybridIdentityMetadatumResultOutput) ToLookupHybridIdentityMetadatumResultOutput() LookupHybridIdentityMetadatumResultOutput {
	return o
}

func (o LookupHybridIdentityMetadatumResultOutput) ToLookupHybridIdentityMetadatumResultOutputWithContext(ctx context.Context) LookupHybridIdentityMetadatumResultOutput {
	return o
}

func (o LookupHybridIdentityMetadatumResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHybridIdentityMetadatumResultOutput) Identity() IdentityResponseOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) IdentityResponse { return v.Identity }).(IdentityResponseOutput)
}

func (o LookupHybridIdentityMetadatumResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHybridIdentityMetadatumResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupHybridIdentityMetadatumResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupHybridIdentityMetadatumResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupHybridIdentityMetadatumResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHybridIdentityMetadatumResultOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridIdentityMetadatumResultOutput{})
}
