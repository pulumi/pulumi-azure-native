// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.V20200214Preview.Inputs
{

    public sealed class ServerPropertiesDelegatedSubnetArgumentsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// delegated subnet arm resource id.
        /// </summary>
        [Input("subnetArmResourceId")]
        public Input<string>? SubnetArmResourceId { get; set; }

        public ServerPropertiesDelegatedSubnetArgumentsArgs()
        {
        }
        public static new ServerPropertiesDelegatedSubnetArgumentsArgs Empty => new ServerPropertiesDelegatedSubnetArgumentsArgs();
    }
}
