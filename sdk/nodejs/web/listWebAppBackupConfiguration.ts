// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Description for Gets the backup configuration of an app.
 *
 * Uses Azure REST API version 2022-09-01.
 *
 * Other available API versions: 2016-08-01, 2020-10-01, 2023-01-01, 2023-12-01, 2024-04-01.
 */
export function listWebAppBackupConfiguration(args: ListWebAppBackupConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<ListWebAppBackupConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:web:listWebAppBackupConfiguration", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListWebAppBackupConfigurationArgs {
    /**
     * Name of the app.
     */
    name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: string;
}

/**
 * Description of a backup which will be performed.
 */
export interface ListWebAppBackupConfigurationResult {
    /**
     * Name of the backup.
     */
    readonly backupName?: string;
    /**
     * Schedule for the backup if it is executed periodically.
     */
    readonly backupSchedule?: outputs.web.BackupScheduleResponse;
    /**
     * Databases included in the backup.
     */
    readonly databases?: outputs.web.DatabaseBackupSettingResponse[];
    /**
     * True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
     */
    readonly enabled?: boolean;
    /**
     * Resource Id.
     */
    readonly id: string;
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * SAS URL to the container.
     */
    readonly storageAccountUrl: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Description for Gets the backup configuration of an app.
 *
 * Uses Azure REST API version 2022-09-01.
 *
 * Other available API versions: 2016-08-01, 2020-10-01, 2023-01-01, 2023-12-01, 2024-04-01.
 */
export function listWebAppBackupConfigurationOutput(args: ListWebAppBackupConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListWebAppBackupConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:web:listWebAppBackupConfiguration", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListWebAppBackupConfigurationOutputArgs {
    /**
     * Name of the app.
     */
    name: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: pulumi.Input<string>;
}
