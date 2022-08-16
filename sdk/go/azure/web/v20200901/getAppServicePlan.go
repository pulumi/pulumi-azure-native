


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServicePlan(ctx *pulumi.Context, args *LookupAppServicePlanArgs, opts ...pulumi.InvokeOption) (*LookupAppServicePlanResult, error) {
	var rv LookupAppServicePlanResult
	err := ctx.Invoke("azure-native:web/v20200901:getAppServicePlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppServicePlanArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAppServicePlanResult struct {
	FreeOfferExpirationTime   *string                            `pulumi:"freeOfferExpirationTime"`
	GeoRegion                 string                             `pulumi:"geoRegion"`
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	HyperV                    *bool                              `pulumi:"hyperV"`
	Id                        string                             `pulumi:"id"`
	IsSpot                    *bool                              `pulumi:"isSpot"`
	IsXenon                   *bool                              `pulumi:"isXenon"`
	Kind                      *string                            `pulumi:"kind"`
	Location                  string                             `pulumi:"location"`
	MaximumElasticWorkerCount *int                               `pulumi:"maximumElasticWorkerCount"`
	MaximumNumberOfWorkers    int                                `pulumi:"maximumNumberOfWorkers"`
	Name                      string                             `pulumi:"name"`
	NumberOfSites             int                                `pulumi:"numberOfSites"`
	PerSiteScaling            *bool                              `pulumi:"perSiteScaling"`
	ProvisioningState         string                             `pulumi:"provisioningState"`
	Reserved                  *bool                              `pulumi:"reserved"`
	ResourceGroup             string                             `pulumi:"resourceGroup"`
	Sku                       *SkuDescriptionResponse            `pulumi:"sku"`
	SpotExpirationTime        *string                            `pulumi:"spotExpirationTime"`
	Status                    string                             `pulumi:"status"`
	Subscription              string                             `pulumi:"subscription"`
	SystemData                SystemDataResponse                 `pulumi:"systemData"`
	Tags                      map[string]string                  `pulumi:"tags"`
	TargetWorkerCount         *int                               `pulumi:"targetWorkerCount"`
	TargetWorkerSizeId        *int                               `pulumi:"targetWorkerSizeId"`
	Type                      string                             `pulumi:"type"`
	WorkerTierName            *string                            `pulumi:"workerTierName"`
}


func (val *LookupAppServicePlanResult) Defaults() *LookupAppServicePlanResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HyperV) {
		hyperV_ := false
		tmp.HyperV = &hyperV_
	}
	if isZero(tmp.IsXenon) {
		isXenon_ := false
		tmp.IsXenon = &isXenon_
	}
	if isZero(tmp.PerSiteScaling) {
		perSiteScaling_ := false
		tmp.PerSiteScaling = &perSiteScaling_
	}
	if isZero(tmp.Reserved) {
		reserved_ := false
		tmp.Reserved = &reserved_
	}
	return &tmp
}

func LookupAppServicePlanOutput(ctx *pulumi.Context, args LookupAppServicePlanOutputArgs, opts ...pulumi.InvokeOption) LookupAppServicePlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppServicePlanResult, error) {
			args := v.(LookupAppServicePlanArgs)
			r, err := LookupAppServicePlan(ctx, &args, opts...)
			var s LookupAppServicePlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppServicePlanResultOutput)
}

type LookupAppServicePlanOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppServicePlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServicePlanArgs)(nil)).Elem()
}


type LookupAppServicePlanResultOutput struct{ *pulumi.OutputState }

func (LookupAppServicePlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppServicePlanResult)(nil)).Elem()
}

func (o LookupAppServicePlanResultOutput) ToLookupAppServicePlanResultOutput() LookupAppServicePlanResultOutput {
	return o
}

func (o LookupAppServicePlanResultOutput) ToLookupAppServicePlanResultOutputWithContext(ctx context.Context) LookupAppServicePlanResultOutput {
	return o
}

func (o LookupAppServicePlanResultOutput) FreeOfferExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.FreeOfferExpirationTime }).(pulumi.StringPtrOutput)
}

func (o LookupAppServicePlanResultOutput) GeoRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.GeoRegion }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *HostingEnvironmentProfileResponse {
		return v.HostingEnvironmentProfile
	}).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o LookupAppServicePlanResultOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.HyperV }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServicePlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) IsSpot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.IsSpot }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServicePlanResultOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServicePlanResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAppServicePlanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) MaximumElasticWorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *int { return v.MaximumElasticWorkerCount }).(pulumi.IntPtrOutput)
}

func (o LookupAppServicePlanResultOutput) MaximumNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) int { return v.MaximumNumberOfWorkers }).(pulumi.IntOutput)
}

func (o LookupAppServicePlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) NumberOfSites() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) int { return v.NumberOfSites }).(pulumi.IntOutput)
}

func (o LookupAppServicePlanResultOutput) PerSiteScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.PerSiteScaling }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServicePlanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *bool { return v.Reserved }).(pulumi.BoolPtrOutput)
}

func (o LookupAppServicePlanResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

func (o LookupAppServicePlanResultOutput) SpotExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.SpotExpirationTime }).(pulumi.StringPtrOutput)
}

func (o LookupAppServicePlanResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) Subscription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Subscription }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAppServicePlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAppServicePlanResultOutput) TargetWorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *int { return v.TargetWorkerCount }).(pulumi.IntPtrOutput)
}

func (o LookupAppServicePlanResultOutput) TargetWorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *int { return v.TargetWorkerSizeId }).(pulumi.IntPtrOutput)
}

func (o LookupAppServicePlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAppServicePlanResultOutput) WorkerTierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppServicePlanResult) *string { return v.WorkerTierName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppServicePlanResultOutput{})
}
