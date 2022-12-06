


package v20220401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigServer struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties ConfigServerPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput             `pulumi:"systemData"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewConfigServer(ctx *pulumi.Context,
	name string, args *ConfigServerArgs, opts ...pulumi.ResourceOption) (*ConfigServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ConfigServer"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ConfigServer"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigServer
	err := ctx.RegisterResource("azure-native:appplatform/v20220401:ConfigServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigServerState, opts ...pulumi.ResourceOption) (*ConfigServer, error) {
	var resource ConfigServer
	err := ctx.ReadResource("azure-native:appplatform/v20220401:ConfigServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configServerState struct {
}

type ConfigServerState struct {
}

func (ConfigServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*configServerState)(nil)).Elem()
}

type configServerArgs struct {
	Properties        *ConfigServerProperties `pulumi:"properties"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	ServiceName       string                  `pulumi:"serviceName"`
}


type ConfigServerArgs struct {
	Properties        ConfigServerPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ConfigServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configServerArgs)(nil)).Elem()
}

type ConfigServerInput interface {
	pulumi.Input

	ToConfigServerOutput() ConfigServerOutput
	ToConfigServerOutputWithContext(ctx context.Context) ConfigServerOutput
}

func (*ConfigServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServer)(nil)).Elem()
}

func (i *ConfigServer) ToConfigServerOutput() ConfigServerOutput {
	return i.ToConfigServerOutputWithContext(context.Background())
}

func (i *ConfigServer) ToConfigServerOutputWithContext(ctx context.Context) ConfigServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerOutput)
}

type ConfigServerOutput struct{ *pulumi.OutputState }

func (ConfigServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServer)(nil)).Elem()
}

func (o ConfigServerOutput) ToConfigServerOutput() ConfigServerOutput {
	return o
}

func (o ConfigServerOutput) ToConfigServerOutputWithContext(ctx context.Context) ConfigServerOutput {
	return o
}

func (o ConfigServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigServerOutput) Properties() ConfigServerPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigServer) ConfigServerPropertiesResponseOutput { return v.Properties }).(ConfigServerPropertiesResponseOutput)
}

func (o ConfigServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigServer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigServerOutput{})
}
