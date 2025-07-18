// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute
{
    public static class GetGalleryInVMAccessControlProfileVersion
    {
        /// <summary>
        /// Retrieves information about a gallery inVMAccessControlProfile version.
        /// 
        /// Uses Azure REST API version 2024-03-03.
        /// </summary>
        public static Task<GetGalleryInVMAccessControlProfileVersionResult> InvokeAsync(GetGalleryInVMAccessControlProfileVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGalleryInVMAccessControlProfileVersionResult>("azure-native:compute:getGalleryInVMAccessControlProfileVersion", args ?? new GetGalleryInVMAccessControlProfileVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about a gallery inVMAccessControlProfile version.
        /// 
        /// Uses Azure REST API version 2024-03-03.
        /// </summary>
        public static Output<GetGalleryInVMAccessControlProfileVersionResult> Invoke(GetGalleryInVMAccessControlProfileVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGalleryInVMAccessControlProfileVersionResult>("azure-native:compute:getGalleryInVMAccessControlProfileVersion", args ?? new GetGalleryInVMAccessControlProfileVersionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about a gallery inVMAccessControlProfile version.
        /// 
        /// Uses Azure REST API version 2024-03-03.
        /// </summary>
        public static Output<GetGalleryInVMAccessControlProfileVersionResult> Invoke(GetGalleryInVMAccessControlProfileVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGalleryInVMAccessControlProfileVersionResult>("azure-native:compute:getGalleryInVMAccessControlProfileVersion", args ?? new GetGalleryInVMAccessControlProfileVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGalleryInVMAccessControlProfileVersionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image Gallery.
        /// </summary>
        [Input("galleryName", required: true)]
        public string GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery inVMAccessControlProfile to be retrieved.
        /// </summary>
        [Input("inVMAccessControlProfileName", required: true)]
        public string InVMAccessControlProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery inVMAccessControlProfile version to be retrieved.
        /// </summary>
        [Input("inVMAccessControlProfileVersionName", required: true)]
        public string InVMAccessControlProfileVersionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetGalleryInVMAccessControlProfileVersionArgs()
        {
        }
        public static new GetGalleryInVMAccessControlProfileVersionArgs Empty => new GetGalleryInVMAccessControlProfileVersionArgs();
    }

    public sealed class GetGalleryInVMAccessControlProfileVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image Gallery.
        /// </summary>
        [Input("galleryName", required: true)]
        public Input<string> GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery inVMAccessControlProfile to be retrieved.
        /// </summary>
        [Input("inVMAccessControlProfileName", required: true)]
        public Input<string> InVMAccessControlProfileName { get; set; } = null!;

        /// <summary>
        /// The name of the gallery inVMAccessControlProfile version to be retrieved.
        /// </summary>
        [Input("inVMAccessControlProfileVersionName", required: true)]
        public Input<string> InVMAccessControlProfileVersionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetGalleryInVMAccessControlProfileVersionInvokeArgs()
        {
        }
        public static new GetGalleryInVMAccessControlProfileVersionInvokeArgs Empty => new GetGalleryInVMAccessControlProfileVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetGalleryInVMAccessControlProfileVersionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// This property allows you to specify if the requests will be allowed to access the host endpoints. Possible values are: 'Allow', 'Deny'.
        /// </summary>
        public readonly string DefaultAccess;
        /// <summary>
        /// If set to true, Virtual Machines deployed from the latest version of the Resource Profile won't use this Profile version.
        /// </summary>
        public readonly bool? ExcludeFromLatest;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// This property allows you to specify whether the access control rules are in Audit mode, in Enforce mode or Disabled. Possible values are: 'Audit', 'Enforce' or 'Disabled'.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The timestamp for when the Resource Profile Version is published.
        /// </summary>
        public readonly string PublishedDate;
        /// <summary>
        /// This is the replication status of the gallery image version.
        /// </summary>
        public readonly Outputs.ReplicationStatusResponse ReplicationStatus;
        /// <summary>
        /// This is the Access Control Rules specification for an inVMAccessControlProfile version.
        /// </summary>
        public readonly Outputs.AccessControlRulesResponse? Rules;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The target regions where the Resource Profile version is going to be replicated to. This property is updatable.
        /// </summary>
        public readonly ImmutableArray<Outputs.TargetRegionResponse> TargetLocations;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGalleryInVMAccessControlProfileVersionResult(
            string azureApiVersion,

            string defaultAccess,

            bool? excludeFromLatest,

            string id,

            string location,

            string mode,

            string name,

            string provisioningState,

            string publishedDate,

            Outputs.ReplicationStatusResponse replicationStatus,

            Outputs.AccessControlRulesResponse? rules,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.TargetRegionResponse> targetLocations,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DefaultAccess = defaultAccess;
            ExcludeFromLatest = excludeFromLatest;
            Id = id;
            Location = location;
            Mode = mode;
            Name = name;
            ProvisioningState = provisioningState;
            PublishedDate = publishedDate;
            ReplicationStatus = replicationStatus;
            Rules = rules;
            SystemData = systemData;
            Tags = tags;
            TargetLocations = targetLocations;
            Type = type;
        }
    }
}
