// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// Connection error
    /// </summary>
    [OutputType]
    public sealed class ConnectionErrorResponse
    {
        /// <summary>
        /// Code of the status
        /// </summary>
        public readonly string? Code;
        /// <summary>
        /// Resource ETag
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Description of the status
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ConnectionErrorResponse(
            string? code,

            string? etag,

            string id,

            string? location,

            string? message,

            string name,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Code = code;
            Etag = etag;
            Id = id;
            Location = location;
            Message = message;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}
