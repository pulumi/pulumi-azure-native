// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.V20250113
{
    public static class GetLicense
    {
        /// <summary>
        /// Retrieves information about the view of a license.
        /// </summary>
        public static Task<GetLicenseResult> InvokeAsync(GetLicenseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLicenseResult>("azure-native:hybridcompute/v20250113:getLicense", args ?? new GetLicenseArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about the view of a license.
        /// </summary>
        public static Output<GetLicenseResult> Invoke(GetLicenseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLicenseResult>("azure-native:hybridcompute/v20250113:getLicense", args ?? new GetLicenseInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about the view of a license.
        /// </summary>
        public static Output<GetLicenseResult> Invoke(GetLicenseInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLicenseResult>("azure-native:hybridcompute/v20250113:getLicense", args ?? new GetLicenseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLicenseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the license.
        /// </summary>
        [Input("licenseName", required: true)]
        public string LicenseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLicenseArgs()
        {
        }
        public static new GetLicenseArgs Empty => new GetLicenseArgs();
    }

    public sealed class GetLicenseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the license.
        /// </summary>
        [Input("licenseName", required: true)]
        public Input<string> LicenseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetLicenseInvokeArgs()
        {
        }
        public static new GetLicenseInvokeArgs Empty => new GetLicenseInvokeArgs();
    }


    [OutputType]
    public sealed class GetLicenseResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Describes the properties of a License.
        /// </summary>
        public readonly Outputs.LicenseDetailsResponse? LicenseDetails;
        /// <summary>
        /// The type of the license resource.
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Describes the tenant id.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLicenseResult(
            string id,

            Outputs.LicenseDetailsResponse? licenseDetails,

            string? licenseType,

            string location,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string? tenantId,

            string type)
        {
            Id = id;
            LicenseDetails = licenseDetails;
            LicenseType = licenseType;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            TenantId = tenantId;
            Type = type;
        }
    }
}
