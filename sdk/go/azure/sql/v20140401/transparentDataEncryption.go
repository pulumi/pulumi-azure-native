


package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TransparentDataEncryption struct {
	pulumi.CustomResourceState

	Location pulumi.StringOutput    `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Status   pulumi.StringPtrOutput `pulumi:"status"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewTransparentDataEncryption(ctx *pulumi.Context,
	name string, args *TransparentDataEncryptionArgs, opts ...pulumi.ResourceOption) (*TransparentDataEncryption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:TransparentDataEncryption"),
		},
	})
	opts = append(opts, aliases)
	var resource TransparentDataEncryption
	err := ctx.RegisterResource("azure-native:sql/v20140401:TransparentDataEncryption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTransparentDataEncryption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransparentDataEncryptionState, opts ...pulumi.ResourceOption) (*TransparentDataEncryption, error) {
	var resource TransparentDataEncryption
	err := ctx.ReadResource("azure-native:sql/v20140401:TransparentDataEncryption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type transparentDataEncryptionState struct {
}

type TransparentDataEncryptionState struct {
}

func (TransparentDataEncryptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*transparentDataEncryptionState)(nil)).Elem()
}

type transparentDataEncryptionArgs struct {
	DatabaseName                  string  `pulumi:"databaseName"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
	ServerName                    string  `pulumi:"serverName"`
	Status                        *string `pulumi:"status"`
	TransparentDataEncryptionName *string `pulumi:"transparentDataEncryptionName"`
}


type TransparentDataEncryptionArgs struct {
	DatabaseName                  pulumi.StringInput
	ResourceGroupName             pulumi.StringInput
	ServerName                    pulumi.StringInput
	Status                        pulumi.StringPtrInput
	TransparentDataEncryptionName pulumi.StringPtrInput
}

func (TransparentDataEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transparentDataEncryptionArgs)(nil)).Elem()
}

type TransparentDataEncryptionInput interface {
	pulumi.Input

	ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput
	ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput
}

func (*TransparentDataEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((**TransparentDataEncryption)(nil)).Elem()
}

func (i *TransparentDataEncryption) ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput {
	return i.ToTransparentDataEncryptionOutputWithContext(context.Background())
}

func (i *TransparentDataEncryption) ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransparentDataEncryptionOutput)
}

type TransparentDataEncryptionOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransparentDataEncryption)(nil)).Elem()
}

func (o TransparentDataEncryptionOutput) ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput {
	return o
}

func (o TransparentDataEncryptionOutput) ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput {
	return o
}

func (o TransparentDataEncryptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TransparentDataEncryption) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o TransparentDataEncryptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TransparentDataEncryption) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TransparentDataEncryptionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransparentDataEncryption) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TransparentDataEncryptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TransparentDataEncryption) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TransparentDataEncryptionOutput{})
}
