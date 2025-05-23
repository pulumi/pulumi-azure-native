// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MySQLDiscovery
{
    public static class GetMySQLSite
    {
        /// <summary>
        /// Gets the MySQLSites resource.
        /// 
        /// Uses Azure REST API version 2024-09-30-preview.
        /// </summary>
        public static Task<GetMySQLSiteResult> InvokeAsync(GetMySQLSiteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMySQLSiteResult>("azure-native:mysqldiscovery:getMySQLSite", args ?? new GetMySQLSiteArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the MySQLSites resource.
        /// 
        /// Uses Azure REST API version 2024-09-30-preview.
        /// </summary>
        public static Output<GetMySQLSiteResult> Invoke(GetMySQLSiteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMySQLSiteResult>("azure-native:mysqldiscovery:getMySQLSite", args ?? new GetMySQLSiteInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the MySQLSites resource.
        /// 
        /// Uses Azure REST API version 2024-09-30-preview.
        /// </summary>
        public static Output<GetMySQLSiteResult> Invoke(GetMySQLSiteInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMySQLSiteResult>("azure-native:mysqldiscovery:getMySQLSite", args ?? new GetMySQLSiteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMySQLSiteArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of Site
        /// </summary>
        [Input("siteName", required: true)]
        public string SiteName { get; set; } = null!;

        public GetMySQLSiteArgs()
        {
        }
        public static new GetMySQLSiteArgs Empty => new GetMySQLSiteArgs();
    }

    public sealed class GetMySQLSiteInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of Site
        /// </summary>
        [Input("siteName", required: true)]
        public Input<string> SiteName { get; set; } = null!;

        public GetMySQLSiteInvokeArgs()
        {
        }
        public static new GetMySQLSiteInvokeArgs Empty => new GetMySQLSiteInvokeArgs();
    }


    [OutputType]
    public sealed class GetMySQLSiteResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The extended location.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse ExtendedLocation;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The mapped master Site Id.
        /// </summary>
        public readonly string MasterSiteId;
        /// <summary>
        /// The mapped migrate project Id.
        /// </summary>
        public readonly string MigrateProjectId;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets the provisioning state.
        /// </summary>
        public readonly string? ProvisioningState;
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

        [OutputConstructor]
        private GetMySQLSiteResult(
            string azureApiVersion,

            Outputs.ExtendedLocationResponse extendedLocation,

            string id,

            string location,

            string masterSiteId,

            string migrateProjectId,

            string name,

            string? provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ExtendedLocation = extendedLocation;
            Id = id;
            Location = location;
            MasterSiteId = masterSiteId;
            MigrateProjectId = migrateProjectId;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
