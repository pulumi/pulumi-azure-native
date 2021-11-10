


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomIPPrefix struct {
	pulumi.CustomResourceState

	AuthorizationMessage  pulumi.StringPtrOutput            `pulumi:"authorizationMessage"`
	ChildCustomIpPrefixes SubResourceResponseArrayOutput    `pulumi:"childCustomIpPrefixes"`
	Cidr                  pulumi.StringPtrOutput            `pulumi:"cidr"`
	CommissionedState     pulumi.StringPtrOutput            `pulumi:"commissionedState"`
	CustomIpPrefixParent  SubResourceResponsePtrOutput      `pulumi:"customIpPrefixParent"`
	Etag                  pulumi.StringOutput               `pulumi:"etag"`
	ExtendedLocation      ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	FailedReason          pulumi.StringOutput               `pulumi:"failedReason"`
	Location              pulumi.StringPtrOutput            `pulumi:"location"`
	Name                  pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput               `pulumi:"provisioningState"`
	PublicIpPrefixes      SubResourceResponseArrayOutput    `pulumi:"publicIpPrefixes"`
	ResourceGuid          pulumi.StringOutput               `pulumi:"resourceGuid"`
	SignedMessage         pulumi.StringPtrOutput            `pulumi:"signedMessage"`
	Tags                  pulumi.StringMapOutput            `pulumi:"tags"`
	Type                  pulumi.StringOutput               `pulumi:"type"`
	Zones                 pulumi.StringArrayOutput          `pulumi:"zones"`
}


func NewCustomIPPrefix(ctx *pulumi.Context,
	name string, args *CustomIPPrefixArgs, opts ...pulumi.ResourceOption) (*CustomIPPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:CustomIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:CustomIPPrefix"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomIPPrefix
	err := ctx.RegisterResource("azure-native:network/v20210501:CustomIPPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomIPPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomIPPrefixState, opts ...pulumi.ResourceOption) (*CustomIPPrefix, error) {
	var resource CustomIPPrefix
	err := ctx.ReadResource("azure-native:network/v20210501:CustomIPPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customIPPrefixState struct {
}

type CustomIPPrefixState struct {
}

func (CustomIPPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*customIPPrefixState)(nil)).Elem()
}

type customIPPrefixArgs struct {
	AuthorizationMessage *string           `pulumi:"authorizationMessage"`
	Cidr                 *string           `pulumi:"cidr"`
	CommissionedState    *string           `pulumi:"commissionedState"`
	CustomIpPrefixName   *string           `pulumi:"customIpPrefixName"`
	CustomIpPrefixParent *SubResource      `pulumi:"customIpPrefixParent"`
	ExtendedLocation     *ExtendedLocation `pulumi:"extendedLocation"`
	Id                   *string           `pulumi:"id"`
	Location             *string           `pulumi:"location"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	SignedMessage        *string           `pulumi:"signedMessage"`
	Tags                 map[string]string `pulumi:"tags"`
	Zones                []string          `pulumi:"zones"`
}


type CustomIPPrefixArgs struct {
	AuthorizationMessage pulumi.StringPtrInput
	Cidr                 pulumi.StringPtrInput
	CommissionedState    pulumi.StringPtrInput
	CustomIpPrefixName   pulumi.StringPtrInput
	CustomIpPrefixParent SubResourcePtrInput
	ExtendedLocation     ExtendedLocationPtrInput
	Id                   pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	SignedMessage        pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	Zones                pulumi.StringArrayInput
}

func (CustomIPPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customIPPrefixArgs)(nil)).Elem()
}

type CustomIPPrefixInput interface {
	pulumi.Input

	ToCustomIPPrefixOutput() CustomIPPrefixOutput
	ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput
}

func (*CustomIPPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomIPPrefix)(nil))
}

func (i *CustomIPPrefix) ToCustomIPPrefixOutput() CustomIPPrefixOutput {
	return i.ToCustomIPPrefixOutputWithContext(context.Background())
}

func (i *CustomIPPrefix) ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomIPPrefixOutput)
}

type CustomIPPrefixOutput struct{ *pulumi.OutputState }

func (CustomIPPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomIPPrefix)(nil))
}

func (o CustomIPPrefixOutput) ToCustomIPPrefixOutput() CustomIPPrefixOutput {
	return o
}

func (o CustomIPPrefixOutput) ToCustomIPPrefixOutputWithContext(ctx context.Context) CustomIPPrefixOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomIPPrefixOutput{})
}
