


package v20220801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CollectorPolicy struct {
	pulumi.CustomResourceState

	EmissionPolicies  EmissionPoliciesPropertiesFormatResponseArrayOutput `pulumi:"emissionPolicies"`
	Etag              pulumi.StringOutput                                 `pulumi:"etag"`
	IngestionPolicy   IngestionPolicyPropertiesFormatResponsePtrOutput    `pulumi:"ingestionPolicy"`
	Location          pulumi.StringPtrOutput                              `pulumi:"location"`
	Name              pulumi.StringOutput                                 `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                                 `pulumi:"provisioningState"`
	SystemData        TrackedResourceResponseSystemDataOutput             `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                              `pulumi:"tags"`
	Type              pulumi.StringOutput                                 `pulumi:"type"`
}


func NewCollectorPolicy(ctx *pulumi.Context,
	name string, args *CollectorPolicyArgs, opts ...pulumi.ResourceOption) (*CollectorPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureTrafficCollectorName == nil {
		return nil, errors.New("invalid value for required argument 'AzureTrafficCollectorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkfunction:CollectorPolicy"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20210901preview:CollectorPolicy"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20220501:CollectorPolicy"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20221101:CollectorPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource CollectorPolicy
	err := ctx.RegisterResource("azure-native:networkfunction/v20220801:CollectorPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCollectorPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CollectorPolicyState, opts ...pulumi.ResourceOption) (*CollectorPolicy, error) {
	var resource CollectorPolicy
	err := ctx.ReadResource("azure-native:networkfunction/v20220801:CollectorPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type collectorPolicyState struct {
}

type CollectorPolicyState struct {
}

func (CollectorPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*collectorPolicyState)(nil)).Elem()
}

type collectorPolicyArgs struct {
	AzureTrafficCollectorName string                             `pulumi:"azureTrafficCollectorName"`
	CollectorPolicyName       *string                            `pulumi:"collectorPolicyName"`
	EmissionPolicies          []EmissionPoliciesPropertiesFormat `pulumi:"emissionPolicies"`
	IngestionPolicy           *IngestionPolicyPropertiesFormat   `pulumi:"ingestionPolicy"`
	Location                  *string                            `pulumi:"location"`
	ResourceGroupName         string                             `pulumi:"resourceGroupName"`
	Tags                      map[string]string                  `pulumi:"tags"`
}


type CollectorPolicyArgs struct {
	AzureTrafficCollectorName pulumi.StringInput
	CollectorPolicyName       pulumi.StringPtrInput
	EmissionPolicies          EmissionPoliciesPropertiesFormatArrayInput
	IngestionPolicy           IngestionPolicyPropertiesFormatPtrInput
	Location                  pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
}

func (CollectorPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*collectorPolicyArgs)(nil)).Elem()
}

type CollectorPolicyInput interface {
	pulumi.Input

	ToCollectorPolicyOutput() CollectorPolicyOutput
	ToCollectorPolicyOutputWithContext(ctx context.Context) CollectorPolicyOutput
}

func (*CollectorPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorPolicy)(nil)).Elem()
}

func (i *CollectorPolicy) ToCollectorPolicyOutput() CollectorPolicyOutput {
	return i.ToCollectorPolicyOutputWithContext(context.Background())
}

func (i *CollectorPolicy) ToCollectorPolicyOutputWithContext(ctx context.Context) CollectorPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorPolicyOutput)
}

type CollectorPolicyOutput struct{ *pulumi.OutputState }

func (CollectorPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorPolicy)(nil)).Elem()
}

func (o CollectorPolicyOutput) ToCollectorPolicyOutput() CollectorPolicyOutput {
	return o
}

func (o CollectorPolicyOutput) ToCollectorPolicyOutputWithContext(ctx context.Context) CollectorPolicyOutput {
	return o
}

func (o CollectorPolicyOutput) EmissionPolicies() EmissionPoliciesPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v *CollectorPolicy) EmissionPoliciesPropertiesFormatResponseArrayOutput {
		return v.EmissionPolicies
	}).(EmissionPoliciesPropertiesFormatResponseArrayOutput)
}

func (o CollectorPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectorPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CollectorPolicyOutput) IngestionPolicy() IngestionPolicyPropertiesFormatResponsePtrOutput {
	return o.ApplyT(func(v *CollectorPolicy) IngestionPolicyPropertiesFormatResponsePtrOutput { return v.IngestionPolicy }).(IngestionPolicyPropertiesFormatResponsePtrOutput)
}

func (o CollectorPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CollectorPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectorPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CollectorPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectorPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CollectorPolicyOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *CollectorPolicy) TrackedResourceResponseSystemDataOutput { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

func (o CollectorPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CollectorPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CollectorPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CollectorPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CollectorPolicyOutput{})
}
