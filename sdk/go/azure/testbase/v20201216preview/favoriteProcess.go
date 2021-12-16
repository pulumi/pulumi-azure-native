


package v20201216preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FavoriteProcess struct {
	pulumi.CustomResourceState

	ActualProcessName pulumi.StringOutput      `pulumi:"actualProcessName"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewFavoriteProcess(ctx *pulumi.Context,
	name string, args *FavoriteProcessArgs, opts ...pulumi.ResourceOption) (*FavoriteProcess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActualProcessName == nil {
		return nil, errors.New("invalid value for required argument 'ActualProcessName'")
	}
	if args.PackageName == nil {
		return nil, errors.New("invalid value for required argument 'PackageName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TestBaseAccountName == nil {
		return nil, errors.New("invalid value for required argument 'TestBaseAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:testbase:FavoriteProcess"),
		},
	})
	opts = append(opts, aliases)
	var resource FavoriteProcess
	err := ctx.RegisterResource("azure-native:testbase/v20201216preview:FavoriteProcess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFavoriteProcess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FavoriteProcessState, opts ...pulumi.ResourceOption) (*FavoriteProcess, error) {
	var resource FavoriteProcess
	err := ctx.ReadResource("azure-native:testbase/v20201216preview:FavoriteProcess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type favoriteProcessState struct {
}

type FavoriteProcessState struct {
}

func (FavoriteProcessState) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteProcessState)(nil)).Elem()
}

type favoriteProcessArgs struct {
	ActualProcessName           string  `pulumi:"actualProcessName"`
	FavoriteProcessResourceName *string `pulumi:"favoriteProcessResourceName"`
	PackageName                 string  `pulumi:"packageName"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
	TestBaseAccountName         string  `pulumi:"testBaseAccountName"`
}


type FavoriteProcessArgs struct {
	ActualProcessName           pulumi.StringInput
	FavoriteProcessResourceName pulumi.StringPtrInput
	PackageName                 pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	TestBaseAccountName         pulumi.StringInput
}

func (FavoriteProcessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteProcessArgs)(nil)).Elem()
}

type FavoriteProcessInput interface {
	pulumi.Input

	ToFavoriteProcessOutput() FavoriteProcessOutput
	ToFavoriteProcessOutputWithContext(ctx context.Context) FavoriteProcessOutput
}

func (*FavoriteProcess) ElementType() reflect.Type {
	return reflect.TypeOf((**FavoriteProcess)(nil)).Elem()
}

func (i *FavoriteProcess) ToFavoriteProcessOutput() FavoriteProcessOutput {
	return i.ToFavoriteProcessOutputWithContext(context.Background())
}

func (i *FavoriteProcess) ToFavoriteProcessOutputWithContext(ctx context.Context) FavoriteProcessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FavoriteProcessOutput)
}

type FavoriteProcessOutput struct{ *pulumi.OutputState }

func (FavoriteProcessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FavoriteProcess)(nil)).Elem()
}

func (o FavoriteProcessOutput) ToFavoriteProcessOutput() FavoriteProcessOutput {
	return o
}

func (o FavoriteProcessOutput) ToFavoriteProcessOutputWithContext(ctx context.Context) FavoriteProcessOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FavoriteProcessOutput{})
}
