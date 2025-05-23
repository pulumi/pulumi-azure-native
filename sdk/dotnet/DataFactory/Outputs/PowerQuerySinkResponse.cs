// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Power query sink.
    /// </summary>
    [OutputType]
    public sealed class PowerQuerySinkResponse
    {
        /// <summary>
        /// Dataset reference.
        /// </summary>
        public readonly Outputs.DatasetReferenceResponse? Dataset;
        /// <summary>
        /// Transformation description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Flowlet Reference
        /// </summary>
        public readonly Outputs.DataFlowReferenceResponse? Flowlet;
        /// <summary>
        /// Linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse? LinkedService;
        /// <summary>
        /// Transformation name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Rejected data linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse? RejectedDataLinkedService;
        /// <summary>
        /// Schema linked service reference.
        /// </summary>
        public readonly Outputs.LinkedServiceReferenceResponse? SchemaLinkedService;
        /// <summary>
        /// sink script.
        /// </summary>
        public readonly string? Script;

        [OutputConstructor]
        private PowerQuerySinkResponse(
            Outputs.DatasetReferenceResponse? dataset,

            string? description,

            Outputs.DataFlowReferenceResponse? flowlet,

            Outputs.LinkedServiceReferenceResponse? linkedService,

            string name,

            Outputs.LinkedServiceReferenceResponse? rejectedDataLinkedService,

            Outputs.LinkedServiceReferenceResponse? schemaLinkedService,

            string? script)
        {
            Dataset = dataset;
            Description = description;
            Flowlet = flowlet;
            LinkedService = linkedService;
            Name = name;
            RejectedDataLinkedService = rejectedDataLinkedService;
            SchemaLinkedService = schemaLinkedService;
            Script = script;
        }
    }
}
