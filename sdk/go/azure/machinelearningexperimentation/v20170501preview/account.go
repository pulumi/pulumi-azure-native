


package v20170501preview

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
			Type: pulumi.String("azure-native:machinelearningexperimentation:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:machinelearningexperimentation/v20170501preview:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:machinelearningexperimentation/v20170501preview:Account", name, id, state, &resource, opts...)
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

func (o AccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o AccountOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o AccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) DiscoveryUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DiscoveryUri }).(pulumi.StringOutput)
}

func (o AccountOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) KeyVaultId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.KeyVaultId }).(pulumi.StringOutput)
}

func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AccountOutput) Seats() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Seats }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) StorageAccount() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *Account) StorageAccountPropertiesResponseOutput { return v.StorageAccount }).(StorageAccountPropertiesResponseOutput)
}

func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AccountOutput) VsoAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.VsoAccountId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
