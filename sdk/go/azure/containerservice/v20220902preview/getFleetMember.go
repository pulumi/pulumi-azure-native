


package v20220902preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFleetMember(ctx *pulumi.Context, args *LookupFleetMemberArgs, opts ...pulumi.InvokeOption) (*LookupFleetMemberResult, error) {
	var rv LookupFleetMemberResult
	err := ctx.Invoke("azure-native:containerservice/v20220902preview:getFleetMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetMemberArgs struct {
	FleetMemberName   string `pulumi:"fleetMemberName"`
	FleetName         string `pulumi:"fleetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFleetMemberResult struct {
	ClusterResourceId *string            `pulumi:"clusterResourceId"`
	Etag              string             `pulumi:"etag"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupFleetMemberOutput(ctx *pulumi.Context, args LookupFleetMemberOutputArgs, opts ...pulumi.InvokeOption) LookupFleetMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFleetMemberResult, error) {
			args := v.(LookupFleetMemberArgs)
			r, err := LookupFleetMember(ctx, &args, opts...)
			var s LookupFleetMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFleetMemberResultOutput)
}

type LookupFleetMemberOutputArgs struct {
	FleetMemberName   pulumi.StringInput `pulumi:"fleetMemberName"`
	FleetName         pulumi.StringInput `pulumi:"fleetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFleetMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetMemberArgs)(nil)).Elem()
}


type LookupFleetMemberResultOutput struct{ *pulumi.OutputState }

func (LookupFleetMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetMemberResult)(nil)).Elem()
}

func (o LookupFleetMemberResultOutput) ToLookupFleetMemberResultOutput() LookupFleetMemberResultOutput {
	return o
}

func (o LookupFleetMemberResultOutput) ToLookupFleetMemberResultOutputWithContext(ctx context.Context) LookupFleetMemberResultOutput {
	return o
}

func (o LookupFleetMemberResultOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) *string { return v.ClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupFleetMemberResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFleetMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFleetMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFleetMemberResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFleetMemberResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFleetMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetMemberResultOutput{})
}
