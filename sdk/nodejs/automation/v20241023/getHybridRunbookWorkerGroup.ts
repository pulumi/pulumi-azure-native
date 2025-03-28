// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieve a hybrid runbook worker group.
 */
export function getHybridRunbookWorkerGroup(args: GetHybridRunbookWorkerGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetHybridRunbookWorkerGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:automation/v20241023:getHybridRunbookWorkerGroup", {
        "automationAccountName": args.automationAccountName,
        "hybridRunbookWorkerGroupName": args.hybridRunbookWorkerGroupName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHybridRunbookWorkerGroupArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: string;
    /**
     * The hybrid runbook worker group name
     */
    hybridRunbookWorkerGroupName: string;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: string;
}

/**
 * Definition of hybrid runbook worker group.
 */
export interface GetHybridRunbookWorkerGroupResult {
    /**
     * Sets the credential of a worker group.
     */
    readonly credential?: outputs.automation.v20241023.RunAsCredentialAssociationPropertyResponse;
    /**
     * Type of the HybridWorkerGroup.
     */
    readonly groupType?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.automation.v20241023.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Retrieve a hybrid runbook worker group.
 */
export function getHybridRunbookWorkerGroupOutput(args: GetHybridRunbookWorkerGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetHybridRunbookWorkerGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:automation/v20241023:getHybridRunbookWorkerGroup", {
        "automationAccountName": args.automationAccountName,
        "hybridRunbookWorkerGroupName": args.hybridRunbookWorkerGroupName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHybridRunbookWorkerGroupOutputArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: pulumi.Input<string>;
    /**
     * The hybrid runbook worker group name
     */
    hybridRunbookWorkerGroupName: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
