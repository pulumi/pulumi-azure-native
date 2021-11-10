


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AFDOrigin struct {
	pulumi.CustomResourceState

	AzureOrigin               ResourceReferenceResponsePtrOutput                     `pulumi:"azureOrigin"`
	DeploymentStatus          pulumi.StringOutput                                    `pulumi:"deploymentStatus"`
	EnabledState              pulumi.StringPtrOutput                                 `pulumi:"enabledState"`
	HostName                  pulumi.StringOutput                                    `pulumi:"hostName"`
	HttpPort                  pulumi.IntPtrOutput                                    `pulumi:"httpPort"`
	HttpsPort                 pulumi.IntPtrOutput                                    `pulumi:"httpsPort"`
	Name                      pulumi.StringOutput                                    `pulumi:"name"`
	OriginHostHeader          pulumi.StringPtrOutput                                 `pulumi:"originHostHeader"`
	Priority                  pulumi.IntPtrOutput                                    `pulumi:"priority"`
	ProvisioningState         pulumi.StringOutput                                    `pulumi:"provisioningState"`
	SharedPrivateLinkResource SharedPrivateLinkResourcePropertiesResponseArrayOutput `pulumi:"sharedPrivateLinkResource"`
	SystemData                SystemDataResponseOutput                               `pulumi:"systemData"`
	Type                      pulumi.StringOutput                                    `pulumi:"type"`
	Weight                    pulumi.IntPtrOutput                                    `pulumi:"weight"`
}


func NewAFDOrigin(ctx *pulumi.Context,
	name string, args *AFDOriginArgs, opts ...pulumi.ResourceOption) (*AFDOrigin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.OriginGroupName == nil {
		return nil, errors.New("invalid value for required argument 'OriginGroupName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:AFDOrigin"),
		},
	})
	opts = append(opts, aliases)
	var resource AFDOrigin
	err := ctx.RegisterResource("azure-native:cdn/v20200901:AFDOrigin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAFDOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDOriginState, opts ...pulumi.ResourceOption) (*AFDOrigin, error) {
	var resource AFDOrigin
	err := ctx.ReadResource("azure-native:cdn/v20200901:AFDOrigin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type afdoriginState struct {
}

type AFDOriginState struct {
}

func (AFDOriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginState)(nil)).Elem()
}

type afdoriginArgs struct {
	AzureOrigin               *ResourceReference                    `pulumi:"azureOrigin"`
	EnabledState              *string                               `pulumi:"enabledState"`
	HostName                  string                                `pulumi:"hostName"`
	HttpPort                  *int                                  `pulumi:"httpPort"`
	HttpsPort                 *int                                  `pulumi:"httpsPort"`
	OriginGroupName           string                                `pulumi:"originGroupName"`
	OriginHostHeader          *string                               `pulumi:"originHostHeader"`
	OriginName                *string                               `pulumi:"originName"`
	Priority                  *int                                  `pulumi:"priority"`
	ProfileName               string                                `pulumi:"profileName"`
	ResourceGroupName         string                                `pulumi:"resourceGroupName"`
	SharedPrivateLinkResource []SharedPrivateLinkResourceProperties `pulumi:"sharedPrivateLinkResource"`
	Weight                    *int                                  `pulumi:"weight"`
}


type AFDOriginArgs struct {
	AzureOrigin               ResourceReferencePtrInput
	EnabledState              pulumi.StringPtrInput
	HostName                  pulumi.StringInput
	HttpPort                  pulumi.IntPtrInput
	HttpsPort                 pulumi.IntPtrInput
	OriginGroupName           pulumi.StringInput
	OriginHostHeader          pulumi.StringPtrInput
	OriginName                pulumi.StringPtrInput
	Priority                  pulumi.IntPtrInput
	ProfileName               pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	SharedPrivateLinkResource SharedPrivateLinkResourcePropertiesArrayInput
	Weight                    pulumi.IntPtrInput
}

func (AFDOriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginArgs)(nil)).Elem()
}

type AFDOriginInput interface {
	pulumi.Input

	ToAFDOriginOutput() AFDOriginOutput
	ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput
}

func (*AFDOrigin) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDOrigin)(nil))
}

func (i *AFDOrigin) ToAFDOriginOutput() AFDOriginOutput {
	return i.ToAFDOriginOutputWithContext(context.Background())
}

func (i *AFDOrigin) ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDOriginOutput)
}

type AFDOriginOutput struct{ *pulumi.OutputState }

func (AFDOriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AFDOrigin)(nil))
}

func (o AFDOriginOutput) ToAFDOriginOutput() AFDOriginOutput {
	return o
}

func (o AFDOriginOutput) ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AFDOriginOutput{})
}
