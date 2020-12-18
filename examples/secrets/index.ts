import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import * as resources from "@pulumi/azure-nextgen/resources/latest";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
    number: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westus2",
    tags: {
        something: new pulumi.Config().requireSecret("message"),
    },
});
