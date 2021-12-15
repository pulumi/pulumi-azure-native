


package v20150501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Component struct {
	pulumi.CustomResourceState

	AppId                      pulumi.StringOutput                          `pulumi:"appId"`
	ApplicationId              pulumi.StringOutput                          `pulumi:"applicationId"`
	ApplicationType            pulumi.StringOutput                          `pulumi:"applicationType"`
	ConnectionString           pulumi.StringOutput                          `pulumi:"connectionString"`
	CreationDate               pulumi.StringOutput                          `pulumi:"creationDate"`
	DisableIpMasking           pulumi.BoolPtrOutput                         `pulumi:"disableIpMasking"`
	FlowType                   pulumi.StringPtrOutput                       `pulumi:"flowType"`
	HockeyAppId                pulumi.StringPtrOutput                       `pulumi:"hockeyAppId"`
	HockeyAppToken             pulumi.StringOutput                          `pulumi:"hockeyAppToken"`
	ImmediatePurgeDataOn30Days pulumi.BoolPtrOutput                         `pulumi:"immediatePurgeDataOn30Days"`
	IngestionMode              pulumi.StringPtrOutput                       `pulumi:"ingestionMode"`
	InstrumentationKey         pulumi.StringOutput                          `pulumi:"instrumentationKey"`
	Kind                       pulumi.StringOutput                          `pulumi:"kind"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateLinkScopedResources PrivateLinkScopedResourceResponseArrayOutput `pulumi:"privateLinkScopedResources"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	RequestSource              pulumi.StringPtrOutput                       `pulumi:"requestSource"`
	RetentionInDays            pulumi.IntPtrOutput                          `pulumi:"retentionInDays"`
	SamplingPercentage         pulumi.Float64PtrOutput                      `pulumi:"samplingPercentage"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	TenantId                   pulumi.StringOutput                          `pulumi:"tenantId"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.ApplicationType) {
		args.ApplicationType = pulumi.String("web")
	}
	if isZero(args.FlowType) {
		args.FlowType = pulumi.StringPtr("Bluefield")
	}
	if isZero(args.IngestionMode) {
		args.IngestionMode = pulumi.StringPtr("ApplicationInsights")
	}
	if isZero(args.RequestSource) {
		args.RequestSource = pulumi.StringPtr("rest")
	}
	if isZero(args.RetentionInDays) {
		args.RetentionInDays = pulumi.IntPtr(90)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:Component"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180501preview:Component"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200202:Component"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20200202preview:Component"),
		},
	})
	opts = append(opts, aliases)
	var resource Component
	err := ctx.RegisterResource("azure-native:insights/v20150501:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentState, opts ...pulumi.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("azure-native:insights/v20150501:Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type componentState struct {
}

type ComponentState struct {
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	ApplicationType            string            `pulumi:"applicationType"`
	DisableIpMasking           *bool             `pulumi:"disableIpMasking"`
	FlowType                   *string           `pulumi:"flowType"`
	HockeyAppId                *string           `pulumi:"hockeyAppId"`
	ImmediatePurgeDataOn30Days *bool             `pulumi:"immediatePurgeDataOn30Days"`
	IngestionMode              *string           `pulumi:"ingestionMode"`
	Kind                       string            `pulumi:"kind"`
	Location                   *string           `pulumi:"location"`
	RequestSource              *string           `pulumi:"requestSource"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	ResourceName               *string           `pulumi:"resourceName"`
	RetentionInDays            *int              `pulumi:"retentionInDays"`
	SamplingPercentage         *float64          `pulumi:"samplingPercentage"`
	Tags                       map[string]string `pulumi:"tags"`
}


type ComponentArgs struct {
	ApplicationType            pulumi.StringInput
	DisableIpMasking           pulumi.BoolPtrInput
	FlowType                   pulumi.StringPtrInput
	HockeyAppId                pulumi.StringPtrInput
	ImmediatePurgeDataOn30Days pulumi.BoolPtrInput
	IngestionMode              pulumi.StringPtrInput
	Kind                       pulumi.StringInput
	Location                   pulumi.StringPtrInput
	RequestSource              pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	ResourceName               pulumi.StringPtrInput
	RetentionInDays            pulumi.IntPtrInput
	SamplingPercentage         pulumi.Float64PtrInput
	Tags                       pulumi.StringMapInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

type ComponentOutput struct{ *pulumi.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ComponentOutput{})
}
