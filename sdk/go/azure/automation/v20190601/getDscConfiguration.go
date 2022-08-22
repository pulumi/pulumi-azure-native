


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDscConfiguration(ctx *pulumi.Context, args *LookupDscConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscConfigurationResult, error) {
	var rv LookupDscConfigurationResult
	err := ctx.Invoke("azure-native:automation/v20190601:getDscConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscConfigurationArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ConfigurationName     string `pulumi:"configurationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDscConfigurationResult struct {
	CreationTime           *string                                      `pulumi:"creationTime"`
	Description            *string                                      `pulumi:"description"`
	Etag                   *string                                      `pulumi:"etag"`
	Id                     string                                       `pulumi:"id"`
	JobCount               *int                                         `pulumi:"jobCount"`
	LastModifiedTime       *string                                      `pulumi:"lastModifiedTime"`
	Location               *string                                      `pulumi:"location"`
	LogVerbose             *bool                                        `pulumi:"logVerbose"`
	Name                   string                                       `pulumi:"name"`
	NodeConfigurationCount *int                                         `pulumi:"nodeConfigurationCount"`
	Parameters             map[string]DscConfigurationParameterResponse `pulumi:"parameters"`
	ProvisioningState      *string                                      `pulumi:"provisioningState"`
	Source                 *ContentSourceResponse                       `pulumi:"source"`
	State                  *string                                      `pulumi:"state"`
	Tags                   map[string]string                            `pulumi:"tags"`
	Type                   string                                       `pulumi:"type"`
}

func LookupDscConfigurationOutput(ctx *pulumi.Context, args LookupDscConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupDscConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDscConfigurationResult, error) {
			args := v.(LookupDscConfigurationArgs)
			r, err := LookupDscConfiguration(ctx, &args, opts...)
			var s LookupDscConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDscConfigurationResultOutput)
}

type LookupDscConfigurationOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ConfigurationName     pulumi.StringInput `pulumi:"configurationName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDscConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscConfigurationArgs)(nil)).Elem()
}


type LookupDscConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDscConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscConfigurationResult)(nil)).Elem()
}

func (o LookupDscConfigurationResultOutput) ToLookupDscConfigurationResultOutput() LookupDscConfigurationResultOutput {
	return o
}

func (o LookupDscConfigurationResultOutput) ToLookupDscConfigurationResultOutputWithContext(ctx context.Context) LookupDscConfigurationResultOutput {
	return o
}

func (o LookupDscConfigurationResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDscConfigurationResultOutput) JobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *int { return v.JobCount }).(pulumi.IntPtrOutput)
}

func (o LookupDscConfigurationResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupDscConfigurationResultOutput) LogVerbose() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *bool { return v.LogVerbose }).(pulumi.BoolPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDscConfigurationResultOutput) NodeConfigurationCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *int { return v.NodeConfigurationCount }).(pulumi.IntPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Parameters() DscConfigurationParameterResponseMapOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) map[string]DscConfigurationParameterResponse { return v.Parameters }).(DscConfigurationParameterResponseMapOutput)
}

func (o LookupDscConfigurationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Source() ContentSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *ContentSourceResponse { return v.Source }).(ContentSourceResponsePtrOutput)
}

func (o LookupDscConfigurationResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupDscConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDscConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDscConfigurationResultOutput{})
}
