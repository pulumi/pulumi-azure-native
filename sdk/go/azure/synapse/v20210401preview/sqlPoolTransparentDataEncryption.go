


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlPoolTransparentDataEncryption struct {
	pulumi.CustomResourceState

	Location pulumi.StringOutput    `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Status   pulumi.StringPtrOutput `pulumi:"status"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewSqlPoolTransparentDataEncryption(ctx *pulumi.Context,
	name string, args *SqlPoolTransparentDataEncryptionArgs, opts ...pulumi.ResourceOption) (*SqlPoolTransparentDataEncryption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlPoolName == nil {
		return nil, errors.New("invalid value for required argument 'SqlPoolName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPoolTransparentDataEncryption"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPoolTransparentDataEncryption
	err := ctx.RegisterResource("azure-native:synapse/v20210401preview:SqlPoolTransparentDataEncryption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlPoolTransparentDataEncryption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolTransparentDataEncryptionState, opts ...pulumi.ResourceOption) (*SqlPoolTransparentDataEncryption, error) {
	var resource SqlPoolTransparentDataEncryption
	err := ctx.ReadResource("azure-native:synapse/v20210401preview:SqlPoolTransparentDataEncryption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlPoolTransparentDataEncryptionState struct {
}

type SqlPoolTransparentDataEncryptionState struct {
}

func (SqlPoolTransparentDataEncryptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolTransparentDataEncryptionState)(nil)).Elem()
}

type sqlPoolTransparentDataEncryptionArgs struct {
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
	SqlPoolName                   string  `pulumi:"sqlPoolName"`
	Status                        *string `pulumi:"status"`
	TransparentDataEncryptionName *string `pulumi:"transparentDataEncryptionName"`
	WorkspaceName                 string  `pulumi:"workspaceName"`
}


type SqlPoolTransparentDataEncryptionArgs struct {
	ResourceGroupName             pulumi.StringInput
	SqlPoolName                   pulumi.StringInput
	Status                        pulumi.StringPtrInput
	TransparentDataEncryptionName pulumi.StringPtrInput
	WorkspaceName                 pulumi.StringInput
}

func (SqlPoolTransparentDataEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolTransparentDataEncryptionArgs)(nil)).Elem()
}

type SqlPoolTransparentDataEncryptionInput interface {
	pulumi.Input

	ToSqlPoolTransparentDataEncryptionOutput() SqlPoolTransparentDataEncryptionOutput
	ToSqlPoolTransparentDataEncryptionOutputWithContext(ctx context.Context) SqlPoolTransparentDataEncryptionOutput
}

func (*SqlPoolTransparentDataEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolTransparentDataEncryption)(nil))
}

func (i *SqlPoolTransparentDataEncryption) ToSqlPoolTransparentDataEncryptionOutput() SqlPoolTransparentDataEncryptionOutput {
	return i.ToSqlPoolTransparentDataEncryptionOutputWithContext(context.Background())
}

func (i *SqlPoolTransparentDataEncryption) ToSqlPoolTransparentDataEncryptionOutputWithContext(ctx context.Context) SqlPoolTransparentDataEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolTransparentDataEncryptionOutput)
}

type SqlPoolTransparentDataEncryptionOutput struct{ *pulumi.OutputState }

func (SqlPoolTransparentDataEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlPoolTransparentDataEncryption)(nil))
}

func (o SqlPoolTransparentDataEncryptionOutput) ToSqlPoolTransparentDataEncryptionOutput() SqlPoolTransparentDataEncryptionOutput {
	return o
}

func (o SqlPoolTransparentDataEncryptionOutput) ToSqlPoolTransparentDataEncryptionOutputWithContext(ctx context.Context) SqlPoolTransparentDataEncryptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlPoolTransparentDataEncryptionOutput{})
}
