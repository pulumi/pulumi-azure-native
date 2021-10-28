


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DdosCustomPolicy struct {
	pulumi.CustomResourceState

	Etag                   pulumi.StringOutput                             `pulumi:"etag"`
	Location               pulumi.StringPtrOutput                          `pulumi:"location"`
	Name                   pulumi.StringOutput                             `pulumi:"name"`
	ProtocolCustomSettings ProtocolCustomSettingsFormatResponseArrayOutput `pulumi:"protocolCustomSettings"`
	ProvisioningState      pulumi.StringOutput                             `pulumi:"provisioningState"`
	PublicIPAddresses      SubResourceResponseArrayOutput                  `pulumi:"publicIPAddresses"`
	ResourceGuid           pulumi.StringOutput                             `pulumi:"resourceGuid"`
	Tags                   pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                   pulumi.StringOutput                             `pulumi:"type"`
}


func NewDdosCustomPolicy(ctx *pulumi.Context,
	name string, args *DdosCustomPolicyArgs, opts ...pulumi.ResourceOption) (*DdosCustomPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:DdosCustomPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:DdosCustomPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource DdosCustomPolicy
	err := ctx.RegisterResource("azure-native:network/v20200701:DdosCustomPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDdosCustomPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosCustomPolicyState, opts ...pulumi.ResourceOption) (*DdosCustomPolicy, error) {
	var resource DdosCustomPolicy
	err := ctx.ReadResource("azure-native:network/v20200701:DdosCustomPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ddosCustomPolicyState struct {
}

type DdosCustomPolicyState struct {
}

func (DdosCustomPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosCustomPolicyState)(nil)).Elem()
}

type ddosCustomPolicyArgs struct {
	DdosCustomPolicyName   *string                        `pulumi:"ddosCustomPolicyName"`
	Id                     *string                        `pulumi:"id"`
	Location               *string                        `pulumi:"location"`
	ProtocolCustomSettings []ProtocolCustomSettingsFormat `pulumi:"protocolCustomSettings"`
	ResourceGroupName      string                         `pulumi:"resourceGroupName"`
	Tags                   map[string]string              `pulumi:"tags"`
}


type DdosCustomPolicyArgs struct {
	DdosCustomPolicyName   pulumi.StringPtrInput
	Id                     pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ProtocolCustomSettings ProtocolCustomSettingsFormatArrayInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (DdosCustomPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosCustomPolicyArgs)(nil)).Elem()
}

type DdosCustomPolicyInput interface {
	pulumi.Input

	ToDdosCustomPolicyOutput() DdosCustomPolicyOutput
	ToDdosCustomPolicyOutputWithContext(ctx context.Context) DdosCustomPolicyOutput
}

func (*DdosCustomPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosCustomPolicy)(nil))
}

func (i *DdosCustomPolicy) ToDdosCustomPolicyOutput() DdosCustomPolicyOutput {
	return i.ToDdosCustomPolicyOutputWithContext(context.Background())
}

func (i *DdosCustomPolicy) ToDdosCustomPolicyOutputWithContext(ctx context.Context) DdosCustomPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosCustomPolicyOutput)
}

type DdosCustomPolicyOutput struct{ *pulumi.OutputState }

func (DdosCustomPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosCustomPolicy)(nil))
}

func (o DdosCustomPolicyOutput) ToDdosCustomPolicyOutput() DdosCustomPolicyOutput {
	return o
}

func (o DdosCustomPolicyOutput) ToDdosCustomPolicyOutputWithContext(ctx context.Context) DdosCustomPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DdosCustomPolicyOutput{})
}
