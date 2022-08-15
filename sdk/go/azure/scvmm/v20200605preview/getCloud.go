


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloud(ctx *pulumi.Context, args *LookupCloudArgs, opts ...pulumi.InvokeOption) (*LookupCloudResult, error) {
	var rv LookupCloudResult
	err := ctx.Invoke("azure-native:scvmm/v20200605preview:getCloud", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudArgs struct {
	CloudName         string `pulumi:"cloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCloudResult struct {
	CloudCapacity      CloudCapacityResponse      `pulumi:"cloudCapacity"`
	CloudName          string                     `pulumi:"cloudName"`
	ExtendedLocation   ExtendedLocationResponse   `pulumi:"extendedLocation"`
	Id                 string                     `pulumi:"id"`
	InventoryItemId    *string                    `pulumi:"inventoryItemId"`
	Location           string                     `pulumi:"location"`
	Name               string                     `pulumi:"name"`
	ProvisioningState  string                     `pulumi:"provisioningState"`
	StorageQoSPolicies []StorageQoSPolicyResponse `pulumi:"storageQoSPolicies"`
	SystemData         SystemDataResponse         `pulumi:"systemData"`
	Tags               map[string]string          `pulumi:"tags"`
	Type               string                     `pulumi:"type"`
	Uuid               *string                    `pulumi:"uuid"`
	VmmServerId        *string                    `pulumi:"vmmServerId"`
}

func LookupCloudOutput(ctx *pulumi.Context, args LookupCloudOutputArgs, opts ...pulumi.InvokeOption) LookupCloudResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudResult, error) {
			args := v.(LookupCloudArgs)
			r, err := LookupCloud(ctx, &args, opts...)
			var s LookupCloudResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudResultOutput)
}

type LookupCloudOutputArgs struct {
	CloudName         pulumi.StringInput `pulumi:"cloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCloudOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudArgs)(nil)).Elem()
}


type LookupCloudResultOutput struct{ *pulumi.OutputState }

func (LookupCloudResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudResult)(nil)).Elem()
}

func (o LookupCloudResultOutput) ToLookupCloudResultOutput() LookupCloudResultOutput {
	return o
}

func (o LookupCloudResultOutput) ToLookupCloudResultOutputWithContext(ctx context.Context) LookupCloudResultOutput {
	return o
}

func (o LookupCloudResultOutput) CloudCapacity() CloudCapacityResponseOutput {
	return o.ApplyT(func(v LookupCloudResult) CloudCapacityResponse { return v.CloudCapacity }).(CloudCapacityResponseOutput)
}

func (o LookupCloudResultOutput) CloudName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.CloudName }).(pulumi.StringOutput)
}

func (o LookupCloudResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupCloudResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o LookupCloudResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupCloudResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCloudResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCloudResultOutput) StorageQoSPolicies() StorageQoSPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupCloudResult) []StorageQoSPolicyResponse { return v.StorageQoSPolicies }).(StorageQoSPolicyResponseArrayOutput)
}

func (o LookupCloudResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCloudResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCloudResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCloudResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCloudResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o LookupCloudResultOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudResult) *string { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudResultOutput{})
}
