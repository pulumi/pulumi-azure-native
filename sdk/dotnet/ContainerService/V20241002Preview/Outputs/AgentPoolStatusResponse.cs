// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Outputs
{

    /// <summary>
    /// Contains read-only information about the Agent Pool.
    /// </summary>
    [OutputType]
    public sealed class AgentPoolStatusResponse
    {
        /// <summary>
        /// Preserves the detailed info of failure. If there was no error, this field is omitted.
        /// </summary>
        public readonly Outputs.CloudErrorBodyResponse ProvisioningError;

        [OutputConstructor]
        private AgentPoolStatusResponse(Outputs.CloudErrorBodyResponse provisioningError)
        {
            ProvisioningError = provisioningError;
        }
    }
}
