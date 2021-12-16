


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FluxConfiguration struct {
	pulumi.CustomResourceState

	Bucket                         BucketDefinitionResponsePtrOutput         `pulumi:"bucket"`
	ComplianceState                pulumi.StringOutput                       `pulumi:"complianceState"`
	ConfigurationProtectedSettings pulumi.StringMapOutput                    `pulumi:"configurationProtectedSettings"`
	ErrorMessage                   pulumi.StringOutput                       `pulumi:"errorMessage"`
	GitRepository                  GitRepositoryDefinitionResponsePtrOutput  `pulumi:"gitRepository"`
	Kustomizations                 KustomizationDefinitionResponseMapOutput  `pulumi:"kustomizations"`
	LastSourceUpdatedAt            pulumi.StringOutput                       `pulumi:"lastSourceUpdatedAt"`
	LastSourceUpdatedCommitId      pulumi.StringOutput                       `pulumi:"lastSourceUpdatedCommitId"`
	Name                           pulumi.StringOutput                       `pulumi:"name"`
	Namespace                      pulumi.StringPtrOutput                    `pulumi:"namespace"`
	ProvisioningState              pulumi.StringOutput                       `pulumi:"provisioningState"`
	RepositoryPublicKey            pulumi.StringOutput                       `pulumi:"repositoryPublicKey"`
	Scope                          pulumi.StringPtrOutput                    `pulumi:"scope"`
	SourceKind                     pulumi.StringPtrOutput                    `pulumi:"sourceKind"`
	Statuses                       ObjectStatusDefinitionResponseArrayOutput `pulumi:"statuses"`
	Suspend                        pulumi.BoolPtrOutput                      `pulumi:"suspend"`
	SystemData                     SystemDataResponseOutput                  `pulumi:"systemData"`
	Type                           pulumi.StringOutput                       `pulumi:"type"`
}


func NewFluxConfiguration(ctx *pulumi.Context,
	name string, args *FluxConfigurationArgs, opts ...pulumi.ResourceOption) (*FluxConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceName'")
	}
	if args.ClusterRp == nil {
		return nil, errors.New("invalid value for required argument 'ClusterRp'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	bucketApplier := func(v BucketDefinition) *BucketDefinition { return v.Defaults() }
	if args.Bucket != nil {
		args.Bucket = args.Bucket.ToBucketDefinitionPtrOutput().Elem().ApplyT(bucketApplier).(BucketDefinitionPtrOutput)
	}
	gitRepositoryApplier := func(v GitRepositoryDefinition) *GitRepositoryDefinition { return v.Defaults() }
	if args.GitRepository != nil {
		args.GitRepository = args.GitRepository.ToGitRepositoryDefinitionPtrOutput().Elem().ApplyT(gitRepositoryApplier).(GitRepositoryDefinitionPtrOutput)
	}
	if isZero(args.Namespace) {
		args.Namespace = pulumi.StringPtr("default")
	}
	if isZero(args.SourceKind) {
		args.SourceKind = pulumi.StringPtr("GitRepository")
	}
	if isZero(args.Suspend) {
		args.Suspend = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration:FluxConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20211101preview:FluxConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource FluxConfiguration
	err := ctx.RegisterResource("azure-native:kubernetesconfiguration/v20220101preview:FluxConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFluxConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FluxConfigurationState, opts ...pulumi.ResourceOption) (*FluxConfiguration, error) {
	var resource FluxConfiguration
	err := ctx.ReadResource("azure-native:kubernetesconfiguration/v20220101preview:FluxConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fluxConfigurationState struct {
}

type FluxConfigurationState struct {
}

func (FluxConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*fluxConfigurationState)(nil)).Elem()
}

type fluxConfigurationArgs struct {
	Bucket                         *BucketDefinition                  `pulumi:"bucket"`
	ClusterName                    string                             `pulumi:"clusterName"`
	ClusterResourceName            string                             `pulumi:"clusterResourceName"`
	ClusterRp                      string                             `pulumi:"clusterRp"`
	ConfigurationProtectedSettings map[string]string                  `pulumi:"configurationProtectedSettings"`
	FluxConfigurationName          *string                            `pulumi:"fluxConfigurationName"`
	GitRepository                  *GitRepositoryDefinition           `pulumi:"gitRepository"`
	Kustomizations                 map[string]KustomizationDefinition `pulumi:"kustomizations"`
	Namespace                      *string                            `pulumi:"namespace"`
	ResourceGroupName              string                             `pulumi:"resourceGroupName"`
	Scope                          *string                            `pulumi:"scope"`
	SourceKind                     *string                            `pulumi:"sourceKind"`
	Suspend                        *bool                              `pulumi:"suspend"`
}


type FluxConfigurationArgs struct {
	Bucket                         BucketDefinitionPtrInput
	ClusterName                    pulumi.StringInput
	ClusterResourceName            pulumi.StringInput
	ClusterRp                      pulumi.StringInput
	ConfigurationProtectedSettings pulumi.StringMapInput
	FluxConfigurationName          pulumi.StringPtrInput
	GitRepository                  GitRepositoryDefinitionPtrInput
	Kustomizations                 KustomizationDefinitionMapInput
	Namespace                      pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	Scope                          pulumi.StringPtrInput
	SourceKind                     pulumi.StringPtrInput
	Suspend                        pulumi.BoolPtrInput
}

func (FluxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fluxConfigurationArgs)(nil)).Elem()
}

type FluxConfigurationInput interface {
	pulumi.Input

	ToFluxConfigurationOutput() FluxConfigurationOutput
	ToFluxConfigurationOutputWithContext(ctx context.Context) FluxConfigurationOutput
}

func (*FluxConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**FluxConfiguration)(nil)).Elem()
}

func (i *FluxConfiguration) ToFluxConfigurationOutput() FluxConfigurationOutput {
	return i.ToFluxConfigurationOutputWithContext(context.Background())
}

func (i *FluxConfiguration) ToFluxConfigurationOutputWithContext(ctx context.Context) FluxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FluxConfigurationOutput)
}

type FluxConfigurationOutput struct{ *pulumi.OutputState }

func (FluxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FluxConfiguration)(nil)).Elem()
}

func (o FluxConfigurationOutput) ToFluxConfigurationOutput() FluxConfigurationOutput {
	return o
}

func (o FluxConfigurationOutput) ToFluxConfigurationOutputWithContext(ctx context.Context) FluxConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FluxConfigurationOutput{})
}
