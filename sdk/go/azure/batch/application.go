


package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Application struct {
	pulumi.CustomResourceState

	AllowUpdates   pulumi.BoolPtrOutput   `pulumi:"allowUpdates"`
	DefaultVersion pulumi.StringPtrOutput `pulumi:"defaultVersion"`
	DisplayName    pulumi.StringPtrOutput `pulumi:"displayName"`
	Etag           pulumi.StringOutput    `pulumi:"etag"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	Type           pulumi.StringOutput    `pulumi:"type"`
}


func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch/v20151201:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170101:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170501:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:Application"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:batch:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:batch:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	AccountName       string  `pulumi:"accountName"`
	AllowUpdates      *bool   `pulumi:"allowUpdates"`
	ApplicationName   *string `pulumi:"applicationName"`
	DefaultVersion    *string `pulumi:"defaultVersion"`
	DisplayName       *string `pulumi:"displayName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ApplicationArgs struct {
	AccountName       pulumi.StringInput
	AllowUpdates      pulumi.BoolPtrInput
	ApplicationName   pulumi.StringPtrInput
	DefaultVersion    pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
