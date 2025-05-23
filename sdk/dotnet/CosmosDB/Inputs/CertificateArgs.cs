// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CosmosDB.Inputs
{

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// PEM formatted public key.
        /// </summary>
        [Input("pem")]
        public Input<string>? Pem { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}
