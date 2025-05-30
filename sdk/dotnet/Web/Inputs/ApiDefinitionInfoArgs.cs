// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// Information about the formal API definition for the app.
    /// </summary>
    public sealed class ApiDefinitionInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the API definition.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ApiDefinitionInfoArgs()
        {
        }
        public static new ApiDefinitionInfoArgs Empty => new ApiDefinitionInfoArgs();
    }
}
