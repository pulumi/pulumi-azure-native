


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BatchEndpoint struct {
	pulumi.CustomResourceState

	Identity   ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	Kind       pulumi.StringPtrOutput            `pulumi:"kind"`
	Location   pulumi.StringOutput               `pulumi:"location"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties BatchEndpointResponseOutput       `pulumi:"properties"`
	SystemData SystemDataResponseOutput          `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput            `pulumi:"tags"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewBatchEndpoint(ctx *pulumi.Context,
	name string, args *BatchEndpointArgs, opts ...pulumi.ResourceOption) (*BatchEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:machinelearningservices:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:BatchEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource BatchEndpoint
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210301preview:BatchEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBatchEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchEndpointState, opts ...pulumi.ResourceOption) (*BatchEndpoint, error) {
	var resource BatchEndpoint
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210301preview:BatchEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type batchEndpointState struct {
}

type BatchEndpointState struct {
}

func (BatchEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchEndpointState)(nil)).Elem()
}

type batchEndpointArgs struct {
	EndpointName      *string           `pulumi:"endpointName"`
	Identity          *ResourceIdentity `pulumi:"identity"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Properties        BatchEndpointType `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	WorkspaceName     string            `pulumi:"workspaceName"`
}


type BatchEndpointArgs struct {
	EndpointName      pulumi.StringPtrInput
	Identity          ResourceIdentityPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        BatchEndpointTypeInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
}

func (BatchEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchEndpointArgs)(nil)).Elem()
}

type BatchEndpointInput interface {
	pulumi.Input

	ToBatchEndpointOutput() BatchEndpointOutput
	ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput
}

func (*BatchEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchEndpoint)(nil)).Elem()
}

func (i *BatchEndpoint) ToBatchEndpointOutput() BatchEndpointOutput {
	return i.ToBatchEndpointOutputWithContext(context.Background())
}

func (i *BatchEndpoint) ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchEndpointOutput)
}

type BatchEndpointOutput struct{ *pulumi.OutputState }

func (BatchEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchEndpoint)(nil)).Elem()
}

func (o BatchEndpointOutput) ToBatchEndpointOutput() BatchEndpointOutput {
	return o
}

func (o BatchEndpointOutput) ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput {
	return o
}

func (o BatchEndpointOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *BatchEndpoint) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

func (o BatchEndpointOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BatchEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BatchEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BatchEndpointOutput) Properties() BatchEndpointResponseOutput {
	return o.ApplyT(func(v *BatchEndpoint) BatchEndpointResponseOutput { return v.Properties }).(BatchEndpointResponseOutput)
}

func (o BatchEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BatchEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BatchEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BatchEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BatchEndpointOutput{})
}
