import * as random from "@pulumi/random";
import * as dbforpostgresql from "@pulumi/azure-nextgen/dbforpostgresql/latest";
import * as resources from "@pulumi/azure-nextgen/resources/latest";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westus2",
});

const server = new dbforpostgresql.Server("server", {
    resourceGroupName: resourceGroup.name,
    serverName: randomString.result,
    location: resourceGroup.location,
    properties: {
        administratorLogin: "cloudsa",
        administratorLoginPassword: `pa$$w0rd`,
        createMode: "Default",
        minimalTlsVersion: dbforpostgresql.MinimalTlsVersionEnum.TLS1_2,
        sslEnforcement: dbforpostgresql.SslEnforcementEnum.Enabled,
        storageProfile: {
            backupRetentionDays: 7,
            geoRedundantBackup: dbforpostgresql.GeoRedundantBackup.Disabled,
            storageMB: 128000,
        },
    },
    sku: {
        capacity: 2,
        family: "Gen5",
        name: "B_Gen5_2",
        tier: dbforpostgresql.SkuTier.Basic,
    },
});

new dbforpostgresql.Configuration("backslash_quote", {
    resourceGroupName: resourceGroup.name,
    serverName: server.name,
    configurationName: "backslash_quote",
    value: "on",
});
