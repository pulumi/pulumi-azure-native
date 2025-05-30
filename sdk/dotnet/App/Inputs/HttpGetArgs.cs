// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Inputs
{

    /// <summary>
    /// Model representing a http get request.
    /// </summary>
    public sealed class HttpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the file that the request should be saved to.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        [Input("headers")]
        private InputList<string>? _headers;

        /// <summary>
        /// List of headers to send with the request.
        /// </summary>
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        /// <summary>
        /// URL to make HTTP GET request against.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public HttpGetArgs()
        {
        }
        public static new HttpGetArgs Empty => new HttpGetArgs();
    }
}
