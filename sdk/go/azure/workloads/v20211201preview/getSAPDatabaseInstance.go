


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSAPDatabaseInstance(ctx *pulumi.Context, args *LookupSAPDatabaseInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSAPDatabaseInstanceResult, error) {
	var rv LookupSAPDatabaseInstanceResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getSAPDatabaseInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSAPDatabaseInstanceArgs struct {
	DatabaseInstanceName   string `pulumi:"databaseInstanceName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SapVirtualInstanceName string `pulumi:"sapVirtualInstanceName"`
}


type LookupSAPDatabaseInstanceResult struct {
	DatabaseSid         string                          `pulumi:"databaseSid"`
	DatabaseType        string                          `pulumi:"databaseType"`
	Errors              SAPVirtualInstanceErrorResponse `pulumi:"errors"`
	Id                  string                          `pulumi:"id"`
	IpAddress           string                          `pulumi:"ipAddress"`
	LoadBalancerDetails LoadBalancerDetailsResponse     `pulumi:"loadBalancerDetails"`
	Location            string                          `pulumi:"location"`
	Name                string                          `pulumi:"name"`
	ProvisioningState   string                          `pulumi:"provisioningState"`
	Status              string                          `pulumi:"status"`
	Subnet              string                          `pulumi:"subnet"`
	SystemData          SystemDataResponse              `pulumi:"systemData"`
	Tags                map[string]string               `pulumi:"tags"`
	Type                string                          `pulumi:"type"`
	VmDetails           []DatabaseVmDetailsResponse     `pulumi:"vmDetails"`
}

func LookupSAPDatabaseInstanceOutput(ctx *pulumi.Context, args LookupSAPDatabaseInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSAPDatabaseInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSAPDatabaseInstanceResult, error) {
			args := v.(LookupSAPDatabaseInstanceArgs)
			r, err := LookupSAPDatabaseInstance(ctx, &args, opts...)
			var s LookupSAPDatabaseInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSAPDatabaseInstanceResultOutput)
}

type LookupSAPDatabaseInstanceOutputArgs struct {
	DatabaseInstanceName   pulumi.StringInput `pulumi:"databaseInstanceName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SapVirtualInstanceName pulumi.StringInput `pulumi:"sapVirtualInstanceName"`
}

func (LookupSAPDatabaseInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPDatabaseInstanceArgs)(nil)).Elem()
}


type LookupSAPDatabaseInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSAPDatabaseInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPDatabaseInstanceResult)(nil)).Elem()
}

func (o LookupSAPDatabaseInstanceResultOutput) ToLookupSAPDatabaseInstanceResultOutput() LookupSAPDatabaseInstanceResultOutput {
	return o
}

func (o LookupSAPDatabaseInstanceResultOutput) ToLookupSAPDatabaseInstanceResultOutputWithContext(ctx context.Context) LookupSAPDatabaseInstanceResultOutput {
	return o
}

func (o LookupSAPDatabaseInstanceResultOutput) DatabaseSid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.DatabaseSid }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.DatabaseType }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) SAPVirtualInstanceErrorResponse { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) LoadBalancerDetails() LoadBalancerDetailsResponseOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) LoadBalancerDetailsResponse { return v.LoadBalancerDetails }).(LoadBalancerDetailsResponseOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.Subnet }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSAPDatabaseInstanceResultOutput) VmDetails() DatabaseVmDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupSAPDatabaseInstanceResult) []DatabaseVmDetailsResponse { return v.VmDetails }).(DatabaseVmDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSAPDatabaseInstanceResultOutput{})
}
