


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type App struct {
	pulumi.CustomResourceState

	Identity   ManagedIdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties AppResourcePropertiesResponseOutput        `pulumi:"properties"`
	SystemData SystemDataResponseOutput                   `pulumi:"systemData"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToAppResourcePropertiesPtrOutput().ApplyT(func(v *AppResourceProperties) *AppResourceProperties { return v.Defaults() }).(AppResourcePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:App"),
		},
	})
	opts = append(opts, aliases)
	var resource App
	err := ctx.RegisterResource("azure-native:appplatform/v20220101preview:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("azure-native:appplatform/v20220101preview:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appState struct {
}

type AppState struct {
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	AppName           *string                    `pulumi:"appName"`
	Identity          *ManagedIdentityProperties `pulumi:"identity"`
	Location          *string                    `pulumi:"location"`
	Properties        *AppResourceProperties     `pulumi:"properties"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	ServiceName       string                     `pulumi:"serviceName"`
}


type AppArgs struct {
	AppName           pulumi.StringPtrInput
	Identity          ManagedIdentityPropertiesPtrInput
	Location          pulumi.StringPtrInput
	Properties        AppResourcePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func (o AppOutput) Identity() ManagedIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *App) ManagedIdentityPropertiesResponsePtrOutput { return v.Identity }).(ManagedIdentityPropertiesResponsePtrOutput)
}

func (o AppOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppOutput) Properties() AppResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *App) AppResourcePropertiesResponseOutput { return v.Properties }).(AppResourcePropertiesResponseOutput)
}

func (o AppOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *App) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}
