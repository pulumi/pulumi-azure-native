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
    /// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp Service Authentication/Authorization.
    /// </summary>
    public sealed class HttpSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration settings of a forward proxy used to make the requests.
        /// </summary>
        [Input("forwardProxy")]
        public Input<Inputs.ForwardProxyArgs>? ForwardProxy { get; set; }

        /// <summary>
        /// &lt;code&gt;false&lt;/code&gt; if the authentication/authorization responses not having the HTTPS scheme are permissible; otherwise, &lt;code&gt;true&lt;/code&gt;.
        /// </summary>
        [Input("requireHttps")]
        public Input<bool>? RequireHttps { get; set; }

        /// <summary>
        /// The configuration settings of the paths HTTP requests.
        /// </summary>
        [Input("routes")]
        public Input<Inputs.HttpSettingsRoutesArgs>? Routes { get; set; }

        public HttpSettingsArgs()
        {
        }
        public static new HttpSettingsArgs Empty => new HttpSettingsArgs();
    }
}
