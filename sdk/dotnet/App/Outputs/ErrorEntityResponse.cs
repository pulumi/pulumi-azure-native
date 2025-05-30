// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Body of the error response returned from the API.
    /// </summary>
    [OutputType]
    public sealed class ErrorEntityResponse
    {
        /// <summary>
        /// Basic error code.
        /// </summary>
        public readonly string? Code;
        /// <summary>
        /// Error Details.
        /// </summary>
        public readonly ImmutableArray<Outputs.ErrorEntityResponse> Details;
        /// <summary>
        /// Type of error.
        /// </summary>
        public readonly string? ExtendedCode;
        /// <summary>
        /// Inner errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.ErrorEntityResponse> InnerErrors;
        /// <summary>
        /// Any details of the error.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Message template.
        /// </summary>
        public readonly string? MessageTemplate;
        /// <summary>
        /// Parameters for the template.
        /// </summary>
        public readonly ImmutableArray<string> Parameters;
        /// <summary>
        /// The error target.
        /// </summary>
        public readonly string? Target;

        [OutputConstructor]
        private ErrorEntityResponse(
            string? code,

            ImmutableArray<Outputs.ErrorEntityResponse> details,

            string? extendedCode,

            ImmutableArray<Outputs.ErrorEntityResponse> innerErrors,

            string? message,

            string? messageTemplate,

            ImmutableArray<string> parameters,

            string? target)
        {
            Code = code;
            Details = details;
            ExtendedCode = extendedCode;
            InnerErrors = innerErrors;
            Message = message;
            MessageTemplate = messageTemplate;
            Parameters = parameters;
            Target = target;
        }
    }
}
