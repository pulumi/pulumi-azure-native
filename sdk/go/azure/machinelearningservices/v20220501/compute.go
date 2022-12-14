


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Compute struct {
	pulumi.CustomResourceState

	Identity   ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput                  `pulumi:"location"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties pulumi.AnyOutput                        `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                    `pulumi:"sku"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewCompute(ctx *pulumi.Context,
	name string, args *ComputeArgs, opts ...pulumi.ResourceOption) (*Compute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20180301preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20181119:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190501:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190601:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20191101:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220101preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Compute"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:Compute"),
		},
	})
	opts = append(opts, aliases)
	var resource Compute
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:Compute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCompute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeState, opts ...pulumi.ResourceOption) (*Compute, error) {
	var resource Compute
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:Compute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type computeState struct {
}

type ComputeState struct {
}

func (ComputeState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeState)(nil)).Elem()
}

type computeArgs struct {
	ComputeName       *string                 `pulumi:"computeName"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	Location          *string                 `pulumi:"location"`
	Properties        interface{}             `pulumi:"properties"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Sku               *Sku                    `pulumi:"sku"`
	Tags              map[string]string       `pulumi:"tags"`
	WorkspaceName     string                  `pulumi:"workspaceName"`
}


type ComputeArgs struct {
	ComputeName       pulumi.StringPtrInput
	Identity          ManagedServiceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
}

func (ComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeArgs)(nil)).Elem()
}

type ComputeInput interface {
	pulumi.Input

	ToComputeOutput() ComputeOutput
	ToComputeOutputWithContext(ctx context.Context) ComputeOutput
}

func (*Compute) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (i *Compute) ToComputeOutput() ComputeOutput {
	return i.ToComputeOutputWithContext(context.Background())
}

func (i *Compute) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeOutput)
}

type ComputeOutput struct{ *pulumi.OutputState }

func (ComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Compute)(nil)).Elem()
}

func (o ComputeOutput) ToComputeOutput() ComputeOutput {
	return o
}

func (o ComputeOutput) ToComputeOutputWithContext(ctx context.Context) ComputeOutput {
	return o
}

func (o ComputeOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Compute) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ComputeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ComputeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Compute) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ComputeOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Compute) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ComputeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Compute) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ComputeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ComputeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Compute) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ComputeOutput{})
}
