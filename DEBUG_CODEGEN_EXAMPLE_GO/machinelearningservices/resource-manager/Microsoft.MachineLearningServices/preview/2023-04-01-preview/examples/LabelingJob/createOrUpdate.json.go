package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewLabelingJob(ctx, "labelingJob", &machinelearningservices.LabelingJobArgs{
Id: pulumi.String("testLabelingJob"),
LabelingJobProperties: &machinelearningservices.LabelingJobTypeArgs{
Description: pulumi.String("string"),
JobInstructions: &machinelearningservices.LabelingJobInstructionsArgs{
Uri: pulumi.String("link/to/instructions"),
},
JobType: pulumi.String("Labeling"),
LabelCategories: machinelearningservices.LabelCategoryMap{
"myCategory1": &machinelearningservices.LabelCategoryArgs{
Classes: interface{}{
MyLabelClass1: &machinelearningservices.LabelClassArgs{
DisplayName: pulumi.String("myLabelClass1"),
Subclasses: nil,
},
MyLabelClass2: interface{}{
DisplayName: pulumi.String("myLabelClass2"),
Subclasses: nil,
},
},
DisplayName: pulumi.String("myCategory1Title"),
MultiSelect: pulumi.String("Disabled"),
},
"myCategory2": &machinelearningservices.LabelCategoryArgs{
Classes: interface{}{
MyLabelClass1: &machinelearningservices.LabelClassArgs{
DisplayName: pulumi.String("myLabelClass1"),
Subclasses: nil,
},
MyLabelClass2: interface{}{
DisplayName: pulumi.String("myLabelClass2"),
Subclasses: nil,
},
},
DisplayName: pulumi.String("myCategory2Title"),
MultiSelect: pulumi.String("Disabled"),
},
},
LabelingJobMediaProperties: machinelearningservices.LabelingJobImageProperties{
MediaType: "Image",
},
MlAssistConfiguration: machinelearningservices.MLAssistConfigurationEnabled{
InferencingComputeBinding: "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/computes/myscoringcompute",
MlAssist: "Enabled",
TrainingComputeBinding: "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1234/providers/Microsoft.MachineLearningServices/workspaces/testworkspace/computes/mytrainingompute",
},
Properties: pulumi.StringMap{
"additionalProp1": pulumi.String("string"),
"additionalProp2": pulumi.String("string"),
"additionalProp3": pulumi.String("string"),
},
Tags: pulumi.StringMap{
"additionalProp1": pulumi.String("string"),
"additionalProp2": pulumi.String("string"),
"additionalProp3": pulumi.String("string"),
},
},
ResourceGroupName: pulumi.String("workspace-1234"),
WorkspaceName: pulumi.String("testworkspace"),
})
if err != nil {
return err
}
return nil
})
}
