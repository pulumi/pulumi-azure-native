// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Inputs
{

    /// <summary>
    /// Immutability Settings at vault level
    /// </summary>
    public sealed class ImmutabilitySettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutability state
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.DataProtection.ImmutabilityState>? State { get; set; }

        public ImmutabilitySettingsArgs()
        {
        }
        public static new ImmutabilitySettingsArgs Empty => new ImmutabilitySettingsArgs();
    }
}
