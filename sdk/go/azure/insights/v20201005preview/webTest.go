


package v20201005preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebTest struct {
	pulumi.CustomResourceState

	Configuration      WebTestPropertiesResponseConfigurationPtrOutput   `pulumi:"configuration"`
	Description        pulumi.StringPtrOutput                            `pulumi:"description"`
	Enabled            pulumi.BoolPtrOutput                              `pulumi:"enabled"`
	Frequency          pulumi.IntPtrOutput                               `pulumi:"frequency"`
	Kind               pulumi.StringPtrOutput                            `pulumi:"kind"`
	Location           pulumi.StringOutput                               `pulumi:"location"`
	Locations          WebTestGeolocationResponseArrayOutput             `pulumi:"locations"`
	Name               pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput                               `pulumi:"provisioningState"`
	Request            WebTestPropertiesResponseRequestPtrOutput         `pulumi:"request"`
	RetryEnabled       pulumi.BoolPtrOutput                              `pulumi:"retryEnabled"`
	SyntheticMonitorId pulumi.StringOutput                               `pulumi:"syntheticMonitorId"`
	Tags               pulumi.StringMapOutput                            `pulumi:"tags"`
	Timeout            pulumi.IntPtrOutput                               `pulumi:"timeout"`
	Type               pulumi.StringOutput                               `pulumi:"type"`
	ValidationRules    WebTestPropertiesResponseValidationRulesPtrOutput `pulumi:"validationRules"`
	WebTestKind        pulumi.StringOutput                               `pulumi:"webTestKind"`
	WebTestName        pulumi.StringOutput                               `pulumi:"webTestName"`
}


func NewWebTest(ctx *pulumi.Context,
	name string, args *WebTestArgs, opts ...pulumi.ResourceOption) (*WebTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SyntheticMonitorId == nil {
		return nil, errors.New("invalid value for required argument 'SyntheticMonitorId'")
	}
	if isZero(args.Frequency) {
		args.Frequency = pulumi.IntPtr(300)
	}
	if isZero(args.Kind) {
		args.Kind = WebTestKind("ping")
	}
	if isZero(args.Timeout) {
		args.Timeout = pulumi.IntPtr(30)
	}
	if isZero(args.WebTestKind) {
		args.WebTestKind = WebTestKindEnum("ping")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:WebTest"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20150501:WebTest"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180501preview:WebTest"),
		},
	})
	opts = append(opts, aliases)
	var resource WebTest
	err := ctx.RegisterResource("azure-native:insights/v20201005preview:WebTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTestState, opts ...pulumi.ResourceOption) (*WebTest, error) {
	var resource WebTest
	err := ctx.ReadResource("azure-native:insights/v20201005preview:WebTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webTestState struct {
}

type WebTestState struct {
}

func (WebTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestState)(nil)).Elem()
}

type webTestArgs struct {
	Configuration      *WebTestPropertiesConfiguration   `pulumi:"configuration"`
	Description        *string                           `pulumi:"description"`
	Enabled            *bool                             `pulumi:"enabled"`
	Frequency          *int                              `pulumi:"frequency"`
	Kind               *WebTestKind                      `pulumi:"kind"`
	Location           *string                           `pulumi:"location"`
	Locations          []WebTestGeolocation              `pulumi:"locations"`
	Request            *WebTestPropertiesRequest         `pulumi:"request"`
	ResourceGroupName  string                            `pulumi:"resourceGroupName"`
	RetryEnabled       *bool                             `pulumi:"retryEnabled"`
	SyntheticMonitorId string                            `pulumi:"syntheticMonitorId"`
	Tags               map[string]string                 `pulumi:"tags"`
	Timeout            *int                              `pulumi:"timeout"`
	ValidationRules    *WebTestPropertiesValidationRules `pulumi:"validationRules"`
	WebTestKind        WebTestKindEnum                   `pulumi:"webTestKind"`
	WebTestName        *string                           `pulumi:"webTestName"`
}


type WebTestArgs struct {
	Configuration      WebTestPropertiesConfigurationPtrInput
	Description        pulumi.StringPtrInput
	Enabled            pulumi.BoolPtrInput
	Frequency          pulumi.IntPtrInput
	Kind               WebTestKindPtrInput
	Location           pulumi.StringPtrInput
	Locations          WebTestGeolocationArrayInput
	Request            WebTestPropertiesRequestPtrInput
	ResourceGroupName  pulumi.StringInput
	RetryEnabled       pulumi.BoolPtrInput
	SyntheticMonitorId pulumi.StringInput
	Tags               pulumi.StringMapInput
	Timeout            pulumi.IntPtrInput
	ValidationRules    WebTestPropertiesValidationRulesPtrInput
	WebTestKind        WebTestKindEnumInput
	WebTestName        pulumi.StringPtrInput
}

func (WebTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestArgs)(nil)).Elem()
}

type WebTestInput interface {
	pulumi.Input

	ToWebTestOutput() WebTestOutput
	ToWebTestOutputWithContext(ctx context.Context) WebTestOutput
}

func (*WebTest) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTest)(nil)).Elem()
}

func (i *WebTest) ToWebTestOutput() WebTestOutput {
	return i.ToWebTestOutputWithContext(context.Background())
}

func (i *WebTest) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestOutput)
}

type WebTestOutput struct{ *pulumi.OutputState }

func (WebTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTest)(nil)).Elem()
}

func (o WebTestOutput) ToWebTestOutput() WebTestOutput {
	return o
}

func (o WebTestOutput) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebTestOutput{})
}
