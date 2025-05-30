// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StandbyPool.Inputs
{

    /// <summary>
    /// Details of the ContainerGroupProperties.
    /// </summary>
    public sealed class ContainerGroupPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies container group profile of standby container groups.
        /// </summary>
        [Input("containerGroupProfile", required: true)]
        public Input<Inputs.ContainerGroupProfileArgs> ContainerGroupProfile { get; set; } = null!;

        [Input("subnetIds")]
        private InputList<Inputs.SubnetArgs>? _subnetIds;

        /// <summary>
        /// Specifies subnet Ids for container group.
        /// </summary>
        public InputList<Inputs.SubnetArgs> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<Inputs.SubnetArgs>());
            set => _subnetIds = value;
        }

        public ContainerGroupPropertiesArgs()
        {
        }
        public static new ContainerGroupPropertiesArgs Empty => new ContainerGroupPropertiesArgs();
    }
}
