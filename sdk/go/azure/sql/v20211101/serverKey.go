


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerKey struct {
	pulumi.CustomResourceState

	AutoRotationEnabled pulumi.BoolOutput   `pulumi:"autoRotationEnabled"`
	CreationDate        pulumi.StringOutput `pulumi:"creationDate"`
	Kind                pulumi.StringOutput `pulumi:"kind"`
	Location            pulumi.StringOutput `pulumi:"location"`
	Name                pulumi.StringOutput `pulumi:"name"`
	Subregion           pulumi.StringOutput `pulumi:"subregion"`
	Thumbprint          pulumi.StringOutput `pulumi:"thumbprint"`
	Type                pulumi.StringOutput `pulumi:"type"`
}


func NewServerKey(ctx *pulumi.Context,
	name string, args *ServerKeyArgs, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerKeyType == nil {
		return nil, errors.New("invalid value for required argument 'ServerKeyType'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerKey"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerKey"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerKey
	err := ctx.RegisterResource("azure-native:sql/v20211101:ServerKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerKeyState, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	var resource ServerKey
	err := ctx.ReadResource("azure-native:sql/v20211101:ServerKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverKeyState struct {
}

type ServerKeyState struct {
}

func (ServerKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyState)(nil)).Elem()
}

type serverKeyArgs struct {
	KeyName           *string `pulumi:"keyName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerKeyType     string  `pulumi:"serverKeyType"`
	ServerName        string  `pulumi:"serverName"`
	Uri               *string `pulumi:"uri"`
}


type ServerKeyArgs struct {
	KeyName           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerKeyType     pulumi.StringInput
	ServerName        pulumi.StringInput
	Uri               pulumi.StringPtrInput
}

func (ServerKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyArgs)(nil)).Elem()
}

type ServerKeyInput interface {
	pulumi.Input

	ToServerKeyOutput() ServerKeyOutput
	ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput
}

func (*ServerKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKey)(nil)).Elem()
}

func (i *ServerKey) ToServerKeyOutput() ServerKeyOutput {
	return i.ToServerKeyOutputWithContext(context.Background())
}

func (i *ServerKey) ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerKeyOutput)
}

type ServerKeyOutput struct{ *pulumi.OutputState }

func (ServerKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKey)(nil)).Elem()
}

func (o ServerKeyOutput) ToServerKeyOutput() ServerKeyOutput {
	return o
}

func (o ServerKeyOutput) ToServerKeyOutputWithContext(ctx context.Context) ServerKeyOutput {
	return o
}

func (o ServerKeyOutput) AutoRotationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.BoolOutput { return v.AutoRotationEnabled }).(pulumi.BoolOutput)
}

func (o ServerKeyOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o ServerKeyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ServerKeyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ServerKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerKeyOutput) Subregion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Subregion }).(pulumi.StringOutput)
}

func (o ServerKeyOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o ServerKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerKeyOutput{})
}
