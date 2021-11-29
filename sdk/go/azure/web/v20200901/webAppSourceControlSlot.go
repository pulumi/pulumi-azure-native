


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSourceControlSlot struct {
	pulumi.CustomResourceState

	Branch                    pulumi.StringPtrOutput   `pulumi:"branch"`
	DeploymentRollbackEnabled pulumi.BoolPtrOutput     `pulumi:"deploymentRollbackEnabled"`
	IsGitHubAction            pulumi.BoolPtrOutput     `pulumi:"isGitHubAction"`
	IsManualIntegration       pulumi.BoolPtrOutput     `pulumi:"isManualIntegration"`
	IsMercurial               pulumi.BoolPtrOutput     `pulumi:"isMercurial"`
	Kind                      pulumi.StringPtrOutput   `pulumi:"kind"`
	Name                      pulumi.StringOutput      `pulumi:"name"`
	RepoUrl                   pulumi.StringPtrOutput   `pulumi:"repoUrl"`
	SystemData                SystemDataResponseOutput `pulumi:"systemData"`
	Type                      pulumi.StringOutput      `pulumi:"type"`
}


func NewWebAppSourceControlSlot(ctx *pulumi.Context,
	name string, args *WebAppSourceControlSlotArgs, opts ...pulumi.ResourceOption) (*WebAppSourceControlSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSourceControlSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSourceControlSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSourceControlSlot
	err := ctx.RegisterResource("azure-native:web/v20200901:WebAppSourceControlSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSourceControlSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSourceControlSlotState, opts ...pulumi.ResourceOption) (*WebAppSourceControlSlot, error) {
	var resource WebAppSourceControlSlot
	err := ctx.ReadResource("azure-native:web/v20200901:WebAppSourceControlSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSourceControlSlotState struct {
}

type WebAppSourceControlSlotState struct {
}

func (WebAppSourceControlSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlSlotState)(nil)).Elem()
}

type webAppSourceControlSlotArgs struct {
	Branch                    *string `pulumi:"branch"`
	DeploymentRollbackEnabled *bool   `pulumi:"deploymentRollbackEnabled"`
	IsGitHubAction            *bool   `pulumi:"isGitHubAction"`
	IsManualIntegration       *bool   `pulumi:"isManualIntegration"`
	IsMercurial               *bool   `pulumi:"isMercurial"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	RepoUrl                   *string `pulumi:"repoUrl"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	Slot                      string  `pulumi:"slot"`
}


type WebAppSourceControlSlotArgs struct {
	Branch                    pulumi.StringPtrInput
	DeploymentRollbackEnabled pulumi.BoolPtrInput
	IsGitHubAction            pulumi.BoolPtrInput
	IsManualIntegration       pulumi.BoolPtrInput
	IsMercurial               pulumi.BoolPtrInput
	Kind                      pulumi.StringPtrInput
	Name                      pulumi.StringInput
	RepoUrl                   pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	Slot                      pulumi.StringInput
}

func (WebAppSourceControlSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSourceControlSlotArgs)(nil)).Elem()
}

type WebAppSourceControlSlotInput interface {
	pulumi.Input

	ToWebAppSourceControlSlotOutput() WebAppSourceControlSlotOutput
	ToWebAppSourceControlSlotOutputWithContext(ctx context.Context) WebAppSourceControlSlotOutput
}

func (*WebAppSourceControlSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSourceControlSlot)(nil))
}

func (i *WebAppSourceControlSlot) ToWebAppSourceControlSlotOutput() WebAppSourceControlSlotOutput {
	return i.ToWebAppSourceControlSlotOutputWithContext(context.Background())
}

func (i *WebAppSourceControlSlot) ToWebAppSourceControlSlotOutputWithContext(ctx context.Context) WebAppSourceControlSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSourceControlSlotOutput)
}

type WebAppSourceControlSlotOutput struct{ *pulumi.OutputState }

func (WebAppSourceControlSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSourceControlSlot)(nil))
}

func (o WebAppSourceControlSlotOutput) ToWebAppSourceControlSlotOutput() WebAppSourceControlSlotOutput {
	return o
}

func (o WebAppSourceControlSlotOutput) ToWebAppSourceControlSlotOutputWithContext(ctx context.Context) WebAppSourceControlSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppSourceControlSlotOutput{})
}
