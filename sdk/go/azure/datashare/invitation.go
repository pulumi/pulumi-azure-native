


package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Invitation struct {
	pulumi.CustomResourceState

	ExpirationDate          pulumi.StringPtrOutput   `pulumi:"expirationDate"`
	InvitationId            pulumi.StringOutput      `pulumi:"invitationId"`
	InvitationStatus        pulumi.StringOutput      `pulumi:"invitationStatus"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	RespondedAt             pulumi.StringOutput      `pulumi:"respondedAt"`
	SentAt                  pulumi.StringOutput      `pulumi:"sentAt"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	TargetActiveDirectoryId pulumi.StringPtrOutput   `pulumi:"targetActiveDirectoryId"`
	TargetEmail             pulumi.StringPtrOutput   `pulumi:"targetEmail"`
	TargetObjectId          pulumi.StringPtrOutput   `pulumi:"targetObjectId"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
	UserEmail               pulumi.StringOutput      `pulumi:"userEmail"`
	UserName                pulumi.StringOutput      `pulumi:"userName"`
}


func NewInvitation(ctx *pulumi.Context,
	name string, args *InvitationArgs, opts ...pulumi.ResourceOption) (*Invitation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:Invitation"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:Invitation"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:Invitation"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:Invitation"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:Invitation"),
		},
	})
	opts = append(opts, aliases)
	var resource Invitation
	err := ctx.RegisterResource("azure-native:datashare:Invitation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInvitation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InvitationState, opts ...pulumi.ResourceOption) (*Invitation, error) {
	var resource Invitation
	err := ctx.ReadResource("azure-native:datashare:Invitation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type invitationState struct {
}

type InvitationState struct {
}

func (InvitationState) ElementType() reflect.Type {
	return reflect.TypeOf((*invitationState)(nil)).Elem()
}

type invitationArgs struct {
	AccountName             string  `pulumi:"accountName"`
	ExpirationDate          *string `pulumi:"expirationDate"`
	InvitationName          *string `pulumi:"invitationName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ShareName               string  `pulumi:"shareName"`
	TargetActiveDirectoryId *string `pulumi:"targetActiveDirectoryId"`
	TargetEmail             *string `pulumi:"targetEmail"`
	TargetObjectId          *string `pulumi:"targetObjectId"`
}


type InvitationArgs struct {
	AccountName             pulumi.StringInput
	ExpirationDate          pulumi.StringPtrInput
	InvitationName          pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ShareName               pulumi.StringInput
	TargetActiveDirectoryId pulumi.StringPtrInput
	TargetEmail             pulumi.StringPtrInput
	TargetObjectId          pulumi.StringPtrInput
}

func (InvitationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*invitationArgs)(nil)).Elem()
}

type InvitationInput interface {
	pulumi.Input

	ToInvitationOutput() InvitationOutput
	ToInvitationOutputWithContext(ctx context.Context) InvitationOutput
}

func (*Invitation) ElementType() reflect.Type {
	return reflect.TypeOf((*Invitation)(nil))
}

func (i *Invitation) ToInvitationOutput() InvitationOutput {
	return i.ToInvitationOutputWithContext(context.Background())
}

func (i *Invitation) ToInvitationOutputWithContext(ctx context.Context) InvitationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InvitationOutput)
}

type InvitationOutput struct{ *pulumi.OutputState }

func (InvitationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Invitation)(nil))
}

func (o InvitationOutput) ToInvitationOutput() InvitationOutput {
	return o
}

func (o InvitationOutput) ToInvitationOutputWithContext(ctx context.Context) InvitationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InvitationOutput{})
}
