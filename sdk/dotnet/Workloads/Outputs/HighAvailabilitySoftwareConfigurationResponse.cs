// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Gets or sets the HA software configuration.
    /// </summary>
    [OutputType]
    public sealed class HighAvailabilitySoftwareConfigurationResponse
    {
        /// <summary>
        /// The fencing client id.
        /// </summary>
        public readonly string FencingClientId;
        /// <summary>
        /// The fencing client id secret/password. The secret should never expire. This will be used pacemaker to start/stop the cluster VMs.
        /// </summary>
        public readonly string FencingClientPassword;

        [OutputConstructor]
        private HighAvailabilitySoftwareConfigurationResponse(
            string fencingClientId,

            string fencingClientPassword)
        {
            FencingClientId = fencingClientId;
            FencingClientPassword = fencingClientPassword;
        }
    }
}
