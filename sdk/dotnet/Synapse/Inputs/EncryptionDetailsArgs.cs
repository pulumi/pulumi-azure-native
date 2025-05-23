// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Inputs
{

    /// <summary>
    /// Details of the encryption associated with the workspace
    /// </summary>
    public sealed class EncryptionDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Customer Managed Key Details
        /// </summary>
        [Input("cmk")]
        public Input<Inputs.CustomerManagedKeyDetailsArgs>? Cmk { get; set; }

        public EncryptionDetailsArgs()
        {
        }
        public static new EncryptionDetailsArgs Empty => new EncryptionDetailsArgs();
    }
}
