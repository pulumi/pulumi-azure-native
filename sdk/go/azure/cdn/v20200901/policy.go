


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Policy struct {
	pulumi.CustomResourceState

	CustomRules       CustomRuleListResponsePtrOutput     `pulumi:"customRules"`
	EndpointLinks     CdnEndpointResponseArrayOutput      `pulumi:"endpointLinks"`
	Etag              pulumi.StringPtrOutput              `pulumi:"etag"`
	Location          pulumi.StringOutput                 `pulumi:"location"`
	ManagedRules      ManagedRuleSetListResponsePtrOutput `pulumi:"managedRules"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	PolicySettings    PolicySettingsResponsePtrOutput     `pulumi:"policySettings"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	RateLimitRules    RateLimitRuleListResponsePtrOutput  `pulumi:"rateLimitRules"`
	ResourceState     pulumi.StringOutput                 `pulumi:"resourceState"`
	Sku               SkuResponseOutput                   `pulumi:"sku"`
	SystemData        SystemDataResponseOutput            `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput              `pulumi:"tags"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
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
			Type: pulumi.String("azure-nextgen:cdn/v20200901:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20190615:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20190615preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200331:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Policy"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200415:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azure-native:cdn/v20200901:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure-native:cdn/v20200901:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	CustomRules       *CustomRuleList     `pulumi:"customRules"`
	Etag              *string             `pulumi:"etag"`
	Location          *string             `pulumi:"location"`
	ManagedRules      *ManagedRuleSetList `pulumi:"managedRules"`
	PolicyName        *string             `pulumi:"policyName"`
	PolicySettings    *PolicySettings     `pulumi:"policySettings"`
	RateLimitRules    *RateLimitRuleList  `pulumi:"rateLimitRules"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Sku               Sku                 `pulumi:"sku"`
	Tags              map[string]string   `pulumi:"tags"`
}


type PolicyArgs struct {
	CustomRules       CustomRuleListPtrInput
	Etag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ManagedRules      ManagedRuleSetListPtrInput
	PolicyName        pulumi.StringPtrInput
	PolicySettings    PolicySettingsPtrInput
	RateLimitRules    RateLimitRuleListPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuInput
	Tags              pulumi.StringMapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterOutputType(PolicyOutput{})
}
