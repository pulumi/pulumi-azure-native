


package v20180101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiIssueAttachment struct {
	pulumi.CustomResourceState

	Content       pulumi.StringOutput `pulumi:"content"`
	ContentFormat pulumi.StringOutput `pulumi:"contentFormat"`
	Name          pulumi.StringOutput `pulumi:"name"`
	Title         pulumi.StringOutput `pulumi:"title"`
	Type          pulumi.StringOutput `pulumi:"type"`
}


func NewApiIssueAttachment(ctx *pulumi.Context,
	name string, args *ApiIssueAttachmentArgs, opts ...pulumi.ResourceOption) (*ApiIssueAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.ContentFormat == nil {
		return nil, errors.New("invalid value for required argument 'ContentFormat'")
	}
	if args.IssueId == nil {
		return nil, errors.New("invalid value for required argument 'IssueId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiIssueAttachment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiIssueAttachment"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiIssueAttachment
	err := ctx.RegisterResource("azure-native:apimanagement/v20180101:ApiIssueAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiIssueAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIssueAttachmentState, opts ...pulumi.ResourceOption) (*ApiIssueAttachment, error) {
	var resource ApiIssueAttachment
	err := ctx.ReadResource("azure-native:apimanagement/v20180101:ApiIssueAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiIssueAttachmentState struct {
}

type ApiIssueAttachmentState struct {
}

func (ApiIssueAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueAttachmentState)(nil)).Elem()
}

type apiIssueAttachmentArgs struct {
	ApiId             string  `pulumi:"apiId"`
	AttachmentId      *string `pulumi:"attachmentId"`
	Content           string  `pulumi:"content"`
	ContentFormat     string  `pulumi:"contentFormat"`
	IssueId           string  `pulumi:"issueId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	Title             string  `pulumi:"title"`
}


type ApiIssueAttachmentArgs struct {
	ApiId             pulumi.StringInput
	AttachmentId      pulumi.StringPtrInput
	Content           pulumi.StringInput
	ContentFormat     pulumi.StringInput
	IssueId           pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Title             pulumi.StringInput
}

func (ApiIssueAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueAttachmentArgs)(nil)).Elem()
}

type ApiIssueAttachmentInput interface {
	pulumi.Input

	ToApiIssueAttachmentOutput() ApiIssueAttachmentOutput
	ToApiIssueAttachmentOutputWithContext(ctx context.Context) ApiIssueAttachmentOutput
}

func (*ApiIssueAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssueAttachment)(nil)).Elem()
}

func (i *ApiIssueAttachment) ToApiIssueAttachmentOutput() ApiIssueAttachmentOutput {
	return i.ToApiIssueAttachmentOutputWithContext(context.Background())
}

func (i *ApiIssueAttachment) ToApiIssueAttachmentOutputWithContext(ctx context.Context) ApiIssueAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIssueAttachmentOutput)
}

type ApiIssueAttachmentOutput struct{ *pulumi.OutputState }

func (ApiIssueAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssueAttachment)(nil)).Elem()
}

func (o ApiIssueAttachmentOutput) ToApiIssueAttachmentOutput() ApiIssueAttachmentOutput {
	return o
}

func (o ApiIssueAttachmentOutput) ToApiIssueAttachmentOutputWithContext(ctx context.Context) ApiIssueAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiIssueAttachmentOutput{})
}
