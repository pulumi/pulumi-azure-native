


package v20200820

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CommunicationService struct {
	pulumi.CustomResourceState

	DataLocation        pulumi.StringOutput      `pulumi:"dataLocation"`
	HostName            pulumi.StringOutput      `pulumi:"hostName"`
	ImmutableResourceId pulumi.StringOutput      `pulumi:"immutableResourceId"`
	Location            pulumi.StringPtrOutput   `pulumi:"location"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	NotificationHubId   pulumi.StringOutput      `pulumi:"notificationHubId"`
	ProvisioningState   pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput   `pulumi:"tags"`
	Type                pulumi.StringOutput      `pulumi:"type"`
	Version             pulumi.StringOutput      `pulumi:"version"`
}


func NewCommunicationService(ctx *pulumi.Context,
	name string, args *CommunicationServiceArgs, opts ...pulumi.ResourceOption) (*CommunicationService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataLocation == nil {
		return nil, errors.New("invalid value for required argument 'DataLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:communication:CommunicationService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20200820preview:CommunicationService"),
		},
	})
	opts = append(opts, aliases)
	var resource CommunicationService
	err := ctx.RegisterResource("azure-native:communication/v20200820:CommunicationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCommunicationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommunicationServiceState, opts ...pulumi.ResourceOption) (*CommunicationService, error) {
	var resource CommunicationService
	err := ctx.ReadResource("azure-native:communication/v20200820:CommunicationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type communicationServiceState struct {
}

type CommunicationServiceState struct {
}

func (CommunicationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationServiceState)(nil)).Elem()
}

type communicationServiceArgs struct {
	CommunicationServiceName *string           `pulumi:"communicationServiceName"`
	DataLocation             string            `pulumi:"dataLocation"`
	Location                 *string           `pulumi:"location"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	Tags                     map[string]string `pulumi:"tags"`
}


type CommunicationServiceArgs struct {
	CommunicationServiceName pulumi.StringPtrInput
	DataLocation             pulumi.StringInput
	Location                 pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
}

func (CommunicationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*communicationServiceArgs)(nil)).Elem()
}

type CommunicationServiceInput interface {
	pulumi.Input

	ToCommunicationServiceOutput() CommunicationServiceOutput
	ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput
}

func (*CommunicationService) ElementType() reflect.Type {
	return reflect.TypeOf((*CommunicationService)(nil))
}

func (i *CommunicationService) ToCommunicationServiceOutput() CommunicationServiceOutput {
	return i.ToCommunicationServiceOutputWithContext(context.Background())
}

func (i *CommunicationService) ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunicationServiceOutput)
}

type CommunicationServiceOutput struct{ *pulumi.OutputState }

func (CommunicationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommunicationService)(nil))
}

func (o CommunicationServiceOutput) ToCommunicationServiceOutput() CommunicationServiceOutput {
	return o
}

func (o CommunicationServiceOutput) ToCommunicationServiceOutputWithContext(ctx context.Context) CommunicationServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CommunicationServiceOutput{})
}
