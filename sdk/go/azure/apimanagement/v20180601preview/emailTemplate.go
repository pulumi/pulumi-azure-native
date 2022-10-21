


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EmailTemplate struct {
	pulumi.CustomResourceState

	Body        pulumi.StringOutput                                          `pulumi:"body"`
	Description pulumi.StringPtrOutput                                       `pulumi:"description"`
	IsDefault   pulumi.BoolOutput                                            `pulumi:"isDefault"`
	Name        pulumi.StringOutput                                          `pulumi:"name"`
	Parameters  EmailTemplateParametersContractPropertiesResponseArrayOutput `pulumi:"parameters"`
	Subject     pulumi.StringOutput                                          `pulumi:"subject"`
	Title       pulumi.StringPtrOutput                                       `pulumi:"title"`
	Type        pulumi.StringOutput                                          `pulumi:"type"`
}


func NewEmailTemplate(ctx *pulumi.Context,
	name string, args *EmailTemplateArgs, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:EmailTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource EmailTemplate
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:EmailTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEmailTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailTemplateState, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	var resource EmailTemplate
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:EmailTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type emailTemplateState struct {
}

type EmailTemplateState struct {
}

func (EmailTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateState)(nil)).Elem()
}

type emailTemplateArgs struct {
	Body              *string                                     `pulumi:"body"`
	Description       *string                                     `pulumi:"description"`
	Parameters        []EmailTemplateParametersContractProperties `pulumi:"parameters"`
	ResourceGroupName string                                      `pulumi:"resourceGroupName"`
	ServiceName       string                                      `pulumi:"serviceName"`
	Subject           *string                                     `pulumi:"subject"`
	TemplateName      *string                                     `pulumi:"templateName"`
	Title             *string                                     `pulumi:"title"`
}


type EmailTemplateArgs struct {
	Body              pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	Parameters        EmailTemplateParametersContractPropertiesArrayInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Subject           pulumi.StringPtrInput
	TemplateName      pulumi.StringPtrInput
	Title             pulumi.StringPtrInput
}

func (EmailTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateArgs)(nil)).Elem()
}

type EmailTemplateInput interface {
	pulumi.Input

	ToEmailTemplateOutput() EmailTemplateOutput
	ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput
}

func (*EmailTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailTemplate)(nil)).Elem()
}

func (i *EmailTemplate) ToEmailTemplateOutput() EmailTemplateOutput {
	return i.ToEmailTemplateOutputWithContext(context.Background())
}

func (i *EmailTemplate) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateOutput)
}

type EmailTemplateOutput struct{ *pulumi.OutputState }

func (EmailTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailTemplate)(nil)).Elem()
}

func (o EmailTemplateOutput) ToEmailTemplateOutput() EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Body }).(pulumi.StringOutput)
}

func (o EmailTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

func (o EmailTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EmailTemplateOutput) Parameters() EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *EmailTemplate) EmailTemplateParametersContractPropertiesResponseArrayOutput {
		return v.Parameters
	}).(EmailTemplateParametersContractPropertiesResponseArrayOutput)
}

func (o EmailTemplateOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

func (o EmailTemplateOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EmailTemplateOutput{})
}
