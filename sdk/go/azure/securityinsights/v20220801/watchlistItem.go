


package v20220801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WatchlistItem struct {
	pulumi.CustomResourceState

	Created           pulumi.StringPtrOutput             `pulumi:"created"`
	CreatedBy         WatchlistUserInfoResponsePtrOutput `pulumi:"createdBy"`
	EntityMapping     pulumi.AnyOutput                   `pulumi:"entityMapping"`
	Etag              pulumi.StringPtrOutput             `pulumi:"etag"`
	IsDeleted         pulumi.BoolPtrOutput               `pulumi:"isDeleted"`
	ItemsKeyValue     pulumi.AnyOutput                   `pulumi:"itemsKeyValue"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	SystemData        SystemDataResponseOutput           `pulumi:"systemData"`
	TenantId          pulumi.StringPtrOutput             `pulumi:"tenantId"`
	Type              pulumi.StringOutput                `pulumi:"type"`
	Updated           pulumi.StringPtrOutput             `pulumi:"updated"`
	UpdatedBy         WatchlistUserInfoResponsePtrOutput `pulumi:"updatedBy"`
	WatchlistItemId   pulumi.StringPtrOutput             `pulumi:"watchlistItemId"`
	WatchlistItemType pulumi.StringPtrOutput             `pulumi:"watchlistItemType"`
}


func NewWatchlistItem(ctx *pulumi.Context,
	name string, args *WatchlistItemArgs, opts ...pulumi.ResourceOption) (*WatchlistItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ItemsKeyValue == nil {
		return nil, errors.New("invalid value for required argument 'ItemsKeyValue'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WatchlistAlias == nil {
		return nil, errors.New("invalid value for required argument 'WatchlistAlias'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:WatchlistItem"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:WatchlistItem"),
		},
	})
	opts = append(opts, aliases)
	var resource WatchlistItem
	err := ctx.RegisterResource("azure-native:securityinsights/v20220801:WatchlistItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWatchlistItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatchlistItemState, opts ...pulumi.ResourceOption) (*WatchlistItem, error) {
	var resource WatchlistItem
	err := ctx.ReadResource("azure-native:securityinsights/v20220801:WatchlistItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type watchlistItemState struct {
}

type WatchlistItemState struct {
}

func (WatchlistItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistItemState)(nil)).Elem()
}

type watchlistItemArgs struct {
	Created           *string            `pulumi:"created"`
	CreatedBy         *WatchlistUserInfo `pulumi:"createdBy"`
	EntityMapping     interface{}        `pulumi:"entityMapping"`
	IsDeleted         *bool              `pulumi:"isDeleted"`
	ItemsKeyValue     interface{}        `pulumi:"itemsKeyValue"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	TenantId          *string            `pulumi:"tenantId"`
	Updated           *string            `pulumi:"updated"`
	UpdatedBy         *WatchlistUserInfo `pulumi:"updatedBy"`
	WatchlistAlias    string             `pulumi:"watchlistAlias"`
	WatchlistItemId   *string            `pulumi:"watchlistItemId"`
	WatchlistItemType *string            `pulumi:"watchlistItemType"`
	WorkspaceName     string             `pulumi:"workspaceName"`
}


type WatchlistItemArgs struct {
	Created           pulumi.StringPtrInput
	CreatedBy         WatchlistUserInfoPtrInput
	EntityMapping     pulumi.Input
	IsDeleted         pulumi.BoolPtrInput
	ItemsKeyValue     pulumi.Input
	ResourceGroupName pulumi.StringInput
	TenantId          pulumi.StringPtrInput
	Updated           pulumi.StringPtrInput
	UpdatedBy         WatchlistUserInfoPtrInput
	WatchlistAlias    pulumi.StringInput
	WatchlistItemId   pulumi.StringPtrInput
	WatchlistItemType pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (WatchlistItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistItemArgs)(nil)).Elem()
}

type WatchlistItemInput interface {
	pulumi.Input

	ToWatchlistItemOutput() WatchlistItemOutput
	ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput
}

func (*WatchlistItem) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistItem)(nil)).Elem()
}

func (i *WatchlistItem) ToWatchlistItemOutput() WatchlistItemOutput {
	return i.ToWatchlistItemOutputWithContext(context.Background())
}

func (i *WatchlistItem) ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistItemOutput)
}

type WatchlistItemOutput struct{ *pulumi.OutputState }

func (WatchlistItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistItem)(nil)).Elem()
}

func (o WatchlistItemOutput) ToWatchlistItemOutput() WatchlistItemOutput {
	return o
}

func (o WatchlistItemOutput) ToWatchlistItemOutputWithContext(ctx context.Context) WatchlistItemOutput {
	return o
}

func (o WatchlistItemOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.Created }).(pulumi.StringPtrOutput)
}

func (o WatchlistItemOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *WatchlistItem) WatchlistUserInfoResponsePtrOutput { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o WatchlistItemOutput) EntityMapping() pulumi.AnyOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.AnyOutput { return v.EntityMapping }).(pulumi.AnyOutput)
}

func (o WatchlistItemOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o WatchlistItemOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.BoolPtrOutput { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

func (o WatchlistItemOutput) ItemsKeyValue() pulumi.AnyOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.AnyOutput { return v.ItemsKeyValue }).(pulumi.AnyOutput)
}

func (o WatchlistItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WatchlistItemOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WatchlistItem) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WatchlistItemOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o WatchlistItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WatchlistItemOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.Updated }).(pulumi.StringPtrOutput)
}

func (o WatchlistItemOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *WatchlistItem) WatchlistUserInfoResponsePtrOutput { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o WatchlistItemOutput) WatchlistItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.WatchlistItemId }).(pulumi.StringPtrOutput)
}

func (o WatchlistItemOutput) WatchlistItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistItem) pulumi.StringPtrOutput { return v.WatchlistItemType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WatchlistItemOutput{})
}
