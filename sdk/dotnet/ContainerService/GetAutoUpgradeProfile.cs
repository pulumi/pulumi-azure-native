// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService
{
    public static class GetAutoUpgradeProfile
    {
        /// <summary>
        /// Get a AutoUpgradeProfile
        /// 
        /// Uses Azure REST API version 2024-05-02-preview.
        /// 
        /// Other available API versions: 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAutoUpgradeProfileResult> InvokeAsync(GetAutoUpgradeProfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutoUpgradeProfileResult>("azure-native:containerservice:getAutoUpgradeProfile", args ?? new GetAutoUpgradeProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Get a AutoUpgradeProfile
        /// 
        /// Uses Azure REST API version 2024-05-02-preview.
        /// 
        /// Other available API versions: 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAutoUpgradeProfileResult> Invoke(GetAutoUpgradeProfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutoUpgradeProfileResult>("azure-native:containerservice:getAutoUpgradeProfile", args ?? new GetAutoUpgradeProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a AutoUpgradeProfile
        /// 
        /// Uses Azure REST API version 2024-05-02-preview.
        /// 
        /// Other available API versions: 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAutoUpgradeProfileResult> Invoke(GetAutoUpgradeProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutoUpgradeProfileResult>("azure-native:containerservice:getAutoUpgradeProfile", args ?? new GetAutoUpgradeProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutoUpgradeProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the AutoUpgradeProfile resource.
        /// </summary>
        [Input("autoUpgradeProfileName", required: true)]
        public string AutoUpgradeProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the Fleet resource.
        /// </summary>
        [Input("fleetName", required: true)]
        public string FleetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAutoUpgradeProfileArgs()
        {
        }
        public static new GetAutoUpgradeProfileArgs Empty => new GetAutoUpgradeProfileArgs();
    }

    public sealed class GetAutoUpgradeProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the AutoUpgradeProfile resource.
        /// </summary>
        [Input("autoUpgradeProfileName", required: true)]
        public Input<string> AutoUpgradeProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the Fleet resource.
        /// </summary>
        [Input("fleetName", required: true)]
        public Input<string> FleetName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAutoUpgradeProfileInvokeArgs()
        {
        }
        public static new GetAutoUpgradeProfileInvokeArgs Empty => new GetAutoUpgradeProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutoUpgradeProfileResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Configures how auto-upgrade will be run.
        /// </summary>
        public readonly string Channel;
        /// <summary>
        /// If set to False: the auto upgrade has effect - target managed clusters will be upgraded on schedule.
        /// If set to True: the auto upgrade has no effect - no upgrade will be run on the target managed clusters.
        /// This is a boolean and not an enum because enabled/disabled are all available states of the auto upgrade profile.
        /// By default, this is set to False.
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
        /// </summary>
        public readonly string ETag;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The node image upgrade to be applied to the target clusters in auto upgrade.
        /// </summary>
        public readonly Outputs.AutoUpgradeNodeImageSelectionResponse? NodeImageSelection;
        /// <summary>
        /// The provisioning state of the AutoUpgradeProfile resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The resource id of the UpdateStrategy resource to reference. If not specified, the auto upgrade will run on all clusters which are members of the fleet.
        /// </summary>
        public readonly string? UpdateStrategyId;

        [OutputConstructor]
        private GetAutoUpgradeProfileResult(
            string azureApiVersion,

            string channel,

            bool? disabled,

            string eTag,

            string id,

            string name,

            Outputs.AutoUpgradeNodeImageSelectionResponse? nodeImageSelection,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type,

            string? updateStrategyId)
        {
            AzureApiVersion = azureApiVersion;
            Channel = channel;
            Disabled = disabled;
            ETag = eTag;
            Id = id;
            Name = name;
            NodeImageSelection = nodeImageSelection;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
            UpdateStrategyId = updateStrategyId;
        }
    }
}
