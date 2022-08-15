


package v20220801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureTrafficCollector struct {
	pulumi.CustomResourceState

	CollectorPolicies CollectorPolicyResponseArrayOutput      `pulumi:"collectorPolicies"`
	Etag              pulumi.StringOutput                     `pulumi:"etag"`
	Location          pulumi.StringPtrOutput                  `pulumi:"location"`
	Name              pulumi.StringOutput                     `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                     `pulumi:"provisioningState"`
	SystemData        TrackedResourceResponseSystemDataOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                  `pulumi:"tags"`
	Type              pulumi.StringOutput                     `pulumi:"type"`
	VirtualHub        ResourceReferenceResponsePtrOutput      `pulumi:"virtualHub"`
}


func NewAzureTrafficCollector(ctx *pulumi.Context,
	name string, args *AzureTrafficCollectorArgs, opts ...pulumi.ResourceOption) (*AzureTrafficCollector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkfunction:AzureTrafficCollector"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20210901preview:AzureTrafficCollector"),
		},
		{
			Type: pulumi.String("azure-native:networkfunction/v20220501:AzureTrafficCollector"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureTrafficCollector
	err := ctx.RegisterResource("azure-native:networkfunction/v20220801:AzureTrafficCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureTrafficCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureTrafficCollectorState, opts ...pulumi.ResourceOption) (*AzureTrafficCollector, error) {
	var resource AzureTrafficCollector
	err := ctx.ReadResource("azure-native:networkfunction/v20220801:AzureTrafficCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azureTrafficCollectorState struct {
}

type AzureTrafficCollectorState struct {
}

func (AzureTrafficCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureTrafficCollectorState)(nil)).Elem()
}

type azureTrafficCollectorArgs struct {
	AzureTrafficCollectorName *string               `pulumi:"azureTrafficCollectorName"`
	CollectorPolicies         []CollectorPolicyType `pulumi:"collectorPolicies"`
	Location                  *string               `pulumi:"location"`
	ResourceGroupName         string                `pulumi:"resourceGroupName"`
	Tags                      map[string]string     `pulumi:"tags"`
}


type AzureTrafficCollectorArgs struct {
	AzureTrafficCollectorName pulumi.StringPtrInput
	CollectorPolicies         CollectorPolicyTypeArrayInput
	Location                  pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
}

func (AzureTrafficCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureTrafficCollectorArgs)(nil)).Elem()
}

type AzureTrafficCollectorInput interface {
	pulumi.Input

	ToAzureTrafficCollectorOutput() AzureTrafficCollectorOutput
	ToAzureTrafficCollectorOutputWithContext(ctx context.Context) AzureTrafficCollectorOutput
}

func (*AzureTrafficCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTrafficCollector)(nil)).Elem()
}

func (i *AzureTrafficCollector) ToAzureTrafficCollectorOutput() AzureTrafficCollectorOutput {
	return i.ToAzureTrafficCollectorOutputWithContext(context.Background())
}

func (i *AzureTrafficCollector) ToAzureTrafficCollectorOutputWithContext(ctx context.Context) AzureTrafficCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTrafficCollectorOutput)
}

type AzureTrafficCollectorOutput struct{ *pulumi.OutputState }

func (AzureTrafficCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTrafficCollector)(nil)).Elem()
}

func (o AzureTrafficCollectorOutput) ToAzureTrafficCollectorOutput() AzureTrafficCollectorOutput {
	return o
}

func (o AzureTrafficCollectorOutput) ToAzureTrafficCollectorOutputWithContext(ctx context.Context) AzureTrafficCollectorOutput {
	return o
}

func (o AzureTrafficCollectorOutput) CollectorPolicies() CollectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) CollectorPolicyResponseArrayOutput { return v.CollectorPolicies }).(CollectorPolicyResponseArrayOutput)
}

func (o AzureTrafficCollectorOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AzureTrafficCollectorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AzureTrafficCollectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureTrafficCollectorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AzureTrafficCollectorOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) TrackedResourceResponseSystemDataOutput { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

func (o AzureTrafficCollectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AzureTrafficCollectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AzureTrafficCollectorOutput) VirtualHub() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *AzureTrafficCollector) ResourceReferenceResponsePtrOutput { return v.VirtualHub }).(ResourceReferenceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureTrafficCollectorOutput{})
}
