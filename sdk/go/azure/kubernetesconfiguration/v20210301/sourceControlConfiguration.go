


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceControlConfiguration struct {
	pulumi.CustomResourceState

	ComplianceStatus               ComplianceStatusResponseOutput          `pulumi:"complianceStatus"`
	ConfigurationProtectedSettings pulumi.StringMapOutput                  `pulumi:"configurationProtectedSettings"`
	EnableHelmOperator             pulumi.BoolPtrOutput                    `pulumi:"enableHelmOperator"`
	HelmOperatorProperties         HelmOperatorPropertiesResponsePtrOutput `pulumi:"helmOperatorProperties"`
	Name                           pulumi.StringOutput                     `pulumi:"name"`
	OperatorInstanceName           pulumi.StringPtrOutput                  `pulumi:"operatorInstanceName"`
	OperatorNamespace              pulumi.StringPtrOutput                  `pulumi:"operatorNamespace"`
	OperatorParams                 pulumi.StringPtrOutput                  `pulumi:"operatorParams"`
	OperatorScope                  pulumi.StringPtrOutput                  `pulumi:"operatorScope"`
	OperatorType                   pulumi.StringPtrOutput                  `pulumi:"operatorType"`
	ProvisioningState              pulumi.StringOutput                     `pulumi:"provisioningState"`
	RepositoryPublicKey            pulumi.StringOutput                     `pulumi:"repositoryPublicKey"`
	RepositoryUrl                  pulumi.StringPtrOutput                  `pulumi:"repositoryUrl"`
	SshKnownHostsContents          pulumi.StringPtrOutput                  `pulumi:"sshKnownHostsContents"`
	SystemData                     SystemDataResponseOutput                `pulumi:"systemData"`
	Type                           pulumi.StringOutput                     `pulumi:"type"`
}


func NewSourceControlConfiguration(ctx *pulumi.Context,
	name string, args *SourceControlConfigurationArgs, opts ...pulumi.ResourceOption) (*SourceControlConfiguration, error) {
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
	if isZero(args.OperatorNamespace) {
		args.OperatorNamespace = pulumi.StringPtr("default")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20191101preview:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20200701preview:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20201001preview:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20210501preview:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20211101preview:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220101preview:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220301:SourceControlConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220701:SourceControlConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource SourceControlConfiguration
	err := ctx.RegisterResource("azure-native:kubernetesconfiguration/v20210301:SourceControlConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSourceControlConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceControlConfigurationState, opts ...pulumi.ResourceOption) (*SourceControlConfiguration, error) {
	var resource SourceControlConfiguration
	err := ctx.ReadResource("azure-native:kubernetesconfiguration/v20210301:SourceControlConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sourceControlConfigurationState struct {
}

type SourceControlConfigurationState struct {
}

func (SourceControlConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlConfigurationState)(nil)).Elem()
}

type sourceControlConfigurationArgs struct {
	ClusterName                    string                  `pulumi:"clusterName"`
	ClusterResourceName            string                  `pulumi:"clusterResourceName"`
	ClusterRp                      string                  `pulumi:"clusterRp"`
	ConfigurationProtectedSettings map[string]string       `pulumi:"configurationProtectedSettings"`
	EnableHelmOperator             *bool                   `pulumi:"enableHelmOperator"`
	HelmOperatorProperties         *HelmOperatorProperties `pulumi:"helmOperatorProperties"`
	OperatorInstanceName           *string                 `pulumi:"operatorInstanceName"`
	OperatorNamespace              *string                 `pulumi:"operatorNamespace"`
	OperatorParams                 *string                 `pulumi:"operatorParams"`
	OperatorScope                  *string                 `pulumi:"operatorScope"`
	OperatorType                   *string                 `pulumi:"operatorType"`
	RepositoryUrl                  *string                 `pulumi:"repositoryUrl"`
	ResourceGroupName              string                  `pulumi:"resourceGroupName"`
	SourceControlConfigurationName *string                 `pulumi:"sourceControlConfigurationName"`
	SshKnownHostsContents          *string                 `pulumi:"sshKnownHostsContents"`
}


type SourceControlConfigurationArgs struct {
	ClusterName                    pulumi.StringInput
	ClusterResourceName            pulumi.StringInput
	ClusterRp                      pulumi.StringInput
	ConfigurationProtectedSettings pulumi.StringMapInput
	EnableHelmOperator             pulumi.BoolPtrInput
	HelmOperatorProperties         HelmOperatorPropertiesPtrInput
	OperatorInstanceName           pulumi.StringPtrInput
	OperatorNamespace              pulumi.StringPtrInput
	OperatorParams                 pulumi.StringPtrInput
	OperatorScope                  pulumi.StringPtrInput
	OperatorType                   pulumi.StringPtrInput
	RepositoryUrl                  pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	SourceControlConfigurationName pulumi.StringPtrInput
	SshKnownHostsContents          pulumi.StringPtrInput
}

func (SourceControlConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlConfigurationArgs)(nil)).Elem()
}

type SourceControlConfigurationInput interface {
	pulumi.Input

	ToSourceControlConfigurationOutput() SourceControlConfigurationOutput
	ToSourceControlConfigurationOutputWithContext(ctx context.Context) SourceControlConfigurationOutput
}

func (*SourceControlConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlConfiguration)(nil)).Elem()
}

func (i *SourceControlConfiguration) ToSourceControlConfigurationOutput() SourceControlConfigurationOutput {
	return i.ToSourceControlConfigurationOutputWithContext(context.Background())
}

func (i *SourceControlConfiguration) ToSourceControlConfigurationOutputWithContext(ctx context.Context) SourceControlConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlConfigurationOutput)
}

type SourceControlConfigurationOutput struct{ *pulumi.OutputState }

func (SourceControlConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlConfiguration)(nil)).Elem()
}

func (o SourceControlConfigurationOutput) ToSourceControlConfigurationOutput() SourceControlConfigurationOutput {
	return o
}

func (o SourceControlConfigurationOutput) ToSourceControlConfigurationOutputWithContext(ctx context.Context) SourceControlConfigurationOutput {
	return o
}

func (o SourceControlConfigurationOutput) ComplianceStatus() ComplianceStatusResponseOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) ComplianceStatusResponseOutput { return v.ComplianceStatus }).(ComplianceStatusResponseOutput)
}

func (o SourceControlConfigurationOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringMapOutput { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

func (o SourceControlConfigurationOutput) EnableHelmOperator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.BoolPtrOutput { return v.EnableHelmOperator }).(pulumi.BoolPtrOutput)
}

func (o SourceControlConfigurationOutput) HelmOperatorProperties() HelmOperatorPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) HelmOperatorPropertiesResponsePtrOutput {
		return v.HelmOperatorProperties
	}).(HelmOperatorPropertiesResponsePtrOutput)
}

func (o SourceControlConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SourceControlConfigurationOutput) OperatorInstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringPtrOutput { return v.OperatorInstanceName }).(pulumi.StringPtrOutput)
}

func (o SourceControlConfigurationOutput) OperatorNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringPtrOutput { return v.OperatorNamespace }).(pulumi.StringPtrOutput)
}

func (o SourceControlConfigurationOutput) OperatorParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringPtrOutput { return v.OperatorParams }).(pulumi.StringPtrOutput)
}

func (o SourceControlConfigurationOutput) OperatorScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringPtrOutput { return v.OperatorScope }).(pulumi.StringPtrOutput)
}

func (o SourceControlConfigurationOutput) OperatorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringPtrOutput { return v.OperatorType }).(pulumi.StringPtrOutput)
}

func (o SourceControlConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SourceControlConfigurationOutput) RepositoryPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringOutput { return v.RepositoryPublicKey }).(pulumi.StringOutput)
}

func (o SourceControlConfigurationOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringPtrOutput { return v.RepositoryUrl }).(pulumi.StringPtrOutput)
}

func (o SourceControlConfigurationOutput) SshKnownHostsContents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringPtrOutput { return v.SshKnownHostsContents }).(pulumi.StringPtrOutput)
}

func (o SourceControlConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SourceControlConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControlConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SourceControlConfigurationOutput{})
}
