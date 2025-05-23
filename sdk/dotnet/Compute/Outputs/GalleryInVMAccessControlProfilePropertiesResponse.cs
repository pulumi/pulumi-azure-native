// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// Describes the properties of a gallery inVMAccessControlProfile.
    /// </summary>
    [OutputType]
    public sealed class GalleryInVMAccessControlProfilePropertiesResponse
    {
        /// <summary>
        /// This property allows you to specify the Endpoint type for which this profile is defining the access control for. Possible values are: 'WireServer' or 'IMDS'
        /// </summary>
        public readonly string ApplicableHostEndpoint;
        /// <summary>
        /// The description of this gallery inVMAccessControlProfile resources. This property is updatable.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// This property allows you to specify the OS type of the VMs/VMSS for which this profile can be used against. Possible values are: 'Windows' or 'Linux'
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private GalleryInVMAccessControlProfilePropertiesResponse(
            string applicableHostEndpoint,

            string? description,

            string osType,

            string provisioningState)
        {
            ApplicableHostEndpoint = applicableHostEndpoint;
            Description = description;
            OsType = osType;
            ProvisioningState = provisioningState;
        }
    }
}
