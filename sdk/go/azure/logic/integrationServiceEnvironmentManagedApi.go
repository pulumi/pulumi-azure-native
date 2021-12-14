


package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationServiceEnvironmentManagedApi struct {
	pulumi.CustomResourceState

	ApiDefinitionUrl              pulumi.StringOutput                                                          `pulumi:"apiDefinitionUrl"`
	ApiDefinitions                ApiResourceDefinitionsResponseOutput                                         `pulumi:"apiDefinitions"`
	BackendService                ApiResourceBackendServiceResponseOutput                                      `pulumi:"backendService"`
	Capabilities                  pulumi.StringArrayOutput                                                     `pulumi:"capabilities"`
	Category                      pulumi.StringOutput                                                          `pulumi:"category"`
	ConnectionParameters          pulumi.MapOutput                                                             `pulumi:"connectionParameters"`
	DeploymentParameters          IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput `pulumi:"deploymentParameters"`
	GeneralInformation            ApiResourceGeneralInformationResponseOutput                                  `pulumi:"generalInformation"`
	IntegrationServiceEnvironment ResourceReferenceResponsePtrOutput                                           `pulumi:"integrationServiceEnvironment"`
	Location                      pulumi.StringPtrOutput                                                       `pulumi:"location"`
	Metadata                      ApiResourceMetadataResponseOutput                                            `pulumi:"metadata"`
	Name                          pulumi.StringOutput                                                          `pulumi:"name"`
	Policies                      ApiResourcePoliciesResponseOutput                                            `pulumi:"policies"`
	ProvisioningState             pulumi.StringOutput                                                          `pulumi:"provisioningState"`
	RuntimeUrls                   pulumi.StringArrayOutput                                                     `pulumi:"runtimeUrls"`
	Tags                          pulumi.StringMapOutput                                                       `pulumi:"tags"`
	Type                          pulumi.StringOutput                                                          `pulumi:"type"`
}


func NewIntegrationServiceEnvironmentManagedApi(ctx *pulumi.Context,
	name string, args *IntegrationServiceEnvironmentManagedApiArgs, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironmentManagedApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationServiceEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationServiceEnvironmentName'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationServiceEnvironmentManagedApi"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationServiceEnvironmentManagedApi
	err := ctx.RegisterResource("azure-native:logic:IntegrationServiceEnvironmentManagedApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationServiceEnvironmentManagedApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationServiceEnvironmentManagedApiState, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironmentManagedApi, error) {
	var resource IntegrationServiceEnvironmentManagedApi
	err := ctx.ReadResource("azure-native:logic:IntegrationServiceEnvironmentManagedApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationServiceEnvironmentManagedApiState struct {
}

type IntegrationServiceEnvironmentManagedApiState struct {
}

func (IntegrationServiceEnvironmentManagedApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentManagedApiState)(nil)).Elem()
}

type integrationServiceEnvironmentManagedApiArgs struct {
	ApiName                           *string                                                      `pulumi:"apiName"`
	DeploymentParameters              *IntegrationServiceEnvironmentManagedApiDeploymentParameters `pulumi:"deploymentParameters"`
	IntegrationServiceEnvironment     *ResourceReference                                           `pulumi:"integrationServiceEnvironment"`
	IntegrationServiceEnvironmentName string                                                       `pulumi:"integrationServiceEnvironmentName"`
	Location                          *string                                                      `pulumi:"location"`
	ResourceGroup                     string                                                       `pulumi:"resourceGroup"`
	Tags                              map[string]string                                            `pulumi:"tags"`
}


type IntegrationServiceEnvironmentManagedApiArgs struct {
	ApiName                           pulumi.StringPtrInput
	DeploymentParameters              IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrInput
	IntegrationServiceEnvironment     ResourceReferencePtrInput
	IntegrationServiceEnvironmentName pulumi.StringInput
	Location                          pulumi.StringPtrInput
	ResourceGroup                     pulumi.StringInput
	Tags                              pulumi.StringMapInput
}

func (IntegrationServiceEnvironmentManagedApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentManagedApiArgs)(nil)).Elem()
}

type IntegrationServiceEnvironmentManagedApiInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentManagedApiOutput() IntegrationServiceEnvironmentManagedApiOutput
	ToIntegrationServiceEnvironmentManagedApiOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiOutput
}

func (*IntegrationServiceEnvironmentManagedApi) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentManagedApi)(nil)).Elem()
}

func (i *IntegrationServiceEnvironmentManagedApi) ToIntegrationServiceEnvironmentManagedApiOutput() IntegrationServiceEnvironmentManagedApiOutput {
	return i.ToIntegrationServiceEnvironmentManagedApiOutputWithContext(context.Background())
}

func (i *IntegrationServiceEnvironmentManagedApi) ToIntegrationServiceEnvironmentManagedApiOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentManagedApiOutput)
}

type IntegrationServiceEnvironmentManagedApiOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentManagedApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentManagedApi)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentManagedApiOutput) ToIntegrationServiceEnvironmentManagedApiOutput() IntegrationServiceEnvironmentManagedApiOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiOutput) ToIntegrationServiceEnvironmentManagedApiOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentManagedApiOutput{})
}
