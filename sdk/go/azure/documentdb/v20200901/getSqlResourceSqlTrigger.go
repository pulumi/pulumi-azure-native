


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSqlResourceSqlTrigger(ctx *pulumi.Context, args *LookupSqlResourceSqlTriggerArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlTriggerResult, error) {
	var rv LookupSqlResourceSqlTriggerResult
	err := ctx.Invoke("azure-native:documentdb/v20200901:getSqlResourceSqlTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlTriggerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TriggerName       string `pulumi:"triggerName"`
}


type LookupSqlResourceSqlTriggerResult struct {
	Id       string                                   `pulumi:"id"`
	Location *string                                  `pulumi:"location"`
	Name     string                                   `pulumi:"name"`
	Resource *SqlTriggerGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                        `pulumi:"tags"`
	Type     string                                   `pulumi:"type"`
}

func LookupSqlResourceSqlTriggerOutput(ctx *pulumi.Context, args LookupSqlResourceSqlTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlTriggerResult, error) {
			args := v.(LookupSqlResourceSqlTriggerArgs)
			r, err := LookupSqlResourceSqlTrigger(ctx, &args, opts...)
			var s LookupSqlResourceSqlTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlTriggerResultOutput)
}

type LookupSqlResourceSqlTriggerOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ContainerName     pulumi.StringInput `pulumi:"containerName"`
	DatabaseName      pulumi.StringInput `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TriggerName       pulumi.StringInput `pulumi:"triggerName"`
}

func (LookupSqlResourceSqlTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlTriggerArgs)(nil)).Elem()
}


type LookupSqlResourceSqlTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlTriggerResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlTriggerResultOutput) ToLookupSqlResourceSqlTriggerResultOutput() LookupSqlResourceSqlTriggerResultOutput {
	return o
}

func (o LookupSqlResourceSqlTriggerResultOutput) ToLookupSqlResourceSqlTriggerResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlTriggerResultOutput {
	return o
}

func (o LookupSqlResourceSqlTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlTriggerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlTriggerResultOutput) Resource() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) *SqlTriggerGetPropertiesResponseResource { return v.Resource }).(SqlTriggerGetPropertiesResponseResourcePtrOutput)
}

func (o LookupSqlResourceSqlTriggerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlResourceSqlTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlTriggerResultOutput{})
}
