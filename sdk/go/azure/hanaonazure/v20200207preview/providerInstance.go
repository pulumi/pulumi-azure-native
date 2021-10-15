


package v20200207preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderInstance struct {
	pulumi.CustomResourceState

	Metadata          pulumi.StringPtrOutput `pulumi:"metadata"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	Properties        pulumi.StringOutput    `pulumi:"properties"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewProviderInstance(ctx *pulumi.Context,
	name string, args *ProviderInstanceArgs, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapMonitorName == nil {
		return nil, errors.New("invalid value for required argument 'SapMonitorName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:hanaonazure/v20200207preview:ProviderInstance"),
		},
		{
			Type: pulumi.String("azure-native:hanaonazure:ProviderInstance"),
		},
		{
			Type: pulumi.String("azure-nextgen:hanaonazure:ProviderInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource ProviderInstance
	err := ctx.RegisterResource("azure-native:hanaonazure/v20200207preview:ProviderInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProviderInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderInstanceState, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	var resource ProviderInstance
	err := ctx.ReadResource("azure-native:hanaonazure/v20200207preview:ProviderInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type providerInstanceState struct {
}

type ProviderInstanceState struct {
}

func (ProviderInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceState)(nil)).Elem()
}

type providerInstanceArgs struct {
	Metadata             *string `pulumi:"metadata"`
	Properties           *string `pulumi:"properties"`
	ProviderInstanceName *string `pulumi:"providerInstanceName"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
	SapMonitorName       string  `pulumi:"sapMonitorName"`
	Type                 *string `pulumi:"type"`
}


type ProviderInstanceArgs struct {
	Metadata             pulumi.StringPtrInput
	Properties           pulumi.StringPtrInput
	ProviderInstanceName pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	SapMonitorName       pulumi.StringInput
	Type                 pulumi.StringPtrInput
}

func (ProviderInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceArgs)(nil)).Elem()
}

type ProviderInstanceInput interface {
	pulumi.Input

	ToProviderInstanceOutput() ProviderInstanceOutput
	ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput
}

func (*ProviderInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderInstance)(nil))
}

func (i *ProviderInstance) ToProviderInstanceOutput() ProviderInstanceOutput {
	return i.ToProviderInstanceOutputWithContext(context.Background())
}

func (i *ProviderInstance) ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderInstanceOutput)
}

type ProviderInstanceOutput struct{ *pulumi.OutputState }

func (ProviderInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderInstance)(nil))
}

func (o ProviderInstanceOutput) ToProviderInstanceOutput() ProviderInstanceOutput {
	return o
}

func (o ProviderInstanceOutput) ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProviderInstanceOutput{})
}
