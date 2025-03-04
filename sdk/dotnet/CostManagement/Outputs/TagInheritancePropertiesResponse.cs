// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Outputs
{

    /// <summary>
    /// The properties of the tag inheritance setting.
    /// </summary>
    [OutputType]
    public sealed class TagInheritancePropertiesResponse
    {
        /// <summary>
        /// When resource has the same tag as subscription or resource group and this property is set to true - the subscription or resource group tag will be applied. If subscription and resource group tags are also the same, subscription tag will be applied.
        /// </summary>
        public readonly bool PreferContainerTags;

        [OutputConstructor]
        private TagInheritancePropertiesResponse(bool preferContainerTags)
        {
            PreferContainerTags = preferContainerTags;
        }
    }
}
