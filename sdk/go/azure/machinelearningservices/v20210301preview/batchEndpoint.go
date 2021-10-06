


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
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:BatchEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:BatchEndpoint"),
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
	return reflect.TypeOf((*BatchEndpoint)(nil))
}

func (i *BatchEndpoint) ToBatchEndpointOutput() BatchEndpointOutput {
	return i.ToBatchEndpointOutputWithContext(context.Background())
}

func (i *BatchEndpoint) ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchEndpointOutput)
}

type BatchEndpointOutput struct{ *pulumi.OutputState }

func (BatchEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchEndpoint)(nil))
}

func (o BatchEndpointOutput) ToBatchEndpointOutput() BatchEndpointOutput {
	return o
}

func (o BatchEndpointOutput) ToBatchEndpointOutputWithContext(ctx context.Context) BatchEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BatchEndpointOutput{})
}
