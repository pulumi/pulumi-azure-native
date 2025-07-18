// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets a job step in a job's current version.
 *
 * Uses Azure REST API version 2023-08-01.
 *
 * Other available API versions: 2017-03-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getJobStep(args: GetJobStepArgs, opts?: pulumi.InvokeOptions): Promise<GetJobStepResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:sql:getJobStep", {
        "jobAgentName": args.jobAgentName,
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
        "stepName": args.stepName,
    }, opts);
}

export interface GetJobStepArgs {
    /**
     * The name of the job agent.
     */
    jobAgentName: string;
    /**
     * The name of the job.
     */
    jobName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: string;
    /**
     * The name of the server.
     */
    serverName: string;
    /**
     * The name of the job step.
     */
    stepName: string;
}

/**
 * A job step.
 */
export interface GetJobStepResult {
    /**
     * The action payload of the job step.
     */
    readonly action: outputs.sql.JobStepActionResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The resource ID of the job credential that will be used to connect to the targets.
     */
    readonly credential?: string;
    /**
     * Execution options for the job step.
     */
    readonly executionOptions?: outputs.sql.JobStepExecutionOptionsResponse;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Output destination properties of the job step.
     */
    readonly output?: outputs.sql.JobStepOutputResponse;
    /**
     * The job step's index within the job. If not specified when creating the job step, it will be created as the last step. If not specified when updating the job step, the step id is not modified.
     */
    readonly stepId?: number;
    /**
     * The resource ID of the target group that the job step will be executed on.
     */
    readonly targetGroup: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Gets a job step in a job's current version.
 *
 * Uses Azure REST API version 2023-08-01.
 *
 * Other available API versions: 2017-03-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getJobStepOutput(args: GetJobStepOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetJobStepResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:sql:getJobStep", {
        "jobAgentName": args.jobAgentName,
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
        "stepName": args.stepName,
    }, opts);
}

export interface GetJobStepOutputArgs {
    /**
     * The name of the job agent.
     */
    jobAgentName: pulumi.Input<string>;
    /**
     * The name of the job.
     */
    jobName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
    /**
     * The name of the job step.
     */
    stepName: pulumi.Input<string>;
}
