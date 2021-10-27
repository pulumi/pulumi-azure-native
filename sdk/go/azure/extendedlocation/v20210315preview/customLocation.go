


package v20210315preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomLocation struct {
	pulumi.CustomResourceState

	Authentication      CustomLocationPropertiesResponseAuthenticationPtrOutput `pulumi:"authentication"`
	ClusterExtensionIds pulumi.StringArrayOutput                                `pulumi:"clusterExtensionIds"`
	DisplayName         pulumi.StringPtrOutput                                  `pulumi:"displayName"`
	HostResourceId      pulumi.StringPtrOutput                                  `pulumi:"hostResourceId"`
	HostType            pulumi.StringPtrOutput                                  `pulumi:"hostType"`
	Location            pulumi.StringOutput                                     `pulumi:"location"`
	Name                pulumi.StringOutput                                     `pulumi:"name"`
	Namespace           pulumi.StringPtrOutput                                  `pulumi:"namespace"`
	ProvisioningState   pulumi.StringPtrOutput                                  `pulumi:"provisioningState"`
	SystemData          SystemDataResponseOutput                                `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                pulumi.StringOutput                                     `pulumi:"type"`
}


func NewCustomLocation(ctx *pulumi.Context,
	name string, args *CustomLocationArgs, opts ...pulumi.ResourceOption) (*CustomLocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:extendedlocation/v20210315preview:CustomLocation"),
		},
		{
			Type: pulumi.String("azure-native:extendedlocation:CustomLocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:extendedlocation:CustomLocation"),
		},
		{
			Type: pulumi.String("azure-native:extendedlocation/v20210815:CustomLocation"),
		},
		{
			Type: pulumi.String("azure-nextgen:extendedlocation/v20210815:CustomLocation"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomLocation
	err := ctx.RegisterResource("azure-native:extendedlocation/v20210315preview:CustomLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomLocationState, opts ...pulumi.ResourceOption) (*CustomLocation, error) {
	var resource CustomLocation
	err := ctx.ReadResource("azure-native:extendedlocation/v20210315preview:CustomLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customLocationState struct {
}

type CustomLocationState struct {
}

func (CustomLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customLocationState)(nil)).Elem()
}

type customLocationArgs struct {
	Authentication      *CustomLocationPropertiesAuthentication `pulumi:"authentication"`
	ClusterExtensionIds []string                                `pulumi:"clusterExtensionIds"`
	DisplayName         *string                                 `pulumi:"displayName"`
	HostResourceId      *string                                 `pulumi:"hostResourceId"`
	HostType            *string                                 `pulumi:"hostType"`
	Location            *string                                 `pulumi:"location"`
	Namespace           *string                                 `pulumi:"namespace"`
	ProvisioningState   *string                                 `pulumi:"provisioningState"`
	ResourceGroupName   string                                  `pulumi:"resourceGroupName"`
	ResourceName        *string                                 `pulumi:"resourceName"`
	Tags                map[string]string                       `pulumi:"tags"`
}


type CustomLocationArgs struct {
	Authentication      CustomLocationPropertiesAuthenticationPtrInput
	ClusterExtensionIds pulumi.StringArrayInput
	DisplayName         pulumi.StringPtrInput
	HostResourceId      pulumi.StringPtrInput
	HostType            pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	Namespace           pulumi.StringPtrInput
	ProvisioningState   pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (CustomLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customLocationArgs)(nil)).Elem()
}

type CustomLocationInput interface {
	pulumi.Input

	ToCustomLocationOutput() CustomLocationOutput
	ToCustomLocationOutputWithContext(ctx context.Context) CustomLocationOutput
}

func (*CustomLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocation)(nil))
}

func (i *CustomLocation) ToCustomLocationOutput() CustomLocationOutput {
	return i.ToCustomLocationOutputWithContext(context.Background())
}

func (i *CustomLocation) ToCustomLocationOutputWithContext(ctx context.Context) CustomLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationOutput)
}

type CustomLocationOutput struct{ *pulumi.OutputState }

func (CustomLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocation)(nil))
}

func (o CustomLocationOutput) ToCustomLocationOutput() CustomLocationOutput {
	return o
}

func (o CustomLocationOutput) ToCustomLocationOutputWithContext(ctx context.Context) CustomLocationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLocationInput)(nil)).Elem(), &CustomLocation{})
	pulumi.RegisterOutputType(CustomLocationOutput{})
}
