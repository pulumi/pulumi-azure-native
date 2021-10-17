


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Deployment struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                     `pulumi:"tags"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:resources/v20200601:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20151101:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20151101:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160201:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20160201:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160701:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20160701:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160901:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20160901:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20170510:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20170510:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180201:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20180201:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180501:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20180501:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190301:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190501:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190510:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190701:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20190801:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20191001:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20200801:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20201001:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20210101:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:Deployment"),
		},
		{
			Type: pulumi.String("azure-nextgen:resources/v20210401:Deployment"),
		},
	})
	opts = append(opts, aliases)
	var resource Deployment
	err := ctx.RegisterResource("azure-native:resources/v20200601:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("azure-native:resources/v20200601:Deployment", name, id, state, &resource, opts...)
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
	DeploymentName    *string              `pulumi:"deploymentName"`
	Location          *string              `pulumi:"location"`
	Properties        DeploymentProperties `pulumi:"properties"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
	Tags              map[string]string    `pulumi:"tags"`
}


type DeploymentArgs struct {
	DeploymentName    pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        DeploymentPropertiesInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
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
	return reflect.TypeOf((*Deployment)(nil))
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
}
