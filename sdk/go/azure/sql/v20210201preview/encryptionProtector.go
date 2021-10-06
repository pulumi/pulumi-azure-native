


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EncryptionProtector struct {
	pulumi.CustomResourceState

	AutoRotationEnabled pulumi.BoolPtrOutput   `pulumi:"autoRotationEnabled"`
	Kind                pulumi.StringOutput    `pulumi:"kind"`
	Location            pulumi.StringOutput    `pulumi:"location"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	ServerKeyName       pulumi.StringPtrOutput `pulumi:"serverKeyName"`
	ServerKeyType       pulumi.StringOutput    `pulumi:"serverKeyType"`
	Subregion           pulumi.StringOutput    `pulumi:"subregion"`
	Thumbprint          pulumi.StringOutput    `pulumi:"thumbprint"`
	Type                pulumi.StringOutput    `pulumi:"type"`
	Uri                 pulumi.StringOutput    `pulumi:"uri"`
}


func NewEncryptionProtector(ctx *pulumi.Context,
	name string, args *EncryptionProtectorArgs, opts ...pulumi.ResourceOption) (*EncryptionProtector, error) {
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
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20150501preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:EncryptionProtector"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:EncryptionProtector"),
		},
	})
	opts = append(opts, aliases)
	var resource EncryptionProtector
	err := ctx.RegisterResource("azure-native:sql/v20210201preview:EncryptionProtector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEncryptionProtector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptionProtectorState, opts ...pulumi.ResourceOption) (*EncryptionProtector, error) {
	var resource EncryptionProtector
	err := ctx.ReadResource("azure-native:sql/v20210201preview:EncryptionProtector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type encryptionProtectorState struct {
}

type EncryptionProtectorState struct {
}

func (EncryptionProtectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionProtectorState)(nil)).Elem()
}

type encryptionProtectorArgs struct {
	AutoRotationEnabled     *bool   `pulumi:"autoRotationEnabled"`
	EncryptionProtectorName *string `pulumi:"encryptionProtectorName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ServerKeyName           *string `pulumi:"serverKeyName"`
	ServerKeyType           string  `pulumi:"serverKeyType"`
	ServerName              string  `pulumi:"serverName"`
}


type EncryptionProtectorArgs struct {
	AutoRotationEnabled     pulumi.BoolPtrInput
	EncryptionProtectorName pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ServerKeyName           pulumi.StringPtrInput
	ServerKeyType           pulumi.StringInput
	ServerName              pulumi.StringInput
}

func (EncryptionProtectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*encryptionProtectorArgs)(nil)).Elem()
}

type EncryptionProtectorInput interface {
	pulumi.Input

	ToEncryptionProtectorOutput() EncryptionProtectorOutput
	ToEncryptionProtectorOutputWithContext(ctx context.Context) EncryptionProtectorOutput
}

func (*EncryptionProtector) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProtector)(nil))
}

func (i *EncryptionProtector) ToEncryptionProtectorOutput() EncryptionProtectorOutput {
	return i.ToEncryptionProtectorOutputWithContext(context.Background())
}

func (i *EncryptionProtector) ToEncryptionProtectorOutputWithContext(ctx context.Context) EncryptionProtectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionProtectorOutput)
}

type EncryptionProtectorOutput struct{ *pulumi.OutputState }

func (EncryptionProtectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProtector)(nil))
}

func (o EncryptionProtectorOutput) ToEncryptionProtectorOutput() EncryptionProtectorOutput {
	return o
}

func (o EncryptionProtectorOutput) ToEncryptionProtectorOutputWithContext(ctx context.Context) EncryptionProtectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EncryptionProtectorOutput{})
}
