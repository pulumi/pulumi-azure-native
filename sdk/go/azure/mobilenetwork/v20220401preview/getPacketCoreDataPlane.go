


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPacketCoreDataPlane(ctx *pulumi.Context, args *LookupPacketCoreDataPlaneArgs, opts ...pulumi.InvokeOption) (*LookupPacketCoreDataPlaneResult, error) {
	var rv LookupPacketCoreDataPlaneResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20220401preview:getPacketCoreDataPlane", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPacketCoreDataPlaneArgs struct {
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	PacketCoreDataPlaneName    string `pulumi:"packetCoreDataPlaneName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupPacketCoreDataPlaneResult struct {
	CreatedAt                *string                     `pulumi:"createdAt"`
	CreatedBy                *string                     `pulumi:"createdBy"`
	CreatedByType            *string                     `pulumi:"createdByType"`
	Id                       string                      `pulumi:"id"`
	LastModifiedAt           *string                     `pulumi:"lastModifiedAt"`
	LastModifiedBy           *string                     `pulumi:"lastModifiedBy"`
	LastModifiedByType       *string                     `pulumi:"lastModifiedByType"`
	Location                 string                      `pulumi:"location"`
	Name                     string                      `pulumi:"name"`
	ProvisioningState        string                      `pulumi:"provisioningState"`
	SystemData               SystemDataResponse          `pulumi:"systemData"`
	Tags                     map[string]string           `pulumi:"tags"`
	Type                     string                      `pulumi:"type"`
	UserPlaneAccessInterface InterfacePropertiesResponse `pulumi:"userPlaneAccessInterface"`
}

func LookupPacketCoreDataPlaneOutput(ctx *pulumi.Context, args LookupPacketCoreDataPlaneOutputArgs, opts ...pulumi.InvokeOption) LookupPacketCoreDataPlaneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPacketCoreDataPlaneResult, error) {
			args := v.(LookupPacketCoreDataPlaneArgs)
			r, err := LookupPacketCoreDataPlane(ctx, &args, opts...)
			var s LookupPacketCoreDataPlaneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPacketCoreDataPlaneResultOutput)
}

type LookupPacketCoreDataPlaneOutputArgs struct {
	PacketCoreControlPlaneName pulumi.StringInput `pulumi:"packetCoreControlPlaneName"`
	PacketCoreDataPlaneName    pulumi.StringInput `pulumi:"packetCoreDataPlaneName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPacketCoreDataPlaneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreDataPlaneArgs)(nil)).Elem()
}


type LookupPacketCoreDataPlaneResultOutput struct{ *pulumi.OutputState }

func (LookupPacketCoreDataPlaneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreDataPlaneResult)(nil)).Elem()
}

func (o LookupPacketCoreDataPlaneResultOutput) ToLookupPacketCoreDataPlaneResultOutput() LookupPacketCoreDataPlaneResultOutput {
	return o
}

func (o LookupPacketCoreDataPlaneResultOutput) ToLookupPacketCoreDataPlaneResultOutputWithContext(ctx context.Context) LookupPacketCoreDataPlaneResultOutput {
	return o
}

func (o LookupPacketCoreDataPlaneResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPacketCoreDataPlaneResultOutput) UserPlaneAccessInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreDataPlaneResult) InterfacePropertiesResponse { return v.UserPlaneAccessInterface }).(InterfacePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPacketCoreDataPlaneResultOutput{})
}
