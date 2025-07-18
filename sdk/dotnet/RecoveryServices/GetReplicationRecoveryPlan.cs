// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices
{
    public static class GetReplicationRecoveryPlan
    {
        /// <summary>
        /// Gets the details of the recovery plan.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetReplicationRecoveryPlanResult> InvokeAsync(GetReplicationRecoveryPlanArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReplicationRecoveryPlanResult>("azure-native:recoveryservices:getReplicationRecoveryPlan", args ?? new GetReplicationRecoveryPlanArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the recovery plan.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetReplicationRecoveryPlanResult> Invoke(GetReplicationRecoveryPlanInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplicationRecoveryPlanResult>("azure-native:recoveryservices:getReplicationRecoveryPlan", args ?? new GetReplicationRecoveryPlanInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the recovery plan.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetReplicationRecoveryPlanResult> Invoke(GetReplicationRecoveryPlanInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplicationRecoveryPlanResult>("azure-native:recoveryservices:getReplicationRecoveryPlan", args ?? new GetReplicationRecoveryPlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplicationRecoveryPlanArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the recovery plan.
        /// </summary>
        [Input("recoveryPlanName", required: true)]
        public string RecoveryPlanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the recovery services vault is present.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the recovery services vault.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetReplicationRecoveryPlanArgs()
        {
        }
        public static new GetReplicationRecoveryPlanArgs Empty => new GetReplicationRecoveryPlanArgs();
    }

    public sealed class GetReplicationRecoveryPlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the recovery plan.
        /// </summary>
        [Input("recoveryPlanName", required: true)]
        public Input<string> RecoveryPlanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the recovery services vault is present.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the recovery services vault.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetReplicationRecoveryPlanInvokeArgs()
        {
        }
        public static new GetReplicationRecoveryPlanInvokeArgs Empty => new GetReplicationRecoveryPlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetReplicationRecoveryPlanResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The custom details.
        /// </summary>
        public readonly Outputs.RecoveryPlanPropertiesResponse Properties;
        /// <summary>
        /// Resource Type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetReplicationRecoveryPlanResult(
            string azureApiVersion,

            string id,

            string? location,

            string name,

            Outputs.RecoveryPlanPropertiesResponse properties,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
