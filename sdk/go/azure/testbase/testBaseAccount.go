


package testbase

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TestBaseAccount struct {
	pulumi.CustomResourceState

	AccessLevel       pulumi.StringOutput              `pulumi:"accessLevel"`
	Etag              pulumi.StringOutput              `pulumi:"etag"`
	Location          pulumi.StringOutput              `pulumi:"location"`
	Name              pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState pulumi.StringOutput              `pulumi:"provisioningState"`
	Sku               TestBaseAccountSKUResponseOutput `pulumi:"sku"`
	SystemData        SystemDataResponseOutput         `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput           `pulumi:"tags"`
	Type              pulumi.StringOutput              `pulumi:"type"`
}


func NewTestBaseAccount(ctx *pulumi.Context,
	name string, args *TestBaseAccountArgs, opts ...pulumi.ResourceOption) (*TestBaseAccount, error) {
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
			Type: pulumi.String("azure-native:testbase/v20201216preview:TestBaseAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource TestBaseAccount
	err := ctx.RegisterResource("azure-native:testbase:TestBaseAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTestBaseAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TestBaseAccountState, opts ...pulumi.ResourceOption) (*TestBaseAccount, error) {
	var resource TestBaseAccount
	err := ctx.ReadResource("azure-native:testbase:TestBaseAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type testBaseAccountState struct {
}

type TestBaseAccountState struct {
}

func (TestBaseAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*testBaseAccountState)(nil)).Elem()
}

type testBaseAccountArgs struct {
	Location            *string            `pulumi:"location"`
	ResourceGroupName   string             `pulumi:"resourceGroupName"`
	Restore             *bool              `pulumi:"restore"`
	Sku                 TestBaseAccountSKU `pulumi:"sku"`
	Tags                map[string]string  `pulumi:"tags"`
	TestBaseAccountName *string            `pulumi:"testBaseAccountName"`
}


type TestBaseAccountArgs struct {
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Restore             pulumi.BoolPtrInput
	Sku                 TestBaseAccountSKUInput
	Tags                pulumi.StringMapInput
	TestBaseAccountName pulumi.StringPtrInput
}

func (TestBaseAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*testBaseAccountArgs)(nil)).Elem()
}

type TestBaseAccountInput interface {
	pulumi.Input

	ToTestBaseAccountOutput() TestBaseAccountOutput
	ToTestBaseAccountOutputWithContext(ctx context.Context) TestBaseAccountOutput
}

func (*TestBaseAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccount)(nil)).Elem()
}

func (i *TestBaseAccount) ToTestBaseAccountOutput() TestBaseAccountOutput {
	return i.ToTestBaseAccountOutputWithContext(context.Background())
}

func (i *TestBaseAccount) ToTestBaseAccountOutputWithContext(ctx context.Context) TestBaseAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountOutput)
}

type TestBaseAccountOutput struct{ *pulumi.OutputState }

func (TestBaseAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccount)(nil)).Elem()
}

func (o TestBaseAccountOutput) ToTestBaseAccountOutput() TestBaseAccountOutput {
	return o
}

func (o TestBaseAccountOutput) ToTestBaseAccountOutputWithContext(ctx context.Context) TestBaseAccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TestBaseAccountOutput{})
}
