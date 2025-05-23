// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Framework specific data for a web application.
    /// </summary>
    public sealed class WebApplicationFrameworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets Name of the framework.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Gets or sets Version of the framework.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public WebApplicationFrameworkArgs()
        {
        }
        public static new WebApplicationFrameworkArgs Empty => new WebApplicationFrameworkArgs();
    }
}
