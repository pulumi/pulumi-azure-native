// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.Inputs
{

    /// <summary>
    /// Describes a license in a hybrid machine.
    /// </summary>
    public sealed class LicenseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the properties of a License.
        /// </summary>
        [Input("licenseDetails")]
        public Input<Inputs.LicenseDetailsArgs>? LicenseDetails { get; set; }

        /// <summary>
        /// The type of the license resource.
        /// </summary>
        [Input("licenseType")]
        public InputUnion<string, Pulumi.AzureNative.HybridCompute.LicenseType>? LicenseType { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Describes the tenant id.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public LicenseArgs()
        {
        }
        public static new LicenseArgs Empty => new LicenseArgs();
    }
}
