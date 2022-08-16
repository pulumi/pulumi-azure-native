


package devspaces

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListControllerConnectionDetails(ctx *pulumi.Context, args *ListControllerConnectionDetailsArgs, opts ...pulumi.InvokeOption) (*ListControllerConnectionDetailsResult, error) {
	var rv ListControllerConnectionDetailsResult
	err := ctx.Invoke("azure-native:devspaces:listControllerConnectionDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListControllerConnectionDetailsArgs struct {
	Name                          string `pulumi:"name"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	TargetContainerHostResourceId string `pulumi:"targetContainerHostResourceId"`
}

type ListControllerConnectionDetailsResult struct {
	ConnectionDetailsList []ControllerConnectionDetailsResponse `pulumi:"connectionDetailsList"`
}

func ListControllerConnectionDetailsOutput(ctx *pulumi.Context, args ListControllerConnectionDetailsOutputArgs, opts ...pulumi.InvokeOption) ListControllerConnectionDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListControllerConnectionDetailsResult, error) {
			args := v.(ListControllerConnectionDetailsArgs)
			r, err := ListControllerConnectionDetails(ctx, &args, opts...)
			var s ListControllerConnectionDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListControllerConnectionDetailsResultOutput)
}

type ListControllerConnectionDetailsOutputArgs struct {
	Name                          pulumi.StringInput `pulumi:"name"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	TargetContainerHostResourceId pulumi.StringInput `pulumi:"targetContainerHostResourceId"`
}

func (ListControllerConnectionDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListControllerConnectionDetailsArgs)(nil)).Elem()
}

type ListControllerConnectionDetailsResultOutput struct{ *pulumi.OutputState }

func (ListControllerConnectionDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListControllerConnectionDetailsResult)(nil)).Elem()
}

func (o ListControllerConnectionDetailsResultOutput) ToListControllerConnectionDetailsResultOutput() ListControllerConnectionDetailsResultOutput {
	return o
}

func (o ListControllerConnectionDetailsResultOutput) ToListControllerConnectionDetailsResultOutputWithContext(ctx context.Context) ListControllerConnectionDetailsResultOutput {
	return o
}

func (o ListControllerConnectionDetailsResultOutput) ConnectionDetailsList() ControllerConnectionDetailsResponseArrayOutput {
	return o.ApplyT(func(v ListControllerConnectionDetailsResult) []ControllerConnectionDetailsResponse {
		return v.ConnectionDetailsList
	}).(ControllerConnectionDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListControllerConnectionDetailsResultOutput{})
}
