


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServicePlan struct {
	pulumi.CustomResourceState

	FreeOfferExpirationTime   pulumi.StringPtrOutput                     `pulumi:"freeOfferExpirationTime"`
	GeoRegion                 pulumi.StringOutput                        `pulumi:"geoRegion"`
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	HyperV                    pulumi.BoolPtrOutput                       `pulumi:"hyperV"`
	IsSpot                    pulumi.BoolPtrOutput                       `pulumi:"isSpot"`
	IsXenon                   pulumi.BoolPtrOutput                       `pulumi:"isXenon"`
	Kind                      pulumi.StringPtrOutput                     `pulumi:"kind"`
	Location                  pulumi.StringOutput                        `pulumi:"location"`
	MaximumElasticWorkerCount pulumi.IntPtrOutput                        `pulumi:"maximumElasticWorkerCount"`
	MaximumNumberOfWorkers    pulumi.IntOutput                           `pulumi:"maximumNumberOfWorkers"`
	Name                      pulumi.StringOutput                        `pulumi:"name"`
	NumberOfSites             pulumi.IntOutput                           `pulumi:"numberOfSites"`
	PerSiteScaling            pulumi.BoolPtrOutput                       `pulumi:"perSiteScaling"`
	ProvisioningState         pulumi.StringOutput                        `pulumi:"provisioningState"`
	Reserved                  pulumi.BoolPtrOutput                       `pulumi:"reserved"`
	ResourceGroup             pulumi.StringOutput                        `pulumi:"resourceGroup"`
	Sku                       SkuDescriptionResponsePtrOutput            `pulumi:"sku"`
	SpotExpirationTime        pulumi.StringPtrOutput                     `pulumi:"spotExpirationTime"`
	Status                    pulumi.StringOutput                        `pulumi:"status"`
	Subscription              pulumi.StringOutput                        `pulumi:"subscription"`
	Tags                      pulumi.StringMapOutput                     `pulumi:"tags"`
	TargetWorkerCount         pulumi.IntPtrOutput                        `pulumi:"targetWorkerCount"`
	TargetWorkerSizeId        pulumi.IntPtrOutput                        `pulumi:"targetWorkerSizeId"`
	Type                      pulumi.StringOutput                        `pulumi:"type"`
	WorkerTierName            pulumi.StringPtrOutput                     `pulumi:"workerTierName"`
}


func NewAppServicePlan(ctx *pulumi.Context,
	name string, args *AppServicePlanArgs, opts ...pulumi.ResourceOption) (*AppServicePlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.HyperV) {
		args.HyperV = pulumi.BoolPtr(false)
	}
	if isZero(args.IsXenon) {
		args.IsXenon = pulumi.BoolPtr(false)
	}
	if isZero(args.PerSiteScaling) {
		args.PerSiteScaling = pulumi.BoolPtr(false)
	}
	if isZero(args.Reserved) {
		args.Reserved = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:AppServicePlan"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServicePlan
	err := ctx.RegisterResource("azure-native:web/v20190801:AppServicePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServicePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServicePlanState, opts ...pulumi.ResourceOption) (*AppServicePlan, error) {
	var resource AppServicePlan
	err := ctx.ReadResource("azure-native:web/v20190801:AppServicePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServicePlanState struct {
}

type AppServicePlanState struct {
}

func (AppServicePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanState)(nil)).Elem()
}

type appServicePlanArgs struct {
	FreeOfferExpirationTime   *string                    `pulumi:"freeOfferExpirationTime"`
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	HyperV                    *bool                      `pulumi:"hyperV"`
	IsSpot                    *bool                      `pulumi:"isSpot"`
	IsXenon                   *bool                      `pulumi:"isXenon"`
	Kind                      *string                    `pulumi:"kind"`
	Location                  *string                    `pulumi:"location"`
	MaximumElasticWorkerCount *int                       `pulumi:"maximumElasticWorkerCount"`
	Name                      *string                    `pulumi:"name"`
	PerSiteScaling            *bool                      `pulumi:"perSiteScaling"`
	Reserved                  *bool                      `pulumi:"reserved"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	Sku                       *SkuDescription            `pulumi:"sku"`
	SpotExpirationTime        *string                    `pulumi:"spotExpirationTime"`
	Tags                      map[string]string          `pulumi:"tags"`
	TargetWorkerCount         *int                       `pulumi:"targetWorkerCount"`
	TargetWorkerSizeId        *int                       `pulumi:"targetWorkerSizeId"`
	WorkerTierName            *string                    `pulumi:"workerTierName"`
}


type AppServicePlanArgs struct {
	FreeOfferExpirationTime   pulumi.StringPtrInput
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	HyperV                    pulumi.BoolPtrInput
	IsSpot                    pulumi.BoolPtrInput
	IsXenon                   pulumi.BoolPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	MaximumElasticWorkerCount pulumi.IntPtrInput
	Name                      pulumi.StringPtrInput
	PerSiteScaling            pulumi.BoolPtrInput
	Reserved                  pulumi.BoolPtrInput
	ResourceGroupName         pulumi.StringInput
	Sku                       SkuDescriptionPtrInput
	SpotExpirationTime        pulumi.StringPtrInput
	Tags                      pulumi.StringMapInput
	TargetWorkerCount         pulumi.IntPtrInput
	TargetWorkerSizeId        pulumi.IntPtrInput
	WorkerTierName            pulumi.StringPtrInput
}

func (AppServicePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanArgs)(nil)).Elem()
}

type AppServicePlanInput interface {
	pulumi.Input

	ToAppServicePlanOutput() AppServicePlanOutput
	ToAppServicePlanOutputWithContext(ctx context.Context) AppServicePlanOutput
}

func (*AppServicePlan) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServicePlan)(nil)).Elem()
}

func (i *AppServicePlan) ToAppServicePlanOutput() AppServicePlanOutput {
	return i.ToAppServicePlanOutputWithContext(context.Background())
}

func (i *AppServicePlan) ToAppServicePlanOutputWithContext(ctx context.Context) AppServicePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServicePlanOutput)
}

type AppServicePlanOutput struct{ *pulumi.OutputState }

func (AppServicePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServicePlan)(nil)).Elem()
}

func (o AppServicePlanOutput) ToAppServicePlanOutput() AppServicePlanOutput {
	return o
}

func (o AppServicePlanOutput) ToAppServicePlanOutputWithContext(ctx context.Context) AppServicePlanOutput {
	return o
}

func (o AppServicePlanOutput) FreeOfferExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringPtrOutput { return v.FreeOfferExpirationTime }).(pulumi.StringPtrOutput)
}

func (o AppServicePlanOutput) GeoRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.GeoRegion }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v *AppServicePlan) HostingEnvironmentProfileResponsePtrOutput { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o AppServicePlanOutput) HyperV() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.BoolPtrOutput { return v.HyperV }).(pulumi.BoolPtrOutput)
}

func (o AppServicePlanOutput) IsSpot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.BoolPtrOutput { return v.IsSpot }).(pulumi.BoolPtrOutput)
}

func (o AppServicePlanOutput) IsXenon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.BoolPtrOutput { return v.IsXenon }).(pulumi.BoolPtrOutput)
}

func (o AppServicePlanOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AppServicePlanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) MaximumElasticWorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.IntPtrOutput { return v.MaximumElasticWorkerCount }).(pulumi.IntPtrOutput)
}

func (o AppServicePlanOutput) MaximumNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.IntOutput { return v.MaximumNumberOfWorkers }).(pulumi.IntOutput)
}

func (o AppServicePlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) NumberOfSites() pulumi.IntOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.IntOutput { return v.NumberOfSites }).(pulumi.IntOutput)
}

func (o AppServicePlanOutput) PerSiteScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.BoolPtrOutput { return v.PerSiteScaling }).(pulumi.BoolPtrOutput)
}

func (o AppServicePlanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.BoolPtrOutput { return v.Reserved }).(pulumi.BoolPtrOutput)
}

func (o AppServicePlanOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *AppServicePlan) SkuDescriptionResponsePtrOutput { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

func (o AppServicePlanOutput) SpotExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringPtrOutput { return v.SpotExpirationTime }).(pulumi.StringPtrOutput)
}

func (o AppServicePlanOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) Subscription() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.Subscription }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AppServicePlanOutput) TargetWorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.IntPtrOutput { return v.TargetWorkerCount }).(pulumi.IntPtrOutput)
}

func (o AppServicePlanOutput) TargetWorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.IntPtrOutput { return v.TargetWorkerSizeId }).(pulumi.IntPtrOutput)
}

func (o AppServicePlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AppServicePlanOutput) WorkerTierName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlan) pulumi.StringPtrOutput { return v.WorkerTierName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServicePlanOutput{})
}
