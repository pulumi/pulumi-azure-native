


package v20221012preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAttachedNetworkByDevCenter(ctx *pulumi.Context, args *LookupAttachedNetworkByDevCenterArgs, opts ...pulumi.InvokeOption) (*LookupAttachedNetworkByDevCenterResult, error) {
	var rv LookupAttachedNetworkByDevCenterResult
	err := ctx.Invoke("azure-native:devcenter/v20221012preview:getAttachedNetworkByDevCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttachedNetworkByDevCenterArgs struct {
	AttachedNetworkConnectionName string `pulumi:"attachedNetworkConnectionName"`
	DevCenterName                 string `pulumi:"devCenterName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupAttachedNetworkByDevCenterResult struct {
	DomainJoinType            string             `pulumi:"domainJoinType"`
	HealthCheckStatus         string             `pulumi:"healthCheckStatus"`
	Id                        string             `pulumi:"id"`
	Name                      string             `pulumi:"name"`
	NetworkConnectionId       string             `pulumi:"networkConnectionId"`
	NetworkConnectionLocation string             `pulumi:"networkConnectionLocation"`
	ProvisioningState         string             `pulumi:"provisioningState"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Type                      string             `pulumi:"type"`
}

func LookupAttachedNetworkByDevCenterOutput(ctx *pulumi.Context, args LookupAttachedNetworkByDevCenterOutputArgs, opts ...pulumi.InvokeOption) LookupAttachedNetworkByDevCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttachedNetworkByDevCenterResult, error) {
			args := v.(LookupAttachedNetworkByDevCenterArgs)
			r, err := LookupAttachedNetworkByDevCenter(ctx, &args, opts...)
			var s LookupAttachedNetworkByDevCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttachedNetworkByDevCenterResultOutput)
}

type LookupAttachedNetworkByDevCenterOutputArgs struct {
	AttachedNetworkConnectionName pulumi.StringInput `pulumi:"attachedNetworkConnectionName"`
	DevCenterName                 pulumi.StringInput `pulumi:"devCenterName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAttachedNetworkByDevCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedNetworkByDevCenterArgs)(nil)).Elem()
}


type LookupAttachedNetworkByDevCenterResultOutput struct{ *pulumi.OutputState }

func (LookupAttachedNetworkByDevCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedNetworkByDevCenterResult)(nil)).Elem()
}

func (o LookupAttachedNetworkByDevCenterResultOutput) ToLookupAttachedNetworkByDevCenterResultOutput() LookupAttachedNetworkByDevCenterResultOutput {
	return o
}

func (o LookupAttachedNetworkByDevCenterResultOutput) ToLookupAttachedNetworkByDevCenterResultOutputWithContext(ctx context.Context) LookupAttachedNetworkByDevCenterResultOutput {
	return o
}

func (o LookupAttachedNetworkByDevCenterResultOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.DomainJoinType }).(pulumi.StringOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) NetworkConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.NetworkConnectionId }).(pulumi.StringOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) NetworkConnectionLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.NetworkConnectionLocation }).(pulumi.StringOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAttachedNetworkByDevCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttachedNetworkByDevCenterResultOutput{})
}
