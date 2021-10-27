


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerFarm struct {
	pulumi.CustomResourceState

	AdminSiteName             pulumi.StringPtrOutput                     `pulumi:"adminSiteName"`
	GeoRegion                 pulumi.StringOutput                        `pulumi:"geoRegion"`
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	Kind                      pulumi.StringPtrOutput                     `pulumi:"kind"`
	Location                  pulumi.StringOutput                        `pulumi:"location"`
	MaximumNumberOfWorkers    pulumi.IntPtrOutput                        `pulumi:"maximumNumberOfWorkers"`
	Name                      pulumi.StringPtrOutput                     `pulumi:"name"`
	NumberOfSites             pulumi.IntOutput                           `pulumi:"numberOfSites"`
	PerSiteScaling            pulumi.BoolPtrOutput                       `pulumi:"perSiteScaling"`
	Reserved                  pulumi.BoolPtrOutput                       `pulumi:"reserved"`
	ResourceGroup             pulumi.StringOutput                        `pulumi:"resourceGroup"`
	Sku                       SkuDescriptionResponsePtrOutput            `pulumi:"sku"`
	Status                    pulumi.StringOutput                        `pulumi:"status"`
	Subscription              pulumi.StringOutput                        `pulumi:"subscription"`
	Tags                      pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                      pulumi.StringPtrOutput                     `pulumi:"type"`
	WorkerTierName            pulumi.StringPtrOutput                     `pulumi:"workerTierName"`
}


func NewServerFarm(ctx *pulumi.Context,
	name string, args *ServerFarmArgs, opts ...pulumi.ResourceOption) (*ServerFarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160901:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:ServerFarm"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:ServerFarm"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerFarm
	err := ctx.RegisterResource("azure-native:web/v20150801:ServerFarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerFarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerFarmState, opts ...pulumi.ResourceOption) (*ServerFarm, error) {
	var resource ServerFarm
	err := ctx.ReadResource("azure-native:web/v20150801:ServerFarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverFarmState struct {
}

type ServerFarmState struct {
}

func (ServerFarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverFarmState)(nil)).Elem()
}

type serverFarmArgs struct {
	AdminSiteName             *string                    `pulumi:"adminSiteName"`
	AllowPendingState         *bool                      `pulumi:"allowPendingState"`
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	Id                        *string                    `pulumi:"id"`
	Kind                      *string                    `pulumi:"kind"`
	Location                  *string                    `pulumi:"location"`
	MaximumNumberOfWorkers    *int                       `pulumi:"maximumNumberOfWorkers"`
	Name                      *string                    `pulumi:"name"`
	PerSiteScaling            *bool                      `pulumi:"perSiteScaling"`
	Reserved                  *bool                      `pulumi:"reserved"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	Sku                       *SkuDescription            `pulumi:"sku"`
	Tags                      map[string]string          `pulumi:"tags"`
	Type                      *string                    `pulumi:"type"`
	WorkerTierName            *string                    `pulumi:"workerTierName"`
}


type ServerFarmArgs struct {
	AdminSiteName             pulumi.StringPtrInput
	AllowPendingState         pulumi.BoolPtrInput
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	Id                        pulumi.StringPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	MaximumNumberOfWorkers    pulumi.IntPtrInput
	Name                      pulumi.StringPtrInput
	PerSiteScaling            pulumi.BoolPtrInput
	Reserved                  pulumi.BoolPtrInput
	ResourceGroupName         pulumi.StringInput
	Sku                       SkuDescriptionPtrInput
	Tags                      pulumi.StringMapInput
	Type                      pulumi.StringPtrInput
	WorkerTierName            pulumi.StringPtrInput
}

func (ServerFarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverFarmArgs)(nil)).Elem()
}

type ServerFarmInput interface {
	pulumi.Input

	ToServerFarmOutput() ServerFarmOutput
	ToServerFarmOutputWithContext(ctx context.Context) ServerFarmOutput
}

func (*ServerFarm) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerFarm)(nil))
}

func (i *ServerFarm) ToServerFarmOutput() ServerFarmOutput {
	return i.ToServerFarmOutputWithContext(context.Background())
}

func (i *ServerFarm) ToServerFarmOutputWithContext(ctx context.Context) ServerFarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerFarmOutput)
}

type ServerFarmOutput struct{ *pulumi.OutputState }

func (ServerFarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerFarm)(nil))
}

func (o ServerFarmOutput) ToServerFarmOutput() ServerFarmOutput {
	return o
}

func (o ServerFarmOutput) ToServerFarmOutputWithContext(ctx context.Context) ServerFarmOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerFarmInput)(nil)).Elem(), &ServerFarm{})
	pulumi.RegisterOutputType(ServerFarmOutput{})
}
