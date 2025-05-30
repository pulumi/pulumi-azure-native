// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApplicationInsights.Outputs
{

    /// <summary>
    /// Localized template data and gallery information.
    /// </summary>
    [OutputType]
    public sealed class WorkbookTemplateLocalizedGalleryResponse
    {
        /// <summary>
        /// Workbook galleries supported by the template.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkbookTemplateGalleryResponse> Galleries;
        /// <summary>
        /// Valid JSON object containing workbook template payload.
        /// </summary>
        public readonly object? TemplateData;

        [OutputConstructor]
        private WorkbookTemplateLocalizedGalleryResponse(
            ImmutableArray<Outputs.WorkbookTemplateGalleryResponse> galleries,

            object? templateData)
        {
            Galleries = galleries;
            TemplateData = templateData;
        }
    }
}
