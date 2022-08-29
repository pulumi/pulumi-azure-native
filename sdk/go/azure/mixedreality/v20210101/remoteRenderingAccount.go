


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RemoteRenderingAccount struct {
	pulumi.CustomResourceState

	AccountDomain      pulumi.StringOutput       `pulumi:"accountDomain"`
	AccountId          pulumi.StringOutput       `pulumi:"accountId"`
	Identity           IdentityResponsePtrOutput `pulumi:"identity"`
	Kind               SkuResponsePtrOutput      `pulumi:"kind"`
	Location           pulumi.StringOutput       `pulumi:"location"`
	Name               pulumi.StringOutput       `pulumi:"name"`
	Plan               IdentityResponsePtrOutput `pulumi:"plan"`
	Sku                SkuResponsePtrOutput      `pulumi:"sku"`
	StorageAccountName pulumi.StringPtrOutput    `pulumi:"storageAccountName"`
	SystemData         SystemDataResponseOutput  `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput    `pulumi:"tags"`
	Type               pulumi.StringOutput       `pulumi:"type"`
}


func NewRemoteRenderingAccount(ctx *pulumi.Context,
	name string, args *RemoteRenderingAccountArgs, opts ...pulumi.ResourceOption) (*RemoteRenderingAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mixedreality:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20191202preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20200406preview:RemoteRenderingAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:RemoteRenderingAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource RemoteRenderingAccount
	err := ctx.RegisterResource("azure-native:mixedreality/v20210101:RemoteRenderingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRemoteRenderingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteRenderingAccountState, opts ...pulumi.ResourceOption) (*RemoteRenderingAccount, error) {
	var resource RemoteRenderingAccount
	err := ctx.ReadResource("azure-native:mixedreality/v20210101:RemoteRenderingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type remoteRenderingAccountState struct {
}

type RemoteRenderingAccountState struct {
}

func (RemoteRenderingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRenderingAccountState)(nil)).Elem()
}

type remoteRenderingAccountArgs struct {
	AccountName        *string           `pulumi:"accountName"`
	Identity           *Identity         `pulumi:"identity"`
	Kind               *Sku              `pulumi:"kind"`
	Location           *string           `pulumi:"location"`
	Plan               *Identity         `pulumi:"plan"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Sku                *Sku              `pulumi:"sku"`
	StorageAccountName *string           `pulumi:"storageAccountName"`
	Tags               map[string]string `pulumi:"tags"`
}


type RemoteRenderingAccountArgs struct {
	AccountName        pulumi.StringPtrInput
	Identity           IdentityPtrInput
	Kind               SkuPtrInput
	Location           pulumi.StringPtrInput
	Plan               IdentityPtrInput
	ResourceGroupName  pulumi.StringInput
	Sku                SkuPtrInput
	StorageAccountName pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (RemoteRenderingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteRenderingAccountArgs)(nil)).Elem()
}

type RemoteRenderingAccountInput interface {
	pulumi.Input

	ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput
	ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput
}

func (*RemoteRenderingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteRenderingAccount)(nil)).Elem()
}

func (i *RemoteRenderingAccount) ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput {
	return i.ToRemoteRenderingAccountOutputWithContext(context.Background())
}

func (i *RemoteRenderingAccount) ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteRenderingAccountOutput)
}

type RemoteRenderingAccountOutput struct{ *pulumi.OutputState }

func (RemoteRenderingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteRenderingAccount)(nil)).Elem()
}

func (o RemoteRenderingAccountOutput) ToRemoteRenderingAccountOutput() RemoteRenderingAccountOutput {
	return o
}

func (o RemoteRenderingAccountOutput) ToRemoteRenderingAccountOutputWithContext(ctx context.Context) RemoteRenderingAccountOutput {
	return o
}

func (o RemoteRenderingAccountOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.AccountDomain }).(pulumi.StringOutput)
}

func (o RemoteRenderingAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o RemoteRenderingAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o RemoteRenderingAccountOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) SkuResponsePtrOutput { return v.Kind }).(SkuResponsePtrOutput)
}

func (o RemoteRenderingAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RemoteRenderingAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RemoteRenderingAccountOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) IdentityResponsePtrOutput { return v.Plan }).(IdentityResponsePtrOutput)
}

func (o RemoteRenderingAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o RemoteRenderingAccountOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringPtrOutput { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

func (o RemoteRenderingAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o RemoteRenderingAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RemoteRenderingAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RemoteRenderingAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RemoteRenderingAccountOutput{})
}
