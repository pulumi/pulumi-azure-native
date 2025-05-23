// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Private endpoint connection definition.
    /// </summary>
    public sealed class RegistryPrivateEndpointConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is the private endpoint connection name created on SRP
        /// Full resource id: /subscriptions/{subId}/resourceGroups/{rgName}/providers/Microsoft.MachineLearningServices/{resourceType}/{resourceName}/registryPrivateEndpointConnections/{peConnectionName}
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Same as workspace location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Properties of the Private Endpoint Connection
        /// </summary>
        [Input("properties")]
        public Input<Inputs.RegistryPrivateEndpointConnectionPropertiesArgs>? Properties { get; set; }

        public RegistryPrivateEndpointConnectionArgs()
        {
        }
        public static new RegistryPrivateEndpointConnectionArgs Empty => new RegistryPrivateEndpointConnectionArgs();
    }
}
