


package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSoftwareUpdateConfigurationByName(ctx *pulumi.Context, args *LookupSoftwareUpdateConfigurationByNameArgs, opts ...pulumi.InvokeOption) (*LookupSoftwareUpdateConfigurationByNameResult, error) {
	var rv LookupSoftwareUpdateConfigurationByNameResult
	err := ctx.Invoke("azure-native:automation:getSoftwareUpdateConfigurationByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSoftwareUpdateConfigurationByNameArgs struct {
	AutomationAccountName           string `pulumi:"automationAccountName"`
	ResourceGroupName               string `pulumi:"resourceGroupName"`
	SoftwareUpdateConfigurationName string `pulumi:"softwareUpdateConfigurationName"`
}


type LookupSoftwareUpdateConfigurationByNameResult struct {
	CreatedBy           string                                    `pulumi:"createdBy"`
	CreationTime        string                                    `pulumi:"creationTime"`
	Error               *ErrorResponseResponse                    `pulumi:"error"`
	Id                  string                                    `pulumi:"id"`
	LastModifiedBy      string                                    `pulumi:"lastModifiedBy"`
	LastModifiedTime    string                                    `pulumi:"lastModifiedTime"`
	Name                string                                    `pulumi:"name"`
	ProvisioningState   string                                    `pulumi:"provisioningState"`
	ScheduleInfo        SUCSchedulePropertiesResponse             `pulumi:"scheduleInfo"`
	Tasks               *SoftwareUpdateConfigurationTasksResponse `pulumi:"tasks"`
	Type                string                                    `pulumi:"type"`
	UpdateConfiguration UpdateConfigurationResponse               `pulumi:"updateConfiguration"`
}


func (val *LookupSoftwareUpdateConfigurationByNameResult) Defaults() *LookupSoftwareUpdateConfigurationByNameResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ScheduleInfo = *tmp.ScheduleInfo.Defaults()

	return &tmp
}

func LookupSoftwareUpdateConfigurationByNameOutput(ctx *pulumi.Context, args LookupSoftwareUpdateConfigurationByNameOutputArgs, opts ...pulumi.InvokeOption) LookupSoftwareUpdateConfigurationByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSoftwareUpdateConfigurationByNameResult, error) {
			args := v.(LookupSoftwareUpdateConfigurationByNameArgs)
			r, err := LookupSoftwareUpdateConfigurationByName(ctx, &args, opts...)
			var s LookupSoftwareUpdateConfigurationByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSoftwareUpdateConfigurationByNameResultOutput)
}

type LookupSoftwareUpdateConfigurationByNameOutputArgs struct {
	AutomationAccountName           pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName               pulumi.StringInput `pulumi:"resourceGroupName"`
	SoftwareUpdateConfigurationName pulumi.StringInput `pulumi:"softwareUpdateConfigurationName"`
}

func (LookupSoftwareUpdateConfigurationByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSoftwareUpdateConfigurationByNameArgs)(nil)).Elem()
}


type LookupSoftwareUpdateConfigurationByNameResultOutput struct{ *pulumi.OutputState }

func (LookupSoftwareUpdateConfigurationByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSoftwareUpdateConfigurationByNameResult)(nil)).Elem()
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ToLookupSoftwareUpdateConfigurationByNameResultOutput() LookupSoftwareUpdateConfigurationByNameResultOutput {
	return o
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ToLookupSoftwareUpdateConfigurationByNameResultOutputWithContext(ctx context.Context) LookupSoftwareUpdateConfigurationByNameResultOutput {
	return o
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) *ErrorResponseResponse { return v.Error }).(ErrorResponseResponsePtrOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) ScheduleInfo() SUCSchedulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) SUCSchedulePropertiesResponse {
		return v.ScheduleInfo
	}).(SUCSchedulePropertiesResponseOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Tasks() SoftwareUpdateConfigurationTasksResponsePtrOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) *SoftwareUpdateConfigurationTasksResponse {
		return v.Tasks
	}).(SoftwareUpdateConfigurationTasksResponsePtrOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSoftwareUpdateConfigurationByNameResultOutput) UpdateConfiguration() UpdateConfigurationResponseOutput {
	return o.ApplyT(func(v LookupSoftwareUpdateConfigurationByNameResult) UpdateConfigurationResponse {
		return v.UpdateConfiguration
	}).(UpdateConfigurationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSoftwareUpdateConfigurationByNameResultOutput{})
}
