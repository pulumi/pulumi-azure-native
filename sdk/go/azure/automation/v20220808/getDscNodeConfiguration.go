


package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDscNodeConfiguration(ctx *pulumi.Context, args *LookupDscNodeConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscNodeConfigurationResult, error) {
	var rv LookupDscNodeConfigurationResult
	err := ctx.Invoke("azure-native:automation/v20220808:getDscNodeConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscNodeConfigurationArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	NodeConfigurationName string `pulumi:"nodeConfigurationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDscNodeConfigurationResult struct {
	Configuration                   *DscConfigurationAssociationPropertyResponse `pulumi:"configuration"`
	CreationTime                    *string                                      `pulumi:"creationTime"`
	Id                              string                                       `pulumi:"id"`
	IncrementNodeConfigurationBuild *bool                                        `pulumi:"incrementNodeConfigurationBuild"`
	LastModifiedTime                *string                                      `pulumi:"lastModifiedTime"`
	Name                            string                                       `pulumi:"name"`
	NodeCount                       *float64                                     `pulumi:"nodeCount"`
	Source                          *string                                      `pulumi:"source"`
	Type                            string                                       `pulumi:"type"`
}

func LookupDscNodeConfigurationOutput(ctx *pulumi.Context, args LookupDscNodeConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupDscNodeConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDscNodeConfigurationResult, error) {
			args := v.(LookupDscNodeConfigurationArgs)
			r, err := LookupDscNodeConfiguration(ctx, &args, opts...)
			var s LookupDscNodeConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDscNodeConfigurationResultOutput)
}

type LookupDscNodeConfigurationOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	NodeConfigurationName pulumi.StringInput `pulumi:"nodeConfigurationName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDscNodeConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscNodeConfigurationArgs)(nil)).Elem()
}


type LookupDscNodeConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDscNodeConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscNodeConfigurationResult)(nil)).Elem()
}

func (o LookupDscNodeConfigurationResultOutput) ToLookupDscNodeConfigurationResultOutput() LookupDscNodeConfigurationResultOutput {
	return o
}

func (o LookupDscNodeConfigurationResultOutput) ToLookupDscNodeConfigurationResultOutputWithContext(ctx context.Context) LookupDscNodeConfigurationResultOutput {
	return o
}

func (o LookupDscNodeConfigurationResultOutput) Configuration() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *DscConfigurationAssociationPropertyResponse {
		return v.Configuration
	}).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

func (o LookupDscNodeConfigurationResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupDscNodeConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDscNodeConfigurationResultOutput) IncrementNodeConfigurationBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *bool { return v.IncrementNodeConfigurationBuild }).(pulumi.BoolPtrOutput)
}

func (o LookupDscNodeConfigurationResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupDscNodeConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDscNodeConfigurationResultOutput) NodeCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *float64 { return v.NodeCount }).(pulumi.Float64PtrOutput)
}

func (o LookupDscNodeConfigurationResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o LookupDscNodeConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDscNodeConfigurationResultOutput{})
}
