


package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FluxConfiguration struct {
	pulumi.CustomResourceState

	AzureBlob                      AzureBlobDefinitionResponsePtrOutput      `pulumi:"azureBlob"`
	Bucket                         BucketDefinitionResponsePtrOutput         `pulumi:"bucket"`
	ComplianceState                pulumi.StringOutput                       `pulumi:"complianceState"`
	ConfigurationProtectedSettings pulumi.StringMapOutput                    `pulumi:"configurationProtectedSettings"`
	ErrorMessage                   pulumi.StringOutput                       `pulumi:"errorMessage"`
	GitRepository                  GitRepositoryDefinitionResponsePtrOutput  `pulumi:"gitRepository"`
	Kustomizations                 KustomizationDefinitionResponseMapOutput  `pulumi:"kustomizations"`
	Name                           pulumi.StringOutput                       `pulumi:"name"`
	Namespace                      pulumi.StringPtrOutput                    `pulumi:"namespace"`
	ProvisioningState              pulumi.StringOutput                       `pulumi:"provisioningState"`
	RepositoryPublicKey            pulumi.StringOutput                       `pulumi:"repositoryPublicKey"`
	Scope                          pulumi.StringPtrOutput                    `pulumi:"scope"`
	SourceKind                     pulumi.StringPtrOutput                    `pulumi:"sourceKind"`
	SourceSyncedCommitId           pulumi.StringOutput                       `pulumi:"sourceSyncedCommitId"`
	SourceUpdatedAt                pulumi.StringOutput                       `pulumi:"sourceUpdatedAt"`
	StatusUpdatedAt                pulumi.StringOutput                       `pulumi:"statusUpdatedAt"`
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
	if args.AzureBlob != nil {
		args.AzureBlob = args.AzureBlob.ToAzureBlobDefinitionPtrOutput().ApplyT(func(v *AzureBlobDefinition) *AzureBlobDefinition { return v.Defaults() }).(AzureBlobDefinitionPtrOutput)
	}
	if args.Bucket != nil {
		args.Bucket = args.Bucket.ToBucketDefinitionPtrOutput().ApplyT(func(v *BucketDefinition) *BucketDefinition { return v.Defaults() }).(BucketDefinitionPtrOutput)
	}
	if args.GitRepository != nil {
		args.GitRepository = args.GitRepository.ToGitRepositoryDefinitionPtrOutput().ApplyT(func(v *GitRepositoryDefinition) *GitRepositoryDefinition { return v.Defaults() }).(GitRepositoryDefinitionPtrOutput)
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
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220101preview:FluxConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220301:FluxConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource FluxConfiguration
	err := ctx.RegisterResource("azure-native:kubernetesconfiguration/v20220701:FluxConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFluxConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FluxConfigurationState, opts ...pulumi.ResourceOption) (*FluxConfiguration, error) {
	var resource FluxConfiguration
	err := ctx.ReadResource("azure-native:kubernetesconfiguration/v20220701:FluxConfiguration", name, id, state, &resource, opts...)
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
	AzureBlob                      *AzureBlobDefinition               `pulumi:"azureBlob"`
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
	AzureBlob                      AzureBlobDefinitionPtrInput
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

func (o FluxConfigurationOutput) AzureBlob() AzureBlobDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *FluxConfiguration) AzureBlobDefinitionResponsePtrOutput { return v.AzureBlob }).(AzureBlobDefinitionResponsePtrOutput)
}

func (o FluxConfigurationOutput) Bucket() BucketDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *FluxConfiguration) BucketDefinitionResponsePtrOutput { return v.Bucket }).(BucketDefinitionResponsePtrOutput)
}

func (o FluxConfigurationOutput) ComplianceState() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.ComplianceState }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringMapOutput { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

func (o FluxConfigurationOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) GitRepository() GitRepositoryDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *FluxConfiguration) GitRepositoryDefinitionResponsePtrOutput { return v.GitRepository }).(GitRepositoryDefinitionResponsePtrOutput)
}

func (o FluxConfigurationOutput) Kustomizations() KustomizationDefinitionResponseMapOutput {
	return o.ApplyT(func(v *FluxConfiguration) KustomizationDefinitionResponseMapOutput { return v.Kustomizations }).(KustomizationDefinitionResponseMapOutput)
}

func (o FluxConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o FluxConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) RepositoryPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.RepositoryPublicKey }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o FluxConfigurationOutput) SourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringPtrOutput { return v.SourceKind }).(pulumi.StringPtrOutput)
}

func (o FluxConfigurationOutput) SourceSyncedCommitId() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.SourceSyncedCommitId }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) SourceUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.SourceUpdatedAt }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) StatusUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.StatusUpdatedAt }).(pulumi.StringOutput)
}

func (o FluxConfigurationOutput) Statuses() ObjectStatusDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *FluxConfiguration) ObjectStatusDefinitionResponseArrayOutput { return v.Statuses }).(ObjectStatusDefinitionResponseArrayOutput)
}

func (o FluxConfigurationOutput) Suspend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.BoolPtrOutput { return v.Suspend }).(pulumi.BoolPtrOutput)
}

func (o FluxConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FluxConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FluxConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FluxConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FluxConfigurationOutput{})
}
