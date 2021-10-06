


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	Identity   ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind       pulumi.StringPtrOutput                  `pulumi:"kind"`
	Location   pulumi.StringOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties MapsAccountPropertiesResponseOutput     `pulumi:"properties"`
	Sku        SkuResponseOutput                       `pulumi:"sku"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:maps/v20210701preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:maps:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20170101preview:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:maps/v20170101preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20180501:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:maps/v20180501:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20200201preview:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:maps/v20200201preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20210201:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:maps/v20210201:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:maps/v20210701preview:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:maps/v20210701preview:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	AccountName       *string                 `pulumi:"accountName"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	Kind              *string                 `pulumi:"kind"`
	Location          *string                 `pulumi:"location"`
	Properties        *MapsAccountProperties  `pulumi:"properties"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Sku               Sku                     `pulumi:"sku"`
	Tags              map[string]string       `pulumi:"tags"`
}


type AccountArgs struct {
	AccountName       pulumi.StringPtrInput
	Identity          ManagedServiceIdentityPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        MapsAccountPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
	Tags              pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
