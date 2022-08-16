


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListLinkerConfigurations(ctx *pulumi.Context, args *ListLinkerConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListLinkerConfigurationsResult, error) {
	var rv ListLinkerConfigurationsResult
	err := ctx.Invoke("azure-native:servicelinker/v20220101preview:listLinkerConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLinkerConfigurationsArgs struct {
	LinkerName  string `pulumi:"linkerName"`
	ResourceUri string `pulumi:"resourceUri"`
}


type ListLinkerConfigurationsResult struct {
	Configurations []SourceConfigurationResponse `pulumi:"configurations"`
}

func ListLinkerConfigurationsOutput(ctx *pulumi.Context, args ListLinkerConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListLinkerConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListLinkerConfigurationsResult, error) {
			args := v.(ListLinkerConfigurationsArgs)
			r, err := ListLinkerConfigurations(ctx, &args, opts...)
			var s ListLinkerConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListLinkerConfigurationsResultOutput)
}

type ListLinkerConfigurationsOutputArgs struct {
	LinkerName  pulumi.StringInput `pulumi:"linkerName"`
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (ListLinkerConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLinkerConfigurationsArgs)(nil)).Elem()
}


type ListLinkerConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListLinkerConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLinkerConfigurationsResult)(nil)).Elem()
}

func (o ListLinkerConfigurationsResultOutput) ToListLinkerConfigurationsResultOutput() ListLinkerConfigurationsResultOutput {
	return o
}

func (o ListLinkerConfigurationsResultOutput) ToListLinkerConfigurationsResultOutputWithContext(ctx context.Context) ListLinkerConfigurationsResultOutput {
	return o
}

func (o ListLinkerConfigurationsResultOutput) Configurations() SourceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListLinkerConfigurationsResult) []SourceConfigurationResponse { return v.Configurations }).(SourceConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLinkerConfigurationsResultOutput{})
}
