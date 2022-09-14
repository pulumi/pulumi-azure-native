


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerAppsSourceControl struct {
	pulumi.CustomResourceState

	Branch                    pulumi.StringPtrOutput                     `pulumi:"branch"`
	GithubActionConfiguration GithubActionConfigurationResponsePtrOutput `pulumi:"githubActionConfiguration"`
	Name                      pulumi.StringOutput                        `pulumi:"name"`
	OperationState            pulumi.StringOutput                        `pulumi:"operationState"`
	RepoUrl                   pulumi.StringPtrOutput                     `pulumi:"repoUrl"`
	SystemData                SystemDataResponseOutput                   `pulumi:"systemData"`
	Type                      pulumi.StringOutput                        `pulumi:"type"`
}


func NewContainerAppsSourceControl(ctx *pulumi.Context,
	name string, args *ContainerAppsSourceControlArgs, opts ...pulumi.ResourceOption) (*ContainerAppsSourceControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerAppName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerAppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ContainerAppsSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220101preview:ContainerAppsSourceControl"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerAppsSourceControl
	err := ctx.RegisterResource("azure-native:app/v20220301:ContainerAppsSourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContainerAppsSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppsSourceControlState, opts ...pulumi.ResourceOption) (*ContainerAppsSourceControl, error) {
	var resource ContainerAppsSourceControl
	err := ctx.ReadResource("azure-native:app/v20220301:ContainerAppsSourceControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type containerAppsSourceControlState struct {
}

type ContainerAppsSourceControlState struct {
}

func (ContainerAppsSourceControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsSourceControlState)(nil)).Elem()
}

type containerAppsSourceControlArgs struct {
	Branch                    *string                    `pulumi:"branch"`
	ContainerAppName          string                     `pulumi:"containerAppName"`
	GithubActionConfiguration *GithubActionConfiguration `pulumi:"githubActionConfiguration"`
	RepoUrl                   *string                    `pulumi:"repoUrl"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	SourceControlName         *string                    `pulumi:"sourceControlName"`
}


type ContainerAppsSourceControlArgs struct {
	Branch                    pulumi.StringPtrInput
	ContainerAppName          pulumi.StringInput
	GithubActionConfiguration GithubActionConfigurationPtrInput
	RepoUrl                   pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	SourceControlName         pulumi.StringPtrInput
}

func (ContainerAppsSourceControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppsSourceControlArgs)(nil)).Elem()
}

type ContainerAppsSourceControlInput interface {
	pulumi.Input

	ToContainerAppsSourceControlOutput() ContainerAppsSourceControlOutput
	ToContainerAppsSourceControlOutputWithContext(ctx context.Context) ContainerAppsSourceControlOutput
}

func (*ContainerAppsSourceControl) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsSourceControl)(nil)).Elem()
}

func (i *ContainerAppsSourceControl) ToContainerAppsSourceControlOutput() ContainerAppsSourceControlOutput {
	return i.ToContainerAppsSourceControlOutputWithContext(context.Background())
}

func (i *ContainerAppsSourceControl) ToContainerAppsSourceControlOutputWithContext(ctx context.Context) ContainerAppsSourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppsSourceControlOutput)
}

type ContainerAppsSourceControlOutput struct{ *pulumi.OutputState }

func (ContainerAppsSourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsSourceControl)(nil)).Elem()
}

func (o ContainerAppsSourceControlOutput) ToContainerAppsSourceControlOutput() ContainerAppsSourceControlOutput {
	return o
}

func (o ContainerAppsSourceControlOutput) ToContainerAppsSourceControlOutputWithContext(ctx context.Context) ContainerAppsSourceControlOutput {
	return o
}

func (o ContainerAppsSourceControlOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringPtrOutput { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsSourceControlOutput) GithubActionConfiguration() GithubActionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) GithubActionConfigurationResponsePtrOutput {
		return v.GithubActionConfiguration
	}).(GithubActionConfigurationResponsePtrOutput)
}

func (o ContainerAppsSourceControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerAppsSourceControlOutput) OperationState() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringOutput { return v.OperationState }).(pulumi.StringOutput)
}

func (o ContainerAppsSourceControlOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringPtrOutput { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsSourceControlOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ContainerAppsSourceControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAppsSourceControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAppsSourceControlOutput{})
}
