// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforMySQL.V20241001Preview.Outputs
{

    /// <summary>
    /// Maintenance policy of a server.
    /// </summary>
    [OutputType]
    public sealed class MaintenancePolicyResponse
    {
        /// <summary>
        /// The patch strategy of this server
        /// </summary>
        public readonly string? PatchStrategy;

        [OutputConstructor]
        private MaintenancePolicyResponse(string? patchStrategy)
        {
            PatchStrategy = patchStrategy;
        }
    }
}
