// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The export policy for a container registry.
    /// </summary>
    [OutputType]
    public sealed class ExportPolicyResponse
    {
        /// <summary>
        /// The value that indicates whether the policy is enabled or not.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ExportPolicyResponse(string? status)
        {
            Status = status;
        }
    }
}
