


package voiceservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Contact struct {
	pulumi.CustomResourceState

	ContactName       pulumi.StringOutput      `pulumi:"contactName"`
	Email             pulumi.StringOutput      `pulumi:"email"`
	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	PhoneNumber       pulumi.StringOutput      `pulumi:"phoneNumber"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Role              pulumi.StringOutput      `pulumi:"role"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewContact(ctx *pulumi.Context,
	name string, args *ContactArgs, opts ...pulumi.ResourceOption) (*Contact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommunicationsGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'CommunicationsGatewayName'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.PhoneNumber == nil {
		return nil, errors.New("invalid value for required argument 'PhoneNumber'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:voiceservices/v20221201preview:Contact"),
		},
	})
	opts = append(opts, aliases)
	var resource Contact
	err := ctx.RegisterResource("azure-native:voiceservices:Contact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactState, opts ...pulumi.ResourceOption) (*Contact, error) {
	var resource Contact
	err := ctx.ReadResource("azure-native:voiceservices:Contact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type contactState struct {
}

type ContactState struct {
}

func (ContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactState)(nil)).Elem()
}

type contactArgs struct {
	CommunicationsGatewayName string            `pulumi:"communicationsGatewayName"`
	ContactName               *string           `pulumi:"contactName"`
	Email                     string            `pulumi:"email"`
	Location                  *string           `pulumi:"location"`
	PhoneNumber               string            `pulumi:"phoneNumber"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	Role                      string            `pulumi:"role"`
	Tags                      map[string]string `pulumi:"tags"`
}


type ContactArgs struct {
	CommunicationsGatewayName pulumi.StringInput
	ContactName               pulumi.StringPtrInput
	Email                     pulumi.StringInput
	Location                  pulumi.StringPtrInput
	PhoneNumber               pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	Role                      pulumi.StringInput
	Tags                      pulumi.StringMapInput
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactArgs)(nil)).Elem()
}

type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(ctx context.Context) ContactOutput
}

func (*Contact) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *Contact) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i *Contact) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

func (o ContactOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ContactName }).(pulumi.StringOutput)
}

func (o ContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o ContactOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContactOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.PhoneNumber }).(pulumi.StringOutput)
}

func (o ContactOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ContactOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ContactOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Contact) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ContactOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ContactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Contact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContactOutput{})
}
