


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BatchDeployment struct {
	pulumi.CustomResourceState

	BatchDeploymentDetails BatchDeploymentResponseOutput           `pulumi:"batchDeploymentDetails"`
	Identity               ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind                   pulumi.StringPtrOutput                  `pulumi:"kind"`
	Location               pulumi.StringOutput                     `pulumi:"location"`
	Name                   pulumi.StringOutput                     `pulumi:"name"`
	Sku                    SkuResponsePtrOutput                    `pulumi:"sku"`
	SystemData             SystemDataResponseOutput                `pulumi:"systemData"`
	Tags                   pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                   pulumi.StringOutput                     `pulumi:"type"`
}


func NewBatchDeployment(ctx *pulumi.Context,
	name string, args *BatchDeploymentArgs, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BatchDeploymentDetails == nil {
		return nil, errors.New("invalid value for required argument 'BatchDeploymentDetails'")
	}
	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.BatchDeploymentDetails = args.BatchDeploymentDetails.ToBatchDeploymentTypeOutput().ApplyT(func(v BatchDeploymentType) BatchDeploymentType { return *v.Defaults() }).(BatchDeploymentTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:BatchDeployment"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:BatchDeployment"),
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
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220201preview:BatchDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBatchDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchDeploymentState, opts ...pulumi.ResourceOption) (*BatchDeployment, error) {
	var resource BatchDeployment
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220201preview:BatchDeployment", name, id, state, &resource, opts...)
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
	BatchDeploymentDetails BatchDeploymentType     `pulumi:"batchDeploymentDetails"`
	DeploymentName         *string                 `pulumi:"deploymentName"`
	EndpointName           string                  `pulumi:"endpointName"`
	Identity               *ManagedServiceIdentity `pulumi:"identity"`
	Kind                   *string                 `pulumi:"kind"`
	Location               *string                 `pulumi:"location"`
	ResourceGroupName      string                  `pulumi:"resourceGroupName"`
	Sku                    *Sku                    `pulumi:"sku"`
	Tags                   map[string]string       `pulumi:"tags"`
	WorkspaceName          string                  `pulumi:"workspaceName"`
}


type BatchDeploymentArgs struct {
	BatchDeploymentDetails BatchDeploymentTypeInput
	DeploymentName         pulumi.StringPtrInput
	EndpointName           pulumi.StringInput
	Identity               ManagedServiceIdentityPtrInput
	Kind                   pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SkuPtrInput
	Tags                   pulumi.StringMapInput
	WorkspaceName          pulumi.StringInput
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

func (o BatchDeploymentOutput) BatchDeploymentDetails() BatchDeploymentResponseOutput {
	return o.ApplyT(func(v *BatchDeployment) BatchDeploymentResponseOutput { return v.BatchDeploymentDetails }).(BatchDeploymentResponseOutput)
}

func (o BatchDeploymentOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *BatchDeployment) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
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

func (o BatchDeploymentOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *BatchDeployment) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
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
