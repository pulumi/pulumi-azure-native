


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatastoreArgs struct {
	DatastoreName     string `pulumi:"datastoreName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatastoreResult struct {
	CustomResourceName string                    `pulumi:"customResourceName"`
	ExtendedLocation   *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                 string                    `pulumi:"id"`
	InventoryItemId    *string                   `pulumi:"inventoryItemId"`
	Kind               *string                   `pulumi:"kind"`
	Location           string                    `pulumi:"location"`
	MoName             string                    `pulumi:"moName"`
	MoRefId            *string                   `pulumi:"moRefId"`
	Name               string                    `pulumi:"name"`
	ProvisioningState  string                    `pulumi:"provisioningState"`
	Statuses           []ResourceStatusResponse  `pulumi:"statuses"`
	SystemData         SystemDataResponse        `pulumi:"systemData"`
	Tags               map[string]string         `pulumi:"tags"`
	Type               string                    `pulumi:"type"`
	Uuid               string                    `pulumi:"uuid"`
	VCenterId          *string                   `pulumi:"vCenterId"`
}

func LookupDatastoreOutput(ctx *pulumi.Context, args LookupDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreResult, error) {
			args := v.(LookupDatastoreArgs)
			r, err := LookupDatastore(ctx, &args, opts...)
			var s LookupDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreResultOutput)
}

type LookupDatastoreOutputArgs struct {
	DatastoreName     pulumi.StringInput `pulumi:"datastoreName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreArgs)(nil)).Elem()
}


type LookupDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreResult)(nil)).Elem()
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutput() LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutputWithContext(ctx context.Context) LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupDatastoreResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupDatastoreResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.MoName }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o LookupDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupDatastoreResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o LookupDatastoreResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatastoreResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDatastoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDatastoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreResultOutput{})
}
