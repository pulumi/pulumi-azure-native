


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AKSService struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties pulumi.AnyOutput          `pulumi:"properties"`
	Sku        SkuResponsePtrOutput      `pulumi:"sku"`
	SystemData SystemDataResponseOutput  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewAKSService(ctx *pulumi.Context,
	name string, args *AKSServiceArgs, opts ...pulumi.ResourceOption) (*AKSService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeType == nil {
		return nil, errors.New("invalid value for required argument 'ComputeType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ComputeType = pulumi.String("AKS")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:AKSService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:AKSService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:AKSService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:AKSService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:AKSService"),
		},
	})
	opts = append(opts, aliases)
	var resource AKSService
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210101:AKSService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAKSService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AKSServiceState, opts ...pulumi.ResourceOption) (*AKSService, error) {
	var resource AKSService
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210101:AKSService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type aksserviceState struct {
}

type AKSServiceState struct {
}

func (AKSServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*aksserviceState)(nil)).Elem()
}

type aksserviceArgs struct {
	AadAuthEnabled                    *bool                                             `pulumi:"aadAuthEnabled"`
	AppInsightsEnabled                *bool                                             `pulumi:"appInsightsEnabled"`
	AuthEnabled                       *bool                                             `pulumi:"authEnabled"`
	AutoScaler                        *AKSServiceCreateRequestAutoScaler                `pulumi:"autoScaler"`
	ComputeName                       *string                                           `pulumi:"computeName"`
	ComputeType                       string                                            `pulumi:"computeType"`
	ContainerResourceRequirements     *ContainerResourceRequirements                    `pulumi:"containerResourceRequirements"`
	DataCollection                    *AKSServiceCreateRequestDataCollection            `pulumi:"dataCollection"`
	Description                       *string                                           `pulumi:"description"`
	EnvironmentImageRequest           *CreateServiceRequestEnvironmentImageRequest      `pulumi:"environmentImageRequest"`
	IsDefault                         *bool                                             `pulumi:"isDefault"`
	Keys                              *CreateServiceRequestKeys                         `pulumi:"keys"`
	KvTags                            map[string]string                                 `pulumi:"kvTags"`
	LivenessProbeRequirements         *AKSServiceCreateRequestLivenessProbeRequirements `pulumi:"livenessProbeRequirements"`
	Location                          *string                                           `pulumi:"location"`
	MaxConcurrentRequestsPerContainer *int                                              `pulumi:"maxConcurrentRequestsPerContainer"`
	MaxQueueWaitMs                    *int                                              `pulumi:"maxQueueWaitMs"`
	Namespace                         *string                                           `pulumi:"namespace"`
	NumReplicas                       *int                                              `pulumi:"numReplicas"`
	Properties                        map[string]string                                 `pulumi:"properties"`
	ResourceGroupName                 string                                            `pulumi:"resourceGroupName"`
	ScoringTimeoutMs                  *int                                              `pulumi:"scoringTimeoutMs"`
	ServiceName                       *string                                           `pulumi:"serviceName"`
	TrafficPercentile                 *float64                                          `pulumi:"trafficPercentile"`
	Type                              *string                                           `pulumi:"type"`
	WorkspaceName                     string                                            `pulumi:"workspaceName"`
}


type AKSServiceArgs struct {
	AadAuthEnabled                    pulumi.BoolPtrInput
	AppInsightsEnabled                pulumi.BoolPtrInput
	AuthEnabled                       pulumi.BoolPtrInput
	AutoScaler                        AKSServiceCreateRequestAutoScalerPtrInput
	ComputeName                       pulumi.StringPtrInput
	ComputeType                       pulumi.StringInput
	ContainerResourceRequirements     ContainerResourceRequirementsPtrInput
	DataCollection                    AKSServiceCreateRequestDataCollectionPtrInput
	Description                       pulumi.StringPtrInput
	EnvironmentImageRequest           CreateServiceRequestEnvironmentImageRequestPtrInput
	IsDefault                         pulumi.BoolPtrInput
	Keys                              CreateServiceRequestKeysPtrInput
	KvTags                            pulumi.StringMapInput
	LivenessProbeRequirements         AKSServiceCreateRequestLivenessProbeRequirementsPtrInput
	Location                          pulumi.StringPtrInput
	MaxConcurrentRequestsPerContainer pulumi.IntPtrInput
	MaxQueueWaitMs                    pulumi.IntPtrInput
	Namespace                         pulumi.StringPtrInput
	NumReplicas                       pulumi.IntPtrInput
	Properties                        pulumi.StringMapInput
	ResourceGroupName                 pulumi.StringInput
	ScoringTimeoutMs                  pulumi.IntPtrInput
	ServiceName                       pulumi.StringPtrInput
	TrafficPercentile                 pulumi.Float64PtrInput
	Type                              pulumi.StringPtrInput
	WorkspaceName                     pulumi.StringInput
}

func (AKSServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aksserviceArgs)(nil)).Elem()
}

type AKSServiceInput interface {
	pulumi.Input

	ToAKSServiceOutput() AKSServiceOutput
	ToAKSServiceOutputWithContext(ctx context.Context) AKSServiceOutput
}

func (*AKSService) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSService)(nil)).Elem()
}

func (i *AKSService) ToAKSServiceOutput() AKSServiceOutput {
	return i.ToAKSServiceOutputWithContext(context.Background())
}

func (i *AKSService) ToAKSServiceOutputWithContext(ctx context.Context) AKSServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSServiceOutput)
}

type AKSServiceOutput struct{ *pulumi.OutputState }

func (AKSServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSService)(nil)).Elem()
}

func (o AKSServiceOutput) ToAKSServiceOutput() AKSServiceOutput {
	return o
}

func (o AKSServiceOutput) ToAKSServiceOutputWithContext(ctx context.Context) AKSServiceOutput {
	return o
}

func (o AKSServiceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *AKSService) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o AKSServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AKSServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AKSService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AKSServiceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *AKSService) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o AKSServiceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *AKSService) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o AKSServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AKSService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AKSServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AKSService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AKSServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AKSService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AKSServiceOutput{})
}
