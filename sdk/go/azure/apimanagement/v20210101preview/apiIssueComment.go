


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiIssueComment struct {
	pulumi.CustomResourceState

	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Text        pulumi.StringOutput    `pulumi:"text"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	UserId      pulumi.StringOutput    `pulumi:"userId"`
}


func NewApiIssueComment(ctx *pulumi.Context,
	name string, args *ApiIssueCommentArgs, opts ...pulumi.ResourceOption) (*ApiIssueComment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
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
	if args.Text == nil {
		return nil, errors.New("invalid value for required argument 'Text'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiIssueComment"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiIssueComment"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiIssueComment
	err := ctx.RegisterResource("azure-native:apimanagement/v20210101preview:ApiIssueComment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiIssueComment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIssueCommentState, opts ...pulumi.ResourceOption) (*ApiIssueComment, error) {
	var resource ApiIssueComment
	err := ctx.ReadResource("azure-native:apimanagement/v20210101preview:ApiIssueComment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiIssueCommentState struct {
}

type ApiIssueCommentState struct {
}

func (ApiIssueCommentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueCommentState)(nil)).Elem()
}

type apiIssueCommentArgs struct {
	ApiId             string  `pulumi:"apiId"`
	CommentId         *string `pulumi:"commentId"`
	CreatedDate       *string `pulumi:"createdDate"`
	IssueId           string  `pulumi:"issueId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	Text              string  `pulumi:"text"`
	UserId            string  `pulumi:"userId"`
}


type ApiIssueCommentArgs struct {
	ApiId             pulumi.StringInput
	CommentId         pulumi.StringPtrInput
	CreatedDate       pulumi.StringPtrInput
	IssueId           pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Text              pulumi.StringInput
	UserId            pulumi.StringInput
}

func (ApiIssueCommentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueCommentArgs)(nil)).Elem()
}

type ApiIssueCommentInput interface {
	pulumi.Input

	ToApiIssueCommentOutput() ApiIssueCommentOutput
	ToApiIssueCommentOutputWithContext(ctx context.Context) ApiIssueCommentOutput
}

func (*ApiIssueComment) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiIssueComment)(nil))
}

func (i *ApiIssueComment) ToApiIssueCommentOutput() ApiIssueCommentOutput {
	return i.ToApiIssueCommentOutputWithContext(context.Background())
}

func (i *ApiIssueComment) ToApiIssueCommentOutputWithContext(ctx context.Context) ApiIssueCommentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIssueCommentOutput)
}

type ApiIssueCommentOutput struct{ *pulumi.OutputState }

func (ApiIssueCommentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiIssueComment)(nil))
}

func (o ApiIssueCommentOutput) ToApiIssueCommentOutput() ApiIssueCommentOutput {
	return o
}

func (o ApiIssueCommentOutput) ToApiIssueCommentOutputWithContext(ctx context.Context) ApiIssueCommentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiIssueCommentOutput{})
}
