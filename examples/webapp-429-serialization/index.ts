import * as resources from "@pulumi/azure-native/resources";
import * as web from "@pulumi/azure-native/web";

// This test verifies that Web Apps sharing the same App Service Plan are serialized (created, updated, and deleted
// one after the other rather than in parallel)to prevent 429 "Cannot acquire exclusive lock" errors when Azure locks
// the App Service Plan during concurrent operations.
//
// Based on GitHub issue #3720: https://github.com/pulumi/pulumi-azure-native/issues/3720
//
// Test scenario:
// Create >4 WebApps in parallel, all referencing the same AppServicePlan.
// On Delete, without serialization, this amount of WebApps will trigger 429 "Cannot acquire exclusive lock" errors.
// The test passing without a 429 error indicates that CRUD operations on WebApps are serialized, not concurrent.

const resourceGroup = new resources.ResourceGroup("repro-429-rg", {
    location: "westus2",
});

const appServicePlan = new web.AppServicePlan("repro-asp", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    kind: "Linux",
    reserved: true,
    sku: {
        name: "B1",
        tier: "Basic",
    },
});

const webAppNames = [
    "webapp-1", "webapp-2", "webapp-3", "webapp-4", "webapp-5",
    "webapp-6", "webapp-7", "webapp-8",
];

const webApps: web.WebApp[] = [];

for (const name of webAppNames) {
    const webApp = new web.WebApp(name, {
        resourceGroupName: resourceGroup.name,
        location: resourceGroup.location,
        serverFarmId: appServicePlan.id,
        siteConfig: {
            linuxFxVersion: "DOCKER|nginx",
        },
        httpsOnly: true,
    });
    webApps.push(webApp);
}

export const resourceGroupName = resourceGroup.name;
export const appServicePlanName = appServicePlan.name;
export const webAppCount = webApps.length;
export const webAppNamesExported = webApps.map(w => w.name);

