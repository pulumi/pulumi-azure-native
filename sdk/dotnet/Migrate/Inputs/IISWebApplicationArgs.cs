// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// IISWeb application.
    /// </summary>
    public sealed class IISWebApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the web application id.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Gets or sets the web application name.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// Gets or sets application scratch path.
        /// </summary>
        [Input("applicationScratchPath")]
        public Input<string>? ApplicationScratchPath { get; set; }

        [Input("applications")]
        private InputList<Inputs.IISApplicationDetailsArgs>? _applications;

        /// <summary>
        /// Gets or sets the list of applications for the IIS web site.
        /// </summary>
        public InputList<Inputs.IISApplicationDetailsArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.IISApplicationDetailsArgs>());
            set => _applications = value;
        }

        [Input("bindings")]
        private InputList<Inputs.BindingArgs>? _bindings;

        /// <summary>
        /// Gets or sets the bindings for the application.
        /// </summary>
        public InputList<Inputs.BindingArgs> Bindings
        {
            get => _bindings ?? (_bindings = new InputList<Inputs.BindingArgs>());
            set => _bindings = value;
        }

        [Input("configurations")]
        private InputList<Inputs.WebApplicationConfigurationArgs>? _configurations;

        /// <summary>
        /// Gets or sets application configuration.
        /// </summary>
        public InputList<Inputs.WebApplicationConfigurationArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.WebApplicationConfigurationArgs>());
            set => _configurations = value;
        }

        [Input("directories")]
        private InputList<Inputs.WebApplicationDirectoryArgs>? _directories;

        /// <summary>
        /// Gets or sets application directories.
        /// </summary>
        public InputList<Inputs.WebApplicationDirectoryArgs> Directories
        {
            get => _directories ?? (_directories = new InputList<Inputs.WebApplicationDirectoryArgs>());
            set => _directories = value;
        }

        [Input("discoveredFrameworks")]
        private InputList<Inputs.WebApplicationFrameworkArgs>? _discoveredFrameworks;

        /// <summary>
        /// Gets or sets the discovered frameworks of application.
        /// </summary>
        public InputList<Inputs.WebApplicationFrameworkArgs> DiscoveredFrameworks
        {
            get => _discoveredFrameworks ?? (_discoveredFrameworks = new InputList<Inputs.WebApplicationFrameworkArgs>());
            set => _discoveredFrameworks = value;
        }

        /// <summary>
        /// Gets or sets the display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// IISWeb server.
        /// </summary>
        [Input("iisWebServer")]
        public Input<Inputs.IISWebServerArgs>? IisWebServer { get; set; }

        /// <summary>
        /// Resource Requirements.
        /// </summary>
        [Input("limits")]
        public Input<Inputs.ResourceRequirementsArgs>? Limits { get; set; }

        /// <summary>
        /// Second level entity for virtual directories.
        /// </summary>
        [Input("path")]
        public Input<Inputs.DirectoryPathArgs>? Path { get; set; }

        /// <summary>
        /// Framework specific data for a web application.
        /// </summary>
        [Input("primaryFramework")]
        public Input<Inputs.WebApplicationFrameworkArgs>? PrimaryFramework { get; set; }

        /// <summary>
        /// Resource Requirements.
        /// </summary>
        [Input("requests")]
        public Input<Inputs.ResourceRequirementsArgs>? Requests { get; set; }

        [Input("virtualApplications")]
        private InputList<Inputs.IISVirtualApplicationDetailsArgs>? _virtualApplications;

        /// <summary>
        /// Gets or sets the list of application units for the web site.
        /// </summary>
        public InputList<Inputs.IISVirtualApplicationDetailsArgs> VirtualApplications
        {
            get => _virtualApplications ?? (_virtualApplications = new InputList<Inputs.IISVirtualApplicationDetailsArgs>());
            set => _virtualApplications = value;
        }

        /// <summary>
        /// Gets or sets the web server id.
        /// </summary>
        [Input("webServerId")]
        public Input<string>? WebServerId { get; set; }

        /// <summary>
        /// Gets or sets the web server name.
        /// </summary>
        [Input("webServerName")]
        public Input<string>? WebServerName { get; set; }

        public IISWebApplicationArgs()
        {
        }
        public static new IISWebApplicationArgs Empty => new IISWebApplicationArgs();
    }
}
