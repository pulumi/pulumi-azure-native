


package v20190901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Query struct {
	pulumi.CustomResourceState

	Author       pulumi.StringOutput                                          `pulumi:"author"`
	Body         pulumi.StringOutput                                          `pulumi:"body"`
	Description  pulumi.StringPtrOutput                                       `pulumi:"description"`
	DisplayName  pulumi.StringOutput                                          `pulumi:"displayName"`
	Name         pulumi.StringOutput                                          `pulumi:"name"`
	Properties   pulumi.AnyOutput                                             `pulumi:"properties"`
	Related      LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput `pulumi:"related"`
	SystemData   SystemDataResponseOutput                                     `pulumi:"systemData"`
	Tags         pulumi.StringArrayMapOutput                                  `pulumi:"tags"`
	TimeCreated  pulumi.StringOutput                                          `pulumi:"timeCreated"`
	TimeModified pulumi.StringOutput                                          `pulumi:"timeModified"`
	Type         pulumi.StringOutput                                          `pulumi:"type"`
}


func NewQuery(ctx *pulumi.Context,
	name string, args *QueryArgs, opts ...pulumi.ResourceOption) (*Query, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.QueryPackName == nil {
		return nil, errors.New("invalid value for required argument 'QueryPackName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:Query"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190901:Query"),
		},
	})
	opts = append(opts, aliases)
	var resource Query
	err := ctx.RegisterResource("azure-native:operationalinsights/v20190901preview:Query", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryState, opts ...pulumi.ResourceOption) (*Query, error) {
	var resource Query
	err := ctx.ReadResource("azure-native:operationalinsights/v20190901preview:Query", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type queryState struct {
}

type QueryState struct {
}

func (QueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryState)(nil)).Elem()
}

type queryArgs struct {
	Body              string                                       `pulumi:"body"`
	Description       *string                                      `pulumi:"description"`
	DisplayName       string                                       `pulumi:"displayName"`
	Id                *string                                      `pulumi:"id"`
	Properties        interface{}                                  `pulumi:"properties"`
	QueryPackName     string                                       `pulumi:"queryPackName"`
	Related           *LogAnalyticsQueryPackQueryPropertiesRelated `pulumi:"related"`
	ResourceGroupName string                                       `pulumi:"resourceGroupName"`
	Tags              map[string][]string                          `pulumi:"tags"`
}


type QueryArgs struct {
	Body              pulumi.StringInput
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringInput
	Id                pulumi.StringPtrInput
	Properties        pulumi.Input
	QueryPackName     pulumi.StringInput
	Related           LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringArrayMapInput
}

func (QueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryArgs)(nil)).Elem()
}

type QueryInput interface {
	pulumi.Input

	ToQueryOutput() QueryOutput
	ToQueryOutputWithContext(ctx context.Context) QueryOutput
}

func (*Query) ElementType() reflect.Type {
	return reflect.TypeOf((**Query)(nil)).Elem()
}

func (i *Query) ToQueryOutput() QueryOutput {
	return i.ToQueryOutputWithContext(context.Background())
}

func (i *Query) ToQueryOutputWithContext(ctx context.Context) QueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryOutput)
}

type QueryOutput struct{ *pulumi.OutputState }

func (QueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Query)(nil)).Elem()
}

func (o QueryOutput) ToQueryOutput() QueryOutput {
	return o
}

func (o QueryOutput) ToQueryOutputWithContext(ctx context.Context) QueryOutput {
	return o
}

func (o QueryOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

func (o QueryOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Body }).(pulumi.StringOutput)
}

func (o QueryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Query) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o QueryOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o QueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueryOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Query) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o QueryOutput) Related() LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o.ApplyT(func(v *Query) LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput { return v.Related }).(LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput)
}

func (o QueryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Query) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o QueryOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *Query) pulumi.StringArrayMapOutput { return v.Tags }).(pulumi.StringArrayMapOutput)
}

func (o QueryOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o QueryOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

func (o QueryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Query) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueryOutput{})
}
