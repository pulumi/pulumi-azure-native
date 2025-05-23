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
    /// Definition of Certificate
    /// </summary>
    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the certificate.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// &lt;p&gt;The Base64-encoded certificate data required to communicate with your cluster. Add this to the &lt;code&gt;certificate-authority-data&lt;/code&gt; section of the &lt;code&gt;kubeconfig&lt;/code&gt; file for your cluster.&lt;/p&gt;
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}
