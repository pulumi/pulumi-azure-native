


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedStorageAccount struct {
	pulumi.CustomResourceState

	DataSourceType    pulumi.StringOutput      `pulumi:"dataSourceType"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	StorageAccountIds pulumi.StringArrayOutput `pulumi:"storageAccountIds"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewLinkedStorageAccount(ctx *pulumi.Context,
	name string, args *LinkedStorageAccountArgs, opts ...pulumi.ResourceOption) (*LinkedStorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:LinkedStorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:LinkedStorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:LinkedStorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedStorageAccount
	err := ctx.RegisterResource("azure-native:operationalinsights:LinkedStorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedStorageAccountState, opts ...pulumi.ResourceOption) (*LinkedStorageAccount, error) {
	var resource LinkedStorageAccount
	err := ctx.ReadResource("azure-native:operationalinsights:LinkedStorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkedStorageAccountState struct {
}

type LinkedStorageAccountState struct {
}

func (LinkedStorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedStorageAccountState)(nil)).Elem()
}

type linkedStorageAccountArgs struct {
	DataSourceType    *string  `pulumi:"dataSourceType"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	StorageAccountIds []string `pulumi:"storageAccountIds"`
	WorkspaceName     string   `pulumi:"workspaceName"`
}


type LinkedStorageAccountArgs struct {
	DataSourceType    pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StorageAccountIds pulumi.StringArrayInput
	WorkspaceName     pulumi.StringInput
}

func (LinkedStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedStorageAccountArgs)(nil)).Elem()
}

type LinkedStorageAccountInput interface {
	pulumi.Input

	ToLinkedStorageAccountOutput() LinkedStorageAccountOutput
	ToLinkedStorageAccountOutputWithContext(ctx context.Context) LinkedStorageAccountOutput
}

func (*LinkedStorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedStorageAccount)(nil)).Elem()
}

func (i *LinkedStorageAccount) ToLinkedStorageAccountOutput() LinkedStorageAccountOutput {
	return i.ToLinkedStorageAccountOutputWithContext(context.Background())
}

func (i *LinkedStorageAccount) ToLinkedStorageAccountOutputWithContext(ctx context.Context) LinkedStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedStorageAccountOutput)
}

type LinkedStorageAccountOutput struct{ *pulumi.OutputState }

func (LinkedStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedStorageAccount)(nil)).Elem()
}

func (o LinkedStorageAccountOutput) ToLinkedStorageAccountOutput() LinkedStorageAccountOutput {
	return o
}

func (o LinkedStorageAccountOutput) ToLinkedStorageAccountOutputWithContext(ctx context.Context) LinkedStorageAccountOutput {
	return o
}

func (o LinkedStorageAccountOutput) DataSourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedStorageAccount) pulumi.StringOutput { return v.DataSourceType }).(pulumi.StringOutput)
}

func (o LinkedStorageAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedStorageAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LinkedStorageAccountOutput) StorageAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LinkedStorageAccount) pulumi.StringArrayOutput { return v.StorageAccountIds }).(pulumi.StringArrayOutput)
}

func (o LinkedStorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedStorageAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedStorageAccountOutput{})
}
