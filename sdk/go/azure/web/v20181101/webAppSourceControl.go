


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSourceControl struct {
	pulumi.CustomResourceState

	Branch                    pulumi.StringPtrOutput `pulumi:"branch"`
	DeploymentRollbackEnabled pulumi.BoolPtrOutput   `pulumi:"deploymentRollbackEnabled"`
	IsManualIntegration       pulumi.BoolPtrOutput   `pulumi:"isManualIntegration"`
	IsMercurial               pulumi.BoolPtrOutput   `pulumi:"isMercurial"`
	Kind                      pulumi.StringPtrOutput `pulumi:"kind"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	RepoUrl                   pulumi.StringPtrOutput `pulumi:"repoUrl"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSourceControl"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:WebAppSourceControl"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSourceControl
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppSourceControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSourceControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSourceControlState, opts ...pulumi.ResourceOption) (*WebAppSourceControl, error) {
	var resource WebAppSourceControl
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppSourceControl", name, id, state, &resource, opts...)
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
	Branch                    *string `pulumi:"branch"`
	DeploymentRollbackEnabled *bool   `pulumi:"deploymentRollbackEnabled"`
	IsManualIntegration       *bool   `pulumi:"isManualIntegration"`
	IsMercurial               *bool   `pulumi:"isMercurial"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	RepoUrl                   *string `pulumi:"repoUrl"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
}


type WebAppSourceControlArgs struct {
	Branch                    pulumi.StringPtrInput
	DeploymentRollbackEnabled pulumi.BoolPtrInput
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
	return reflect.TypeOf((*WebAppSourceControl)(nil))
}

func (i *WebAppSourceControl) ToWebAppSourceControlOutput() WebAppSourceControlOutput {
	return i.ToWebAppSourceControlOutputWithContext(context.Background())
}

func (i *WebAppSourceControl) ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSourceControlOutput)
}

type WebAppSourceControlOutput struct{ *pulumi.OutputState }

func (WebAppSourceControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppSourceControl)(nil))
}

func (o WebAppSourceControlOutput) ToWebAppSourceControlOutput() WebAppSourceControlOutput {
	return o
}

func (o WebAppSourceControlOutput) ToWebAppSourceControlOutputWithContext(ctx context.Context) WebAppSourceControlOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAppSourceControlInput)(nil)).Elem(), &WebAppSourceControl{})
	pulumi.RegisterOutputType(WebAppSourceControlOutput{})
}
