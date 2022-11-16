


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Deployment struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties DeploymentResourcePropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                       `pulumi:"sku"`
	SystemData SystemDataResponseOutput                   `pulumi:"systemData"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToDeploymentResourcePropertiesPtrOutput().ApplyT(func(v *DeploymentResourceProperties) *DeploymentResourceProperties { return v.Defaults() }).(DeploymentResourcePropertiesPtrOutput)
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:Deployment"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Deployment"),
		},
	})
	opts = append(opts, aliases)
	var resource Deployment
	err := ctx.RegisterResource("azure-native:appplatform/v20220501preview:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("azure-native:appplatform/v20220501preview:Deployment", name, id, state, &resource, opts...)
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
	AppName           string                        `pulumi:"appName"`
	DeploymentName    *string                       `pulumi:"deploymentName"`
	Properties        *DeploymentResourceProperties `pulumi:"properties"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	ServiceName       string                        `pulumi:"serviceName"`
	Sku               *Sku                          `pulumi:"sku"`
}


type DeploymentArgs struct {
	AppName           pulumi.StringInput
	DeploymentName    pulumi.StringPtrInput
	Properties        DeploymentResourcePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Sku               SkuPtrInput
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

func (o DeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentOutput) Properties() DeploymentResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *Deployment) DeploymentResourcePropertiesResponseOutput { return v.Properties }).(DeploymentResourcePropertiesResponseOutput)
}

func (o DeploymentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Deployment) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
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
