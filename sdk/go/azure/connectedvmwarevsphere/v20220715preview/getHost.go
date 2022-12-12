


package v20220715preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHost(ctx *pulumi.Context, args *LookupHostArgs, opts ...pulumi.InvokeOption) (*LookupHostResult, error) {
	var rv LookupHostResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220715preview:getHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostArgs struct {
	HostName          string `pulumi:"hostName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHostResult struct {
	CustomResourceName string                    `pulumi:"customResourceName"`
	DatastoreIds       []string                  `pulumi:"datastoreIds"`
	ExtendedLocation   *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                 string                    `pulumi:"id"`
	InventoryItemId    *string                   `pulumi:"inventoryItemId"`
	Kind               *string                   `pulumi:"kind"`
	Location           string                    `pulumi:"location"`
	MoName             string                    `pulumi:"moName"`
	MoRefId            *string                   `pulumi:"moRefId"`
	Name               string                    `pulumi:"name"`
	NetworkIds         []string                  `pulumi:"networkIds"`
	ProvisioningState  string                    `pulumi:"provisioningState"`
	Statuses           []ResourceStatusResponse  `pulumi:"statuses"`
	SystemData         SystemDataResponse        `pulumi:"systemData"`
	Tags               map[string]string         `pulumi:"tags"`
	Type               string                    `pulumi:"type"`
	Uuid               string                    `pulumi:"uuid"`
	VCenterId          *string                   `pulumi:"vCenterId"`
}

func LookupHostOutput(ctx *pulumi.Context, args LookupHostOutputArgs, opts ...pulumi.InvokeOption) LookupHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostResult, error) {
			args := v.(LookupHostArgs)
			r, err := LookupHost(ctx, &args, opts...)
			var s LookupHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostResultOutput)
}

type LookupHostOutputArgs struct {
	HostName          pulumi.StringInput `pulumi:"hostName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostArgs)(nil)).Elem()
}


type LookupHostResultOutput struct{ *pulumi.OutputState }

func (LookupHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostResult)(nil)).Elem()
}

func (o LookupHostResultOutput) ToLookupHostResultOutput() LookupHostResultOutput {
	return o
}

func (o LookupHostResultOutput) ToLookupHostResultOutputWithContext(ctx context.Context) LookupHostResultOutput {
	return o
}

func (o LookupHostResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) DatastoreIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostResult) []string { return v.DatastoreIds }).(pulumi.StringArrayOutput)
}

func (o LookupHostResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupHostResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupHostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupHostResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupHostResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.MoName }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o LookupHostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostResult) []string { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

func (o LookupHostResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupHostResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o LookupHostResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHostResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupHostResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHostResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupHostResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostResultOutput{})
}
