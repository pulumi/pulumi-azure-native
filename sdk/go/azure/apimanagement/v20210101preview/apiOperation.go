


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiOperation struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput               `pulumi:"description"`
	DisplayName        pulumi.StringOutput                  `pulumi:"displayName"`
	Method             pulumi.StringOutput                  `pulumi:"method"`
	Name               pulumi.StringOutput                  `pulumi:"name"`
	Policies           pulumi.StringPtrOutput               `pulumi:"policies"`
	Request            RequestContractResponsePtrOutput     `pulumi:"request"`
	Responses          ResponseContractResponseArrayOutput  `pulumi:"responses"`
	TemplateParameters ParameterContractResponseArrayOutput `pulumi:"templateParameters"`
	Type               pulumi.StringOutput                  `pulumi:"type"`
	UrlTemplate        pulumi.StringOutput                  `pulumi:"urlTemplate"`
}


func NewApiOperation(ctx *pulumi.Context,
	name string, args *ApiOperationArgs, opts ...pulumi.ResourceOption) (*ApiOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Method == nil {
		return nil, errors.New("invalid value for required argument 'Method'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.UrlTemplate == nil {
		return nil, errors.New("invalid value for required argument 'UrlTemplate'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiOperation"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:ApiOperation"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiOperation
	err := ctx.RegisterResource("azure-native:apimanagement/v20210101preview:ApiOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationState, opts ...pulumi.ResourceOption) (*ApiOperation, error) {
	var resource ApiOperation
	err := ctx.ReadResource("azure-native:apimanagement/v20210101preview:ApiOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiOperationState struct {
}

type ApiOperationState struct {
}

func (ApiOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationState)(nil)).Elem()
}

type apiOperationArgs struct {
	ApiId              string              `pulumi:"apiId"`
	Description        *string             `pulumi:"description"`
	DisplayName        string              `pulumi:"displayName"`
	Method             string              `pulumi:"method"`
	OperationId        *string             `pulumi:"operationId"`
	Policies           *string             `pulumi:"policies"`
	Request            *RequestContract    `pulumi:"request"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	Responses          []ResponseContract  `pulumi:"responses"`
	ServiceName        string              `pulumi:"serviceName"`
	TemplateParameters []ParameterContract `pulumi:"templateParameters"`
	UrlTemplate        string              `pulumi:"urlTemplate"`
}


type ApiOperationArgs struct {
	ApiId              pulumi.StringInput
	Description        pulumi.StringPtrInput
	DisplayName        pulumi.StringInput
	Method             pulumi.StringInput
	OperationId        pulumi.StringPtrInput
	Policies           pulumi.StringPtrInput
	Request            RequestContractPtrInput
	ResourceGroupName  pulumi.StringInput
	Responses          ResponseContractArrayInput
	ServiceName        pulumi.StringInput
	TemplateParameters ParameterContractArrayInput
	UrlTemplate        pulumi.StringInput
}

func (ApiOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationArgs)(nil)).Elem()
}

type ApiOperationInput interface {
	pulumi.Input

	ToApiOperationOutput() ApiOperationOutput
	ToApiOperationOutputWithContext(ctx context.Context) ApiOperationOutput
}

func (*ApiOperation) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOperation)(nil))
}

func (i *ApiOperation) ToApiOperationOutput() ApiOperationOutput {
	return i.ToApiOperationOutputWithContext(context.Background())
}

func (i *ApiOperation) ToApiOperationOutputWithContext(ctx context.Context) ApiOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationOutput)
}

type ApiOperationOutput struct{ *pulumi.OutputState }

func (ApiOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiOperation)(nil))
}

func (o ApiOperationOutput) ToApiOperationOutput() ApiOperationOutput {
	return o
}

func (o ApiOperationOutput) ToApiOperationOutputWithContext(ctx context.Context) ApiOperationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiOperationOutput{})
}
