


package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountSession struct {
	pulumi.CustomResourceState

	ChangedTime pulumi.StringOutput    `pulumi:"changedTime"`
	Content     pulumi.AnyOutput       `pulumi:"content"`
	CreatedTime pulumi.StringOutput    `pulumi:"createdTime"`
	Location    pulumi.StringPtrOutput `pulumi:"location"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Tags        pulumi.StringMapOutput `pulumi:"tags"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewIntegrationAccountSession(ctx *pulumi.Context,
	name string, args *IntegrationAccountSessionArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountSession, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountSession"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountSession"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountSession"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountSession
	err := ctx.RegisterResource("azure-native:logic/v20190501:IntegrationAccountSession", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountSession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountSessionState, opts ...pulumi.ResourceOption) (*IntegrationAccountSession, error) {
	var resource IntegrationAccountSession
	err := ctx.ReadResource("azure-native:logic/v20190501:IntegrationAccountSession", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountSessionState struct {
}

type IntegrationAccountSessionState struct {
}

func (IntegrationAccountSessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSessionState)(nil)).Elem()
}

type integrationAccountSessionArgs struct {
	Content                interface{}       `pulumi:"content"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SessionName            *string           `pulumi:"sessionName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type IntegrationAccountSessionArgs struct {
	Content                pulumi.Input
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SessionName            pulumi.StringPtrInput
	Tags                   pulumi.StringMapInput
}

func (IntegrationAccountSessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSessionArgs)(nil)).Elem()
}

type IntegrationAccountSessionInput interface {
	pulumi.Input

	ToIntegrationAccountSessionOutput() IntegrationAccountSessionOutput
	ToIntegrationAccountSessionOutputWithContext(ctx context.Context) IntegrationAccountSessionOutput
}

func (*IntegrationAccountSession) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSession)(nil)).Elem()
}

func (i *IntegrationAccountSession) ToIntegrationAccountSessionOutput() IntegrationAccountSessionOutput {
	return i.ToIntegrationAccountSessionOutputWithContext(context.Background())
}

func (i *IntegrationAccountSession) ToIntegrationAccountSessionOutputWithContext(ctx context.Context) IntegrationAccountSessionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSessionOutput)
}

type IntegrationAccountSessionOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSessionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSession)(nil)).Elem()
}

func (o IntegrationAccountSessionOutput) ToIntegrationAccountSessionOutput() IntegrationAccountSessionOutput {
	return o
}

func (o IntegrationAccountSessionOutput) ToIntegrationAccountSessionOutputWithContext(ctx context.Context) IntegrationAccountSessionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountSessionOutput{})
}
