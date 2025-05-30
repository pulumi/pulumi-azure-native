// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices
{
    public static class GetLabPlan
    {
        /// <summary>
        /// Retrieves the properties of a Lab Plan.
        /// 
        /// Uses Azure REST API version 2023-06-07.
        /// 
        /// Other available API versions: 2021-10-01-preview, 2021-11-15-preview, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native labservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetLabPlanResult> InvokeAsync(GetLabPlanArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLabPlanResult>("azure-native:labservices:getLabPlan", args ?? new GetLabPlanArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of a Lab Plan.
        /// 
        /// Uses Azure REST API version 2023-06-07.
        /// 
        /// Other available API versions: 2021-10-01-preview, 2021-11-15-preview, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native labservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLabPlanResult> Invoke(GetLabPlanInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLabPlanResult>("azure-native:labservices:getLabPlan", args ?? new GetLabPlanInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the properties of a Lab Plan.
        /// 
        /// Uses Azure REST API version 2023-06-07.
        /// 
        /// Other available API versions: 2021-10-01-preview, 2021-11-15-preview, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native labservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetLabPlanResult> Invoke(GetLabPlanInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLabPlanResult>("azure-native:labservices:getLabPlan", args ?? new GetLabPlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLabPlanArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the lab plan that uniquely identifies it within containing resource group. Used in resource URIs and in UI.
        /// </summary>
        [Input("labPlanName", required: true)]
        public string LabPlanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLabPlanArgs()
        {
        }
        public static new GetLabPlanArgs Empty => new GetLabPlanArgs();
    }

    public sealed class GetLabPlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the lab plan that uniquely identifies it within containing resource group. Used in resource URIs and in UI.
        /// </summary>
        [Input("labPlanName", required: true)]
        public Input<string> LabPlanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetLabPlanInvokeArgs()
        {
        }
        public static new GetLabPlanInvokeArgs Empty => new GetLabPlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetLabPlanResult
    {
        /// <summary>
        /// The allowed regions for the lab creator to use when creating labs using this lab plan.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRegions;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The default lab shutdown profile. This can be changed on a lab resource and only provides a default profile.
        /// </summary>
        public readonly Outputs.AutoShutdownProfileResponse? DefaultAutoShutdownProfile;
        /// <summary>
        /// The default lab connection profile. This can be changed on a lab resource and only provides a default profile.
        /// </summary>
        public readonly Outputs.ConnectionProfileResponse? DefaultConnectionProfile;
        /// <summary>
        /// The lab plan network profile. To enforce lab network policies they must be defined here and cannot be changed when there are existing labs associated with this lab plan.
        /// </summary>
        public readonly Outputs.LabPlanNetworkProfileResponse? DefaultNetworkProfile;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Managed Identity Information
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// Base Url of the lms instance this lab plan can link lab rosters against.
        /// </summary>
        public readonly string? LinkedLmsInstance;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Current provisioning state of the lab plan.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Error details of last operation done on lab plan.
        /// </summary>
        public readonly Outputs.ResourceOperationErrorResponse ResourceOperationError;
        /// <summary>
        /// Resource ID of the Shared Image Gallery attached to this lab plan. When saving a lab template virtual machine image it will be persisted in this gallery. Shared images from the gallery can be made available to use when creating new labs.
        /// </summary>
        public readonly string? SharedGalleryId;
        /// <summary>
        /// Support contact information and instructions for users of the lab plan. This information is displayed to lab owners and virtual machine users for all labs in the lab plan.
        /// </summary>
        public readonly Outputs.SupportInfoResponse? SupportInfo;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the lab plan.
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
        private GetLabPlanResult(
            ImmutableArray<string> allowedRegions,

            string azureApiVersion,

            Outputs.AutoShutdownProfileResponse? defaultAutoShutdownProfile,

            Outputs.ConnectionProfileResponse? defaultConnectionProfile,

            Outputs.LabPlanNetworkProfileResponse? defaultNetworkProfile,

            string id,

            Outputs.IdentityResponse? identity,

            string? linkedLmsInstance,

            string location,

            string name,

            string provisioningState,

            Outputs.ResourceOperationErrorResponse resourceOperationError,

            string? sharedGalleryId,

            Outputs.SupportInfoResponse? supportInfo,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AllowedRegions = allowedRegions;
            AzureApiVersion = azureApiVersion;
            DefaultAutoShutdownProfile = defaultAutoShutdownProfile;
            DefaultConnectionProfile = defaultConnectionProfile;
            DefaultNetworkProfile = defaultNetworkProfile;
            Id = id;
            Identity = identity;
            LinkedLmsInstance = linkedLmsInstance;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceOperationError = resourceOperationError;
            SharedGalleryId = sharedGalleryId;
            SupportInfo = supportInfo;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
