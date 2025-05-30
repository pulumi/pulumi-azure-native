// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureFleet.Outputs
{

    /// <summary>
    /// Specifies the required information to reference a compute gallery application
    /// version
    /// </summary>
    [OutputType]
    public sealed class VMGalleryApplicationResponse
    {
        /// <summary>
        /// Optional, Specifies the uri to an azure blob that will replace the default
        /// configuration for the package if provided
        /// </summary>
        public readonly string? ConfigurationReference;
        /// <summary>
        /// If set to true, when a new Gallery Application version is available in PIR/SIG,
        /// it will be automatically updated for the VM/VMSS
        /// </summary>
        public readonly bool? EnableAutomaticUpgrade;
        /// <summary>
        /// Optional, Specifies the order in which the packages have to be installed
        /// </summary>
        public readonly int? Order;
        /// <summary>
        /// Specifies the GalleryApplicationVersion resource id on the form of
        /// /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{application}/versions/{version}
        /// </summary>
        public readonly string PackageReferenceId;
        /// <summary>
        /// Optional, Specifies a passthrough value for more generic context.
        /// </summary>
        public readonly string? Tags;
        /// <summary>
        /// Optional, If true, any failure for any operation in the VmApplication will fail
        /// the deployment
        /// </summary>
        public readonly bool? TreatFailureAsDeploymentFailure;

        [OutputConstructor]
        private VMGalleryApplicationResponse(
            string? configurationReference,

            bool? enableAutomaticUpgrade,

            int? order,

            string packageReferenceId,

            string? tags,

            bool? treatFailureAsDeploymentFailure)
        {
            ConfigurationReference = configurationReference;
            EnableAutomaticUpgrade = enableAutomaticUpgrade;
            Order = order;
            PackageReferenceId = packageReferenceId;
            Tags = tags;
            TreatFailureAsDeploymentFailure = treatFailureAsDeploymentFailure;
        }
    }
}
