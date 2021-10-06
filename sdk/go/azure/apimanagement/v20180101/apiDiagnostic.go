


package v20180101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiDiagnostic struct {
	pulumi.CustomResourceState

	Enabled pulumi.BoolOutput   `pulumi:"enabled"`
	Name    pulumi.StringOutput `pulumi:"name"`
	Type    pulumi.StringOutput `pulumi:"type"`
}


func NewApiDiagnostic(ctx *pulumi.Context,
	name string, args *ApiDiagnosticArgs, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:ApiDiagnostic"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiDiagnostic
	err := ctx.RegisterResource("azure-native:apimanagement/v20180101:ApiDiagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDiagnosticState, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	var resource ApiDiagnostic
	err := ctx.ReadResource("azure-native:apimanagement/v20180101:ApiDiagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiDiagnosticState struct {
}

type ApiDiagnosticState struct {
}

func (ApiDiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticState)(nil)).Elem()
}

type apiDiagnosticArgs struct {
	ApiId             string  `pulumi:"apiId"`
	DiagnosticId      *string `pulumi:"diagnosticId"`
	Enabled           bool    `pulumi:"enabled"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ApiDiagnosticArgs struct {
	ApiId             pulumi.StringInput
	DiagnosticId      pulumi.StringPtrInput
	Enabled           pulumi.BoolInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ApiDiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticArgs)(nil)).Elem()
}

type ApiDiagnosticInput interface {
	pulumi.Input

	ToApiDiagnosticOutput() ApiDiagnosticOutput
	ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput
}

func (*ApiDiagnostic) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDiagnostic)(nil))
}

func (i *ApiDiagnostic) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return i.ToApiDiagnosticOutputWithContext(context.Background())
}

func (i *ApiDiagnostic) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticOutput)
}

type ApiDiagnosticOutput struct{ *pulumi.OutputState }

func (ApiDiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDiagnostic)(nil))
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return o
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiDiagnosticOutput{})
}
