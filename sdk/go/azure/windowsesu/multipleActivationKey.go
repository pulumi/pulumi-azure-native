


package windowsesu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MultipleActivationKey struct {
	pulumi.CustomResourceState

	AgreementNumber       pulumi.StringPtrOutput `pulumi:"agreementNumber"`
	ExpirationDate        pulumi.StringOutput    `pulumi:"expirationDate"`
	InstalledServerNumber pulumi.IntPtrOutput    `pulumi:"installedServerNumber"`
	IsEligible            pulumi.BoolPtrOutput   `pulumi:"isEligible"`
	Location              pulumi.StringOutput    `pulumi:"location"`
	MultipleActivationKey pulumi.StringOutput    `pulumi:"multipleActivationKey"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	OsType                pulumi.StringPtrOutput `pulumi:"osType"`
	ProvisioningState     pulumi.StringOutput    `pulumi:"provisioningState"`
	SupportType           pulumi.StringPtrOutput `pulumi:"supportType"`
	Tags                  pulumi.StringMapOutput `pulumi:"tags"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewMultipleActivationKey(ctx *pulumi.Context,
	name string, args *MultipleActivationKeyArgs, opts ...pulumi.ResourceOption) (*MultipleActivationKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SupportType == nil {
		args.SupportType = pulumi.StringPtr("SupplementalServicing")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:windowsesu:MultipleActivationKey"),
		},
		{
			Type: pulumi.String("azure-native:windowsesu/v20190916preview:MultipleActivationKey"),
		},
		{
			Type: pulumi.String("azure-nextgen:windowsesu/v20190916preview:MultipleActivationKey"),
		},
	})
	opts = append(opts, aliases)
	var resource MultipleActivationKey
	err := ctx.RegisterResource("azure-native:windowsesu:MultipleActivationKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMultipleActivationKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultipleActivationKeyState, opts ...pulumi.ResourceOption) (*MultipleActivationKey, error) {
	var resource MultipleActivationKey
	err := ctx.ReadResource("azure-native:windowsesu:MultipleActivationKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type multipleActivationKeyState struct {
}

type MultipleActivationKeyState struct {
}

func (MultipleActivationKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*multipleActivationKeyState)(nil)).Elem()
}

type multipleActivationKeyArgs struct {
	AgreementNumber           *string           `pulumi:"agreementNumber"`
	InstalledServerNumber     *int              `pulumi:"installedServerNumber"`
	IsEligible                *bool             `pulumi:"isEligible"`
	Location                  *string           `pulumi:"location"`
	MultipleActivationKeyName *string           `pulumi:"multipleActivationKeyName"`
	OsType                    *string           `pulumi:"osType"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	SupportType               *string           `pulumi:"supportType"`
	Tags                      map[string]string `pulumi:"tags"`
}


type MultipleActivationKeyArgs struct {
	AgreementNumber           pulumi.StringPtrInput
	InstalledServerNumber     pulumi.IntPtrInput
	IsEligible                pulumi.BoolPtrInput
	Location                  pulumi.StringPtrInput
	MultipleActivationKeyName pulumi.StringPtrInput
	OsType                    pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	SupportType               pulumi.StringPtrInput
	Tags                      pulumi.StringMapInput
}

func (MultipleActivationKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multipleActivationKeyArgs)(nil)).Elem()
}

type MultipleActivationKeyInput interface {
	pulumi.Input

	ToMultipleActivationKeyOutput() MultipleActivationKeyOutput
	ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput
}

func (*MultipleActivationKey) ElementType() reflect.Type {
	return reflect.TypeOf((*MultipleActivationKey)(nil))
}

func (i *MultipleActivationKey) ToMultipleActivationKeyOutput() MultipleActivationKeyOutput {
	return i.ToMultipleActivationKeyOutputWithContext(context.Background())
}

func (i *MultipleActivationKey) ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultipleActivationKeyOutput)
}

type MultipleActivationKeyOutput struct{ *pulumi.OutputState }

func (MultipleActivationKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultipleActivationKey)(nil))
}

func (o MultipleActivationKeyOutput) ToMultipleActivationKeyOutput() MultipleActivationKeyOutput {
	return o
}

func (o MultipleActivationKeyOutput) ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MultipleActivationKeyOutput{})
}
