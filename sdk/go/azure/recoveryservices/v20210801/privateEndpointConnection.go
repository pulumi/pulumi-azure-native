


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput                  `pulumi:"eTag"`
	Location   pulumi.StringPtrOutput                  `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties PrivateEndpointConnectionResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20200202:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210801:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:recoveryservices/v20210801:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionState struct {
}

type PrivateEndpointConnectionState struct {
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	ETag                          *string                        `pulumi:"eTag"`
	Location                      *string                        `pulumi:"location"`
	PrivateEndpointConnectionName *string                        `pulumi:"privateEndpointConnectionName"`
	Properties                    *PrivateEndpointConnectionType `pulumi:"properties"`
	ResourceGroupName             string                         `pulumi:"resourceGroupName"`
	Tags                          map[string]string              `pulumi:"tags"`
	VaultName                     string                         `pulumi:"vaultName"`
}


type PrivateEndpointConnectionArgs struct {
	ETag                          pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	PrivateEndpointConnectionName pulumi.StringPtrInput
	Properties                    PrivateEndpointConnectionTypePtrInput
	ResourceGroupName             pulumi.StringInput
	Tags                          pulumi.StringMapInput
	VaultName                     pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}

type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput
}

func (*PrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionOutput) Properties() PrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointConnectionResponseOutput { return v.Properties }).(PrivateEndpointConnectionResponseOutput)
}

func (o PrivateEndpointConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
