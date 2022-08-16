


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQuery(ctx *pulumi.Context, args *LookupQueryArgs, opts ...pulumi.InvokeOption) (*LookupQueryResult, error) {
	var rv LookupQueryResult
	err := ctx.Invoke("azure-native:operationalinsights:getQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueryArgs struct {
	Id                string `pulumi:"id"`
	QueryPackName     string `pulumi:"queryPackName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueryResult struct {
	Author       string                                               `pulumi:"author"`
	Body         string                                               `pulumi:"body"`
	Description  *string                                              `pulumi:"description"`
	DisplayName  string                                               `pulumi:"displayName"`
	Id           string                                               `pulumi:"id"`
	Name         string                                               `pulumi:"name"`
	Properties   interface{}                                          `pulumi:"properties"`
	Related      *LogAnalyticsQueryPackQueryPropertiesResponseRelated `pulumi:"related"`
	SystemData   SystemDataResponse                                   `pulumi:"systemData"`
	Tags         map[string][]string                                  `pulumi:"tags"`
	TimeCreated  string                                               `pulumi:"timeCreated"`
	TimeModified string                                               `pulumi:"timeModified"`
	Type         string                                               `pulumi:"type"`
}

func LookupQueryOutput(ctx *pulumi.Context, args LookupQueryOutputArgs, opts ...pulumi.InvokeOption) LookupQueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueryResult, error) {
			args := v.(LookupQueryArgs)
			r, err := LookupQuery(ctx, &args, opts...)
			var s LookupQueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueryResultOutput)
}

type LookupQueryOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	QueryPackName     pulumi.StringInput `pulumi:"queryPackName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueryArgs)(nil)).Elem()
}


type LookupQueryResultOutput struct{ *pulumi.OutputState }

func (LookupQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueryResult)(nil)).Elem()
}

func (o LookupQueryResultOutput) ToLookupQueryResultOutput() LookupQueryResultOutput {
	return o
}

func (o LookupQueryResultOutput) ToLookupQueryResultOutputWithContext(ctx context.Context) LookupQueryResultOutput {
	return o
}

func (o LookupQueryResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.Author }).(pulumi.StringOutput)
}

func (o LookupQueryResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.Body }).(pulumi.StringOutput)
}

func (o LookupQueryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupQueryResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupQueryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueryResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupQueryResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupQueryResultOutput) Related() LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o.ApplyT(func(v LookupQueryResult) *LogAnalyticsQueryPackQueryPropertiesResponseRelated { return v.Related }).(LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput)
}

func (o LookupQueryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupQueryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupQueryResultOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v LookupQueryResult) map[string][]string { return v.Tags }).(pulumi.StringArrayMapOutput)
}

func (o LookupQueryResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupQueryResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupQueryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueryResultOutput{})
}
