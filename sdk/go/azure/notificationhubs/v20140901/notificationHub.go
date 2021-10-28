


package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NotificationHub struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                  `pulumi:"location"`
	Name       pulumi.StringPtrOutput                  `pulumi:"name"`
	Properties NotificationHubPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                  `pulumi:"tags"`
	Type       pulumi.StringPtrOutput                  `pulumi:"type"`
}


func NewNotificationHub(ctx *pulumi.Context,
	name string, args *NotificationHubArgs, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:notificationhubs/v20140901:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:notificationhubs:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:notificationhubs/v20160301:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20170401:NotificationHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:notificationhubs/v20170401:NotificationHub"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationHub
	err := ctx.RegisterResource("azure-native:notificationhubs/v20140901:NotificationHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNotificationHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubState, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	var resource NotificationHub
	err := ctx.ReadResource("azure-native:notificationhubs/v20140901:NotificationHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type notificationHubState struct {
}

type NotificationHubState struct {
}

func (NotificationHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubState)(nil)).Elem()
}

type notificationHubArgs struct {
	Location            *string                   `pulumi:"location"`
	NamespaceName       string                    `pulumi:"namespaceName"`
	NotificationHubName *string                   `pulumi:"notificationHubName"`
	Properties          NotificationHubProperties `pulumi:"properties"`
	ResourceGroupName   string                    `pulumi:"resourceGroupName"`
	Tags                map[string]string         `pulumi:"tags"`
}


type NotificationHubArgs struct {
	Location            pulumi.StringPtrInput
	NamespaceName       pulumi.StringInput
	NotificationHubName pulumi.StringPtrInput
	Properties          NotificationHubPropertiesInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (NotificationHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubArgs)(nil)).Elem()
}

type NotificationHubInput interface {
	pulumi.Input

	ToNotificationHubOutput() NotificationHubOutput
	ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput
}

func (*NotificationHub) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationHub)(nil))
}

func (i *NotificationHub) ToNotificationHubOutput() NotificationHubOutput {
	return i.ToNotificationHubOutputWithContext(context.Background())
}

func (i *NotificationHub) ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubOutput)
}

type NotificationHubOutput struct{ *pulumi.OutputState }

func (NotificationHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationHub)(nil))
}

func (o NotificationHubOutput) ToNotificationHubOutput() NotificationHubOutput {
	return o
}

func (o NotificationHubOutput) ToNotificationHubOutputWithContext(ctx context.Context) NotificationHubOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotificationHubOutput{})
}
