


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Dashboard struct {
	pulumi.CustomResourceState

	Lenses   DashboardLensResponseMapOutput `pulumi:"lenses"`
	Location pulumi.StringOutput            `pulumi:"location"`
	Metadata pulumi.MapOutput               `pulumi:"metadata"`
	Name     pulumi.StringOutput            `pulumi:"name"`
	Tags     pulumi.StringMapOutput         `pulumi:"tags"`
	Type     pulumi.StringOutput            `pulumi:"type"`
}


func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:portal:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20181001preview:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20190101preview:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20200901preview:Dashboard"),
		},
	})
	opts = append(opts, aliases)
	var resource Dashboard
	err := ctx.RegisterResource("azure-native:portal/v20150801preview:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("azure-native:portal/v20150801preview:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dashboardState struct {
}

type DashboardState struct {
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	DashboardName     *string                  `pulumi:"dashboardName"`
	Lenses            map[string]DashboardLens `pulumi:"lenses"`
	Location          *string                  `pulumi:"location"`
	Metadata          map[string]interface{}   `pulumi:"metadata"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Tags              map[string]string        `pulumi:"tags"`
}


type DashboardArgs struct {
	DashboardName     pulumi.StringPtrInput
	Lenses            DashboardLensMapInput
	Location          pulumi.StringPtrInput
	Metadata          pulumi.MapInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

func (o DashboardOutput) Lenses() DashboardLensResponseMapOutput {
	return o.ApplyT(func(v *Dashboard) DashboardLensResponseMapOutput { return v.Lenses }).(DashboardLensResponseMapOutput)
}

func (o DashboardOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DashboardOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

func (o DashboardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DashboardOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DashboardOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DashboardOutput{})
}
