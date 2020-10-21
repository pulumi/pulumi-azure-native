import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const hostingPlanNameParam = config.require("hostingPlanNameParam");
const locationParam = config.require("locationParam");
const numberOfWorkersParam = config.require("numberOfWorkersParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const skuCodeParam = config.require("skuCodeParam");
const skuParam = config.require("skuParam");
const workerSizeIdParam = config.require("workerSizeIdParam");
const workerSizeParam = config.require("workerSizeParam");
const serverfarmResource = new azure_nextgen.web.v20190801.AppServicePlan("serverfarmResource", {
    location: locationParam,
    name: hostingPlanNameParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        Name: skuCodeParam,
        Tier: skuParam,
    },
});
