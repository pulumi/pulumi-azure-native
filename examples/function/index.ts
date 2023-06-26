import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure-native";


const resourceGroup = new azure.resources.ResourceGroup("resourceGroup", {
    resourceGroupName: "my-functions",
});

new azure.web.crede
