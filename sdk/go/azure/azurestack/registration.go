


package azurestack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Registration struct {
	pulumi.CustomResourceState

	BillingModel pulumi.StringPtrOutput `pulumi:"billingModel"`
	CloudId      pulumi.StringPtrOutput `pulumi:"cloudId"`
	Etag         pulumi.StringPtrOutput `pulumi:"etag"`
	Location     pulumi.StringOutput    `pulumi:"location"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	ObjectId     pulumi.StringPtrOutput `pulumi:"objectId"`
	Tags         pulumi.StringMapOutput `pulumi:"tags"`
	Type         pulumi.StringOutput    `pulumi:"type"`
}


func NewRegistration(ctx *pulumi.Context,
	name string, args *RegistrationArgs, opts ...pulumi.ResourceOption) (*Registration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistrationToken == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationToken'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestack/v20160101:Registration"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20170601:Registration"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20200601preview:Registration"),
		},
		{
			Type: pulumi.String("azure-native:azurestack/v20220601:Registration"),
		},
	})
	opts = append(opts, aliases)
	var resource Registration
	err := ctx.RegisterResource("azure-native:azurestack:Registration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationState, opts ...pulumi.ResourceOption) (*Registration, error) {
	var resource Registration
	err := ctx.ReadResource("azure-native:azurestack:Registration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registrationState struct {
}

type RegistrationState struct {
}

func (RegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationState)(nil)).Elem()
}

type registrationArgs struct {
	Location          *string `pulumi:"location"`
	RegistrationName  *string `pulumi:"registrationName"`
	RegistrationToken string  `pulumi:"registrationToken"`
	ResourceGroup     string  `pulumi:"resourceGroup"`
}


type RegistrationArgs struct {
	Location          pulumi.StringPtrInput
	RegistrationName  pulumi.StringPtrInput
	RegistrationToken pulumi.StringInput
	ResourceGroup     pulumi.StringInput
}

func (RegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationArgs)(nil)).Elem()
}

type RegistrationInput interface {
	pulumi.Input

	ToRegistrationOutput() RegistrationOutput
	ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput
}

func (*Registration) ElementType() reflect.Type {
	return reflect.TypeOf((**Registration)(nil)).Elem()
}

func (i *Registration) ToRegistrationOutput() RegistrationOutput {
	return i.ToRegistrationOutputWithContext(context.Background())
}

func (i *Registration) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationOutput)
}

type RegistrationOutput struct{ *pulumi.OutputState }

func (RegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registration)(nil)).Elem()
}

func (o RegistrationOutput) ToRegistrationOutput() RegistrationOutput {
	return o
}

func (o RegistrationOutput) ToRegistrationOutputWithContext(ctx context.Context) RegistrationOutput {
	return o
}

func (o RegistrationOutput) BillingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.BillingModel }).(pulumi.StringPtrOutput)
}

func (o RegistrationOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.CloudId }).(pulumi.StringPtrOutput)
}

func (o RegistrationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RegistrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o RegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegistrationOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringPtrOutput { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o RegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Registration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistrationOutput{})
}
