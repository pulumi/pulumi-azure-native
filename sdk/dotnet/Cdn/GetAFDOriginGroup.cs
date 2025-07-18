// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn
{
    public static class GetAFDOriginGroup
    {
        /// <summary>
        /// Gets an existing origin group within a profile.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAFDOriginGroupResult> InvokeAsync(GetAFDOriginGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAFDOriginGroupResult>("azure-native:cdn:getAFDOriginGroup", args ?? new GetAFDOriginGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing origin group within a profile.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAFDOriginGroupResult> Invoke(GetAFDOriginGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAFDOriginGroupResult>("azure-native:cdn:getAFDOriginGroup", args ?? new GetAFDOriginGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing origin group within a profile.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAFDOriginGroupResult> Invoke(GetAFDOriginGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAFDOriginGroupResult>("azure-native:cdn:getAFDOriginGroup", args ?? new GetAFDOriginGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAFDOriginGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the origin group which is unique within the endpoint.
        /// </summary>
        [Input("originGroupName", required: true)]
        public string OriginGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAFDOriginGroupArgs()
        {
        }
        public static new GetAFDOriginGroupArgs Empty => new GetAFDOriginGroupArgs();
    }

    public sealed class GetAFDOriginGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the origin group which is unique within the endpoint.
        /// </summary>
        [Input("originGroupName", required: true)]
        public Input<string> OriginGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAFDOriginGroupInvokeArgs()
        {
        }
        public static new GetAFDOriginGroupInvokeArgs Empty => new GetAFDOriginGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetAFDOriginGroupResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        public readonly string DeploymentStatus;
        /// <summary>
        /// Health probe settings to the origin that is used to determine the health of the origin.
        /// </summary>
        public readonly Outputs.HealthProbeParametersResponse? HealthProbeSettings;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Load balancing settings for a backend pool
        /// </summary>
        public readonly Outputs.LoadBalancingSettingsParametersResponse? LoadBalancingSettings;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the profile which holds the origin group.
        /// </summary>
        public readonly string ProfileName;
        /// <summary>
        /// Provisioning status
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string? SessionAffinityState;
        /// <summary>
        /// Read only system data
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
        /// </summary>
        public readonly int? TrafficRestorationTimeToHealedOrNewEndpointsInMinutes;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAFDOriginGroupResult(
            string azureApiVersion,

            string deploymentStatus,

            Outputs.HealthProbeParametersResponse? healthProbeSettings,

            string id,

            Outputs.LoadBalancingSettingsParametersResponse? loadBalancingSettings,

            string name,

            string profileName,

            string provisioningState,

            string? sessionAffinityState,

            Outputs.SystemDataResponse systemData,

            int? trafficRestorationTimeToHealedOrNewEndpointsInMinutes,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DeploymentStatus = deploymentStatus;
            HealthProbeSettings = healthProbeSettings;
            Id = id;
            LoadBalancingSettings = loadBalancingSettings;
            Name = name;
            ProfileName = profileName;
            ProvisioningState = provisioningState;
            SessionAffinityState = sessionAffinityState;
            SystemData = systemData;
            TrafficRestorationTimeToHealedOrNewEndpointsInMinutes = trafficRestorationTimeToHealedOrNewEndpointsInMinutes;
            Type = type;
        }
    }
}
