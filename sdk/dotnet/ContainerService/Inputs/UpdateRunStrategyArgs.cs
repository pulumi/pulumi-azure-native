// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// Defines the update sequence of the clusters via stages and groups.
    /// 
    /// Stages within a run are executed sequentially one after another.
    /// Groups within a stage are executed in parallel.
    /// Member clusters within a group are updated sequentially one after another.
    /// 
    /// A valid strategy contains no duplicate groups within or across stages.
    /// </summary>
    public sealed class UpdateRunStrategyArgs : global::Pulumi.ResourceArgs
    {
        [Input("stages", required: true)]
        private InputList<Inputs.UpdateStageArgs>? _stages;

        /// <summary>
        /// The list of stages that compose this update run. Min size: 1.
        /// </summary>
        public InputList<Inputs.UpdateStageArgs> Stages
        {
            get => _stages ?? (_stages = new InputList<Inputs.UpdateStageArgs>());
            set => _stages = value;
        }

        public UpdateRunStrategyArgs()
        {
        }
        public static new UpdateRunStrategyArgs Empty => new UpdateRunStrategyArgs();
    }
}
