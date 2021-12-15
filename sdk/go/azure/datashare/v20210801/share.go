


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Share struct {
	pulumi.CustomResourceState

	CreatedAt         pulumi.StringOutput      `pulumi:"createdAt"`
	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	ShareKind         pulumi.StringPtrOutput   `pulumi:"shareKind"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Terms             pulumi.StringPtrOutput   `pulumi:"terms"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	UserEmail         pulumi.StringOutput      `pulumi:"userEmail"`
	UserName          pulumi.StringOutput      `pulumi:"userName"`
}


func NewShare(ctx *pulumi.Context,
	name string, args *ShareArgs, opts ...pulumi.ResourceOption) (*Share, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:Share"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:Share"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:Share"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:Share"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:Share"),
		},
	})
	opts = append(opts, aliases)
	var resource Share
	err := ctx.RegisterResource("azure-native:datashare/v20210801:Share", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareState, opts ...pulumi.ResourceOption) (*Share, error) {
	var resource Share
	err := ctx.ReadResource("azure-native:datashare/v20210801:Share", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type shareState struct {
}

type ShareState struct {
}

func (ShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareState)(nil)).Elem()
}

type shareArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Description       *string `pulumi:"description"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareKind         *string `pulumi:"shareKind"`
	ShareName         *string `pulumi:"shareName"`
	Terms             *string `pulumi:"terms"`
}


type ShareArgs struct {
	AccountName       pulumi.StringInput
	Description       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ShareKind         pulumi.StringPtrInput
	ShareName         pulumi.StringPtrInput
	Terms             pulumi.StringPtrInput
}

func (ShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareArgs)(nil)).Elem()
}

type ShareInput interface {
	pulumi.Input

	ToShareOutput() ShareOutput
	ToShareOutputWithContext(ctx context.Context) ShareOutput
}

func (*Share) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (i *Share) ToShareOutput() ShareOutput {
	return i.ToShareOutputWithContext(context.Background())
}

func (i *Share) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareOutput)
}

type ShareOutput struct{ *pulumi.OutputState }

func (ShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (o ShareOutput) ToShareOutput() ShareOutput {
	return o
}

func (o ShareOutput) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ShareOutput{})
}
