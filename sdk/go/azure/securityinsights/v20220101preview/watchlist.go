


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Watchlist struct {
	pulumi.CustomResourceState

	ContentType         pulumi.StringPtrOutput             `pulumi:"contentType"`
	Created             pulumi.StringPtrOutput             `pulumi:"created"`
	CreatedBy           WatchlistUserInfoResponsePtrOutput `pulumi:"createdBy"`
	DefaultDuration     pulumi.StringPtrOutput             `pulumi:"defaultDuration"`
	Description         pulumi.StringPtrOutput             `pulumi:"description"`
	DisplayName         pulumi.StringOutput                `pulumi:"displayName"`
	Etag                pulumi.StringPtrOutput             `pulumi:"etag"`
	IsDeleted           pulumi.BoolPtrOutput               `pulumi:"isDeleted"`
	ItemsSearchKey      pulumi.StringOutput                `pulumi:"itemsSearchKey"`
	Labels              pulumi.StringArrayOutput           `pulumi:"labels"`
	Name                pulumi.StringOutput                `pulumi:"name"`
	NumberOfLinesToSkip pulumi.IntPtrOutput                `pulumi:"numberOfLinesToSkip"`
	Provider            pulumi.StringOutput                `pulumi:"provider"`
	ProvisioningState   pulumi.StringOutput                `pulumi:"provisioningState"`
	RawContent          pulumi.StringPtrOutput             `pulumi:"rawContent"`
	SasUri              pulumi.StringPtrOutput             `pulumi:"sasUri"`
	Source              pulumi.StringPtrOutput             `pulumi:"source"`
	SourceType          pulumi.StringPtrOutput             `pulumi:"sourceType"`
	SystemData          SystemDataResponseOutput           `pulumi:"systemData"`
	TenantId            pulumi.StringPtrOutput             `pulumi:"tenantId"`
	Type                pulumi.StringOutput                `pulumi:"type"`
	Updated             pulumi.StringPtrOutput             `pulumi:"updated"`
	UpdatedBy           WatchlistUserInfoResponsePtrOutput `pulumi:"updatedBy"`
	UploadStatus        pulumi.StringPtrOutput             `pulumi:"uploadStatus"`
	WatchlistAlias      pulumi.StringPtrOutput             `pulumi:"watchlistAlias"`
	WatchlistId         pulumi.StringPtrOutput             `pulumi:"watchlistId"`
	WatchlistType       pulumi.StringPtrOutput             `pulumi:"watchlistType"`
}


func NewWatchlist(ctx *pulumi.Context,
	name string, args *WatchlistArgs, opts ...pulumi.ResourceOption) (*Watchlist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ItemsSearchKey == nil {
		return nil, errors.New("invalid value for required argument 'ItemsSearchKey'")
	}
	if args.Provider == nil {
		return nil, errors.New("invalid value for required argument 'Provider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Watchlist"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Watchlist"),
		},
	})
	opts = append(opts, aliases)
	var resource Watchlist
	err := ctx.RegisterResource("azure-native:securityinsights/v20220101preview:Watchlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWatchlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatchlistState, opts ...pulumi.ResourceOption) (*Watchlist, error) {
	var resource Watchlist
	err := ctx.ReadResource("azure-native:securityinsights/v20220101preview:Watchlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type watchlistState struct {
}

type WatchlistState struct {
}

func (WatchlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistState)(nil)).Elem()
}

type watchlistArgs struct {
	ContentType         *string            `pulumi:"contentType"`
	Created             *string            `pulumi:"created"`
	CreatedBy           *WatchlistUserInfo `pulumi:"createdBy"`
	DefaultDuration     *string            `pulumi:"defaultDuration"`
	Description         *string            `pulumi:"description"`
	DisplayName         string             `pulumi:"displayName"`
	IsDeleted           *bool              `pulumi:"isDeleted"`
	ItemsSearchKey      string             `pulumi:"itemsSearchKey"`
	Labels              []string           `pulumi:"labels"`
	NumberOfLinesToSkip *int               `pulumi:"numberOfLinesToSkip"`
	Provider            string             `pulumi:"provider"`
	RawContent          *string            `pulumi:"rawContent"`
	ResourceGroupName   string             `pulumi:"resourceGroupName"`
	SasUri              *string            `pulumi:"sasUri"`
	Source              *string            `pulumi:"source"`
	SourceType          *string            `pulumi:"sourceType"`
	TenantId            *string            `pulumi:"tenantId"`
	Updated             *string            `pulumi:"updated"`
	UpdatedBy           *WatchlistUserInfo `pulumi:"updatedBy"`
	UploadStatus        *string            `pulumi:"uploadStatus"`
	WatchlistAlias      *string            `pulumi:"watchlistAlias"`
	WatchlistId         *string            `pulumi:"watchlistId"`
	WatchlistType       *string            `pulumi:"watchlistType"`
	WorkspaceName       string             `pulumi:"workspaceName"`
}


type WatchlistArgs struct {
	ContentType         pulumi.StringPtrInput
	Created             pulumi.StringPtrInput
	CreatedBy           WatchlistUserInfoPtrInput
	DefaultDuration     pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	DisplayName         pulumi.StringInput
	IsDeleted           pulumi.BoolPtrInput
	ItemsSearchKey      pulumi.StringInput
	Labels              pulumi.StringArrayInput
	NumberOfLinesToSkip pulumi.IntPtrInput
	Provider            pulumi.StringInput
	RawContent          pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	SasUri              pulumi.StringPtrInput
	Source              pulumi.StringPtrInput
	SourceType          pulumi.StringPtrInput
	TenantId            pulumi.StringPtrInput
	Updated             pulumi.StringPtrInput
	UpdatedBy           WatchlistUserInfoPtrInput
	UploadStatus        pulumi.StringPtrInput
	WatchlistAlias      pulumi.StringPtrInput
	WatchlistId         pulumi.StringPtrInput
	WatchlistType       pulumi.StringPtrInput
	WorkspaceName       pulumi.StringInput
}

func (WatchlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watchlistArgs)(nil)).Elem()
}

type WatchlistInput interface {
	pulumi.Input

	ToWatchlistOutput() WatchlistOutput
	ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput
}

func (*Watchlist) ElementType() reflect.Type {
	return reflect.TypeOf((**Watchlist)(nil)).Elem()
}

func (i *Watchlist) ToWatchlistOutput() WatchlistOutput {
	return i.ToWatchlistOutputWithContext(context.Background())
}

func (i *Watchlist) ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistOutput)
}

type WatchlistOutput struct{ *pulumi.OutputState }

func (WatchlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Watchlist)(nil)).Elem()
}

func (o WatchlistOutput) ToWatchlistOutput() WatchlistOutput {
	return o
}

func (o WatchlistOutput) ToWatchlistOutputWithContext(ctx context.Context) WatchlistOutput {
	return o
}

func (o WatchlistOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Created }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) CreatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *Watchlist) WatchlistUserInfoResponsePtrOutput { return v.CreatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o WatchlistOutput) DefaultDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.DefaultDuration }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o WatchlistOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) IsDeleted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.BoolPtrOutput { return v.IsDeleted }).(pulumi.BoolPtrOutput)
}

func (o WatchlistOutput) ItemsSearchKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.ItemsSearchKey }).(pulumi.StringOutput)
}

func (o WatchlistOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o WatchlistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WatchlistOutput) NumberOfLinesToSkip() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.IntPtrOutput { return v.NumberOfLinesToSkip }).(pulumi.IntPtrOutput)
}

func (o WatchlistOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.Provider }).(pulumi.StringOutput)
}

func (o WatchlistOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WatchlistOutput) RawContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.RawContent }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) SasUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.SasUri }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Watchlist) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WatchlistOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WatchlistOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.Updated }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) UpdatedBy() WatchlistUserInfoResponsePtrOutput {
	return o.ApplyT(func(v *Watchlist) WatchlistUserInfoResponsePtrOutput { return v.UpdatedBy }).(WatchlistUserInfoResponsePtrOutput)
}

func (o WatchlistOutput) UploadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.UploadStatus }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) WatchlistAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.WatchlistAlias }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) WatchlistId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.WatchlistId }).(pulumi.StringPtrOutput)
}

func (o WatchlistOutput) WatchlistType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watchlist) pulumi.StringPtrOutput { return v.WatchlistType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WatchlistOutput{})
}
