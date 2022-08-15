


package v20150320

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SavedSearch struct {
	pulumi.CustomResourceState

	Category    pulumi.StringOutput     `pulumi:"category"`
	DisplayName pulumi.StringOutput     `pulumi:"displayName"`
	ETag        pulumi.StringPtrOutput  `pulumi:"eTag"`
	Name        pulumi.StringOutput     `pulumi:"name"`
	Query       pulumi.StringOutput     `pulumi:"query"`
	Tags        TagResponseArrayOutput  `pulumi:"tags"`
	Type        pulumi.StringOutput     `pulumi:"type"`
	Version     pulumi.Float64PtrOutput `pulumi:"version"`
}


func NewSavedSearch(ctx *pulumi.Context,
	name string, args *SavedSearchArgs, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
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
			Type: pulumi.String("azure-native:operationalinsights:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:SavedSearch"),
		},
	})
	opts = append(opts, aliases)
	var resource SavedSearch
	err := ctx.RegisterResource("azure-native:operationalinsights/v20150320:SavedSearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSavedSearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SavedSearchState, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	var resource SavedSearch
	err := ctx.ReadResource("azure-native:operationalinsights/v20150320:SavedSearch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type savedSearchState struct {
}

type SavedSearchState struct {
}

func (SavedSearchState) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchState)(nil)).Elem()
}

type savedSearchArgs struct {
	Category          string   `pulumi:"category"`
	DisplayName       string   `pulumi:"displayName"`
	ETag              *string  `pulumi:"eTag"`
	Query             string   `pulumi:"query"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	SavedSearchId     *string  `pulumi:"savedSearchId"`
	Tags              []Tag    `pulumi:"tags"`
	Version           *float64 `pulumi:"version"`
	WorkspaceName     string   `pulumi:"workspaceName"`
}


type SavedSearchArgs struct {
	Category          pulumi.StringInput
	DisplayName       pulumi.StringInput
	ETag              pulumi.StringPtrInput
	Query             pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SavedSearchId     pulumi.StringPtrInput
	Tags              TagArrayInput
	Version           pulumi.Float64PtrInput
	WorkspaceName     pulumi.StringInput
}

func (SavedSearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchArgs)(nil)).Elem()
}

type SavedSearchInput interface {
	pulumi.Input

	ToSavedSearchOutput() SavedSearchOutput
	ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput
}

func (*SavedSearch) ElementType() reflect.Type {
	return reflect.TypeOf((**SavedSearch)(nil)).Elem()
}

func (i *SavedSearch) ToSavedSearchOutput() SavedSearchOutput {
	return i.ToSavedSearchOutputWithContext(context.Background())
}

func (i *SavedSearch) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedSearchOutput)
}

type SavedSearchOutput struct{ *pulumi.OutputState }

func (SavedSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SavedSearch)(nil)).Elem()
}

func (o SavedSearchOutput) ToSavedSearchOutput() SavedSearchOutput {
	return o
}

func (o SavedSearchOutput) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return o
}

func (o SavedSearchOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o SavedSearchOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SavedSearchOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o SavedSearchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SavedSearchOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

func (o SavedSearchOutput) Tags() TagResponseArrayOutput {
	return o.ApplyT(func(v *SavedSearch) TagResponseArrayOutput { return v.Tags }).(TagResponseArrayOutput)
}

func (o SavedSearchOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SavedSearchOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SavedSearch) pulumi.Float64PtrOutput { return v.Version }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SavedSearchOutput{})
}
