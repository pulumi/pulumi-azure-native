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
    /// Operation response details.
    /// </summary>
    public sealed class ResponseContractArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Operation response description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("headers")]
        private InputList<Inputs.ParameterContractArgs>? _headers;

        /// <summary>
        /// Collection of operation response headers.
        /// </summary>
        public InputList<Inputs.ParameterContractArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.ParameterContractArgs>());
            set => _headers = value;
        }

        [Input("representations")]
        private InputList<Inputs.RepresentationContractArgs>? _representations;

        /// <summary>
        /// Collection of operation response representations.
        /// </summary>
        public InputList<Inputs.RepresentationContractArgs> Representations
        {
            get => _representations ?? (_representations = new InputList<Inputs.RepresentationContractArgs>());
            set => _representations = value;
        }

        /// <summary>
        /// Operation response HTTP status code.
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<int> StatusCode { get; set; } = null!;

        public ResponseContractArgs()
        {
        }
        public static new ResponseContractArgs Empty => new ResponseContractArgs();
    }
}
