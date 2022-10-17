


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiRelease struct {
	pulumi.CustomResourceState

	ApiId           pulumi.StringPtrOutput `pulumi:"apiId"`
	CreatedDateTime pulumi.StringOutput    `pulumi:"createdDateTime"`
	Name            pulumi.StringOutput    `pulumi:"name"`
	Notes           pulumi.StringPtrOutput `pulumi:"notes"`
	Type            pulumi.StringOutput    `pulumi:"type"`
	UpdatedDateTime pulumi.StringOutput    `pulumi:"updatedDateTime"`
}


func NewApiRelease(ctx *pulumi.Context,
	name string, args *ApiReleaseArgs, opts ...pulumi.ResourceOption) (*ApiRelease, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiRelease"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiRelease
	err := ctx.RegisterResource("azure-native:apimanagement/v20210101preview:ApiRelease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiReleaseState, opts ...pulumi.ResourceOption) (*ApiRelease, error) {
	var resource ApiRelease
	err := ctx.ReadResource("azure-native:apimanagement/v20210101preview:ApiRelease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiReleaseState struct {
}

type ApiReleaseState struct {
}

func (ApiReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiReleaseState)(nil)).Elem()
}

type apiReleaseArgs struct {
	ApiId             string  `pulumi:"apiId"`
	Notes             *string `pulumi:"notes"`
	ReleaseId         *string `pulumi:"releaseId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ApiReleaseArgs struct {
	ApiId             pulumi.StringInput
	Notes             pulumi.StringPtrInput
	ReleaseId         pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ApiReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiReleaseArgs)(nil)).Elem()
}

type ApiReleaseInput interface {
	pulumi.Input

	ToApiReleaseOutput() ApiReleaseOutput
	ToApiReleaseOutputWithContext(ctx context.Context) ApiReleaseOutput
}

func (*ApiRelease) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiRelease)(nil)).Elem()
}

func (i *ApiRelease) ToApiReleaseOutput() ApiReleaseOutput {
	return i.ToApiReleaseOutputWithContext(context.Background())
}

func (i *ApiRelease) ToApiReleaseOutputWithContext(ctx context.Context) ApiReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiReleaseOutput)
}

type ApiReleaseOutput struct{ *pulumi.OutputState }

func (ApiReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiRelease)(nil)).Elem()
}

func (o ApiReleaseOutput) ToApiReleaseOutput() ApiReleaseOutput {
	return o
}

func (o ApiReleaseOutput) ToApiReleaseOutputWithContext(ctx context.Context) ApiReleaseOutput {
	return o
}

func (o ApiReleaseOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiRelease) pulumi.StringPtrOutput { return v.ApiId }).(pulumi.StringPtrOutput)
}

func (o ApiReleaseOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiRelease) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o ApiReleaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiRelease) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiReleaseOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiRelease) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o ApiReleaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiRelease) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApiReleaseOutput) UpdatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiRelease) pulumi.StringOutput { return v.UpdatedDateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiReleaseOutput{})
}
