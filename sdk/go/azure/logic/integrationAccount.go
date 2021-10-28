


package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccount struct {
	pulumi.CustomResourceState

	IntegrationServiceEnvironment ResourceReferenceResponsePtrOutput     `pulumi:"integrationServiceEnvironment"`
	Location                      pulumi.StringPtrOutput                 `pulumi:"location"`
	Name                          pulumi.StringOutput                    `pulumi:"name"`
	Sku                           IntegrationAccountSkuResponsePtrOutput `pulumi:"sku"`
	State                         pulumi.StringPtrOutput                 `pulumi:"state"`
	Tags                          pulumi.StringMapOutput                 `pulumi:"tags"`
	Type                          pulumi.StringOutput                    `pulumi:"type"`
}


func NewIntegrationAccount(ctx *pulumi.Context,
	name string, args *IntegrationAccountArgs, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:logic:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccount
	err := ctx.RegisterResource("azure-native:logic:IntegrationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountState, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	var resource IntegrationAccount
	err := ctx.ReadResource("azure-native:logic:IntegrationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountState struct {
}

type IntegrationAccountState struct {
}

func (IntegrationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountState)(nil)).Elem()
}

type integrationAccountArgs struct {
	IntegrationAccountName        *string                `pulumi:"integrationAccountName"`
	IntegrationServiceEnvironment *ResourceReference     `pulumi:"integrationServiceEnvironment"`
	Location                      *string                `pulumi:"location"`
	ResourceGroupName             string                 `pulumi:"resourceGroupName"`
	Sku                           *IntegrationAccountSku `pulumi:"sku"`
	State                         *string                `pulumi:"state"`
	Tags                          map[string]string      `pulumi:"tags"`
}


type IntegrationAccountArgs struct {
	IntegrationAccountName        pulumi.StringPtrInput
	IntegrationServiceEnvironment ResourceReferencePtrInput
	Location                      pulumi.StringPtrInput
	ResourceGroupName             pulumi.StringInput
	Sku                           IntegrationAccountSkuPtrInput
	State                         pulumi.StringPtrInput
	Tags                          pulumi.StringMapInput
}

func (IntegrationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountArgs)(nil)).Elem()
}

type IntegrationAccountInput interface {
	pulumi.Input

	ToIntegrationAccountOutput() IntegrationAccountOutput
	ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput
}

func (*IntegrationAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccount)(nil))
}

func (i *IntegrationAccount) ToIntegrationAccountOutput() IntegrationAccountOutput {
	return i.ToIntegrationAccountOutputWithContext(context.Background())
}

func (i *IntegrationAccount) ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountOutput)
}

type IntegrationAccountOutput struct{ *pulumi.OutputState }

func (IntegrationAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccount)(nil))
}

func (o IntegrationAccountOutput) ToIntegrationAccountOutput() IntegrationAccountOutput {
	return o
}

func (o IntegrationAccountOutput) ToIntegrationAccountOutputWithContext(ctx context.Context) IntegrationAccountOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAccountInput)(nil)).Elem(), &IntegrationAccount{})
	pulumi.RegisterOutputType(IntegrationAccountOutput{})
}
