


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServicePlanRouteForVnet struct {
	pulumi.CustomResourceState

	EndAddress   pulumi.StringPtrOutput `pulumi:"endAddress"`
	Kind         pulumi.StringPtrOutput `pulumi:"kind"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	RouteType    pulumi.StringPtrOutput `pulumi:"routeType"`
	StartAddress pulumi.StringPtrOutput `pulumi:"startAddress"`
	Type         pulumi.StringOutput    `pulumi:"type"`
}


func NewAppServicePlanRouteForVnet(ctx *pulumi.Context,
	name string, args *AppServicePlanRouteForVnetArgs, opts ...pulumi.ResourceOption) (*AppServicePlanRouteForVnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VnetName == nil {
		return nil, errors.New("invalid value for required argument 'VnetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160901:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:AppServicePlanRouteForVnet"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:AppServicePlanRouteForVnet"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServicePlanRouteForVnet
	err := ctx.RegisterResource("azure-native:web/v20220301:AppServicePlanRouteForVnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServicePlanRouteForVnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServicePlanRouteForVnetState, opts ...pulumi.ResourceOption) (*AppServicePlanRouteForVnet, error) {
	var resource AppServicePlanRouteForVnet
	err := ctx.ReadResource("azure-native:web/v20220301:AppServicePlanRouteForVnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServicePlanRouteForVnetState struct {
}

type AppServicePlanRouteForVnetState struct {
}

func (AppServicePlanRouteForVnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanRouteForVnetState)(nil)).Elem()
}

type appServicePlanRouteForVnetArgs struct {
	EndAddress        *string `pulumi:"endAddress"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RouteName         *string `pulumi:"routeName"`
	RouteType         *string `pulumi:"routeType"`
	StartAddress      *string `pulumi:"startAddress"`
	VnetName          string  `pulumi:"vnetName"`
}


type AppServicePlanRouteForVnetArgs struct {
	EndAddress        pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RouteName         pulumi.StringPtrInput
	RouteType         pulumi.StringPtrInput
	StartAddress      pulumi.StringPtrInput
	VnetName          pulumi.StringInput
}

func (AppServicePlanRouteForVnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanRouteForVnetArgs)(nil)).Elem()
}

type AppServicePlanRouteForVnetInput interface {
	pulumi.Input

	ToAppServicePlanRouteForVnetOutput() AppServicePlanRouteForVnetOutput
	ToAppServicePlanRouteForVnetOutputWithContext(ctx context.Context) AppServicePlanRouteForVnetOutput
}

func (*AppServicePlanRouteForVnet) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServicePlanRouteForVnet)(nil)).Elem()
}

func (i *AppServicePlanRouteForVnet) ToAppServicePlanRouteForVnetOutput() AppServicePlanRouteForVnetOutput {
	return i.ToAppServicePlanRouteForVnetOutputWithContext(context.Background())
}

func (i *AppServicePlanRouteForVnet) ToAppServicePlanRouteForVnetOutputWithContext(ctx context.Context) AppServicePlanRouteForVnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServicePlanRouteForVnetOutput)
}

type AppServicePlanRouteForVnetOutput struct{ *pulumi.OutputState }

func (AppServicePlanRouteForVnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServicePlanRouteForVnet)(nil)).Elem()
}

func (o AppServicePlanRouteForVnetOutput) ToAppServicePlanRouteForVnetOutput() AppServicePlanRouteForVnetOutput {
	return o
}

func (o AppServicePlanRouteForVnetOutput) ToAppServicePlanRouteForVnetOutputWithContext(ctx context.Context) AppServicePlanRouteForVnetOutput {
	return o
}

func (o AppServicePlanRouteForVnetOutput) EndAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlanRouteForVnet) pulumi.StringPtrOutput { return v.EndAddress }).(pulumi.StringPtrOutput)
}

func (o AppServicePlanRouteForVnetOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlanRouteForVnet) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AppServicePlanRouteForVnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlanRouteForVnet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppServicePlanRouteForVnetOutput) RouteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlanRouteForVnet) pulumi.StringPtrOutput { return v.RouteType }).(pulumi.StringPtrOutput)
}

func (o AppServicePlanRouteForVnetOutput) StartAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServicePlanRouteForVnet) pulumi.StringPtrOutput { return v.StartAddress }).(pulumi.StringPtrOutput)
}

func (o AppServicePlanRouteForVnetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServicePlanRouteForVnet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServicePlanRouteForVnetOutput{})
}
