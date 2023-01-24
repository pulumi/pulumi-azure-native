


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArcSetting(ctx *pulumi.Context, args *LookupArcSettingArgs, opts ...pulumi.InvokeOption) (*LookupArcSettingResult, error) {
	var rv LookupArcSettingResult
	err := ctx.Invoke("azure-native:azurestackhci/v20221201:getArcSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArcSettingArgs struct {
	ArcSettingName    string `pulumi:"arcSettingName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupArcSettingResult struct {
	AggregateState              string                              `pulumi:"aggregateState"`
	ArcApplicationClientId      *string                             `pulumi:"arcApplicationClientId"`
	ArcApplicationObjectId      *string                             `pulumi:"arcApplicationObjectId"`
	ArcApplicationTenantId      *string                             `pulumi:"arcApplicationTenantId"`
	ArcInstanceResourceGroup    *string                             `pulumi:"arcInstanceResourceGroup"`
	ArcServicePrincipalObjectId *string                             `pulumi:"arcServicePrincipalObjectId"`
	ConnectivityProperties      []ArcConnectivityPropertiesResponse `pulumi:"connectivityProperties"`
	Id                          string                              `pulumi:"id"`
	Name                        string                              `pulumi:"name"`
	PerNodeDetails              []PerNodeStateResponse              `pulumi:"perNodeDetails"`
	ProvisioningState           string                              `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse                  `pulumi:"systemData"`
	Type                        string                              `pulumi:"type"`
}

func LookupArcSettingOutput(ctx *pulumi.Context, args LookupArcSettingOutputArgs, opts ...pulumi.InvokeOption) LookupArcSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupArcSettingResult, error) {
			args := v.(LookupArcSettingArgs)
			r, err := LookupArcSetting(ctx, &args, opts...)
			var s LookupArcSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupArcSettingResultOutput)
}

type LookupArcSettingOutputArgs struct {
	ArcSettingName    pulumi.StringInput `pulumi:"arcSettingName"`
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupArcSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcSettingArgs)(nil)).Elem()
}


type LookupArcSettingResultOutput struct{ *pulumi.OutputState }

func (LookupArcSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArcSettingResult)(nil)).Elem()
}

func (o LookupArcSettingResultOutput) ToLookupArcSettingResultOutput() LookupArcSettingResultOutput {
	return o
}

func (o LookupArcSettingResultOutput) ToLookupArcSettingResultOutputWithContext(ctx context.Context) LookupArcSettingResultOutput {
	return o
}

func (o LookupArcSettingResultOutput) AggregateState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.AggregateState }).(pulumi.StringOutput)
}

func (o LookupArcSettingResultOutput) ArcApplicationClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcApplicationClientId }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) ArcApplicationObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcApplicationObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) ArcApplicationTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcApplicationTenantId }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) ArcInstanceResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcInstanceResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) ArcServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.ArcServicePrincipalObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) ConnectivityProperties() ArcConnectivityPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupArcSettingResult) []ArcConnectivityPropertiesResponse { return v.ConnectivityProperties }).(ArcConnectivityPropertiesResponseArrayOutput)
}

func (o LookupArcSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupArcSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupArcSettingResultOutput) PerNodeDetails() PerNodeStateResponseArrayOutput {
	return o.ApplyT(func(v LookupArcSettingResult) []PerNodeStateResponse { return v.PerNodeDetails }).(PerNodeStateResponseArrayOutput)
}

func (o LookupArcSettingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupArcSettingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupArcSettingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupArcSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArcSettingResultOutput{})
}
