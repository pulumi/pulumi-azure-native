// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SqlVirtualMachine.Inputs
{

    /// <summary>
    /// Availability group replica configuration.
    /// </summary>
    public sealed class AgReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Replica commit mode in availability group.
        /// </summary>
        [Input("commit")]
        public InputUnion<string, Pulumi.AzureNative.SqlVirtualMachine.Commit>? Commit { get; set; }

        /// <summary>
        /// Replica failover mode in availability group.
        /// </summary>
        [Input("failover")]
        public InputUnion<string, Pulumi.AzureNative.SqlVirtualMachine.Failover>? Failover { get; set; }

        /// <summary>
        /// Replica readable secondary mode in availability group.
        /// </summary>
        [Input("readableSecondary")]
        public InputUnion<string, Pulumi.AzureNative.SqlVirtualMachine.ReadableSecondary>? ReadableSecondary { get; set; }

        /// <summary>
        /// Replica Role in availability group.
        /// </summary>
        [Input("role")]
        public InputUnion<string, Pulumi.AzureNative.SqlVirtualMachine.Role>? Role { get; set; }

        /// <summary>
        /// Sql VirtualMachine Instance Id.
        /// </summary>
        [Input("sqlVirtualMachineInstanceId")]
        public Input<string>? SqlVirtualMachineInstanceId { get; set; }

        public AgReplicaArgs()
        {
        }
        public static new AgReplicaArgs Empty => new AgReplicaArgs();
    }
}
