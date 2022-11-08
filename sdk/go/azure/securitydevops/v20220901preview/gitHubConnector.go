


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GitHubConnector struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties GitHubConnectorPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewGitHubConnector(ctx *pulumi.Context,
	name string, args *GitHubConnectorArgs, opts ...pulumi.ResourceOption) (*GitHubConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securitydevops:GitHubConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource GitHubConnector
	err := ctx.RegisterResource("azure-native:securitydevops/v20220901preview:GitHubConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGitHubConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitHubConnectorState, opts ...pulumi.ResourceOption) (*GitHubConnector, error) {
	var resource GitHubConnector
	err := ctx.ReadResource("azure-native:securitydevops/v20220901preview:GitHubConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gitHubConnectorState struct {
}

type GitHubConnectorState struct {
}

func (GitHubConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitHubConnectorState)(nil)).Elem()
}

type gitHubConnectorArgs struct {
	GitHubConnectorName *string                    `pulumi:"gitHubConnectorName"`
	Location            *string                    `pulumi:"location"`
	Properties          *GitHubConnectorProperties `pulumi:"properties"`
	ResourceGroupName   string                     `pulumi:"resourceGroupName"`
	Tags                map[string]string          `pulumi:"tags"`
}


type GitHubConnectorArgs struct {
	GitHubConnectorName pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	Properties          GitHubConnectorPropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (GitHubConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitHubConnectorArgs)(nil)).Elem()
}

type GitHubConnectorInput interface {
	pulumi.Input

	ToGitHubConnectorOutput() GitHubConnectorOutput
	ToGitHubConnectorOutputWithContext(ctx context.Context) GitHubConnectorOutput
}

func (*GitHubConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubConnector)(nil)).Elem()
}

func (i *GitHubConnector) ToGitHubConnectorOutput() GitHubConnectorOutput {
	return i.ToGitHubConnectorOutputWithContext(context.Background())
}

func (i *GitHubConnector) ToGitHubConnectorOutputWithContext(ctx context.Context) GitHubConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubConnectorOutput)
}

type GitHubConnectorOutput struct{ *pulumi.OutputState }

func (GitHubConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubConnector)(nil)).Elem()
}

func (o GitHubConnectorOutput) ToGitHubConnectorOutput() GitHubConnectorOutput {
	return o
}

func (o GitHubConnectorOutput) ToGitHubConnectorOutputWithContext(ctx context.Context) GitHubConnectorOutput {
	return o
}

func (o GitHubConnectorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GitHubConnector) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GitHubConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GitHubConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GitHubConnectorOutput) Properties() GitHubConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v *GitHubConnector) GitHubConnectorPropertiesResponseOutput { return v.Properties }).(GitHubConnectorPropertiesResponseOutput)
}

func (o GitHubConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GitHubConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GitHubConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GitHubConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GitHubConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GitHubConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GitHubConnectorOutput{})
}
