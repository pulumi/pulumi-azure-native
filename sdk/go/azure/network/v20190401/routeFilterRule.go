


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RouteFilterRule struct {
	pulumi.CustomResourceState

	Access              pulumi.StringOutput      `pulumi:"access"`
	Communities         pulumi.StringArrayOutput `pulumi:"communities"`
	Etag                pulumi.StringOutput      `pulumi:"etag"`
	Location            pulumi.StringPtrOutput   `pulumi:"location"`
	Name                pulumi.StringPtrOutput   `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput      `pulumi:"provisioningState"`
	RouteFilterRuleType pulumi.StringOutput      `pulumi:"routeFilterRuleType"`
}


func NewRouteFilterRule(ctx *pulumi.Context,
	name string, args *RouteFilterRuleArgs, opts ...pulumi.ResourceOption) (*RouteFilterRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Access == nil {
		return nil, errors.New("invalid value for required argument 'Access'")
	}
	if args.Communities == nil {
		return nil, errors.New("invalid value for required argument 'Communities'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RouteFilterName == nil {
		return nil, errors.New("invalid value for required argument 'RouteFilterName'")
	}
	if args.RouteFilterRuleType == nil {
		return nil, errors.New("invalid value for required argument 'RouteFilterRuleType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:RouteFilterRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:RouteFilterRule"),
		},
	})
	opts = append(opts, aliases)
	var resource RouteFilterRule
	err := ctx.RegisterResource("azure-native:network/v20190401:RouteFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRouteFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteFilterRuleState, opts ...pulumi.ResourceOption) (*RouteFilterRule, error) {
	var resource RouteFilterRule
	err := ctx.ReadResource("azure-native:network/v20190401:RouteFilterRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type routeFilterRuleState struct {
}

type RouteFilterRuleState struct {
}

func (RouteFilterRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterRuleState)(nil)).Elem()
}

type routeFilterRuleArgs struct {
	Access              string   `pulumi:"access"`
	Communities         []string `pulumi:"communities"`
	Id                  *string  `pulumi:"id"`
	Location            *string  `pulumi:"location"`
	Name                *string  `pulumi:"name"`
	ResourceGroupName   string   `pulumi:"resourceGroupName"`
	RouteFilterName     string   `pulumi:"routeFilterName"`
	RouteFilterRuleType string   `pulumi:"routeFilterRuleType"`
	RuleName            *string  `pulumi:"ruleName"`
}


type RouteFilterRuleArgs struct {
	Access              pulumi.StringInput
	Communities         pulumi.StringArrayInput
	Id                  pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	RouteFilterName     pulumi.StringInput
	RouteFilterRuleType pulumi.StringInput
	RuleName            pulumi.StringPtrInput
}

func (RouteFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeFilterRuleArgs)(nil)).Elem()
}

type RouteFilterRuleInput interface {
	pulumi.Input

	ToRouteFilterRuleOutput() RouteFilterRuleOutput
	ToRouteFilterRuleOutputWithContext(ctx context.Context) RouteFilterRuleOutput
}

func (*RouteFilterRule) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterRule)(nil)).Elem()
}

func (i *RouteFilterRule) ToRouteFilterRuleOutput() RouteFilterRuleOutput {
	return i.ToRouteFilterRuleOutputWithContext(context.Background())
}

func (i *RouteFilterRule) ToRouteFilterRuleOutputWithContext(ctx context.Context) RouteFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterRuleOutput)
}

type RouteFilterRuleOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterRule)(nil)).Elem()
}

func (o RouteFilterRuleOutput) ToRouteFilterRuleOutput() RouteFilterRuleOutput {
	return o
}

func (o RouteFilterRuleOutput) ToRouteFilterRuleOutputWithContext(ctx context.Context) RouteFilterRuleOutput {
	return o
}

func (o RouteFilterRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilterRule) pulumi.StringOutput { return v.Access }).(pulumi.StringOutput)
}

func (o RouteFilterRuleOutput) Communities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteFilterRule) pulumi.StringArrayOutput { return v.Communities }).(pulumi.StringArrayOutput)
}

func (o RouteFilterRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilterRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteFilterRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilterRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteFilterRuleOutput) RouteFilterRuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteFilterRule) pulumi.StringOutput { return v.RouteFilterRuleType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteFilterRuleOutput{})
}
