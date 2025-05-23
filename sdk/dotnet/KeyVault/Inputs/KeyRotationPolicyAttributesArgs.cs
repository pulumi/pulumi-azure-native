// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.KeyVault.Inputs
{

    public sealed class KeyRotationPolicyAttributesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The expiration time for the new key version. It should be in ISO8601 format. Eg: 'P90D', 'P1Y'.
        /// </summary>
        [Input("expiryTime")]
        public Input<string>? ExpiryTime { get; set; }

        public KeyRotationPolicyAttributesArgs()
        {
        }
        public static new KeyRotationPolicyAttributesArgs Empty => new KeyRotationPolicyAttributesArgs();
    }
}
