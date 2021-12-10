


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Bookmark struct {
	pulumi.CustomResourceState

	Created        pulumi.StringPtrOutput        `pulumi:"created"`
	CreatedBy      UserInfoResponsePtrOutput     `pulumi:"createdBy"`
	DisplayName    pulumi.StringOutput           `pulumi:"displayName"`
	Etag           pulumi.StringPtrOutput        `pulumi:"etag"`
	EventTime      pulumi.StringPtrOutput        `pulumi:"eventTime"`
	IncidentInfo   IncidentInfoResponsePtrOutput `pulumi:"incidentInfo"`
	Labels         pulumi.StringArrayOutput      `pulumi:"labels"`
	Name           pulumi.StringOutput           `pulumi:"name"`
	Notes          pulumi.StringPtrOutput        `pulumi:"notes"`
	Query          pulumi.StringOutput           `pulumi:"query"`
	QueryEndTime   pulumi.StringPtrOutput        `pulumi:"queryEndTime"`
	QueryResult    pulumi.StringPtrOutput        `pulumi:"queryResult"`
	QueryStartTime pulumi.StringPtrOutput        `pulumi:"queryStartTime"`
	SystemData     SystemDataResponseOutput      `pulumi:"systemData"`
	Type           pulumi.StringOutput           `pulumi:"type"`
	Updated        pulumi.StringPtrOutput        `pulumi:"updated"`
	UpdatedBy      UserInfoResponsePtrOutput     `pulumi:"updatedBy"`
}


func NewBookmark(ctx *pulumi.Context,
	name string, args *BookmarkArgs, opts ...pulumi.ResourceOption) (*Bookmark, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Bookmark"),
		},
	})
	opts = append(opts, aliases)
	var resource Bookmark
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:Bookmark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBookmark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BookmarkState, opts ...pulumi.ResourceOption) (*Bookmark, error) {
	var resource Bookmark
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:Bookmark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type bookmarkState struct {
}

type BookmarkState struct {
}

func (BookmarkState) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkState)(nil)).Elem()
}

type bookmarkArgs struct {
	BookmarkId        *string       `pulumi:"bookmarkId"`
	Created           *string       `pulumi:"created"`
	CreatedBy         *UserInfo     `pulumi:"createdBy"`
	DisplayName       string        `pulumi:"displayName"`
	Etag              *string       `pulumi:"etag"`
	EventTime         *string       `pulumi:"eventTime"`
	IncidentInfo      *IncidentInfo `pulumi:"incidentInfo"`
	Labels            []string      `pulumi:"labels"`
	Notes             *string       `pulumi:"notes"`
	Query             string        `pulumi:"query"`
	QueryEndTime      *string       `pulumi:"queryEndTime"`
	QueryResult       *string       `pulumi:"queryResult"`
	QueryStartTime    *string       `pulumi:"queryStartTime"`
	ResourceGroupName string        `pulumi:"resourceGroupName"`
	Updated           *string       `pulumi:"updated"`
	UpdatedBy         *UserInfo     `pulumi:"updatedBy"`
	WorkspaceName     string        `pulumi:"workspaceName"`
}


type BookmarkArgs struct {
	BookmarkId        pulumi.StringPtrInput
	Created           pulumi.StringPtrInput
	CreatedBy         UserInfoPtrInput
	DisplayName       pulumi.StringInput
	Etag              pulumi.StringPtrInput
	EventTime         pulumi.StringPtrInput
	IncidentInfo      IncidentInfoPtrInput
	Labels            pulumi.StringArrayInput
	Notes             pulumi.StringPtrInput
	Query             pulumi.StringInput
	QueryEndTime      pulumi.StringPtrInput
	QueryResult       pulumi.StringPtrInput
	QueryStartTime    pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Updated           pulumi.StringPtrInput
	UpdatedBy         UserInfoPtrInput
	WorkspaceName     pulumi.StringInput
}

func (BookmarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkArgs)(nil)).Elem()
}

type BookmarkInput interface {
	pulumi.Input

	ToBookmarkOutput() BookmarkOutput
	ToBookmarkOutputWithContext(ctx context.Context) BookmarkOutput
}

func (*Bookmark) ElementType() reflect.Type {
	return reflect.TypeOf((*Bookmark)(nil))
}

func (i *Bookmark) ToBookmarkOutput() BookmarkOutput {
	return i.ToBookmarkOutputWithContext(context.Background())
}

func (i *Bookmark) ToBookmarkOutputWithContext(ctx context.Context) BookmarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkOutput)
}

type BookmarkOutput struct{ *pulumi.OutputState }

func (BookmarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bookmark)(nil))
}

func (o BookmarkOutput) ToBookmarkOutput() BookmarkOutput {
	return o
}

func (o BookmarkOutput) ToBookmarkOutputWithContext(ctx context.Context) BookmarkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BookmarkOutput{})
}
