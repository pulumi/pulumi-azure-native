// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources.Inputs
{

    /// <summary>
    /// Entity representing the reference to the template.
    /// </summary>
    public sealed class DeploymentStacksTemplateLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If included, must match the ContentVersion in the template.
        /// </summary>
        [Input("contentVersion")]
        public Input<string>? ContentVersion { get; set; }

        /// <summary>
        /// The resourceId of a Template Spec. Use either the id or uri property, but not both.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The query string (for example, a SAS token) to be used with the templateLink URI.
        /// </summary>
        [Input("queryString")]
        public Input<string>? QueryString { get; set; }

        /// <summary>
        /// The relativePath property can be used to deploy a linked template at a location relative to the parent. If the parent template was linked with a TemplateSpec, this will reference an artifact in the TemplateSpec.  If the parent was linked with a URI, the child deployment will be a combination of the parent and relativePath URIs.
        /// </summary>
        [Input("relativePath")]
        public Input<string>? RelativePath { get; set; }

        /// <summary>
        /// The URI of the template to deploy. Use either the uri or id property, but not both.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public DeploymentStacksTemplateLinkArgs()
        {
        }
        public static new DeploymentStacksTemplateLinkArgs Empty => new DeploymentStacksTemplateLinkArgs();
    }
}
