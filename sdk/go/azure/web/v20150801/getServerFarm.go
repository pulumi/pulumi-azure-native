


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupServerFarm(ctx *pulumi.Context, args *LookupServerFarmArgs, opts ...pulumi.InvokeOption) (*LookupServerFarmResult, error) {
	var rv LookupServerFarmResult
	err := ctx.Invoke("azure-native:web/v20150801:getServerFarm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerFarmArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupServerFarmResult struct {
	AdminSiteName             *string                            `pulumi:"adminSiteName"`
	GeoRegion                 string                             `pulumi:"geoRegion"`
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	Id                        *string                            `pulumi:"id"`
	Kind                      *string                            `pulumi:"kind"`
	Location                  string                             `pulumi:"location"`
	MaximumNumberOfWorkers    *int                               `pulumi:"maximumNumberOfWorkers"`
	Name                      *string                            `pulumi:"name"`
	NumberOfSites             int                                `pulumi:"numberOfSites"`
	PerSiteScaling            *bool                              `pulumi:"perSiteScaling"`
	Reserved                  *bool                              `pulumi:"reserved"`
	ResourceGroup             string                             `pulumi:"resourceGroup"`
	Sku                       *SkuDescriptionResponse            `pulumi:"sku"`
	Status                    string                             `pulumi:"status"`
	Subscription              string                             `pulumi:"subscription"`
	Tags                      map[string]string                  `pulumi:"tags"`
	Type                      *string                            `pulumi:"type"`
	WorkerTierName            *string                            `pulumi:"workerTierName"`
}

func LookupServerFarmOutput(ctx *pulumi.Context, args LookupServerFarmOutputArgs, opts ...pulumi.InvokeOption) LookupServerFarmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerFarmResult, error) {
			args := v.(LookupServerFarmArgs)
			r, err := LookupServerFarm(ctx, &args, opts...)
			var s LookupServerFarmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerFarmResultOutput)
}

type LookupServerFarmOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServerFarmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerFarmArgs)(nil)).Elem()
}


type LookupServerFarmResultOutput struct{ *pulumi.OutputState }

func (LookupServerFarmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerFarmResult)(nil)).Elem()
}

func (o LookupServerFarmResultOutput) ToLookupServerFarmResultOutput() LookupServerFarmResultOutput {
	return o
}

func (o LookupServerFarmResultOutput) ToLookupServerFarmResultOutputWithContext(ctx context.Context) LookupServerFarmResultOutput {
	return o
}

func (o LookupServerFarmResultOutput) AdminSiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *string { return v.AdminSiteName }).(pulumi.StringPtrOutput)
}

func (o LookupServerFarmResultOutput) GeoRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerFarmResult) string { return v.GeoRegion }).(pulumi.StringOutput)
}

func (o LookupServerFarmResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o LookupServerFarmResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServerFarmResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupServerFarmResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerFarmResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServerFarmResultOutput) MaximumNumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *int { return v.MaximumNumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupServerFarmResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupServerFarmResultOutput) NumberOfSites() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServerFarmResult) int { return v.NumberOfSites }).(pulumi.IntOutput)
}

func (o LookupServerFarmResultOutput) PerSiteScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *bool { return v.PerSiteScaling }).(pulumi.BoolPtrOutput)
}

func (o LookupServerFarmResultOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *bool { return v.Reserved }).(pulumi.BoolPtrOutput)
}

func (o LookupServerFarmResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerFarmResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o LookupServerFarmResultOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

func (o LookupServerFarmResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerFarmResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupServerFarmResultOutput) Subscription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerFarmResult) string { return v.Subscription }).(pulumi.StringOutput)
}

func (o LookupServerFarmResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerFarmResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServerFarmResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupServerFarmResultOutput) WorkerTierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerFarmResult) *string { return v.WorkerTierName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerFarmResultOutput{})
}
