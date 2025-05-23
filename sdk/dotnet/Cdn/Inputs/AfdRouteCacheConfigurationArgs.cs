// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.Inputs
{

    /// <summary>
    /// Caching settings for a caching-type route. To disable caching, do not provide a cacheConfiguration object.
    /// </summary>
    public sealed class AfdRouteCacheConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// compression settings.
        /// </summary>
        [Input("compressionSettings")]
        public Input<Inputs.CompressionSettingsArgs>? CompressionSettings { get; set; }

        /// <summary>
        /// query parameters to include or exclude (comma separated).
        /// </summary>
        [Input("queryParameters")]
        public Input<string>? QueryParameters { get; set; }

        /// <summary>
        /// Defines how Frontdoor caches requests that include query strings. You can ignore any query strings when caching, ignore specific query strings, cache every request with a unique URL, or cache specific query strings.
        /// </summary>
        [Input("queryStringCachingBehavior")]
        public InputUnion<string, Pulumi.AzureNative.Cdn.AfdQueryStringCachingBehavior>? QueryStringCachingBehavior { get; set; }

        public AfdRouteCacheConfigurationArgs()
        {
        }
        public static new AfdRouteCacheConfigurationArgs Empty => new AfdRouteCacheConfigurationArgs();
    }
}
