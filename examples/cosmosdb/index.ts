import * as random from "@pulumi/random";
import * as resources from "@pulumi/azurerm/resources/latest";
import * as cosmosdb from "./component";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westeurope",
    tags: {
        Owner: "mikhail"
    }
});

export const cosmosdbAccount = new cosmosdb.DatabaseAccount(randomString.result, {
    resourceGroup,
    api: "Sql",
    consisencyPolicy: { defaultConsistencyLevel: "Session" },
    location: { 
        type: "single",
        region: "westeurope",
    },    
});
