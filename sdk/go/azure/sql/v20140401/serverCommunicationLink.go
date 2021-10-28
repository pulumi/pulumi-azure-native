


package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerCommunicationLink struct {
	pulumi.CustomResourceState

	Kind          pulumi.StringOutput `pulumi:"kind"`
	Location      pulumi.StringOutput `pulumi:"location"`
	Name          pulumi.StringOutput `pulumi:"name"`
	PartnerServer pulumi.StringOutput `pulumi:"partnerServer"`
	State         pulumi.StringOutput `pulumi:"state"`
	Type          pulumi.StringOutput `pulumi:"type"`
}


func NewServerCommunicationLink(ctx *pulumi.Context,
	name string, args *ServerCommunicationLinkArgs, opts ...pulumi.ResourceOption) (*ServerCommunicationLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerServer == nil {
		return nil, errors.New("invalid value for required argument 'PartnerServer'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:ServerCommunicationLink"),
		},
		{
			Type: pulumi.String("azure-native:sql:ServerCommunicationLink"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:ServerCommunicationLink"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerCommunicationLink
	err := ctx.RegisterResource("azure-native:sql/v20140401:ServerCommunicationLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerCommunicationLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerCommunicationLinkState, opts ...pulumi.ResourceOption) (*ServerCommunicationLink, error) {
	var resource ServerCommunicationLink
	err := ctx.ReadResource("azure-native:sql/v20140401:ServerCommunicationLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverCommunicationLinkState struct {
}

type ServerCommunicationLinkState struct {
}

func (ServerCommunicationLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCommunicationLinkState)(nil)).Elem()
}

type serverCommunicationLinkArgs struct {
	CommunicationLinkName *string `pulumi:"communicationLinkName"`
	PartnerServer         string  `pulumi:"partnerServer"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ServerName            string  `pulumi:"serverName"`
}


type ServerCommunicationLinkArgs struct {
	CommunicationLinkName pulumi.StringPtrInput
	PartnerServer         pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ServerName            pulumi.StringInput
}

func (ServerCommunicationLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCommunicationLinkArgs)(nil)).Elem()
}

type ServerCommunicationLinkInput interface {
	pulumi.Input

	ToServerCommunicationLinkOutput() ServerCommunicationLinkOutput
	ToServerCommunicationLinkOutputWithContext(ctx context.Context) ServerCommunicationLinkOutput
}

func (*ServerCommunicationLink) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCommunicationLink)(nil))
}

func (i *ServerCommunicationLink) ToServerCommunicationLinkOutput() ServerCommunicationLinkOutput {
	return i.ToServerCommunicationLinkOutputWithContext(context.Background())
}

func (i *ServerCommunicationLink) ToServerCommunicationLinkOutputWithContext(ctx context.Context) ServerCommunicationLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCommunicationLinkOutput)
}

type ServerCommunicationLinkOutput struct{ *pulumi.OutputState }

func (ServerCommunicationLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerCommunicationLink)(nil))
}

func (o ServerCommunicationLinkOutput) ToServerCommunicationLinkOutput() ServerCommunicationLinkOutput {
	return o
}

func (o ServerCommunicationLinkOutput) ToServerCommunicationLinkOutputWithContext(ctx context.Context) ServerCommunicationLinkOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCommunicationLinkInput)(nil)).Elem(), &ServerCommunicationLink{})
	pulumi.RegisterOutputType(ServerCommunicationLinkOutput{})
}
