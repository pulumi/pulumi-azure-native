// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App
{
    public static class GetManagedEnvironment
    {
        /// <summary>
        /// Get the properties of a Managed Environment used to host container apps.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetManagedEnvironmentResult> InvokeAsync(GetManagedEnvironmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagedEnvironmentResult>("azure-native:app:getManagedEnvironment", args ?? new GetManagedEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of a Managed Environment used to host container apps.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetManagedEnvironmentResult> Invoke(GetManagedEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedEnvironmentResult>("azure-native:app:getManagedEnvironment", args ?? new GetManagedEnvironmentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the properties of a Managed Environment used to host container apps.
        /// 
        /// Uses Azure REST API version 2024-03-01.
        /// 
        /// Other available API versions: 2022-10-01, 2022-11-01-preview, 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview, 2023-11-02-preview, 2024-02-02-preview, 2024-08-02-preview, 2024-10-02-preview, 2025-01-01, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetManagedEnvironmentResult> Invoke(GetManagedEnvironmentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedEnvironmentResult>("azure-native:app:getManagedEnvironment", args ?? new GetManagedEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedEnvironmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public string EnvironmentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagedEnvironmentArgs()
        {
        }
        public static new GetManagedEnvironmentArgs Empty => new GetManagedEnvironmentArgs();
    }

    public sealed class GetManagedEnvironmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetManagedEnvironmentInvokeArgs()
        {
        }
        public static new GetManagedEnvironmentInvokeArgs Empty => new GetManagedEnvironmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedEnvironmentResult
    {
        /// <summary>
        /// Cluster configuration which enables the log daemon to export app logs to configured destination.
        /// </summary>
        public readonly Outputs.AppLogsConfigurationResponse? AppLogsConfiguration;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Custom domain configuration for the environment
        /// </summary>
        public readonly Outputs.CustomDomainConfigurationResponse? CustomDomainConfiguration;
        /// <summary>
        /// Application Insights connection string used by Dapr to export Service to Service communication telemetry
        /// </summary>
        public readonly string? DaprAIConnectionString;
        /// <summary>
        /// Azure Monitor instrumentation key used by Dapr to export Service to Service communication telemetry
        /// </summary>
        public readonly string? DaprAIInstrumentationKey;
        /// <summary>
        /// The configuration of Dapr component.
        /// </summary>
        public readonly Outputs.DaprConfigurationResponse? DaprConfiguration;
        /// <summary>
        /// Default Domain Name for the cluster
        /// </summary>
        public readonly string DefaultDomain;
        /// <summary>
        /// Any errors that occurred during deployment or deployment validation
        /// </summary>
        public readonly string DeploymentErrors;
        /// <summary>
        /// The endpoint of the eventstream of the Environment.
        /// </summary>
        public readonly string EventStreamEndpoint;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the platform-managed resource group created for the Managed Environment to host infrastructure resources. If a subnet ID is provided, this resource group will be created in the same subscription as the subnet.
        /// </summary>
        public readonly string? InfrastructureResourceGroup;
        /// <summary>
        /// The configuration of Keda component.
        /// </summary>
        public readonly Outputs.KedaConfigurationResponse? KedaConfiguration;
        /// <summary>
        /// Kind of the Environment.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Peer authentication settings for the Managed Environment
        /// </summary>
        public readonly Outputs.ManagedEnvironmentResponsePeerAuthentication? PeerAuthentication;
        /// <summary>
        /// Peer traffic settings for the Managed Environment
        /// </summary>
        public readonly Outputs.ManagedEnvironmentResponsePeerTrafficConfiguration? PeerTrafficConfiguration;
        /// <summary>
        /// Provisioning state of the Environment.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Static IP of the Environment
        /// </summary>
        public readonly string StaticIp;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Vnet configuration for the environment
        /// </summary>
        public readonly Outputs.VnetConfigurationResponse? VnetConfiguration;
        /// <summary>
        /// Workload profiles configured for the Managed Environment.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkloadProfileResponse> WorkloadProfiles;
        /// <summary>
        /// Whether or not this Managed Environment is zone-redundant.
        /// </summary>
        public readonly bool? ZoneRedundant;

        [OutputConstructor]
        private GetManagedEnvironmentResult(
            Outputs.AppLogsConfigurationResponse? appLogsConfiguration,

            string azureApiVersion,

            Outputs.CustomDomainConfigurationResponse? customDomainConfiguration,

            string? daprAIConnectionString,

            string? daprAIInstrumentationKey,

            Outputs.DaprConfigurationResponse? daprConfiguration,

            string defaultDomain,

            string deploymentErrors,

            string eventStreamEndpoint,

            string id,

            string? infrastructureResourceGroup,

            Outputs.KedaConfigurationResponse? kedaConfiguration,

            string? kind,

            string location,

            string name,

            Outputs.ManagedEnvironmentResponsePeerAuthentication? peerAuthentication,

            Outputs.ManagedEnvironmentResponsePeerTrafficConfiguration? peerTrafficConfiguration,

            string provisioningState,

            string staticIp,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.VnetConfigurationResponse? vnetConfiguration,

            ImmutableArray<Outputs.WorkloadProfileResponse> workloadProfiles,

            bool? zoneRedundant)
        {
            AppLogsConfiguration = appLogsConfiguration;
            AzureApiVersion = azureApiVersion;
            CustomDomainConfiguration = customDomainConfiguration;
            DaprAIConnectionString = daprAIConnectionString;
            DaprAIInstrumentationKey = daprAIInstrumentationKey;
            DaprConfiguration = daprConfiguration;
            DefaultDomain = defaultDomain;
            DeploymentErrors = deploymentErrors;
            EventStreamEndpoint = eventStreamEndpoint;
            Id = id;
            InfrastructureResourceGroup = infrastructureResourceGroup;
            KedaConfiguration = kedaConfiguration;
            Kind = kind;
            Location = location;
            Name = name;
            PeerAuthentication = peerAuthentication;
            PeerTrafficConfiguration = peerTrafficConfiguration;
            ProvisioningState = provisioningState;
            StaticIp = staticIp;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VnetConfiguration = vnetConfiguration;
            WorkloadProfiles = workloadProfiles;
            ZoneRedundant = zoneRedundant;
        }
    }
}
