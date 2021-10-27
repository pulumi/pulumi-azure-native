


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivityCustomEntityQuery struct {
	pulumi.CustomResourceState

	Content                 pulumi.StringPtrOutput                                           `pulumi:"content"`
	CreatedTimeUtc          pulumi.StringOutput                                              `pulumi:"createdTimeUtc"`
	Description             pulumi.StringPtrOutput                                           `pulumi:"description"`
	Enabled                 pulumi.BoolPtrOutput                                             `pulumi:"enabled"`
	EntitiesFilter          pulumi.StringArrayMapOutput                                      `pulumi:"entitiesFilter"`
	Etag                    pulumi.StringPtrOutput                                           `pulumi:"etag"`
	InputEntityType         pulumi.StringPtrOutput                                           `pulumi:"inputEntityType"`
	Kind                    pulumi.StringOutput                                              `pulumi:"kind"`
	LastModifiedTimeUtc     pulumi.StringOutput                                              `pulumi:"lastModifiedTimeUtc"`
	Name                    pulumi.StringOutput                                              `pulumi:"name"`
	QueryDefinitions        ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput `pulumi:"queryDefinitions"`
	RequiredInputFieldsSets pulumi.StringArrayArrayOutput                                    `pulumi:"requiredInputFieldsSets"`
	SystemData              SystemDataResponseOutput                                         `pulumi:"systemData"`
	TemplateName            pulumi.StringPtrOutput                                           `pulumi:"templateName"`
	Title                   pulumi.StringPtrOutput                                           `pulumi:"title"`
	Type                    pulumi.StringOutput                                              `pulumi:"type"`
}


func NewActivityCustomEntityQuery(ctx *pulumi.Context,
	name string, args *ActivityCustomEntityQueryArgs, opts ...pulumi.ResourceOption) (*ActivityCustomEntityQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Activity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:ActivityCustomEntityQuery"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:ActivityCustomEntityQuery"),
		},
	})
	opts = append(opts, aliases)
	var resource ActivityCustomEntityQuery
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:ActivityCustomEntityQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetActivityCustomEntityQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityCustomEntityQueryState, opts ...pulumi.ResourceOption) (*ActivityCustomEntityQuery, error) {
	var resource ActivityCustomEntityQuery
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:ActivityCustomEntityQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type activityCustomEntityQueryState struct {
}

type ActivityCustomEntityQueryState struct {
}

func (ActivityCustomEntityQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityCustomEntityQueryState)(nil)).Elem()
}

type activityCustomEntityQueryArgs struct {
	Content                             *string                                          `pulumi:"content"`
	Description                         *string                                          `pulumi:"description"`
	Enabled                             *bool                                            `pulumi:"enabled"`
	EntitiesFilter                      map[string][]string                              `pulumi:"entitiesFilter"`
	EntityQueryId                       *string                                          `pulumi:"entityQueryId"`
	Etag                                *string                                          `pulumi:"etag"`
	InputEntityType                     *string                                          `pulumi:"inputEntityType"`
	Kind                                string                                           `pulumi:"kind"`
	OperationalInsightsResourceProvider string                                           `pulumi:"operationalInsightsResourceProvider"`
	QueryDefinitions                    *ActivityEntityQueriesPropertiesQueryDefinitions `pulumi:"queryDefinitions"`
	RequiredInputFieldsSets             [][]string                                       `pulumi:"requiredInputFieldsSets"`
	ResourceGroupName                   string                                           `pulumi:"resourceGroupName"`
	TemplateName                        *string                                          `pulumi:"templateName"`
	Title                               *string                                          `pulumi:"title"`
	WorkspaceName                       string                                           `pulumi:"workspaceName"`
}


type ActivityCustomEntityQueryArgs struct {
	Content                             pulumi.StringPtrInput
	Description                         pulumi.StringPtrInput
	Enabled                             pulumi.BoolPtrInput
	EntitiesFilter                      pulumi.StringArrayMapInput
	EntityQueryId                       pulumi.StringPtrInput
	Etag                                pulumi.StringPtrInput
	InputEntityType                     pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	QueryDefinitions                    ActivityEntityQueriesPropertiesQueryDefinitionsPtrInput
	RequiredInputFieldsSets             pulumi.StringArrayArrayInput
	ResourceGroupName                   pulumi.StringInput
	TemplateName                        pulumi.StringPtrInput
	Title                               pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
}

func (ActivityCustomEntityQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityCustomEntityQueryArgs)(nil)).Elem()
}

type ActivityCustomEntityQueryInput interface {
	pulumi.Input

	ToActivityCustomEntityQueryOutput() ActivityCustomEntityQueryOutput
	ToActivityCustomEntityQueryOutputWithContext(ctx context.Context) ActivityCustomEntityQueryOutput
}

func (*ActivityCustomEntityQuery) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityCustomEntityQuery)(nil))
}

func (i *ActivityCustomEntityQuery) ToActivityCustomEntityQueryOutput() ActivityCustomEntityQueryOutput {
	return i.ToActivityCustomEntityQueryOutputWithContext(context.Background())
}

func (i *ActivityCustomEntityQuery) ToActivityCustomEntityQueryOutputWithContext(ctx context.Context) ActivityCustomEntityQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityCustomEntityQueryOutput)
}

type ActivityCustomEntityQueryOutput struct{ *pulumi.OutputState }

func (ActivityCustomEntityQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityCustomEntityQuery)(nil))
}

func (o ActivityCustomEntityQueryOutput) ToActivityCustomEntityQueryOutput() ActivityCustomEntityQueryOutput {
	return o
}

func (o ActivityCustomEntityQueryOutput) ToActivityCustomEntityQueryOutputWithContext(ctx context.Context) ActivityCustomEntityQueryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActivityCustomEntityQueryInput)(nil)).Elem(), &ActivityCustomEntityQuery{})
	pulumi.RegisterOutputType(ActivityCustomEntityQueryOutput{})
}
