// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Outputs
{

    /// <summary>
    /// Adhoc trigger context
    /// </summary>
    [OutputType]
    public sealed class AdhocBasedTriggerContextResponse
    {
        /// <summary>
        /// Type of the specific object - used for deserializing
        /// Expected value is 'AdhocBasedTriggerContext'.
        /// </summary>
        public readonly string ObjectType;
        /// <summary>
        /// Tagging Criteria containing retention tag for adhoc backup.
        /// </summary>
        public readonly Outputs.AdhocBasedTaggingCriteriaResponse TaggingCriteria;

        [OutputConstructor]
        private AdhocBasedTriggerContextResponse(
            string objectType,

            Outputs.AdhocBasedTaggingCriteriaResponse taggingCriteria)
        {
            ObjectType = objectType;
            TaggingCriteria = taggingCriteria;
        }
    }
}
