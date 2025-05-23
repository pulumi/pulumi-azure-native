// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Inputs
{

    /// <summary>
    /// Http message diagnostic settings.
    /// </summary>
    public sealed class HttpMessageDiagnosticArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Body logging settings.
        /// </summary>
        [Input("body")]
        public Input<Inputs.BodyDiagnosticSettingsArgs>? Body { get; set; }

        /// <summary>
        /// Data masking settings.
        /// </summary>
        [Input("dataMasking")]
        public Input<Inputs.DataMaskingArgs>? DataMasking { get; set; }

        [Input("headers")]
        private InputList<string>? _headers;

        /// <summary>
        /// Array of HTTP Headers to log.
        /// </summary>
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        public HttpMessageDiagnosticArgs()
        {
        }
        public static new HttpMessageDiagnosticArgs Empty => new HttpMessageDiagnosticArgs();
    }
}
