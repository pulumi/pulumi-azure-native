// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Inputs
{

    /// <summary>
    /// Options for controlling the validation of TLS endpoints.
    /// </summary>
    public sealed class TlsValidationOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set to 'true' causes the certificate subject name validation to be skipped. Default is 'false'.
        /// </summary>
        [Input("ignoreHostname")]
        public Input<string>? IgnoreHostname { get; set; }

        /// <summary>
        /// When set to 'true' causes the certificate chain trust validation to be skipped. Default is 'false'.
        /// </summary>
        [Input("ignoreSignature")]
        public Input<string>? IgnoreSignature { get; set; }

        public TlsValidationOptionsArgs()
        {
        }
        public static new TlsValidationOptionsArgs Empty => new TlsValidationOptionsArgs();
    }
}
