import * as resources from "@pulumi/azure-nextgen/resources";
import * as cosmosdb from "./component";

const resourceGroup = new resources.ResourceGroup("rg", {
    location: "westeurope",
});

export const cosmosdbAccount = new cosmosdb.DatabaseAccount("da", {
    resourceGroup,
    api: "Sql",
    consisencyPolicy: { defaultConsistencyLevel: "Session" },
    location: { 
        type: "single",
        region: "westeurope",
    },    
});
