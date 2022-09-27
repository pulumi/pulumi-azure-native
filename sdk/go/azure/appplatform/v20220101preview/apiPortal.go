


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiPortal struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties ApiPortalPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput              `pulumi:"sku"`
	SystemData SystemDataResponseOutput          `pulumi:"systemData"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewApiPortal(ctx *pulumi.Context,
	name string, args *ApiPortalArgs, opts ...pulumi.ResourceOption) (*ApiPortal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToApiPortalPropertiesPtrOutput().ApplyT(func(v *ApiPortalProperties) *ApiPortalProperties { return v.Defaults() }).(ApiPortalPropertiesPtrOutput)
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ApiPortal"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ApiPortal"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiPortal
	err := ctx.RegisterResource("azure-native:appplatform/v20220101preview:ApiPortal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiPortal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiPortalState, opts ...pulumi.ResourceOption) (*ApiPortal, error) {
	var resource ApiPortal
	err := ctx.ReadResource("azure-native:appplatform/v20220101preview:ApiPortal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiPortalState struct {
}

type ApiPortalState struct {
}

func (ApiPortalState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalState)(nil)).Elem()
}

type apiPortalArgs struct {
	ApiPortalName     *string              `pulumi:"apiPortalName"`
	Properties        *ApiPortalProperties `pulumi:"properties"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
	ServiceName       string               `pulumi:"serviceName"`
	Sku               *Sku                 `pulumi:"sku"`
}


type ApiPortalArgs struct {
	ApiPortalName     pulumi.StringPtrInput
	Properties        ApiPortalPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Sku               SkuPtrInput
}

func (ApiPortalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalArgs)(nil)).Elem()
}

type ApiPortalInput interface {
	pulumi.Input

	ToApiPortalOutput() ApiPortalOutput
	ToApiPortalOutputWithContext(ctx context.Context) ApiPortalOutput
}

func (*ApiPortal) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortal)(nil)).Elem()
}

func (i *ApiPortal) ToApiPortalOutput() ApiPortalOutput {
	return i.ToApiPortalOutputWithContext(context.Background())
}

func (i *ApiPortal) ToApiPortalOutputWithContext(ctx context.Context) ApiPortalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalOutput)
}

type ApiPortalOutput struct{ *pulumi.OutputState }

func (ApiPortalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortal)(nil)).Elem()
}

func (o ApiPortalOutput) ToApiPortalOutput() ApiPortalOutput {
	return o
}

func (o ApiPortalOutput) ToApiPortalOutputWithContext(ctx context.Context) ApiPortalOutput {
	return o
}

func (o ApiPortalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiPortalOutput) Properties() ApiPortalPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiPortal) ApiPortalPropertiesResponseOutput { return v.Properties }).(ApiPortalPropertiesResponseOutput)
}

func (o ApiPortalOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ApiPortal) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ApiPortalOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApiPortal) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApiPortalOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortal) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiPortalOutput{})
}
