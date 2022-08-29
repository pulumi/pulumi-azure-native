


package web

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSourceControl struct {
	pulumi.CustomResourceState

	Branch                    pulumi.StringPtrOutput                     `pulumi:"branch"`
	DeploymentRollbackEnabled pulumi.BoolPtrOutput                       `pulumi:"deploymentRollbackEnabled"`
	GitHubActionConfiguration GitHubActionConfigurationResponsePtrOutput `pulumi:"gitHubActionConfiguration"`
	IsGitHubAction            pulumi.BoolPtrOutput                       `pulumi:"isGitHubAction"`
	IsManualIntegration       pulumi.BoolPtrOutput                       `pulumi:"isManualIntegration"`
	IsMercurial               pulumi.BoolPtrOutput                       `pulumi:"isMercurial"`
	Kind                      pulumi.StringPtrOutput                     `pulumi:"kind"`
	Name                      pulumi.StringOutput                        `pulumi:"name"`
	RepoUrl                   pulumi.StringPtrOutput                     `pulumi:"repoUrl"`
	Type                      pulumi.StringOutput                        `pulumi:"type"`
}


func NewWebAppSourceControl(ctx *pulumi.Context,
	name string, args *WebAppSourceControlArgs, opts ...pulumi.ResourceOption) (*WebAppSourceControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSourceControl"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSourceControl
	err := ctx.RegisterResource("azure-native:web:WebAppSourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSourceControlState, opts ...pulumi.ResourceOption) (*WebAppSourceControl, error) {
	var resource WebAppSourceControl
	err := ctx.ReadResource("azure-native:web:WebAppSourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSourceControlState struct {
}

type WebAppSourceControlState struct {
}

func (WebAppSourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlState)(nil)).Elem()
}

type webAppSourceControlArgs struct {
	Branch                    *string                    `pulumi:"branch"`
	DeploymentRollbackEnabled *bool                      `pulumi:"deploymentRollbackEnabled"`
	GitHubActionConfiguration *GitHubActionConfiguration `pulumi:"gitHubActionConfiguration"`
	IsGitHubAction            *bool                      `pulumi:"isGitHubAction"`
	IsManualIntegration       *bool                      `pulumi:"isManualIntegration"`
	IsMercurial               *bool                      `pulumi:"isMercurial"`
	Kind                      *string                    `pulumi:"kind"`
	Name                      string                     `pulumi:"name"`
	RepoUrl                   *string                    `pulumi:"repoUrl"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
}


type WebAppSourceControlArgs struct {
	Branch                    pulumi.StringPtrInput
	DeploymentRollbackEnabled pulumi.BoolPtrInput
	GitHubActionConfiguration GitHubActionConfigurationPtrInput
	IsGitHubAction            pulumi.BoolPtrInput
	IsManualIntegration       pulumi.BoolPtrInput
	IsMercurial               pulumi.BoolPtrInput
	Kind                      pulumi.StringPtrInput
	Name                      pulumi.StringInput
	RepoUrl                   pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
}

func (WebAppSourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlArgs)(nil)).Elem()
}

type WebAppSourceControlInput interface {
	pulumi.Input

	ToWebAppSourceControlOutput() WebAppSourceControlOutput
	ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput
}

func (*WebAppSourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSourceControl)(nil)).Elem()
}

func (i *WebAppSourceControl) ToWebAppSourceControlOutput() WebAppSourceControlOutput {
	return i.ToWebAppSourceControlOutputWithContext(context.Background())
}

func (i *WebAppSourceControl) ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSourceControlOutput)
}

type WebAppSourceControlOutput struct{ *pulumi.OutputState }

func (WebAppSourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSourceControl)(nil)).Elem()
}

func (o WebAppSourceControlOutput) ToWebAppSourceControlOutput() WebAppSourceControlOutput {
	return o
}

func (o WebAppSourceControlOutput) ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput {
	return o
}

func (o WebAppSourceControlOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o WebAppSourceControlOutput) DeploymentRollbackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.DeploymentRollbackEnabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppSourceControlOutput) GitHubActionConfiguration() GitHubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) GitHubActionConfigurationResponsePtrOutput {
		return v.GitHubActionConfiguration
	}).(GitHubActionConfigurationResponsePtrOutput)
}

func (o WebAppSourceControlOutput) IsGitHubAction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.IsGitHubAction }).(pulumi.BoolPtrOutput)
}

func (o WebAppSourceControlOutput) IsManualIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.IsManualIntegration }).(pulumi.BoolPtrOutput)
}

func (o WebAppSourceControlOutput) IsMercurial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.BoolPtrOutput { return v.IsMercurial }).(pulumi.BoolPtrOutput)
}

func (o WebAppSourceControlOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppSourceControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppSourceControlOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringPtrOutput { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o WebAppSourceControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSourceControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSourceControlOutput{})
}
