


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomerSubscription struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	TenantId   pulumi.StringPtrOutput   `pulumi:"tenantId"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewCustomerSubscription(ctx *pulumi.Context,
	name string, args *CustomerSubscriptionArgs, opts ...pulumi.ResourceOption) (*CustomerSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistrationName == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationName'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:azurestack/v20200601preview:CustomerSubscription"),
		},
		{
			Type: pulumi.String("azure-native:azurestack:CustomerSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:azurestack:CustomerSubscription"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20170601:CustomerSubscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:azurestack/v20170601:CustomerSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomerSubscription
	err := ctx.RegisterResource("azure-native:azurestack/v20200601preview:CustomerSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomerSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerSubscriptionState, opts ...pulumi.ResourceOption) (*CustomerSubscription, error) {
	var resource CustomerSubscription
	err := ctx.ReadResource("azure-native:azurestack/v20200601preview:CustomerSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customerSubscriptionState struct {
}

type CustomerSubscriptionState struct {
}

func (CustomerSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerSubscriptionState)(nil)).Elem()
}

type customerSubscriptionArgs struct {
	CustomerSubscriptionName *string `pulumi:"customerSubscriptionName"`
	Etag                     *string `pulumi:"etag"`
	RegistrationName         string  `pulumi:"registrationName"`
	ResourceGroup            string  `pulumi:"resourceGroup"`
	TenantId                 *string `pulumi:"tenantId"`
}


type CustomerSubscriptionArgs struct {
	CustomerSubscriptionName pulumi.StringPtrInput
	Etag                     pulumi.StringPtrInput
	RegistrationName         pulumi.StringInput
	ResourceGroup            pulumi.StringInput
	TenantId                 pulumi.StringPtrInput
}

func (CustomerSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerSubscriptionArgs)(nil)).Elem()
}

type CustomerSubscriptionInput interface {
	pulumi.Input

	ToCustomerSubscriptionOutput() CustomerSubscriptionOutput
	ToCustomerSubscriptionOutputWithContext(ctx context.Context) CustomerSubscriptionOutput
}

func (*CustomerSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscription)(nil))
}

func (i *CustomerSubscription) ToCustomerSubscriptionOutput() CustomerSubscriptionOutput {
	return i.ToCustomerSubscriptionOutputWithContext(context.Background())
}

func (i *CustomerSubscription) ToCustomerSubscriptionOutputWithContext(ctx context.Context) CustomerSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionOutput)
}

type CustomerSubscriptionOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscription)(nil))
}

func (o CustomerSubscriptionOutput) ToCustomerSubscriptionOutput() CustomerSubscriptionOutput {
	return o
}

func (o CustomerSubscriptionOutput) ToCustomerSubscriptionOutputWithContext(ctx context.Context) CustomerSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomerSubscriptionOutput{})
}
