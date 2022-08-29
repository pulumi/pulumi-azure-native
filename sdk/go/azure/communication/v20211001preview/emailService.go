


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EmailService struct {
	pulumi.CustomResourceState

	DataLocation      pulumi.StringOutput      `pulumi:"dataLocation"`
	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewEmailService(ctx *pulumi.Context,
	name string, args *EmailServiceArgs, opts ...pulumi.ResourceOption) (*EmailService, error) {
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
			Type: pulumi.String("azure-native:communication:EmailService"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20220701preview:EmailService"),
		},
	})
	opts = append(opts, aliases)
	var resource EmailService
	err := ctx.RegisterResource("azure-native:communication/v20211001preview:EmailService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEmailService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailServiceState, opts ...pulumi.ResourceOption) (*EmailService, error) {
	var resource EmailService
	err := ctx.ReadResource("azure-native:communication/v20211001preview:EmailService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type emailServiceState struct {
}

type EmailServiceState struct {
}

func (EmailServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailServiceState)(nil)).Elem()
}

type emailServiceArgs struct {
	DataLocation      string            `pulumi:"dataLocation"`
	EmailServiceName  *string           `pulumi:"emailServiceName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type EmailServiceArgs struct {
	DataLocation      pulumi.StringInput
	EmailServiceName  pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (EmailServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailServiceArgs)(nil)).Elem()
}

type EmailServiceInput interface {
	pulumi.Input

	ToEmailServiceOutput() EmailServiceOutput
	ToEmailServiceOutputWithContext(ctx context.Context) EmailServiceOutput
}

func (*EmailService) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailService)(nil)).Elem()
}

func (i *EmailService) ToEmailServiceOutput() EmailServiceOutput {
	return i.ToEmailServiceOutputWithContext(context.Background())
}

func (i *EmailService) ToEmailServiceOutputWithContext(ctx context.Context) EmailServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailServiceOutput)
}

type EmailServiceOutput struct{ *pulumi.OutputState }

func (EmailServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailService)(nil)).Elem()
}

func (o EmailServiceOutput) ToEmailServiceOutput() EmailServiceOutput {
	return o
}

func (o EmailServiceOutput) ToEmailServiceOutputWithContext(ctx context.Context) EmailServiceOutput {
	return o
}

func (o EmailServiceOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailService) pulumi.StringOutput { return v.DataLocation }).(pulumi.StringOutput)
}

func (o EmailServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o EmailServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EmailServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EmailServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EmailService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EmailServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EmailService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EmailServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EmailServiceOutput{})
}
