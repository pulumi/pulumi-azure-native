


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SpatialAnchorsAccount struct {
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


func NewSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, args *SpatialAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mixedreality:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20190228preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20191202preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20200501:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:SpatialAnchorsAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource SpatialAnchorsAccount
	err := ctx.RegisterResource("azure-native:mixedreality/v20210101:SpatialAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpatialAnchorsAccountState, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	var resource SpatialAnchorsAccount
	err := ctx.ReadResource("azure-native:mixedreality/v20210101:SpatialAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type spatialAnchorsAccountState struct {
}

type SpatialAnchorsAccountState struct {
}

func (SpatialAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountState)(nil)).Elem()
}

type spatialAnchorsAccountArgs struct {
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


type SpatialAnchorsAccountArgs struct {
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

func (SpatialAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountArgs)(nil)).Elem()
}

type SpatialAnchorsAccountInput interface {
	pulumi.Input

	ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput
	ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput
}

func (*SpatialAnchorsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil)).Elem()
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return i.ToSpatialAnchorsAccountOutputWithContext(context.Background())
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountOutput)
}

type SpatialAnchorsAccountOutput struct{ *pulumi.OutputState }

func (SpatialAnchorsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialAnchorsAccount)(nil)).Elem()
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) AccountDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.AccountDomain }).(pulumi.StringOutput)
}

func (o SpatialAnchorsAccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o SpatialAnchorsAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o SpatialAnchorsAccountOutput) Kind() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) SkuResponsePtrOutput { return v.Kind }).(SkuResponsePtrOutput)
}

func (o SpatialAnchorsAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SpatialAnchorsAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SpatialAnchorsAccountOutput) Plan() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) IdentityResponsePtrOutput { return v.Plan }).(IdentityResponsePtrOutput)
}

func (o SpatialAnchorsAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o SpatialAnchorsAccountOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringPtrOutput { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

func (o SpatialAnchorsAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SpatialAnchorsAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SpatialAnchorsAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SpatialAnchorsAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SpatialAnchorsAccountOutput{})
}
