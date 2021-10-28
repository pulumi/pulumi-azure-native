


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Incident struct {
	pulumi.CustomResourceState

	AdditionalData         IncidentAdditionalDataResponseOutput `pulumi:"additionalData"`
	Classification         pulumi.StringPtrOutput               `pulumi:"classification"`
	ClassificationComment  pulumi.StringPtrOutput               `pulumi:"classificationComment"`
	ClassificationReason   pulumi.StringPtrOutput               `pulumi:"classificationReason"`
	CreatedTimeUtc         pulumi.StringOutput                  `pulumi:"createdTimeUtc"`
	Description            pulumi.StringPtrOutput               `pulumi:"description"`
	Etag                   pulumi.StringPtrOutput               `pulumi:"etag"`
	FirstActivityTimeUtc   pulumi.StringPtrOutput               `pulumi:"firstActivityTimeUtc"`
	IncidentNumber         pulumi.IntOutput                     `pulumi:"incidentNumber"`
	IncidentUrl            pulumi.StringOutput                  `pulumi:"incidentUrl"`
	Labels                 IncidentLabelResponseArrayOutput     `pulumi:"labels"`
	LastActivityTimeUtc    pulumi.StringPtrOutput               `pulumi:"lastActivityTimeUtc"`
	LastModifiedTimeUtc    pulumi.StringOutput                  `pulumi:"lastModifiedTimeUtc"`
	Name                   pulumi.StringOutput                  `pulumi:"name"`
	Owner                  IncidentOwnerInfoResponsePtrOutput   `pulumi:"owner"`
	ProviderIncidentId     pulumi.StringPtrOutput               `pulumi:"providerIncidentId"`
	ProviderName           pulumi.StringPtrOutput               `pulumi:"providerName"`
	RelatedAnalyticRuleIds pulumi.StringArrayOutput             `pulumi:"relatedAnalyticRuleIds"`
	Severity               pulumi.StringOutput                  `pulumi:"severity"`
	Status                 pulumi.StringOutput                  `pulumi:"status"`
	Title                  pulumi.StringOutput                  `pulumi:"title"`
	Type                   pulumi.StringOutput                  `pulumi:"type"`
}


func NewIncident(ctx *pulumi.Context,
	name string, args *IncidentArgs, opts ...pulumi.ResourceOption) (*Incident, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:Incident"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Incident"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Incident"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:Incident"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210401:Incident"),
		},
	})
	opts = append(opts, aliases)
	var resource Incident
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:Incident", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIncident(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentState, opts ...pulumi.ResourceOption) (*Incident, error) {
	var resource Incident
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:Incident", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type incidentState struct {
}

type IncidentState struct {
}

func (IncidentState) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentState)(nil)).Elem()
}

type incidentArgs struct {
	Classification                      *string            `pulumi:"classification"`
	ClassificationComment               *string            `pulumi:"classificationComment"`
	ClassificationReason                *string            `pulumi:"classificationReason"`
	Description                         *string            `pulumi:"description"`
	Etag                                *string            `pulumi:"etag"`
	FirstActivityTimeUtc                *string            `pulumi:"firstActivityTimeUtc"`
	IncidentId                          *string            `pulumi:"incidentId"`
	Labels                              []IncidentLabel    `pulumi:"labels"`
	LastActivityTimeUtc                 *string            `pulumi:"lastActivityTimeUtc"`
	OperationalInsightsResourceProvider string             `pulumi:"operationalInsightsResourceProvider"`
	Owner                               *IncidentOwnerInfo `pulumi:"owner"`
	ProviderIncidentId                  *string            `pulumi:"providerIncidentId"`
	ProviderName                        *string            `pulumi:"providerName"`
	ResourceGroupName                   string             `pulumi:"resourceGroupName"`
	Severity                            string             `pulumi:"severity"`
	Status                              string             `pulumi:"status"`
	Title                               string             `pulumi:"title"`
	WorkspaceName                       string             `pulumi:"workspaceName"`
}


type IncidentArgs struct {
	Classification                      pulumi.StringPtrInput
	ClassificationComment               pulumi.StringPtrInput
	ClassificationReason                pulumi.StringPtrInput
	Description                         pulumi.StringPtrInput
	Etag                                pulumi.StringPtrInput
	FirstActivityTimeUtc                pulumi.StringPtrInput
	IncidentId                          pulumi.StringPtrInput
	Labels                              IncidentLabelArrayInput
	LastActivityTimeUtc                 pulumi.StringPtrInput
	OperationalInsightsResourceProvider pulumi.StringInput
	Owner                               IncidentOwnerInfoPtrInput
	ProviderIncidentId                  pulumi.StringPtrInput
	ProviderName                        pulumi.StringPtrInput
	ResourceGroupName                   pulumi.StringInput
	Severity                            pulumi.StringInput
	Status                              pulumi.StringInput
	Title                               pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
}

func (IncidentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*incidentArgs)(nil)).Elem()
}

type IncidentInput interface {
	pulumi.Input

	ToIncidentOutput() IncidentOutput
	ToIncidentOutputWithContext(ctx context.Context) IncidentOutput
}

func (*Incident) ElementType() reflect.Type {
	return reflect.TypeOf((*Incident)(nil))
}

func (i *Incident) ToIncidentOutput() IncidentOutput {
	return i.ToIncidentOutputWithContext(context.Background())
}

func (i *Incident) ToIncidentOutputWithContext(ctx context.Context) IncidentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOutput)
}

type IncidentOutput struct{ *pulumi.OutputState }

func (IncidentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Incident)(nil))
}

func (o IncidentOutput) ToIncidentOutput() IncidentOutput {
	return o
}

func (o IncidentOutput) ToIncidentOutputWithContext(ctx context.Context) IncidentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentInput)(nil)).Elem(), &Incident{})
	pulumi.RegisterOutputType(IncidentOutput{})
}
