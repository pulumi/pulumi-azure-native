


package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Origin struct {
	pulumi.CustomResourceState

	Enabled                    pulumi.BoolPtrOutput     `pulumi:"enabled"`
	HostName                   pulumi.StringOutput      `pulumi:"hostName"`
	HttpPort                   pulumi.IntPtrOutput      `pulumi:"httpPort"`
	HttpsPort                  pulumi.IntPtrOutput      `pulumi:"httpsPort"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	OriginHostHeader           pulumi.StringPtrOutput   `pulumi:"originHostHeader"`
	Priority                   pulumi.IntPtrOutput      `pulumi:"priority"`
	PrivateEndpointStatus      pulumi.StringOutput      `pulumi:"privateEndpointStatus"`
	PrivateLinkAlias           pulumi.StringPtrOutput   `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage pulumi.StringPtrOutput   `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        pulumi.StringPtrOutput   `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      pulumi.StringPtrOutput   `pulumi:"privateLinkResourceId"`
	ProvisioningState          pulumi.StringOutput      `pulumi:"provisioningState"`
	ResourceState              pulumi.StringOutput      `pulumi:"resourceState"`
	SystemData                 SystemDataResponseOutput `pulumi:"systemData"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
	Weight                     pulumi.IntPtrOutput      `pulumi:"weight"`
}


func NewOrigin(ctx *pulumi.Context,
	name string, args *OriginArgs, opts ...pulumi.ResourceOption) (*Origin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:cdn:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20150601:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20160402:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20191231:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200331:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200415:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Origin"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200901:Origin"),
		},
	})
	opts = append(opts, aliases)
	var resource Origin
	err := ctx.RegisterResource("azure-native:cdn:Origin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginState, opts ...pulumi.ResourceOption) (*Origin, error) {
	var resource Origin
	err := ctx.ReadResource("azure-native:cdn:Origin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type originState struct {
}

type OriginState struct {
}

func (OriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*originState)(nil)).Elem()
}

type originArgs struct {
	Enabled                    *bool   `pulumi:"enabled"`
	EndpointName               string  `pulumi:"endpointName"`
	HostName                   string  `pulumi:"hostName"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	OriginHostHeader           *string `pulumi:"originHostHeader"`
	OriginName                 *string `pulumi:"originName"`
	Priority                   *int    `pulumi:"priority"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string `pulumi:"privateLinkResourceId"`
	ProfileName                string  `pulumi:"profileName"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	Weight                     *int    `pulumi:"weight"`
}


type OriginArgs struct {
	Enabled                    pulumi.BoolPtrInput
	EndpointName               pulumi.StringInput
	HostName                   pulumi.StringInput
	HttpPort                   pulumi.IntPtrInput
	HttpsPort                  pulumi.IntPtrInput
	OriginHostHeader           pulumi.StringPtrInput
	OriginName                 pulumi.StringPtrInput
	Priority                   pulumi.IntPtrInput
	PrivateLinkAlias           pulumi.StringPtrInput
	PrivateLinkApprovalMessage pulumi.StringPtrInput
	PrivateLinkLocation        pulumi.StringPtrInput
	PrivateLinkResourceId      pulumi.StringPtrInput
	ProfileName                pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	Weight                     pulumi.IntPtrInput
}

func (OriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originArgs)(nil)).Elem()
}

type OriginInput interface {
	pulumi.Input

	ToOriginOutput() OriginOutput
	ToOriginOutputWithContext(ctx context.Context) OriginOutput
}

func (*Origin) ElementType() reflect.Type {
	return reflect.TypeOf((*Origin)(nil))
}

func (i *Origin) ToOriginOutput() OriginOutput {
	return i.ToOriginOutputWithContext(context.Background())
}

func (i *Origin) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginOutput)
}

type OriginOutput struct{ *pulumi.OutputState }

func (OriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Origin)(nil))
}

func (o OriginOutput) ToOriginOutput() OriginOutput {
	return o
}

func (o OriginOutput) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OriginInput)(nil)).Elem(), &Origin{})
	pulumi.RegisterOutputType(OriginOutput{})
}
