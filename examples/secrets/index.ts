import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg", {
    tags: {
        something: new pulumi.Config().requireSecret("message"),
    },
});
