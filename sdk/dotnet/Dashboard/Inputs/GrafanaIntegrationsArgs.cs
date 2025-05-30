// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Dashboard.Inputs
{

    /// <summary>
    /// GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards, alerting defaults) for common monitoring scenarios.
    /// </summary>
    public sealed class GrafanaIntegrationsArgs : global::Pulumi.ResourceArgs
    {
        [Input("azureMonitorWorkspaceIntegrations")]
        private InputList<Inputs.AzureMonitorWorkspaceIntegrationArgs>? _azureMonitorWorkspaceIntegrations;
        public InputList<Inputs.AzureMonitorWorkspaceIntegrationArgs> AzureMonitorWorkspaceIntegrations
        {
            get => _azureMonitorWorkspaceIntegrations ?? (_azureMonitorWorkspaceIntegrations = new InputList<Inputs.AzureMonitorWorkspaceIntegrationArgs>());
            set => _azureMonitorWorkspaceIntegrations = value;
        }

        public GrafanaIntegrationsArgs()
        {
        }
        public static new GrafanaIntegrationsArgs Empty => new GrafanaIntegrationsArgs();
    }
}
