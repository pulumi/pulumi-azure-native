// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.V20210101Preview
{
    public static class GetAddon
    {
        /// <summary>
        /// An addon resource
        /// </summary>
        public static Task<GetAddonResult> InvokeAsync(GetAddonArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAddonResult>("azure-native:avs/v20210101preview:getAddon", args ?? new GetAddonArgs(), options.WithDefaults());

        /// <summary>
        /// An addon resource
        /// </summary>
        public static Output<GetAddonResult> Invoke(GetAddonInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddonResult>("azure-native:avs/v20210101preview:getAddon", args ?? new GetAddonInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// An addon resource
        /// </summary>
        public static Output<GetAddonResult> Invoke(GetAddonInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddonResult>("azure-native:avs/v20210101preview:getAddon", args ?? new GetAddonInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddonArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the addon for the private cloud
        /// </summary>
        [Input("addonName", required: true)]
        public string AddonName { get; set; } = null!;

        /// <summary>
        /// Name of the private cloud
        /// </summary>
        [Input("privateCloudName", required: true)]
        public string PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAddonArgs()
        {
        }
        public static new GetAddonArgs Empty => new GetAddonArgs();
    }

    public sealed class GetAddonInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the addon for the private cloud
        /// </summary>
        [Input("addonName", required: true)]
        public Input<string> AddonName { get; set; } = null!;

        /// <summary>
        /// Name of the private cloud
        /// </summary>
        [Input("privateCloudName", required: true)]
        public Input<string> PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAddonInvokeArgs()
        {
        }
        public static new GetAddonInvokeArgs Empty => new GetAddonInvokeArgs();
    }


    [OutputType]
    public sealed class GetAddonResult
    {
        /// <summary>
        /// The type of private cloud addon
        /// </summary>
        public readonly string? AddonType;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The SRM license
        /// </summary>
        public readonly string? LicenseKey;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The state of the addon provisioning
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAddonResult(
            string? addonType,

            string id,

            string? licenseKey,

            string name,

            string provisioningState,

            string type)
        {
            AddonType = addonType;
            Id = id;
            LicenseKey = licenseKey;
            Name = name;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
