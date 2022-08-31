


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Deployment struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                `pulumi:"etag"`
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties DeploymentPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput           `pulumi:"systemData"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:cognitiveservices/v20211001:Deployment"),
		},
	})
	opts = append(opts, aliases)
	var resource Deployment
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20220301:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("azure-native:cognitiveservices/v20220301:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deploymentState struct {
}

type DeploymentState struct {
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	AccountName       string                `pulumi:"accountName"`
	DeploymentName    *string               `pulumi:"deploymentName"`
	Properties        *DeploymentProperties `pulumi:"properties"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
}


type DeploymentArgs struct {
	AccountName       pulumi.StringInput
	DeploymentName    pulumi.StringPtrInput
	Properties        DeploymentPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func (o DeploymentOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentOutput) Properties() DeploymentPropertiesResponseOutput {
	return o.ApplyT(func(v *Deployment) DeploymentPropertiesResponseOutput { return v.Properties }).(DeploymentPropertiesResponseOutput)
}

func (o DeploymentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Deployment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
}
