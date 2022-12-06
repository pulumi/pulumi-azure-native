


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetvirtualharddiskRetrieve(ctx *pulumi.Context, args *GetvirtualharddiskRetrieveArgs, opts ...pulumi.InvokeOption) (*GetvirtualharddiskRetrieveResult, error) {
	var rv GetvirtualharddiskRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210701preview:getvirtualharddiskRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetvirtualharddiskRetrieveArgs struct {
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	VirtualharddisksName string `pulumi:"virtualharddisksName"`
}


type GetvirtualharddiskRetrieveResult struct {
	BlockSizeBytes      *int                          `pulumi:"blockSizeBytes"`
	DiskSizeBytes       *float64                      `pulumi:"diskSizeBytes"`
	Dynamic             *bool                         `pulumi:"dynamic"`
	ExtendedLocation    *ExtendedLocationResponse     `pulumi:"extendedLocation"`
	Id                  string                        `pulumi:"id"`
	Location            string                        `pulumi:"location"`
	LogicalSectorBytes  *int                          `pulumi:"logicalSectorBytes"`
	Name                string                        `pulumi:"name"`
	PhysicalSectorBytes *int                          `pulumi:"physicalSectorBytes"`
	ProvisioningState   string                        `pulumi:"provisioningState"`
	ResourceName        *string                       `pulumi:"resourceName"`
	Status              VirtualHardDiskStatusResponse `pulumi:"status"`
	SystemData          SystemDataResponse            `pulumi:"systemData"`
	Tags                map[string]string             `pulumi:"tags"`
	Type                string                        `pulumi:"type"`
}

func GetvirtualharddiskRetrieveOutput(ctx *pulumi.Context, args GetvirtualharddiskRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetvirtualharddiskRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetvirtualharddiskRetrieveResult, error) {
			args := v.(GetvirtualharddiskRetrieveArgs)
			r, err := GetvirtualharddiskRetrieve(ctx, &args, opts...)
			var s GetvirtualharddiskRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetvirtualharddiskRetrieveResultOutput)
}

type GetvirtualharddiskRetrieveOutputArgs struct {
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualharddisksName pulumi.StringInput `pulumi:"virtualharddisksName"`
}

func (GetvirtualharddiskRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualharddiskRetrieveArgs)(nil)).Elem()
}


type GetvirtualharddiskRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetvirtualharddiskRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualharddiskRetrieveResult)(nil)).Elem()
}

func (o GetvirtualharddiskRetrieveResultOutput) ToGetvirtualharddiskRetrieveResultOutput() GetvirtualharddiskRetrieveResultOutput {
	return o
}

func (o GetvirtualharddiskRetrieveResultOutput) ToGetvirtualharddiskRetrieveResultOutputWithContext(ctx context.Context) GetvirtualharddiskRetrieveResultOutput {
	return o
}

func (o GetvirtualharddiskRetrieveResultOutput) BlockSizeBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) *int { return v.BlockSizeBytes }).(pulumi.IntPtrOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) DiskSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) *float64 { return v.DiskSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) Dynamic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) *bool { return v.Dynamic }).(pulumi.BoolPtrOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) LogicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) *int { return v.LogicalSectorBytes }).(pulumi.IntPtrOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) PhysicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) *int { return v.PhysicalSectorBytes }).(pulumi.IntPtrOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) Status() VirtualHardDiskStatusResponseOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) VirtualHardDiskStatusResponse { return v.Status }).(VirtualHardDiskStatusResponseOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetvirtualharddiskRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualharddiskRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetvirtualharddiskRetrieveResultOutput{})
}
