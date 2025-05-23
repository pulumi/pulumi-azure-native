// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConnectedCache.Inputs
{

    /// <summary>
    /// ProxyUrl configuration of cache node
    /// </summary>
    public sealed class ProxyUrlConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host Proxy Address configuration along with port number. This can be a proxy or ip address. ex: xx.xx.xx.xxxx:80 or host name http://exampleproxy.com:80
        /// </summary>
        [Input("proxyUrl")]
        public Input<string>? ProxyUrl { get; set; }

        public ProxyUrlConfigurationArgs()
        {
        }
        public static new ProxyUrlConfigurationArgs Empty => new ProxyUrlConfigurationArgs();
    }
}
