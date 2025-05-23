// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// Immutability Settings of vault
    /// </summary>
    [OutputType]
    public sealed class ImmutabilitySettingsResponse
    {
        public readonly string? State;

        [OutputConstructor]
        private ImmutabilitySettingsResponse(string? state)
        {
            State = state;
        }
    }
}
