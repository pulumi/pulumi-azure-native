


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput             `pulumi:"etag"`
	Identity   IdentityResponsePtrOutput       `pulumi:"identity"`
	Kind       pulumi.StringPtrOutput          `pulumi:"kind"`
	Location   pulumi.StringPtrOutput          `pulumi:"location"`
	Name       pulumi.StringOutput             `pulumi:"name"`
	Properties AccountPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput            `pulumi:"sku"`
	SystemData SystemDataResponseOutput        `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput          `pulumi:"tags"`
	Type       pulumi.StringOutput             `pulumi:"type"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToAccountPropertiesPtrOutput().ApplyT(func(v *AccountProperties) *AccountProperties { return v.Defaults() }).(AccountPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20160201preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20170418:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20210430:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20220301:Account"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20221001:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20211001:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:cognitiveservices/v20211001:Account", name, id, state, &resource, opts...)
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
	AccountName       *string            `pulumi:"accountName"`
	Identity          *Identity          `pulumi:"identity"`
	Kind              *string            `pulumi:"kind"`
	Location          *string            `pulumi:"location"`
	Properties        *AccountProperties `pulumi:"properties"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	Sku               *Sku               `pulumi:"sku"`
	Tags              map[string]string  `pulumi:"tags"`
}


type AccountArgs struct {
	AccountName       pulumi.StringPtrInput
	Identity          IdentityPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        AccountPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
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

func (o AccountOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o AccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccountOutput) Properties() AccountPropertiesResponseOutput {
	return o.ApplyT(func(v *Account) AccountPropertiesResponseOutput { return v.Properties }).(AccountPropertiesResponseOutput)
}

func (o AccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Account) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o AccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Account) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
