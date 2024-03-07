package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsmq/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := iotoperationsmq.NewDataLakeConnectorTopicMap(ctx, "dataLakeConnectorTopicMap", &iotoperationsmq.DataLakeConnectorTopicMapArgs{
DataLakeConnectorName: pulumi.String("E9gU89-1QnIG7-IP8qOQLV-"),
DataLakeConnectorRef: pulumi.String("zirczjfua"),
ExtendedLocation: &iotoperationsmq.ExtendedLocationPropertyArgs{
Name: pulumi.String("an"),
Type: pulumi.String("CustomLocation"),
},
Location: pulumi.String("wjmgeh"),
Mapping: iotoperationsmq.DataLakeConnectorMapResponse{
AllowedLatencySecs: pulumi.Int(25407),
ClientId: pulumi.String("gyjduryceozwqyjdrxhrtwuw"),
MaxMessagesPerBatch: pulumi.Float64(1581641880),
MessagePayloadType: pulumi.String("pwhpwzqlzlhdrozpqziipvjqrnipo"),
MqttSourceTopic: pulumi.String("mtac"),
Qos: pulumi.Int(2),
Table: interface{}{
Schema: iotoperationsmq.DeltaTableSchemaArray{
&iotoperationsmq.DeltaTableSchemaArgs{
Format: pulumi.String("boolean"),
Mapping: pulumi.String("lfnyjp"),
Name: pulumi.String("hejqncdsueoerueffbaqix"),
Optional: pulumi.Bool(true),
},
},
TableName: pulumi.String("qyvzigmefvxwyjqksofyrstn"),
TablePath: pulumi.String("cyrkvxsjhbhwjegmhzyixchkdnxe"),
},
},
MqName: pulumi.String("rc-8Z--2m-MU"),
ResourceGroupName: pulumi.String("rgiotoperationsmq"),
Tags: nil,
TopicMapName: pulumi.String("17-"),
})
if err != nil {
return err
}
return nil
})
}
