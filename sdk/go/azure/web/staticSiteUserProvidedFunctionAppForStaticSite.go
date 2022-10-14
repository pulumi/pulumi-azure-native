


package web

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticSiteUserProvidedFunctionAppForStaticSite struct {
	pulumi.CustomResourceState

	CreatedOn             pulumi.StringOutput    `pulumi:"createdOn"`
	FunctionAppRegion     pulumi.StringPtrOutput `pulumi:"functionAppRegion"`
	FunctionAppResourceId pulumi.StringPtrOutput `pulumi:"functionAppResourceId"`
	Kind                  pulumi.StringPtrOutput `pulumi:"kind"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewStaticSiteUserProvidedFunctionAppForStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteUserProvidedFunctionAppForStaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSiteUserProvidedFunctionAppForStaticSite
	err := ctx.RegisterResource("azure-native:web:StaticSiteUserProvidedFunctionAppForStaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticSiteUserProvidedFunctionAppForStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteUserProvidedFunctionAppForStaticSiteState, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSite, error) {
	var resource StaticSiteUserProvidedFunctionAppForStaticSite
	err := ctx.ReadResource("azure-native:web:StaticSiteUserProvidedFunctionAppForStaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticSiteUserProvidedFunctionAppForStaticSiteState struct {
}

type StaticSiteUserProvidedFunctionAppForStaticSiteState struct {
}

func (StaticSiteUserProvidedFunctionAppForStaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteState)(nil)).Elem()
}

type staticSiteUserProvidedFunctionAppForStaticSiteArgs struct {
	FunctionAppName       *string `pulumi:"functionAppName"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	IsForced              *bool   `pulumi:"isForced"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type StaticSiteUserProvidedFunctionAppForStaticSiteArgs struct {
	FunctionAppName       pulumi.StringPtrInput
	FunctionAppRegion     pulumi.StringPtrInput
	FunctionAppResourceId pulumi.StringPtrInput
	IsForced              pulumi.BoolPtrInput
	Kind                  pulumi.StringPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (StaticSiteUserProvidedFunctionAppForStaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteArgs)(nil)).Elem()
}

type StaticSiteUserProvidedFunctionAppForStaticSiteInput interface {
	pulumi.Input

	ToStaticSiteUserProvidedFunctionAppForStaticSiteOutput() StaticSiteUserProvidedFunctionAppForStaticSiteOutput
	ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteOutput
}

func (*StaticSiteUserProvidedFunctionAppForStaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSite)(nil)).Elem()
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSite) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutput() StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return i.ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(context.Background())
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSite) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteUserProvidedFunctionAppForStaticSiteOutput)
}

type StaticSiteUserProvidedFunctionAppForStaticSiteOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppForStaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSite)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutput() StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringPtrOutput {
		return v.FunctionAppRegion
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringPtrOutput {
		return v.FunctionAppResourceId
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppForStaticSiteOutput{})
}
