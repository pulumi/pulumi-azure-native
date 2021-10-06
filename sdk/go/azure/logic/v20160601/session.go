


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Session struct {
	pulumi.CustomResourceState

	ChangedTime pulumi.StringOutput    `pulumi:"changedTime"`
	Content     pulumi.AnyOutput       `pulumi:"content"`
	CreatedTime pulumi.StringOutput    `pulumi:"createdTime"`
	Location    pulumi.StringPtrOutput `pulumi:"location"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Tags        pulumi.StringMapOutput `pulumi:"tags"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewSession(ctx *pulumi.Context,
	name string, args *SessionArgs, opts ...pulumi.ResourceOption) (*Session, error) {
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
			Type: pulumi.String("azure-nextgen:logic/v20160601:Session"),
		},
		{
			Type: pulumi.String("azure-native:logic:Session"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:Session"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Session"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:Session"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Session"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:Session"),
		},
	})
	opts = append(opts, aliases)
	var resource Session
	err := ctx.RegisterResource("azure-native:logic/v20160601:Session", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SessionState, opts ...pulumi.ResourceOption) (*Session, error) {
	var resource Session
	err := ctx.ReadResource("azure-native:logic/v20160601:Session", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sessionState struct {
}

type SessionState struct {
}

func (SessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionState)(nil)).Elem()
}

type sessionArgs struct {
	Content                interface{}       `pulumi:"content"`
	IntegrationAccountName string            `pulumi:"integrationAccountName"`
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SessionName            *string           `pulumi:"sessionName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type SessionArgs struct {
	Content                pulumi.Input
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SessionName            pulumi.StringPtrInput
	Tags                   pulumi.StringMapInput
}

func (SessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionArgs)(nil)).Elem()
}

type SessionInput interface {
	pulumi.Input

	ToSessionOutput() SessionOutput
	ToSessionOutputWithContext(ctx context.Context) SessionOutput
}

func (*Session) ElementType() reflect.Type {
	return reflect.TypeOf((*Session)(nil))
}

func (i *Session) ToSessionOutput() SessionOutput {
	return i.ToSessionOutputWithContext(context.Background())
}

func (i *Session) ToSessionOutputWithContext(ctx context.Context) SessionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SessionOutput)
}

type SessionOutput struct{ *pulumi.OutputState }

func (SessionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Session)(nil))
}

func (o SessionOutput) ToSessionOutput() SessionOutput {
	return o
}

func (o SessionOutput) ToSessionOutputWithContext(ctx context.Context) SessionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SessionOutput{})
}
