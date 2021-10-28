


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SavedSearch struct {
	pulumi.CustomResourceState

	Category           pulumi.StringOutput     `pulumi:"category"`
	DisplayName        pulumi.StringOutput     `pulumi:"displayName"`
	Etag               pulumi.StringPtrOutput  `pulumi:"etag"`
	FunctionAlias      pulumi.StringPtrOutput  `pulumi:"functionAlias"`
	FunctionParameters pulumi.StringPtrOutput  `pulumi:"functionParameters"`
	Name               pulumi.StringOutput     `pulumi:"name"`
	Query              pulumi.StringOutput     `pulumi:"query"`
	Tags               TagResponseArrayOutput  `pulumi:"tags"`
	Type               pulumi.StringOutput     `pulumi:"type"`
	Version            pulumi.Float64PtrOutput `pulumi:"version"`
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
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200301preview:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20150320:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20150320:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200801:SavedSearch"),
		},
	})
	opts = append(opts, aliases)
	var resource SavedSearch
	err := ctx.RegisterResource("azure-native:operationalinsights/v20200301preview:SavedSearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSavedSearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SavedSearchState, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	var resource SavedSearch
	err := ctx.ReadResource("azure-native:operationalinsights/v20200301preview:SavedSearch", name, id, state, &resource, opts...)
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
	Category           string   `pulumi:"category"`
	DisplayName        string   `pulumi:"displayName"`
	Etag               *string  `pulumi:"etag"`
	FunctionAlias      *string  `pulumi:"functionAlias"`
	FunctionParameters *string  `pulumi:"functionParameters"`
	Query              string   `pulumi:"query"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SavedSearchId      *string  `pulumi:"savedSearchId"`
	Tags               []Tag    `pulumi:"tags"`
	Version            *float64 `pulumi:"version"`
	WorkspaceName      string   `pulumi:"workspaceName"`
}


type SavedSearchArgs struct {
	Category           pulumi.StringInput
	DisplayName        pulumi.StringInput
	Etag               pulumi.StringPtrInput
	FunctionAlias      pulumi.StringPtrInput
	FunctionParameters pulumi.StringPtrInput
	Query              pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	SavedSearchId      pulumi.StringPtrInput
	Tags               TagArrayInput
	Version            pulumi.Float64PtrInput
	WorkspaceName      pulumi.StringInput
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
	return reflect.TypeOf((*SavedSearch)(nil))
}

func (i *SavedSearch) ToSavedSearchOutput() SavedSearchOutput {
	return i.ToSavedSearchOutputWithContext(context.Background())
}

func (i *SavedSearch) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedSearchOutput)
}

type SavedSearchOutput struct{ *pulumi.OutputState }

func (SavedSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SavedSearch)(nil))
}

func (o SavedSearchOutput) ToSavedSearchOutput() SavedSearchOutput {
	return o
}

func (o SavedSearchOutput) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SavedSearchInput)(nil)).Elem(), &SavedSearch{})
	pulumi.RegisterOutputType(SavedSearchOutput{})
}
