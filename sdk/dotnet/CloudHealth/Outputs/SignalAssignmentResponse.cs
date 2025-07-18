// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CloudHealth.Outputs
{

    /// <summary>
    /// Group of signal definition assignments
    /// </summary>
    [OutputType]
    public sealed class SignalAssignmentResponse
    {
        /// <summary>
        /// Signal definitions referenced by their names. All definitions are combined with an AND operator.
        /// </summary>
        public readonly ImmutableArray<string> SignalDefinitions;

        [OutputConstructor]
        private SignalAssignmentResponse(ImmutableArray<string> signalDefinitions)
        {
            SignalDefinitions = signalDefinitions;
        }
    }
}
