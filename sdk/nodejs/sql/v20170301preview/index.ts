// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { BackupLongTermRetentionPolicyArgs } from "./backupLongTermRetentionPolicy";
export type BackupLongTermRetentionPolicy = import("./backupLongTermRetentionPolicy").BackupLongTermRetentionPolicy;
export const BackupLongTermRetentionPolicy: typeof import("./backupLongTermRetentionPolicy").BackupLongTermRetentionPolicy = null as any;
utilities.lazyLoad(exports, ["BackupLongTermRetentionPolicy"], () => require("./backupLongTermRetentionPolicy"));

export { GetBackupLongTermRetentionPolicyArgs, GetBackupLongTermRetentionPolicyResult, GetBackupLongTermRetentionPolicyOutputArgs } from "./getBackupLongTermRetentionPolicy";
export const getBackupLongTermRetentionPolicy: typeof import("./getBackupLongTermRetentionPolicy").getBackupLongTermRetentionPolicy = null as any;
export const getBackupLongTermRetentionPolicyOutput: typeof import("./getBackupLongTermRetentionPolicy").getBackupLongTermRetentionPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getBackupLongTermRetentionPolicy","getBackupLongTermRetentionPolicyOutput"], () => require("./getBackupLongTermRetentionPolicy"));

export { GetServerSecurityAlertPolicyArgs, GetServerSecurityAlertPolicyResult, GetServerSecurityAlertPolicyOutputArgs } from "./getServerSecurityAlertPolicy";
export const getServerSecurityAlertPolicy: typeof import("./getServerSecurityAlertPolicy").getServerSecurityAlertPolicy = null as any;
export const getServerSecurityAlertPolicyOutput: typeof import("./getServerSecurityAlertPolicy").getServerSecurityAlertPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getServerSecurityAlertPolicy","getServerSecurityAlertPolicyOutput"], () => require("./getServerSecurityAlertPolicy"));

export { ServerSecurityAlertPolicyArgs } from "./serverSecurityAlertPolicy";
export type ServerSecurityAlertPolicy = import("./serverSecurityAlertPolicy").ServerSecurityAlertPolicy;
export const ServerSecurityAlertPolicy: typeof import("./serverSecurityAlertPolicy").ServerSecurityAlertPolicy = null as any;
utilities.lazyLoad(exports, ["ServerSecurityAlertPolicy"], () => require("./serverSecurityAlertPolicy"));


// Export enums:
export * from "../../types/enums/sql/v20170301preview";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:sql/v20170301preview:BackupLongTermRetentionPolicy":
                return new BackupLongTermRetentionPolicy(name, <any>undefined, { urn })
            case "azure-native:sql/v20170301preview:ServerSecurityAlertPolicy":
                return new ServerSecurityAlertPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "sql/v20170301preview", _module)
