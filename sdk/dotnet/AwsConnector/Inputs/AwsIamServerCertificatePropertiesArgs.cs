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
    /// Definition of awsIamServerCertificate
    /// </summary>
    public sealed class AwsIamServerCertificatePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the server certificate
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Property certificateBody
        /// </summary>
        [Input("certificateBody")]
        public Input<string>? CertificateBody { get; set; }

        /// <summary>
        /// Property certificateChain
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// Property path
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Property privateKey
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// Property serverCertificateName
        /// </summary>
        [Input("serverCertificateName")]
        public Input<string>? ServerCertificateName { get; set; }

        [Input("tags")]
        private InputList<Inputs.TagArgs>? _tags;

        /// <summary>
        /// Property tags
        /// </summary>
        public InputList<Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagArgs>());
            set => _tags = value;
        }

        public AwsIamServerCertificatePropertiesArgs()
        {
        }
        public static new AwsIamServerCertificatePropertiesArgs Empty => new AwsIamServerCertificatePropertiesArgs();
    }
}
