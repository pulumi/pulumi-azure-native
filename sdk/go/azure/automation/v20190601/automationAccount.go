


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutomationAccount struct {
	pulumi.CustomResourceState

	CreationTime     pulumi.StringOutput    `pulumi:"creationTime"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	Etag             pulumi.StringPtrOutput `pulumi:"etag"`
	LastModifiedBy   pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	LastModifiedTime pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	Location         pulumi.StringPtrOutput `pulumi:"location"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Sku              SkuResponsePtrOutput   `pulumi:"sku"`
	State            pulumi.StringOutput    `pulumi:"state"`
	Tags             pulumi.StringMapOutput `pulumi:"tags"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewAutomationAccount(ctx *pulumi.Context,
	name string, args *AutomationAccountArgs, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20200113preview:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20210622:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20210622:AutomationAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource AutomationAccount
	err := ctx.RegisterResource("azure-native:automation/v20190601:AutomationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAutomationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationAccountState, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	var resource AutomationAccount
	err := ctx.ReadResource("azure-native:automation/v20190601:AutomationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type automationAccountState struct {
}

type AutomationAccountState struct {
}

func (AutomationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountState)(nil)).Elem()
}

type automationAccountArgs struct {
	AutomationAccountName *string           `pulumi:"automationAccountName"`
	Location              *string           `pulumi:"location"`
	Name                  *string           `pulumi:"name"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	Sku                   *Sku              `pulumi:"sku"`
	Tags                  map[string]string `pulumi:"tags"`
}


type AutomationAccountArgs struct {
	AutomationAccountName pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Sku                   SkuPtrInput
	Tags                  pulumi.StringMapInput
}

func (AutomationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountArgs)(nil)).Elem()
}

type AutomationAccountInput interface {
	pulumi.Input

	ToAutomationAccountOutput() AutomationAccountOutput
	ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput
}

func (*AutomationAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationAccount)(nil))
}

func (i *AutomationAccount) ToAutomationAccountOutput() AutomationAccountOutput {
	return i.ToAutomationAccountOutputWithContext(context.Background())
}

func (i *AutomationAccount) ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationAccountOutput)
}

type AutomationAccountOutput struct{ *pulumi.OutputState }

func (AutomationAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationAccount)(nil))
}

func (o AutomationAccountOutput) ToAutomationAccountOutput() AutomationAccountOutput {
	return o
}

func (o AutomationAccountOutput) ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AutomationAccountOutput{})
}
