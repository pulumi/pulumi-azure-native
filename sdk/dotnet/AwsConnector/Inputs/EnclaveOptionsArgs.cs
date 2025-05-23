// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of EnclaveOptions
    /// </summary>
    public sealed class EnclaveOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;If this parameter is set to &lt;code&gt;true&lt;/code&gt;, the instance is enabled for Amazon Web Services Nitro Enclaves; otherwise, it is not enabled for Amazon Web Services Nitro Enclaves.&lt;/p&gt;
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public EnclaveOptionsArgs()
        {
        }
        public static new EnclaveOptionsArgs Empty => new EnclaveOptionsArgs();
    }
}
