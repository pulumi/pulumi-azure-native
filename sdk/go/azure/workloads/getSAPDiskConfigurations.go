


package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSAPDiskConfigurations(ctx *pulumi.Context, args *GetSAPDiskConfigurationsArgs, opts ...pulumi.InvokeOption) (*GetSAPDiskConfigurationsResult, error) {
	var rv GetSAPDiskConfigurationsResult
	err := ctx.Invoke("azure-native:workloads:getSAPDiskConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSAPDiskConfigurationsArgs struct {
	AppLocation    string `pulumi:"appLocation"`
	DatabaseType   string `pulumi:"databaseType"`
	DbVmSku        string `pulumi:"dbVmSku"`
	DeploymentType string `pulumi:"deploymentType"`
	Environment    string `pulumi:"environment"`
	Location       string `pulumi:"location"`
	SapProduct     string `pulumi:"sapProduct"`
}


type GetSAPDiskConfigurationsResult struct {
	DiskConfigurations []SAPDiskConfigurationResponse `pulumi:"diskConfigurations"`
}

func GetSAPDiskConfigurationsOutput(ctx *pulumi.Context, args GetSAPDiskConfigurationsOutputArgs, opts ...pulumi.InvokeOption) GetSAPDiskConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSAPDiskConfigurationsResult, error) {
			args := v.(GetSAPDiskConfigurationsArgs)
			r, err := GetSAPDiskConfigurations(ctx, &args, opts...)
			var s GetSAPDiskConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSAPDiskConfigurationsResultOutput)
}

type GetSAPDiskConfigurationsOutputArgs struct {
	AppLocation    pulumi.StringInput `pulumi:"appLocation"`
	DatabaseType   pulumi.StringInput `pulumi:"databaseType"`
	DbVmSku        pulumi.StringInput `pulumi:"dbVmSku"`
	DeploymentType pulumi.StringInput `pulumi:"deploymentType"`
	Environment    pulumi.StringInput `pulumi:"environment"`
	Location       pulumi.StringInput `pulumi:"location"`
	SapProduct     pulumi.StringInput `pulumi:"sapProduct"`
}

func (GetSAPDiskConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPDiskConfigurationsArgs)(nil)).Elem()
}


type GetSAPDiskConfigurationsResultOutput struct{ *pulumi.OutputState }

func (GetSAPDiskConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSAPDiskConfigurationsResult)(nil)).Elem()
}

func (o GetSAPDiskConfigurationsResultOutput) ToGetSAPDiskConfigurationsResultOutput() GetSAPDiskConfigurationsResultOutput {
	return o
}

func (o GetSAPDiskConfigurationsResultOutput) ToGetSAPDiskConfigurationsResultOutputWithContext(ctx context.Context) GetSAPDiskConfigurationsResultOutput {
	return o
}

func (o GetSAPDiskConfigurationsResultOutput) DiskConfigurations() SAPDiskConfigurationResponseArrayOutput {
	return o.ApplyT(func(v GetSAPDiskConfigurationsResult) []SAPDiskConfigurationResponse { return v.DiskConfigurations }).(SAPDiskConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSAPDiskConfigurationsResultOutput{})
}
