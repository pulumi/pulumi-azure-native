


package machinelearningexperimentation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	AccountId         pulumi.StringOutput                    `pulumi:"accountId"`
	CreationDate      pulumi.StringOutput                    `pulumi:"creationDate"`
	Description       pulumi.StringPtrOutput                 `pulumi:"description"`
	DiscoveryUri      pulumi.StringOutput                    `pulumi:"discoveryUri"`
	FriendlyName      pulumi.StringPtrOutput                 `pulumi:"friendlyName"`
	KeyVaultId        pulumi.StringOutput                    `pulumi:"keyVaultId"`
	Location          pulumi.StringOutput                    `pulumi:"location"`
	Name              pulumi.StringOutput                    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                    `pulumi:"provisioningState"`
	Seats             pulumi.StringPtrOutput                 `pulumi:"seats"`
	StorageAccount    StorageAccountPropertiesResponseOutput `pulumi:"storageAccount"`
	Tags              pulumi.StringMapOutput                 `pulumi:"tags"`
	Type              pulumi.StringOutput                    `pulumi:"type"`
	VsoAccountId      pulumi.StringOutput                    `pulumi:"vsoAccountId"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccount == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccount'")
	}
	if args.VsoAccountId == nil {
		return nil, errors.New("invalid value for required argument 'VsoAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningexperimentation/v20170501preview:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:machinelearningexperimentation:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:machinelearningexperimentation:Account", name, id, state, &resource, opts...)
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
	AccountName       *string                  `pulumi:"accountName"`
	Description       *string                  `pulumi:"description"`
	FriendlyName      *string                  `pulumi:"friendlyName"`
	KeyVaultId        string                   `pulumi:"keyVaultId"`
	Location          *string                  `pulumi:"location"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Seats             *string                  `pulumi:"seats"`
	StorageAccount    StorageAccountProperties `pulumi:"storageAccount"`
	Tags              map[string]string        `pulumi:"tags"`
	VsoAccountId      string                   `pulumi:"vsoAccountId"`
}


type AccountArgs struct {
	AccountName       pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	FriendlyName      pulumi.StringPtrInput
	KeyVaultId        pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Seats             pulumi.StringPtrInput
	StorageAccount    StorageAccountPropertiesInput
	Tags              pulumi.StringMapInput
	VsoAccountId      pulumi.StringInput
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
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
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
