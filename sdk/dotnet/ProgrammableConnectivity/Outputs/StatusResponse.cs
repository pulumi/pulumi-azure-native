// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProgrammableConnectivity.Outputs
{

    /// <summary>
    /// Description of the current status of the OperatorApiConnection resource.
    /// </summary>
    [OutputType]
    public sealed class StatusResponse
    {
        /// <summary>
        /// Explanation of the current state of the OperatorApiConnection resource.
        /// </summary>
        public readonly string? Reason;
        /// <summary>
        /// Current state of the OperatorApiConnection resource.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private StatusResponse(
            string? reason,

            string? state)
        {
            Reason = reason;
            State = state;
        }
    }
}
