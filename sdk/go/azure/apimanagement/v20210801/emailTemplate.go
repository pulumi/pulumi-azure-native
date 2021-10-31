


package v20210801

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
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:EmailTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:EmailTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource EmailTemplate
	err := ctx.RegisterResource("azure-native:apimanagement/v20210801:EmailTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEmailTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailTemplateState, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	var resource EmailTemplate
	err := ctx.ReadResource("azure-native:apimanagement/v20210801:EmailTemplate", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*EmailTemplate)(nil))
}

func (i *EmailTemplate) ToEmailTemplateOutput() EmailTemplateOutput {
	return i.ToEmailTemplateOutputWithContext(context.Background())
}

func (i *EmailTemplate) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateOutput)
}

type EmailTemplateOutput struct{ *pulumi.OutputState }

func (EmailTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailTemplate)(nil))
}

func (o EmailTemplateOutput) ToEmailTemplateOutput() EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EmailTemplateOutput{})
}
