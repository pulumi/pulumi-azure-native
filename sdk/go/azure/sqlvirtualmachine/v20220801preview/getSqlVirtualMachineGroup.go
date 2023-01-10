


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlVirtualMachineGroup(ctx *pulumi.Context, args *LookupSqlVirtualMachineGroupArgs, opts ...pulumi.InvokeOption) (*LookupSqlVirtualMachineGroupResult, error) {
	var rv LookupSqlVirtualMachineGroupResult
	err := ctx.Invoke("azure-native:sqlvirtualmachine/v20220801preview:getSqlVirtualMachineGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlVirtualMachineGroupArgs struct {
	ResourceGroupName          string `pulumi:"resourceGroupName"`
	SqlVirtualMachineGroupName string `pulumi:"sqlVirtualMachineGroupName"`
}


type LookupSqlVirtualMachineGroupResult struct {
	ClusterConfiguration string                     `pulumi:"clusterConfiguration"`
	ClusterManagerType   string                     `pulumi:"clusterManagerType"`
	Id                   string                     `pulumi:"id"`
	Location             string                     `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	ProvisioningState    string                     `pulumi:"provisioningState"`
	ScaleType            string                     `pulumi:"scaleType"`
	SqlImageOffer        *string                    `pulumi:"sqlImageOffer"`
	SqlImageSku          *string                    `pulumi:"sqlImageSku"`
	SystemData           SystemDataResponse         `pulumi:"systemData"`
	Tags                 map[string]string          `pulumi:"tags"`
	Type                 string                     `pulumi:"type"`
	WsfcDomainProfile    *WsfcDomainProfileResponse `pulumi:"wsfcDomainProfile"`
}

func LookupSqlVirtualMachineGroupOutput(ctx *pulumi.Context, args LookupSqlVirtualMachineGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSqlVirtualMachineGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlVirtualMachineGroupResult, error) {
			args := v.(LookupSqlVirtualMachineGroupArgs)
			r, err := LookupSqlVirtualMachineGroup(ctx, &args, opts...)
			var s LookupSqlVirtualMachineGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlVirtualMachineGroupResultOutput)
}

type LookupSqlVirtualMachineGroupOutputArgs struct {
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlVirtualMachineGroupName pulumi.StringInput `pulumi:"sqlVirtualMachineGroupName"`
}

func (LookupSqlVirtualMachineGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlVirtualMachineGroupArgs)(nil)).Elem()
}


type LookupSqlVirtualMachineGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSqlVirtualMachineGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlVirtualMachineGroupResult)(nil)).Elem()
}

func (o LookupSqlVirtualMachineGroupResultOutput) ToLookupSqlVirtualMachineGroupResultOutput() LookupSqlVirtualMachineGroupResultOutput {
	return o
}

func (o LookupSqlVirtualMachineGroupResultOutput) ToLookupSqlVirtualMachineGroupResultOutputWithContext(ctx context.Context) LookupSqlVirtualMachineGroupResultOutput {
	return o
}

func (o LookupSqlVirtualMachineGroupResultOutput) ClusterConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.ClusterConfiguration }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) ClusterManagerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.ClusterManagerType }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) ScaleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.ScaleType }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) SqlImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) *string { return v.SqlImageOffer }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) SqlImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) *string { return v.SqlImageSku }).(pulumi.StringPtrOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSqlVirtualMachineGroupResultOutput) WsfcDomainProfile() WsfcDomainProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineGroupResult) *WsfcDomainProfileResponse { return v.WsfcDomainProfile }).(WsfcDomainProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlVirtualMachineGroupResultOutput{})
}
