// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// A WebLinkedService that uses anonymous authentication to communicate with an HTTP endpoint.
    /// </summary>
    public sealed class WebAnonymousAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of authentication used to connect to the web table source.
        /// Expected value is 'Anonymous'.
        /// </summary>
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        /// <summary>
        /// The URL of the web service endpoint, e.g. https://www.microsoft.com . Type: string (or Expression with resultType string).
        /// </summary>
        [Input("url", required: true)]
        public Input<object> Url { get; set; } = null!;

        public WebAnonymousAuthenticationArgs()
        {
        }
        public static new WebAnonymousAuthenticationArgs Empty => new WebAnonymousAuthenticationArgs();
    }
}
