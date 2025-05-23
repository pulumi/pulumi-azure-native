// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Inputs
{

    /// <summary>
    /// Specifies a CORS rule for the Blob service.
    /// </summary>
    public sealed class CorsRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedHeaders", required: true)]
        private InputList<string>? _allowedHeaders;

        /// <summary>
        /// Required if CorsRule element is present. A list of headers allowed to be part of the cross-origin request.
        /// </summary>
        public InputList<string> AllowedHeaders
        {
            get => _allowedHeaders ?? (_allowedHeaders = new InputList<string>());
            set => _allowedHeaders = value;
        }

        [Input("allowedMethods", required: true)]
        private InputList<Union<string, Pulumi.AzureNative.Storage.AllowedMethods>>? _allowedMethods;

        /// <summary>
        /// Required if CorsRule element is present. A list of HTTP methods that are allowed to be executed by the origin.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Storage.AllowedMethods>> AllowedMethods
        {
            get => _allowedMethods ?? (_allowedMethods = new InputList<Union<string, Pulumi.AzureNative.Storage.AllowedMethods>>());
            set => _allowedMethods = value;
        }

        [Input("allowedOrigins", required: true)]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// Required if CorsRule element is present. A list of origin domains that will be allowed via CORS, or "*" to allow all domains
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        [Input("exposedHeaders", required: true)]
        private InputList<string>? _exposedHeaders;

        /// <summary>
        /// Required if CorsRule element is present. A list of response headers to expose to CORS clients.
        /// </summary>
        public InputList<string> ExposedHeaders
        {
            get => _exposedHeaders ?? (_exposedHeaders = new InputList<string>());
            set => _exposedHeaders = value;
        }

        /// <summary>
        /// Required if CorsRule element is present. The number of seconds that the client/browser should cache a preflight response.
        /// </summary>
        [Input("maxAgeInSeconds", required: true)]
        public Input<int> MaxAgeInSeconds { get; set; } = null!;

        public CorsRuleArgs()
        {
        }
        public static new CorsRuleArgs Empty => new CorsRuleArgs();
    }
}
