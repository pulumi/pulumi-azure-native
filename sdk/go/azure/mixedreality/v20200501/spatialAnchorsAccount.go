


package v20200501

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
			Type: pulumi.String("azure-nextgen:mixedreality/v20200501:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20190228preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20190228preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20191202preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20191202preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210101:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20210101:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:SpatialAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20210301preview:SpatialAnchorsAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource SpatialAnchorsAccount
	err := ctx.RegisterResource("azure-native:mixedreality/v20200501:SpatialAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpatialAnchorsAccountState, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	var resource SpatialAnchorsAccount
	err := ctx.ReadResource("azure-native:mixedreality/v20200501:SpatialAnchorsAccount", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*SpatialAnchorsAccount)(nil))
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return i.ToSpatialAnchorsAccountOutputWithContext(context.Background())
}

func (i *SpatialAnchorsAccount) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialAnchorsAccountOutput)
}

type SpatialAnchorsAccountOutput struct{ *pulumi.OutputState }

func (SpatialAnchorsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialAnchorsAccount)(nil))
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutput() SpatialAnchorsAccountOutput {
	return o
}

func (o SpatialAnchorsAccountOutput) ToSpatialAnchorsAccountOutputWithContext(ctx context.Context) SpatialAnchorsAccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SpatialAnchorsAccountOutput{})
}
