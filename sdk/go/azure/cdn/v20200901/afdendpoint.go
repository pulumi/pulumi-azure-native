


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AFDEndpoint struct {
	pulumi.CustomResourceState

	DeploymentStatus             pulumi.StringOutput      `pulumi:"deploymentStatus"`
	EnabledState                 pulumi.StringPtrOutput   `pulumi:"enabledState"`
	HostName                     pulumi.StringOutput      `pulumi:"hostName"`
	Location                     pulumi.StringOutput      `pulumi:"location"`
	Name                         pulumi.StringOutput      `pulumi:"name"`
	OriginResponseTimeoutSeconds pulumi.IntPtrOutput      `pulumi:"originResponseTimeoutSeconds"`
	ProvisioningState            pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData                   SystemDataResponseOutput `pulumi:"systemData"`
	Tags                         pulumi.StringMapOutput   `pulumi:"tags"`
	Type                         pulumi.StringOutput      `pulumi:"type"`
}


func NewAFDEndpoint(ctx *pulumi.Context,
	name string, args *AFDEndpointArgs, opts ...pulumi.ResourceOption) (*AFDEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200901:AFDEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn:AFDEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn:AFDEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource AFDEndpoint
	err := ctx.RegisterResource("azure-native:cdn/v20200901:AFDEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAFDEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDEndpointState, opts ...pulumi.ResourceOption) (*AFDEndpoint, error) {
	var resource AFDEndpoint
	err := ctx.ReadResource("azure-native:cdn/v20200901:AFDEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type afdendpointState struct {
}

type AFDEndpointState struct {
}

func (AFDEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdendpointState)(nil)).Elem()
}

type afdendpointArgs struct {
	EnabledState                 *string           `pulumi:"enabledState"`
	EndpointName                 *string           `pulumi:"endpointName"`
	Location                     *string           `pulumi:"location"`
	OriginResponseTimeoutSeconds *int              `pulumi:"originResponseTimeoutSeconds"`
	ProfileName                  string            `pulumi:"profileName"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type AFDEndpointArgs struct {
	EnabledState                 pulumi.StringPtrInput
	EndpointName                 pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	OriginResponseTimeoutSeconds pulumi.IntPtrInput
	ProfileName                  pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (AFDEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdendpointArgs)(nil)).Elem()
}

type AFDEndpointInput interface {
	pulumi.Input

	ToAFDEndpointOutput() AFDEndpointOutput
	ToAFDEndpointOutputWithContext(ctx context.Context) AFDEndpointOutput
}

func (*AFDEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDEndpoint)(nil))
}

func (i *AFDEndpoint) ToAFDEndpointOutput() AFDEndpointOutput {
	return i.ToAFDEndpointOutputWithContext(context.Background())
}

func (i *AFDEndpoint) ToAFDEndpointOutputWithContext(ctx context.Context) AFDEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDEndpointOutput)
}

type AFDEndpointOutput struct{ *pulumi.OutputState }

func (AFDEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDEndpoint)(nil))
}

func (o AFDEndpointOutput) ToAFDEndpointOutput() AFDEndpointOutput {
	return o
}

func (o AFDEndpointOutput) ToAFDEndpointOutputWithContext(ctx context.Context) AFDEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AFDEndpointOutput{})
}
