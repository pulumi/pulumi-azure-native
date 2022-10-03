


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BatchDeployment struct {
	pulumi.CustomResourceState

	Identity   ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind       pulumi.StringPtrOutput            `pulumi:"kind"`
	Location   pulumi.StringOutput               `pulumi:"location"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties BatchDeploymentResponseOutput     `pulumi:"properties"`
	SystemData SystemDataResponseOutput          `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput            `pulumi:"tags"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewBatchDeployment(ctx *pulumi.Context,
	name string, args *BatchDeploymentArgs, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:BatchDeployment"),
		},
	})
	opts = append(opts, aliases)
	var resource BatchDeployment
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:BatchDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBatchDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchDeploymentState, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	var resource BatchDeployment
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:BatchDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type batchDeploymentState struct {
}

type BatchDeploymentState struct {
}

func (BatchDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchDeploymentState)(nil)).Elem()
}

type batchDeploymentArgs struct {
	DeploymentName    *string             `pulumi:"deploymentName"`
	EndpointName      string              `pulumi:"endpointName"`
	Identity          *ResourceIdentity   `pulumi:"identity"`
	Kind              *string             `pulumi:"kind"`
	Location          *string             `pulumi:"location"`
	Properties        BatchDeploymentType `pulumi:"properties"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Tags              map[string]string   `pulumi:"tags"`
	WorkspaceName     string              `pulumi:"workspaceName"`
}


type BatchDeploymentArgs struct {
	DeploymentName    pulumi.StringPtrInput
	EndpointName      pulumi.StringInput
	Identity          ResourceIdentityPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        BatchDeploymentTypeInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
}

func (BatchDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchDeploymentArgs)(nil)).Elem()
}

type BatchDeploymentInput interface {
	pulumi.Input

	ToBatchDeploymentOutput() BatchDeploymentOutput
	ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput
}

func (*BatchDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchDeployment)(nil)).Elem()
}

func (i *BatchDeployment) ToBatchDeploymentOutput() BatchDeploymentOutput {
	return i.ToBatchDeploymentOutputWithContext(context.Background())
}

func (i *BatchDeployment) ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchDeploymentOutput)
}

type BatchDeploymentOutput struct{ *pulumi.OutputState }

func (BatchDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchDeployment)(nil)).Elem()
}

func (o BatchDeploymentOutput) ToBatchDeploymentOutput() BatchDeploymentOutput {
	return o
}

func (o BatchDeploymentOutput) ToBatchDeploymentOutputWithContext(ctx context.Context) BatchDeploymentOutput {
	return o
}

func (o BatchDeploymentOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *BatchDeployment) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o BatchDeploymentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BatchDeploymentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BatchDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BatchDeploymentOutput) Properties() BatchDeploymentResponseOutput {
	return o.ApplyT(func(v *BatchDeployment) BatchDeploymentResponseOutput { return v.Properties }).(BatchDeploymentResponseOutput)
}

func (o BatchDeploymentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BatchDeployment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BatchDeploymentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BatchDeploymentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchDeployment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BatchDeploymentOutput{})
}
