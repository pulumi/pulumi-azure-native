


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotSecuritySolution struct {
	pulumi.CustomResourceState

	AutoDiscoveredResources      pulumi.StringArrayOutput                                 `pulumi:"autoDiscoveredResources"`
	DisabledDataSources          pulumi.StringArrayOutput                                 `pulumi:"disabledDataSources"`
	DisplayName                  pulumi.StringOutput                                      `pulumi:"displayName"`
	Export                       pulumi.StringArrayOutput                                 `pulumi:"export"`
	IotHubs                      pulumi.StringArrayOutput                                 `pulumi:"iotHubs"`
	Location                     pulumi.StringPtrOutput                                   `pulumi:"location"`
	Name                         pulumi.StringOutput                                      `pulumi:"name"`
	RecommendationsConfiguration RecommendationConfigurationPropertiesResponseArrayOutput `pulumi:"recommendationsConfiguration"`
	Status                       pulumi.StringPtrOutput                                   `pulumi:"status"`
	Tags                         pulumi.StringMapOutput                                   `pulumi:"tags"`
	Type                         pulumi.StringOutput                                      `pulumi:"type"`
	UserDefinedResources         UserDefinedResourcesPropertiesResponsePtrOutput          `pulumi:"userDefinedResources"`
	Workspace                    pulumi.StringOutput                                      `pulumi:"workspace"`
}


func NewIotSecuritySolution(ctx *pulumi.Context,
	name string, args *IotSecuritySolutionArgs, opts ...pulumi.ResourceOption) (*IotSecuritySolution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IotHubs == nil {
		return nil, errors.New("invalid value for required argument 'IotHubs'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Workspace == nil {
		return nil, errors.New("invalid value for required argument 'Workspace'")
	}
	if args.Status == nil {
		args.Status = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:security/v20170801preview:IotSecuritySolution"),
		},
		{
			Type: pulumi.String("azure-native:security:IotSecuritySolution"),
		},
		{
			Type: pulumi.String("azure-nextgen:security:IotSecuritySolution"),
		},
		{
			Type: pulumi.String("azure-native:security/v20190801:IotSecuritySolution"),
		},
		{
			Type: pulumi.String("azure-nextgen:security/v20190801:IotSecuritySolution"),
		},
	})
	opts = append(opts, aliases)
	var resource IotSecuritySolution
	err := ctx.RegisterResource("azure-native:security/v20170801preview:IotSecuritySolution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotSecuritySolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotSecuritySolutionState, opts ...pulumi.ResourceOption) (*IotSecuritySolution, error) {
	var resource IotSecuritySolution
	err := ctx.ReadResource("azure-native:security/v20170801preview:IotSecuritySolution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotSecuritySolutionState struct {
}

type IotSecuritySolutionState struct {
}

func (IotSecuritySolutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotSecuritySolutionState)(nil)).Elem()
}

type iotSecuritySolutionArgs struct {
	DisabledDataSources          []string                                `pulumi:"disabledDataSources"`
	DisplayName                  string                                  `pulumi:"displayName"`
	Export                       []string                                `pulumi:"export"`
	IotHubs                      []string                                `pulumi:"iotHubs"`
	Location                     *string                                 `pulumi:"location"`
	RecommendationsConfiguration []RecommendationConfigurationProperties `pulumi:"recommendationsConfiguration"`
	ResourceGroupName            string                                  `pulumi:"resourceGroupName"`
	SolutionName                 *string                                 `pulumi:"solutionName"`
	Status                       *string                                 `pulumi:"status"`
	Tags                         map[string]string                       `pulumi:"tags"`
	UserDefinedResources         *UserDefinedResourcesProperties         `pulumi:"userDefinedResources"`
	Workspace                    string                                  `pulumi:"workspace"`
}


type IotSecuritySolutionArgs struct {
	DisabledDataSources          pulumi.StringArrayInput
	DisplayName                  pulumi.StringInput
	Export                       pulumi.StringArrayInput
	IotHubs                      pulumi.StringArrayInput
	Location                     pulumi.StringPtrInput
	RecommendationsConfiguration RecommendationConfigurationPropertiesArrayInput
	ResourceGroupName            pulumi.StringInput
	SolutionName                 pulumi.StringPtrInput
	Status                       pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
	UserDefinedResources         UserDefinedResourcesPropertiesPtrInput
	Workspace                    pulumi.StringInput
}

func (IotSecuritySolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotSecuritySolutionArgs)(nil)).Elem()
}

type IotSecuritySolutionInput interface {
	pulumi.Input

	ToIotSecuritySolutionOutput() IotSecuritySolutionOutput
	ToIotSecuritySolutionOutputWithContext(ctx context.Context) IotSecuritySolutionOutput
}

func (*IotSecuritySolution) ElementType() reflect.Type {
	return reflect.TypeOf((*IotSecuritySolution)(nil))
}

func (i *IotSecuritySolution) ToIotSecuritySolutionOutput() IotSecuritySolutionOutput {
	return i.ToIotSecuritySolutionOutputWithContext(context.Background())
}

func (i *IotSecuritySolution) ToIotSecuritySolutionOutputWithContext(ctx context.Context) IotSecuritySolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotSecuritySolutionOutput)
}

type IotSecuritySolutionOutput struct{ *pulumi.OutputState }

func (IotSecuritySolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotSecuritySolution)(nil))
}

func (o IotSecuritySolutionOutput) ToIotSecuritySolutionOutput() IotSecuritySolutionOutput {
	return o
}

func (o IotSecuritySolutionOutput) ToIotSecuritySolutionOutputWithContext(ctx context.Context) IotSecuritySolutionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IotSecuritySolutionOutput{})
}
