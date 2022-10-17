


package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiVersionSet struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName       pulumi.StringOutput    `pulumi:"displayName"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	Type              pulumi.StringOutput    `pulumi:"type"`
	VersionHeaderName pulumi.StringPtrOutput `pulumi:"versionHeaderName"`
	VersionQueryName  pulumi.StringPtrOutput `pulumi:"versionQueryName"`
	VersioningScheme  pulumi.StringOutput    `pulumi:"versioningScheme"`
}


func NewApiVersionSet(ctx *pulumi.Context,
	name string, args *ApiVersionSetArgs, opts ...pulumi.ResourceOption) (*ApiVersionSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.VersioningScheme == nil {
		return nil, errors.New("invalid value for required argument 'VersioningScheme'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiVersionSet"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiVersionSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiVersionSet
	err := ctx.RegisterResource("azure-native:apimanagement/v20200601preview:ApiVersionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiVersionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiVersionSetState, opts ...pulumi.ResourceOption) (*ApiVersionSet, error) {
	var resource ApiVersionSet
	err := ctx.ReadResource("azure-native:apimanagement/v20200601preview:ApiVersionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiVersionSetState struct {
}

type ApiVersionSetState struct {
}

func (ApiVersionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiVersionSetState)(nil)).Elem()
}

type apiVersionSetArgs struct {
	Description       *string `pulumi:"description"`
	DisplayName       string  `pulumi:"displayName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	VersionQueryName  *string `pulumi:"versionQueryName"`
	VersionSetId      *string `pulumi:"versionSetId"`
	VersioningScheme  string  `pulumi:"versioningScheme"`
}


type ApiVersionSetArgs struct {
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	VersionHeaderName pulumi.StringPtrInput
	VersionQueryName  pulumi.StringPtrInput
	VersionSetId      pulumi.StringPtrInput
	VersioningScheme  pulumi.StringInput
}

func (ApiVersionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiVersionSetArgs)(nil)).Elem()
}

type ApiVersionSetInput interface {
	pulumi.Input

	ToApiVersionSetOutput() ApiVersionSetOutput
	ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput
}

func (*ApiVersionSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSet)(nil)).Elem()
}

func (i *ApiVersionSet) ToApiVersionSetOutput() ApiVersionSetOutput {
	return i.ToApiVersionSetOutputWithContext(context.Background())
}

func (i *ApiVersionSet) ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetOutput)
}

type ApiVersionSetOutput struct{ *pulumi.OutputState }

func (ApiVersionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSet)(nil)).Elem()
}

func (o ApiVersionSetOutput) ToApiVersionSetOutput() ApiVersionSetOutput {
	return o
}

func (o ApiVersionSetOutput) ToApiVersionSetOutputWithContext(ctx context.Context) ApiVersionSetOutput {
	return o
}

func (o ApiVersionSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiVersionSet) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ApiVersionSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiVersionSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiVersionSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiVersionSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApiVersionSetOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSet) pulumi.StringPtrOutput { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSet) pulumi.StringPtrOutput { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetOutput) VersioningScheme() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiVersionSet) pulumi.StringOutput { return v.VersioningScheme }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiVersionSetOutput{})
}
