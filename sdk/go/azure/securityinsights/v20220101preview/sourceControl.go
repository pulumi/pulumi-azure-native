


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SourceControl struct {
	pulumi.CustomResourceState

	ContentTypes           pulumi.StringArrayOutput                `pulumi:"contentTypes"`
	Description            pulumi.StringPtrOutput                  `pulumi:"description"`
	DisplayName            pulumi.StringOutput                     `pulumi:"displayName"`
	Etag                   pulumi.StringPtrOutput                  `pulumi:"etag"`
	LastDeploymentInfo     DeploymentInfoResponsePtrOutput         `pulumi:"lastDeploymentInfo"`
	Name                   pulumi.StringOutput                     `pulumi:"name"`
	RepoType               pulumi.StringOutput                     `pulumi:"repoType"`
	Repository             RepositoryResponseOutput                `pulumi:"repository"`
	RepositoryResourceInfo RepositoryResourceInfoResponsePtrOutput `pulumi:"repositoryResourceInfo"`
	SystemData             SystemDataResponseOutput                `pulumi:"systemData"`
	Type                   pulumi.StringOutput                     `pulumi:"type"`
	Version                pulumi.StringPtrOutput                  `pulumi:"version"`
}


func NewSourceControl(ctx *pulumi.Context,
	name string, args *SourceControlArgs, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentTypes == nil {
		return nil, errors.New("invalid value for required argument 'ContentTypes'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.RepoType == nil {
		return nil, errors.New("invalid value for required argument 'RepoType'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:SourceControl"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:SourceControl"),
		},
	})
	opts = append(opts, aliases)
	var resource SourceControl
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:SourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceControlState, opts ...pulumi.ResourceOption) (*SourceControl, error) {
	var resource SourceControl
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:SourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sourceControlState struct {
}

type SourceControlState struct {
}

func (SourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlState)(nil)).Elem()
}

type sourceControlArgs struct {
	ContentTypes           []string                `pulumi:"contentTypes"`
	Description            *string                 `pulumi:"description"`
	DisplayName            string                  `pulumi:"displayName"`
	Id                     *string                 `pulumi:"id"`
	LastDeploymentInfo     *DeploymentInfo         `pulumi:"lastDeploymentInfo"`
	RepoType               string                  `pulumi:"repoType"`
	Repository             Repository              `pulumi:"repository"`
	RepositoryResourceInfo *RepositoryResourceInfo `pulumi:"repositoryResourceInfo"`
	ResourceGroupName      string                  `pulumi:"resourceGroupName"`
	SourceControlId        *string                 `pulumi:"sourceControlId"`
	Version                *string                 `pulumi:"version"`
	WorkspaceName          string                  `pulumi:"workspaceName"`
}


type SourceControlArgs struct {
	ContentTypes           pulumi.StringArrayInput
	Description            pulumi.StringPtrInput
	DisplayName            pulumi.StringInput
	Id                     pulumi.StringPtrInput
	LastDeploymentInfo     DeploymentInfoPtrInput
	RepoType               pulumi.StringInput
	Repository             RepositoryInput
	RepositoryResourceInfo RepositoryResourceInfoPtrInput
	ResourceGroupName      pulumi.StringInput
	SourceControlId        pulumi.StringPtrInput
	Version                pulumi.StringPtrInput
	WorkspaceName          pulumi.StringInput
}

func (SourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceControlArgs)(nil)).Elem()
}

type SourceControlInput interface {
	pulumi.Input

	ToSourceControlOutput() SourceControlOutput
	ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput
}

func (*SourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControl)(nil)).Elem()
}

func (i *SourceControl) ToSourceControlOutput() SourceControlOutput {
	return i.ToSourceControlOutputWithContext(context.Background())
}

func (i *SourceControl) ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlOutput)
}

type SourceControlOutput struct{ *pulumi.OutputState }

func (SourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControl)(nil)).Elem()
}

func (o SourceControlOutput) ToSourceControlOutput() SourceControlOutput {
	return o
}

func (o SourceControlOutput) ToSourceControlOutputWithContext(ctx context.Context) SourceControlOutput {
	return o
}

func (o SourceControlOutput) ContentTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringArrayOutput { return v.ContentTypes }).(pulumi.StringArrayOutput)
}

func (o SourceControlOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SourceControlOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SourceControlOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SourceControlOutput) LastDeploymentInfo() DeploymentInfoResponsePtrOutput {
	return o.ApplyT(func(v *SourceControl) DeploymentInfoResponsePtrOutput { return v.LastDeploymentInfo }).(DeploymentInfoResponsePtrOutput)
}

func (o SourceControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SourceControlOutput) RepoType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.RepoType }).(pulumi.StringOutput)
}

func (o SourceControlOutput) Repository() RepositoryResponseOutput {
	return o.ApplyT(func(v *SourceControl) RepositoryResponseOutput { return v.Repository }).(RepositoryResponseOutput)
}

func (o SourceControlOutput) RepositoryResourceInfo() RepositoryResourceInfoResponsePtrOutput {
	return o.ApplyT(func(v *SourceControl) RepositoryResourceInfoResponsePtrOutput { return v.RepositoryResourceInfo }).(RepositoryResourceInfoResponsePtrOutput)
}

func (o SourceControlOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SourceControl) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SourceControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SourceControlOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControl) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SourceControlOutput{})
}
