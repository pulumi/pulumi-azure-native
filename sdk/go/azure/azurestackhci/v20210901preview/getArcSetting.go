


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupArcSetting(ctx *pulumi.Context, args *LookupArcSettingArgs, opts ...pulumi.InvokeOption) (*LookupArcSettingResult, error) {
	var rv LookupArcSettingResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getArcSetting", args, &rv, opts...)
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
	AggregateState           string                 `pulumi:"aggregateState"`
	ArcInstanceResourceGroup string                 `pulumi:"arcInstanceResourceGroup"`
	CreatedAt                *string                `pulumi:"createdAt"`
	CreatedBy                *string                `pulumi:"createdBy"`
	CreatedByType            *string                `pulumi:"createdByType"`
	Id                       string                 `pulumi:"id"`
	LastModifiedAt           *string                `pulumi:"lastModifiedAt"`
	LastModifiedBy           *string                `pulumi:"lastModifiedBy"`
	LastModifiedByType       *string                `pulumi:"lastModifiedByType"`
	Name                     string                 `pulumi:"name"`
	PerNodeDetails           []PerNodeStateResponse `pulumi:"perNodeDetails"`
	ProvisioningState        string                 `pulumi:"provisioningState"`
	Type                     string                 `pulumi:"type"`
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

func (o LookupArcSettingResultOutput) ArcInstanceResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.ArcInstanceResourceGroup }).(pulumi.StringOutput)
}

func (o LookupArcSettingResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupArcSettingResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupArcSettingResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArcSettingResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
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

func (o LookupArcSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArcSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArcSettingResultOutput{})
}
