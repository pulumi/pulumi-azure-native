


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiTagDescription struct {
	pulumi.CustomResourceState

	Description             pulumi.StringPtrOutput `pulumi:"description"`
	DisplayName             pulumi.StringPtrOutput `pulumi:"displayName"`
	ExternalDocsDescription pulumi.StringPtrOutput `pulumi:"externalDocsDescription"`
	ExternalDocsUrl         pulumi.StringPtrOutput `pulumi:"externalDocsUrl"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	TagId                   pulumi.StringPtrOutput `pulumi:"tagId"`
	Type                    pulumi.StringOutput    `pulumi:"type"`
}


func NewApiTagDescription(ctx *pulumi.Context,
	name string, args *ApiTagDescriptionArgs, opts ...pulumi.ResourceOption) (*ApiTagDescription, error) {
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
			Type: pulumi.String("azure-native:apimanagement:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiTagDescription"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiTagDescription"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiTagDescription
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201preview:ApiTagDescription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiTagDescription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiTagDescriptionState, opts ...pulumi.ResourceOption) (*ApiTagDescription, error) {
	var resource ApiTagDescription
	err := ctx.ReadResource("azure-native:apimanagement/v20191201preview:ApiTagDescription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiTagDescriptionState struct {
}

type ApiTagDescriptionState struct {
}

func (ApiTagDescriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiTagDescriptionState)(nil)).Elem()
}

type apiTagDescriptionArgs struct {
	ApiId                   string  `pulumi:"apiId"`
	Description             *string `pulumi:"description"`
	ExternalDocsDescription *string `pulumi:"externalDocsDescription"`
	ExternalDocsUrl         *string `pulumi:"externalDocsUrl"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ServiceName             string  `pulumi:"serviceName"`
	TagDescriptionId        *string `pulumi:"tagDescriptionId"`
}


type ApiTagDescriptionArgs struct {
	ApiId                   pulumi.StringInput
	Description             pulumi.StringPtrInput
	ExternalDocsDescription pulumi.StringPtrInput
	ExternalDocsUrl         pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ServiceName             pulumi.StringInput
	TagDescriptionId        pulumi.StringPtrInput
}

func (ApiTagDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiTagDescriptionArgs)(nil)).Elem()
}

type ApiTagDescriptionInput interface {
	pulumi.Input

	ToApiTagDescriptionOutput() ApiTagDescriptionOutput
	ToApiTagDescriptionOutputWithContext(ctx context.Context) ApiTagDescriptionOutput
}

func (*ApiTagDescription) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiTagDescription)(nil)).Elem()
}

func (i *ApiTagDescription) ToApiTagDescriptionOutput() ApiTagDescriptionOutput {
	return i.ToApiTagDescriptionOutputWithContext(context.Background())
}

func (i *ApiTagDescription) ToApiTagDescriptionOutputWithContext(ctx context.Context) ApiTagDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiTagDescriptionOutput)
}

type ApiTagDescriptionOutput struct{ *pulumi.OutputState }

func (ApiTagDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiTagDescription)(nil)).Elem()
}

func (o ApiTagDescriptionOutput) ToApiTagDescriptionOutput() ApiTagDescriptionOutput {
	return o
}

func (o ApiTagDescriptionOutput) ToApiTagDescriptionOutputWithContext(ctx context.Context) ApiTagDescriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiTagDescriptionOutput{})
}
