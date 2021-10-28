


package mixedreality

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ObjectAnchorsAccount struct {
	pulumi.CustomResourceState

	AccountDomain      pulumi.StringOutput                           `pulumi:"accountDomain"`
	AccountId          pulumi.StringOutput                           `pulumi:"accountId"`
	Identity           ObjectAnchorsAccountResponseIdentityPtrOutput `pulumi:"identity"`
	Kind               SkuResponsePtrOutput                          `pulumi:"kind"`
	Location           pulumi.StringOutput                           `pulumi:"location"`
	Name               pulumi.StringOutput                           `pulumi:"name"`
	Plan               IdentityResponsePtrOutput                     `pulumi:"plan"`
	Sku                SkuResponsePtrOutput                          `pulumi:"sku"`
	StorageAccountName pulumi.StringPtrOutput                        `pulumi:"storageAccountName"`
	SystemData         SystemDataResponseOutput                      `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput                        `pulumi:"tags"`
	Type               pulumi.StringOutput                           `pulumi:"type"`
}


func NewObjectAnchorsAccount(ctx *pulumi.Context,
	name string, args *ObjectAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*ObjectAnchorsAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:mixedreality:ObjectAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-native:mixedreality/v20210301preview:ObjectAnchorsAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:mixedreality/v20210301preview:ObjectAnchorsAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource ObjectAnchorsAccount
	err := ctx.RegisterResource("azure-native:mixedreality:ObjectAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetObjectAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectAnchorsAccountState, opts ...pulumi.ResourceOption) (*ObjectAnchorsAccount, error) {
	var resource ObjectAnchorsAccount
	err := ctx.ReadResource("azure-native:mixedreality:ObjectAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type objectAnchorsAccountState struct {
}

type ObjectAnchorsAccountState struct {
}

func (ObjectAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAnchorsAccountState)(nil)).Elem()
}

type objectAnchorsAccountArgs struct {
	AccountName        *string                       `pulumi:"accountName"`
	Identity           *ObjectAnchorsAccountIdentity `pulumi:"identity"`
	Kind               *Sku                          `pulumi:"kind"`
	Location           *string                       `pulumi:"location"`
	Plan               *Identity                     `pulumi:"plan"`
	ResourceGroupName  string                        `pulumi:"resourceGroupName"`
	Sku                *Sku                          `pulumi:"sku"`
	StorageAccountName *string                       `pulumi:"storageAccountName"`
	Tags               map[string]string             `pulumi:"tags"`
}


type ObjectAnchorsAccountArgs struct {
	AccountName        pulumi.StringPtrInput
	Identity           ObjectAnchorsAccountIdentityPtrInput
	Kind               SkuPtrInput
	Location           pulumi.StringPtrInput
	Plan               IdentityPtrInput
	ResourceGroupName  pulumi.StringInput
	Sku                SkuPtrInput
	StorageAccountName pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (ObjectAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAnchorsAccountArgs)(nil)).Elem()
}

type ObjectAnchorsAccountInput interface {
	pulumi.Input

	ToObjectAnchorsAccountOutput() ObjectAnchorsAccountOutput
	ToObjectAnchorsAccountOutputWithContext(ctx context.Context) ObjectAnchorsAccountOutput
}

func (*ObjectAnchorsAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectAnchorsAccount)(nil))
}

func (i *ObjectAnchorsAccount) ToObjectAnchorsAccountOutput() ObjectAnchorsAccountOutput {
	return i.ToObjectAnchorsAccountOutputWithContext(context.Background())
}

func (i *ObjectAnchorsAccount) ToObjectAnchorsAccountOutputWithContext(ctx context.Context) ObjectAnchorsAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAnchorsAccountOutput)
}

type ObjectAnchorsAccountOutput struct{ *pulumi.OutputState }

func (ObjectAnchorsAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectAnchorsAccount)(nil))
}

func (o ObjectAnchorsAccountOutput) ToObjectAnchorsAccountOutput() ObjectAnchorsAccountOutput {
	return o
}

func (o ObjectAnchorsAccountOutput) ToObjectAnchorsAccountOutputWithContext(ctx context.Context) ObjectAnchorsAccountOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAnchorsAccountInput)(nil)).Elem(), &ObjectAnchorsAccount{})
	pulumi.RegisterOutputType(ObjectAnchorsAccountOutput{})
}
