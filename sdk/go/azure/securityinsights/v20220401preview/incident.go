


package v20220401preview

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
	SystemData             SystemDataResponseOutput             `pulumi:"systemData"`
	TeamInformation        TeamInformationResponsePtrOutput     `pulumi:"teamInformation"`
	Title                  pulumi.StringOutput                  `pulumi:"title"`
	Type                   pulumi.StringOutput                  `pulumi:"type"`
}


func NewIncident(ctx *pulumi.Context,
	name string, args *IncidentArgs, opts ...pulumi.ResourceOption) (*Incident, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:securityinsights:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210401:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Incident"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Incident"),
		},
	})
	opts = append(opts, aliases)
	var resource Incident
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:Incident", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIncident(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IncidentState, opts ...pulumi.ResourceOption) (*Incident, error) {
	var resource Incident
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:Incident", name, id, state, &resource, opts...)
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
	Classification        *string            `pulumi:"classification"`
	ClassificationComment *string            `pulumi:"classificationComment"`
	ClassificationReason  *string            `pulumi:"classificationReason"`
	Description           *string            `pulumi:"description"`
	FirstActivityTimeUtc  *string            `pulumi:"firstActivityTimeUtc"`
	IncidentId            *string            `pulumi:"incidentId"`
	Labels                []IncidentLabel    `pulumi:"labels"`
	LastActivityTimeUtc   *string            `pulumi:"lastActivityTimeUtc"`
	Owner                 *IncidentOwnerInfo `pulumi:"owner"`
	ProviderIncidentId    *string            `pulumi:"providerIncidentId"`
	ProviderName          *string            `pulumi:"providerName"`
	ResourceGroupName     string             `pulumi:"resourceGroupName"`
	Severity              string             `pulumi:"severity"`
	Status                string             `pulumi:"status"`
	Title                 string             `pulumi:"title"`
	WorkspaceName         string             `pulumi:"workspaceName"`
}


type IncidentArgs struct {
	Classification        pulumi.StringPtrInput
	ClassificationComment pulumi.StringPtrInput
	ClassificationReason  pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	FirstActivityTimeUtc  pulumi.StringPtrInput
	IncidentId            pulumi.StringPtrInput
	Labels                IncidentLabelArrayInput
	LastActivityTimeUtc   pulumi.StringPtrInput
	Owner                 IncidentOwnerInfoPtrInput
	ProviderIncidentId    pulumi.StringPtrInput
	ProviderName          pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Severity              pulumi.StringInput
	Status                pulumi.StringInput
	Title                 pulumi.StringInput
	WorkspaceName         pulumi.StringInput
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
	return reflect.TypeOf((**Incident)(nil)).Elem()
}

func (i *Incident) ToIncidentOutput() IncidentOutput {
	return i.ToIncidentOutputWithContext(context.Background())
}

func (i *Incident) ToIncidentOutputWithContext(ctx context.Context) IncidentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOutput)
}

type IncidentOutput struct{ *pulumi.OutputState }

func (IncidentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Incident)(nil)).Elem()
}

func (o IncidentOutput) ToIncidentOutput() IncidentOutput {
	return o
}

func (o IncidentOutput) ToIncidentOutputWithContext(ctx context.Context) IncidentOutput {
	return o
}

func (o IncidentOutput) AdditionalData() IncidentAdditionalDataResponseOutput {
	return o.ApplyT(func(v *Incident) IncidentAdditionalDataResponseOutput { return v.AdditionalData }).(IncidentAdditionalDataResponseOutput)
}

func (o IncidentOutput) Classification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.Classification }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) ClassificationComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ClassificationComment }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) ClassificationReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ClassificationReason }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o IncidentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) FirstActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.FirstActivityTimeUtc }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) IncidentNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Incident) pulumi.IntOutput { return v.IncidentNumber }).(pulumi.IntOutput)
}

func (o IncidentOutput) IncidentUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.IncidentUrl }).(pulumi.StringOutput)
}

func (o IncidentOutput) Labels() IncidentLabelResponseArrayOutput {
	return o.ApplyT(func(v *Incident) IncidentLabelResponseArrayOutput { return v.Labels }).(IncidentLabelResponseArrayOutput)
}

func (o IncidentOutput) LastActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.LastActivityTimeUtc }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o IncidentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IncidentOutput) Owner() IncidentOwnerInfoResponsePtrOutput {
	return o.ApplyT(func(v *Incident) IncidentOwnerInfoResponsePtrOutput { return v.Owner }).(IncidentOwnerInfoResponsePtrOutput)
}

func (o IncidentOutput) ProviderIncidentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ProviderIncidentId }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) ProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringPtrOutput { return v.ProviderName }).(pulumi.StringPtrOutput)
}

func (o IncidentOutput) RelatedAnalyticRuleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringArrayOutput { return v.RelatedAnalyticRuleIds }).(pulumi.StringArrayOutput)
}

func (o IncidentOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o IncidentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o IncidentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Incident) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o IncidentOutput) TeamInformation() TeamInformationResponsePtrOutput {
	return o.ApplyT(func(v *Incident) TeamInformationResponsePtrOutput { return v.TeamInformation }).(TeamInformationResponsePtrOutput)
}

func (o IncidentOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

func (o IncidentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Incident) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IncidentOutput{})
}
