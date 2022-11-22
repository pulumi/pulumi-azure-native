


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskPool(ctx *pulumi.Context, args *LookupDiskPoolArgs, opts ...pulumi.InvokeOption) (*LookupDiskPoolResult, error) {
	var rv LookupDiskPoolResult
	err := ctx.Invoke("azure-native:storagepool/v20210801:getDiskPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskPoolArgs struct {
	DiskPoolName      string `pulumi:"diskPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskPoolResult struct {
	AdditionalCapabilities []string               `pulumi:"additionalCapabilities"`
	AvailabilityZones      []string               `pulumi:"availabilityZones"`
	Disks                  []DiskResponse         `pulumi:"disks"`
	Id                     string                 `pulumi:"id"`
	Location               string                 `pulumi:"location"`
	ManagedBy              string                 `pulumi:"managedBy"`
	ManagedByExtended      []string               `pulumi:"managedByExtended"`
	Name                   string                 `pulumi:"name"`
	ProvisioningState      string                 `pulumi:"provisioningState"`
	Status                 string                 `pulumi:"status"`
	SubnetId               string                 `pulumi:"subnetId"`
	SystemData             SystemMetadataResponse `pulumi:"systemData"`
	Tags                   map[string]string      `pulumi:"tags"`
	Tier                   *string                `pulumi:"tier"`
	Type                   string                 `pulumi:"type"`
}

func LookupDiskPoolOutput(ctx *pulumi.Context, args LookupDiskPoolOutputArgs, opts ...pulumi.InvokeOption) LookupDiskPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskPoolResult, error) {
			args := v.(LookupDiskPoolArgs)
			r, err := LookupDiskPool(ctx, &args, opts...)
			var s LookupDiskPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskPoolResultOutput)
}

type LookupDiskPoolOutputArgs struct {
	DiskPoolName      pulumi.StringInput `pulumi:"diskPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskPoolArgs)(nil)).Elem()
}


type LookupDiskPoolResultOutput struct{ *pulumi.OutputState }

func (LookupDiskPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskPoolResult)(nil)).Elem()
}

func (o LookupDiskPoolResultOutput) ToLookupDiskPoolResultOutput() LookupDiskPoolResultOutput {
	return o
}

func (o LookupDiskPoolResultOutput) ToLookupDiskPoolResultOutputWithContext(ctx context.Context) LookupDiskPoolResultOutput {
	return o
}

func (o LookupDiskPoolResultOutput) AdditionalCapabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) []string { return v.AdditionalCapabilities }).(pulumi.StringArrayOutput)
}

func (o LookupDiskPoolResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupDiskPoolResultOutput) Disks() DiskResponseArrayOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) []DiskResponse { return v.Disks }).(DiskResponseArrayOutput)
}

func (o LookupDiskPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDiskPoolResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupDiskPoolResultOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) []string { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

func (o LookupDiskPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDiskPoolResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupDiskPoolResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupDiskPoolResultOutput) SystemData() SystemMetadataResponseOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) SystemMetadataResponse { return v.SystemData }).(SystemMetadataResponseOutput)
}

func (o LookupDiskPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDiskPoolResultOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func (o LookupDiskPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskPoolResultOutput{})
}
