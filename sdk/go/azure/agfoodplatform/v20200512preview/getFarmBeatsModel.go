


package v20200512preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFarmBeatsModel(ctx *pulumi.Context, args *LookupFarmBeatsModelArgs, opts ...pulumi.InvokeOption) (*LookupFarmBeatsModelResult, error) {
	var rv LookupFarmBeatsModelResult
	err := ctx.Invoke("azure-native:agfoodplatform/v20200512preview:getFarmBeatsModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFarmBeatsModelArgs struct {
	FarmBeatsResourceName string `pulumi:"farmBeatsResourceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupFarmBeatsModelResult struct {
	Id                string             `pulumi:"id"`
	InstanceUri       string             `pulumi:"instanceUri"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	Sku               *SkuResponse       `pulumi:"sku"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupFarmBeatsModelOutput(ctx *pulumi.Context, args LookupFarmBeatsModelOutputArgs, opts ...pulumi.InvokeOption) LookupFarmBeatsModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFarmBeatsModelResult, error) {
			args := v.(LookupFarmBeatsModelArgs)
			r, err := LookupFarmBeatsModel(ctx, &args, opts...)
			var s LookupFarmBeatsModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFarmBeatsModelResultOutput)
}

type LookupFarmBeatsModelOutputArgs struct {
	FarmBeatsResourceName pulumi.StringInput `pulumi:"farmBeatsResourceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFarmBeatsModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFarmBeatsModelArgs)(nil)).Elem()
}


type LookupFarmBeatsModelResultOutput struct{ *pulumi.OutputState }

func (LookupFarmBeatsModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFarmBeatsModelResult)(nil)).Elem()
}

func (o LookupFarmBeatsModelResultOutput) ToLookupFarmBeatsModelResultOutput() LookupFarmBeatsModelResultOutput {
	return o
}

func (o LookupFarmBeatsModelResultOutput) ToLookupFarmBeatsModelResultOutputWithContext(ctx context.Context) LookupFarmBeatsModelResultOutput {
	return o
}

func (o LookupFarmBeatsModelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFarmBeatsModelResultOutput) InstanceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.InstanceUri }).(pulumi.StringOutput)
}

func (o LookupFarmBeatsModelResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupFarmBeatsModelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFarmBeatsModelResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFarmBeatsModelResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupFarmBeatsModelResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFarmBeatsModelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFarmBeatsModelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFarmBeatsModelResultOutput{})
}
