


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TransparentDataEncryption struct {
	pulumi.CustomResourceState

	Name  pulumi.StringOutput `pulumi:"name"`
	State pulumi.StringOutput `pulumi:"state"`
	Type  pulumi.StringOutput `pulumi:"type"`
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
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:TransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:TransparentDataEncryption"),
		},
	})
	opts = append(opts, aliases)
	var resource TransparentDataEncryption
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:TransparentDataEncryption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTransparentDataEncryption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransparentDataEncryptionState, opts ...pulumi.ResourceOption) (*TransparentDataEncryption, error) {
	var resource TransparentDataEncryption
	err := ctx.ReadResource("azure-native:sql/v20200801preview:TransparentDataEncryption", name, id, state, &resource, opts...)
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
	DatabaseName      string                             `pulumi:"databaseName"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
	ServerName        string                             `pulumi:"serverName"`
	State             TransparentDataEncryptionStateEnum `pulumi:"state"`
	TdeName           *string                            `pulumi:"tdeName"`
}


type TransparentDataEncryptionArgs struct {
	DatabaseName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	State             TransparentDataEncryptionStateEnumInput
	TdeName           pulumi.StringPtrInput
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
	return reflect.TypeOf((*TransparentDataEncryption)(nil))
}

func (i *TransparentDataEncryption) ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput {
	return i.ToTransparentDataEncryptionOutputWithContext(context.Background())
}

func (i *TransparentDataEncryption) ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransparentDataEncryptionOutput)
}

type TransparentDataEncryptionOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransparentDataEncryption)(nil))
}

func (o TransparentDataEncryptionOutput) ToTransparentDataEncryptionOutput() TransparentDataEncryptionOutput {
	return o
}

func (o TransparentDataEncryptionOutput) ToTransparentDataEncryptionOutputWithContext(ctx context.Context) TransparentDataEncryptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransparentDataEncryptionOutput{})
}
