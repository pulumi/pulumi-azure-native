// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift
{
    public static class GetOpenShiftCluster
    {
        /// <summary>
        /// The operation returns properties of a OpenShift cluster.
        /// 
        /// Uses Azure REST API version 2023-11-22.
        /// 
        /// Other available API versions: 2022-09-04, 2023-04-01, 2023-07-01-preview, 2023-09-04, 2024-08-12-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redhatopenshift [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetOpenShiftClusterResult> InvokeAsync(GetOpenShiftClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOpenShiftClusterResult>("azure-native:redhatopenshift:getOpenShiftCluster", args ?? new GetOpenShiftClusterArgs(), options.WithDefaults());

        /// <summary>
        /// The operation returns properties of a OpenShift cluster.
        /// 
        /// Uses Azure REST API version 2023-11-22.
        /// 
        /// Other available API versions: 2022-09-04, 2023-04-01, 2023-07-01-preview, 2023-09-04, 2024-08-12-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redhatopenshift [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetOpenShiftClusterResult> Invoke(GetOpenShiftClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOpenShiftClusterResult>("azure-native:redhatopenshift:getOpenShiftCluster", args ?? new GetOpenShiftClusterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The operation returns properties of a OpenShift cluster.
        /// 
        /// Uses Azure REST API version 2023-11-22.
        /// 
        /// Other available API versions: 2022-09-04, 2023-04-01, 2023-07-01-preview, 2023-09-04, 2024-08-12-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redhatopenshift [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetOpenShiftClusterResult> Invoke(GetOpenShiftClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOpenShiftClusterResult>("azure-native:redhatopenshift:getOpenShiftCluster", args ?? new GetOpenShiftClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOpenShiftClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the OpenShift cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetOpenShiftClusterArgs()
        {
        }
        public static new GetOpenShiftClusterArgs Empty => new GetOpenShiftClusterArgs();
    }

    public sealed class GetOpenShiftClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the OpenShift cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetOpenShiftClusterInvokeArgs()
        {
        }
        public static new GetOpenShiftClusterInvokeArgs Empty => new GetOpenShiftClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetOpenShiftClusterResult
    {
        /// <summary>
        /// The cluster API server profile.
        /// </summary>
        public readonly Outputs.APIServerProfileResponse? ApiserverProfile;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The cluster profile.
        /// </summary>
        public readonly Outputs.ClusterProfileResponse? ClusterProfile;
        /// <summary>
        /// The console profile.
        /// </summary>
        public readonly Outputs.ConsoleProfileResponse? ConsoleProfile;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The cluster ingress profiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.IngressProfileResponse> IngressProfiles;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The cluster master profile.
        /// </summary>
        public readonly Outputs.MasterProfileResponse? MasterProfile;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The cluster network profile.
        /// </summary>
        public readonly Outputs.NetworkProfileResponse? NetworkProfile;
        /// <summary>
        /// The cluster provisioning state.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The cluster service principal profile.
        /// </summary>
        public readonly Outputs.ServicePrincipalProfileResponse? ServicePrincipalProfile;
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
        /// The cluster worker profiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkerProfileResponse> WorkerProfiles;
        /// <summary>
        /// The cluster worker profiles status.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkerProfileResponse> WorkerProfilesStatus;

        [OutputConstructor]
        private GetOpenShiftClusterResult(
            Outputs.APIServerProfileResponse? apiserverProfile,

            string azureApiVersion,

            Outputs.ClusterProfileResponse? clusterProfile,

            Outputs.ConsoleProfileResponse? consoleProfile,

            string id,

            ImmutableArray<Outputs.IngressProfileResponse> ingressProfiles,

            string location,

            Outputs.MasterProfileResponse? masterProfile,

            string name,

            Outputs.NetworkProfileResponse? networkProfile,

            string? provisioningState,

            Outputs.ServicePrincipalProfileResponse? servicePrincipalProfile,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.WorkerProfileResponse> workerProfiles,

            ImmutableArray<Outputs.WorkerProfileResponse> workerProfilesStatus)
        {
            ApiserverProfile = apiserverProfile;
            AzureApiVersion = azureApiVersion;
            ClusterProfile = clusterProfile;
            ConsoleProfile = consoleProfile;
            Id = id;
            IngressProfiles = ingressProfiles;
            Location = location;
            MasterProfile = masterProfile;
            Name = name;
            NetworkProfile = networkProfile;
            ProvisioningState = provisioningState;
            ServicePrincipalProfile = servicePrincipalProfile;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            WorkerProfiles = workerProfiles;
            WorkerProfilesStatus = workerProfilesStatus;
        }
    }
}
