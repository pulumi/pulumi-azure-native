// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Inputs
{

    /// <summary>
    /// The SAP Software configuration Input when the software is installed externally outside the service.
    /// </summary>
    public sealed class ExternalInstallationSoftwareConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the virtual machine containing the central server instance.
        /// </summary>
        [Input("centralServerVmId")]
        public Input<string>? CentralServerVmId { get; set; }

        /// <summary>
        /// The SAP software installation Type.
        /// Expected value is 'External'.
        /// </summary>
        [Input("softwareInstallationType", required: true)]
        public Input<string> SoftwareInstallationType { get; set; } = null!;

        public ExternalInstallationSoftwareConfigurationArgs()
        {
        }
        public static new ExternalInstallationSoftwareConfigurationArgs Empty => new ExternalInstallationSoftwareConfigurationArgs();
    }
}
