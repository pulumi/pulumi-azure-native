


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpoint struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                     `pulumi:"etag"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties PrivateEndpointPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewPrivateEndpoint(ctx *pulumi.Context,
	name string, args *PrivateEndpointArgs, opts ...pulumi.ResourceOption) (*PrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:streamanalytics:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20200301:PrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpoint
	err := ctx.RegisterResource("azure-native:streamanalytics/v20200301preview:PrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointState, opts ...pulumi.ResourceOption) (*PrivateEndpoint, error) {
	var resource PrivateEndpoint
	err := ctx.ReadResource("azure-native:streamanalytics/v20200301preview:PrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointState struct {
}

type PrivateEndpointState struct {
}

func (PrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointState)(nil)).Elem()
}

type privateEndpointArgs struct {
	ClusterName         string                     `pulumi:"clusterName"`
	PrivateEndpointName *string                    `pulumi:"privateEndpointName"`
	Properties          *PrivateEndpointProperties `pulumi:"properties"`
	ResourceGroupName   string                     `pulumi:"resourceGroupName"`
}


type PrivateEndpointArgs struct {
	ClusterName         pulumi.StringInput
	PrivateEndpointName pulumi.StringPtrInput
	Properties          PrivateEndpointPropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointArgs)(nil)).Elem()
}

type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput
}

func (*PrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil))
}

func (i *PrivateEndpoint) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i *PrivateEndpoint) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil))
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
}
