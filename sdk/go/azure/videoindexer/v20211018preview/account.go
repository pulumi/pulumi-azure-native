


package v20211018preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	AccountId         pulumi.StringPtrOutput                      `pulumi:"accountId"`
	AccountName       pulumi.StringOutput                         `pulumi:"accountName"`
	Identity          ManagedServiceIdentityResponsePtrOutput     `pulumi:"identity"`
	Location          pulumi.StringOutput                         `pulumi:"location"`
	MediaServices     MediaServicesForPutRequestResponsePtrOutput `pulumi:"mediaServices"`
	Name              pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                         `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                    `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                      `pulumi:"tags"`
	TenantId          pulumi.StringOutput                         `pulumi:"tenantId"`
	Type              pulumi.StringOutput                         `pulumi:"type"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AccountId) {
		args.AccountId = pulumi.StringPtr("00000000-0000-0000-0000-000000000000")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoindexer:Account"),
		},
		{
			Type: pulumi.String("azure-native:videoindexer/v20211027preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:videoindexer/v20211110preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:videoindexer/v20220413preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:videoindexer/v20220801:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:videoindexer/v20211018preview:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:videoindexer/v20211018preview:Account", name, id, state, &resource, opts...)
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
	AccountId         *string                     `pulumi:"accountId"`
	AccountName       *string                     `pulumi:"accountName"`
	Identity          *ManagedServiceIdentity     `pulumi:"identity"`
	Location          *string                     `pulumi:"location"`
	MediaServices     *MediaServicesForPutRequest `pulumi:"mediaServices"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	Tags              map[string]string           `pulumi:"tags"`
}


type AccountArgs struct {
	AccountId         pulumi.StringPtrInput
	AccountName       pulumi.StringPtrInput
	Identity          ManagedServiceIdentityPtrInput
	Location          pulumi.StringPtrInput
	MediaServices     MediaServicesForPutRequestPtrInput
	ResourceGroupName pulumi.StringInput
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

func (o AccountOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountName }).(pulumi.StringOutput)
}

func (o AccountOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AccountOutput) MediaServices() MediaServicesForPutRequestResponsePtrOutput {
	return o.ApplyT(func(v *Account) MediaServicesForPutRequestResponsePtrOutput { return v.MediaServices }).(MediaServicesForPutRequestResponsePtrOutput)
}

func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Account) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AccountOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
