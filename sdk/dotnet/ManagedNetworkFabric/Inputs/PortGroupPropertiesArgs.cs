// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Inputs
{

    /// <summary>
    /// Port Group properties.
    /// </summary>
    public sealed class PortGroupPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the port group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ports")]
        private InputList<string>? _ports;

        /// <summary>
        /// List of the ports that need to be matched.
        /// </summary>
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        public PortGroupPropertiesArgs()
        {
        }
        public static new PortGroupPropertiesArgs Empty => new PortGroupPropertiesArgs();
    }
}
