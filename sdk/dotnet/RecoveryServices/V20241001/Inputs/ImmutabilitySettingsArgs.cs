// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.V20241001.Inputs
{

    /// <summary>
    /// Immutability Settings of vault
    /// </summary>
    public sealed class ImmutabilitySettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.V20241001.ImmutabilityState>? State { get; set; }

        public ImmutabilitySettingsArgs()
        {
        }
        public static new ImmutabilitySettingsArgs Empty => new ImmutabilitySettingsArgs();
    }
}
