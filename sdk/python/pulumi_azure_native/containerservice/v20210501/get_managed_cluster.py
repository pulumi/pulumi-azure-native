# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs

__all__ = [
    'GetManagedClusterResult',
    'AwaitableGetManagedClusterResult',
    'get_managed_cluster',
    'get_managed_cluster_output',
]

@pulumi.output_type
class GetManagedClusterResult:
    """
    Managed cluster.
    """
    def __init__(__self__, aad_profile=None, addon_profiles=None, agent_pool_profiles=None, api_server_access_profile=None, auto_scaler_profile=None, auto_upgrade_profile=None, azure_portal_fqdn=None, disable_local_accounts=None, disk_encryption_set_id=None, dns_prefix=None, enable_pod_security_policy=None, enable_rbac=None, extended_location=None, fqdn=None, fqdn_subdomain=None, http_proxy_config=None, id=None, identity=None, identity_profile=None, kubernetes_version=None, linux_profile=None, location=None, max_agent_pools=None, name=None, network_profile=None, node_resource_group=None, pod_identity_profile=None, power_state=None, private_fqdn=None, private_link_resources=None, provisioning_state=None, service_principal_profile=None, sku=None, tags=None, type=None, windows_profile=None):
        if aad_profile and not isinstance(aad_profile, dict):
            raise TypeError("Expected argument 'aad_profile' to be a dict")
        pulumi.set(__self__, "aad_profile", aad_profile)
        if addon_profiles and not isinstance(addon_profiles, dict):
            raise TypeError("Expected argument 'addon_profiles' to be a dict")
        pulumi.set(__self__, "addon_profiles", addon_profiles)
        if agent_pool_profiles and not isinstance(agent_pool_profiles, list):
            raise TypeError("Expected argument 'agent_pool_profiles' to be a list")
        pulumi.set(__self__, "agent_pool_profiles", agent_pool_profiles)
        if api_server_access_profile and not isinstance(api_server_access_profile, dict):
            raise TypeError("Expected argument 'api_server_access_profile' to be a dict")
        pulumi.set(__self__, "api_server_access_profile", api_server_access_profile)
        if auto_scaler_profile and not isinstance(auto_scaler_profile, dict):
            raise TypeError("Expected argument 'auto_scaler_profile' to be a dict")
        pulumi.set(__self__, "auto_scaler_profile", auto_scaler_profile)
        if auto_upgrade_profile and not isinstance(auto_upgrade_profile, dict):
            raise TypeError("Expected argument 'auto_upgrade_profile' to be a dict")
        pulumi.set(__self__, "auto_upgrade_profile", auto_upgrade_profile)
        if azure_portal_fqdn and not isinstance(azure_portal_fqdn, str):
            raise TypeError("Expected argument 'azure_portal_fqdn' to be a str")
        pulumi.set(__self__, "azure_portal_fqdn", azure_portal_fqdn)
        if disable_local_accounts and not isinstance(disable_local_accounts, bool):
            raise TypeError("Expected argument 'disable_local_accounts' to be a bool")
        pulumi.set(__self__, "disable_local_accounts", disable_local_accounts)
        if disk_encryption_set_id and not isinstance(disk_encryption_set_id, str):
            raise TypeError("Expected argument 'disk_encryption_set_id' to be a str")
        pulumi.set(__self__, "disk_encryption_set_id", disk_encryption_set_id)
        if dns_prefix and not isinstance(dns_prefix, str):
            raise TypeError("Expected argument 'dns_prefix' to be a str")
        pulumi.set(__self__, "dns_prefix", dns_prefix)
        if enable_pod_security_policy and not isinstance(enable_pod_security_policy, bool):
            raise TypeError("Expected argument 'enable_pod_security_policy' to be a bool")
        pulumi.set(__self__, "enable_pod_security_policy", enable_pod_security_policy)
        if enable_rbac and not isinstance(enable_rbac, bool):
            raise TypeError("Expected argument 'enable_rbac' to be a bool")
        pulumi.set(__self__, "enable_rbac", enable_rbac)
        if extended_location and not isinstance(extended_location, dict):
            raise TypeError("Expected argument 'extended_location' to be a dict")
        pulumi.set(__self__, "extended_location", extended_location)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if fqdn_subdomain and not isinstance(fqdn_subdomain, str):
            raise TypeError("Expected argument 'fqdn_subdomain' to be a str")
        pulumi.set(__self__, "fqdn_subdomain", fqdn_subdomain)
        if http_proxy_config and not isinstance(http_proxy_config, dict):
            raise TypeError("Expected argument 'http_proxy_config' to be a dict")
        pulumi.set(__self__, "http_proxy_config", http_proxy_config)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if identity_profile and not isinstance(identity_profile, dict):
            raise TypeError("Expected argument 'identity_profile' to be a dict")
        pulumi.set(__self__, "identity_profile", identity_profile)
        if kubernetes_version and not isinstance(kubernetes_version, str):
            raise TypeError("Expected argument 'kubernetes_version' to be a str")
        pulumi.set(__self__, "kubernetes_version", kubernetes_version)
        if linux_profile and not isinstance(linux_profile, dict):
            raise TypeError("Expected argument 'linux_profile' to be a dict")
        pulumi.set(__self__, "linux_profile", linux_profile)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if max_agent_pools and not isinstance(max_agent_pools, int):
            raise TypeError("Expected argument 'max_agent_pools' to be a int")
        pulumi.set(__self__, "max_agent_pools", max_agent_pools)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_profile and not isinstance(network_profile, dict):
            raise TypeError("Expected argument 'network_profile' to be a dict")
        pulumi.set(__self__, "network_profile", network_profile)
        if node_resource_group and not isinstance(node_resource_group, str):
            raise TypeError("Expected argument 'node_resource_group' to be a str")
        pulumi.set(__self__, "node_resource_group", node_resource_group)
        if pod_identity_profile and not isinstance(pod_identity_profile, dict):
            raise TypeError("Expected argument 'pod_identity_profile' to be a dict")
        pulumi.set(__self__, "pod_identity_profile", pod_identity_profile)
        if power_state and not isinstance(power_state, dict):
            raise TypeError("Expected argument 'power_state' to be a dict")
        pulumi.set(__self__, "power_state", power_state)
        if private_fqdn and not isinstance(private_fqdn, str):
            raise TypeError("Expected argument 'private_fqdn' to be a str")
        pulumi.set(__self__, "private_fqdn", private_fqdn)
        if private_link_resources and not isinstance(private_link_resources, list):
            raise TypeError("Expected argument 'private_link_resources' to be a list")
        pulumi.set(__self__, "private_link_resources", private_link_resources)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if service_principal_profile and not isinstance(service_principal_profile, dict):
            raise TypeError("Expected argument 'service_principal_profile' to be a dict")
        pulumi.set(__self__, "service_principal_profile", service_principal_profile)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if windows_profile and not isinstance(windows_profile, dict):
            raise TypeError("Expected argument 'windows_profile' to be a dict")
        pulumi.set(__self__, "windows_profile", windows_profile)

    @property
    @pulumi.getter(name="aadProfile")
    def aad_profile(self) -> Optional['outputs.ManagedClusterAADProfileResponse']:
        """
        The Azure Active Directory configuration.
        """
        return pulumi.get(self, "aad_profile")

    @property
    @pulumi.getter(name="addonProfiles")
    def addon_profiles(self) -> Optional[Mapping[str, 'outputs.ManagedClusterAddonProfileResponse']]:
        """
        The profile of managed cluster add-on.
        """
        return pulumi.get(self, "addon_profiles")

    @property
    @pulumi.getter(name="agentPoolProfiles")
    def agent_pool_profiles(self) -> Optional[Sequence['outputs.ManagedClusterAgentPoolProfileResponse']]:
        """
        The agent pool properties.
        """
        return pulumi.get(self, "agent_pool_profiles")

    @property
    @pulumi.getter(name="apiServerAccessProfile")
    def api_server_access_profile(self) -> Optional['outputs.ManagedClusterAPIServerAccessProfileResponse']:
        """
        The access profile for managed cluster API server.
        """
        return pulumi.get(self, "api_server_access_profile")

    @property
    @pulumi.getter(name="autoScalerProfile")
    def auto_scaler_profile(self) -> Optional['outputs.ManagedClusterPropertiesResponseAutoScalerProfile']:
        """
        Parameters to be applied to the cluster-autoscaler when enabled
        """
        return pulumi.get(self, "auto_scaler_profile")

    @property
    @pulumi.getter(name="autoUpgradeProfile")
    def auto_upgrade_profile(self) -> Optional['outputs.ManagedClusterAutoUpgradeProfileResponse']:
        """
        The auto upgrade configuration.
        """
        return pulumi.get(self, "auto_upgrade_profile")

    @property
    @pulumi.getter(name="azurePortalFQDN")
    def azure_portal_fqdn(self) -> str:
        """
        The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some responses, which Kubernetes APIServer doesn't handle by default. This special FQDN supports CORS, allowing the Azure Portal to function properly.
        """
        return pulumi.get(self, "azure_portal_fqdn")

    @property
    @pulumi.getter(name="disableLocalAccounts")
    def disable_local_accounts(self) -> Optional[bool]:
        """
        If set to true, getting static credentials will be disabled for this cluster. This must only be used on Managed Clusters that are AAD enabled. For more details see [disable local accounts](https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview).
        """
        return pulumi.get(self, "disable_local_accounts")

    @property
    @pulumi.getter(name="diskEncryptionSetID")
    def disk_encryption_set_id(self) -> Optional[str]:
        """
        This is of the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}'
        """
        return pulumi.get(self, "disk_encryption_set_id")

    @property
    @pulumi.getter(name="dnsPrefix")
    def dns_prefix(self) -> Optional[str]:
        """
        This cannot be updated once the Managed Cluster has been created.
        """
        return pulumi.get(self, "dns_prefix")

    @property
    @pulumi.getter(name="enablePodSecurityPolicy")
    def enable_pod_security_policy(self) -> Optional[bool]:
        """
        (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
        """
        return pulumi.get(self, "enable_pod_security_policy")

    @property
    @pulumi.getter(name="enableRBAC")
    def enable_rbac(self) -> Optional[bool]:
        """
        Whether to enable Kubernetes Role-Based Access Control.
        """
        return pulumi.get(self, "enable_rbac")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> Optional['outputs.ExtendedLocationResponse']:
        """
        The extended location of the Virtual Machine.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        The FQDN of the master pool.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="fqdnSubdomain")
    def fqdn_subdomain(self) -> Optional[str]:
        """
        This cannot be updated once the Managed Cluster has been created.
        """
        return pulumi.get(self, "fqdn_subdomain")

    @property
    @pulumi.getter(name="httpProxyConfig")
    def http_proxy_config(self) -> Optional['outputs.ManagedClusterHTTPProxyConfigResponse']:
        """
        Configurations for provisioning the cluster with HTTP proxy servers.
        """
        return pulumi.get(self, "http_proxy_config")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ManagedClusterIdentityResponse']:
        """
        The identity of the managed cluster, if configured.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="identityProfile")
    def identity_profile(self) -> Optional[Mapping[str, 'outputs.ManagedClusterPropertiesResponseIdentityProfile']]:
        """
        Identities associated with the cluster.
        """
        return pulumi.get(self, "identity_profile")

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> Optional[str]:
        """
        When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades must be performed sequentially by major version number. For example, upgrades between 1.14.x -> 1.15.x or 1.15.x -> 1.16.x are allowed, however 1.14.x -> 1.16.x is not allowed. See [upgrading an AKS cluster](https://docs.microsoft.com/azure/aks/upgrade-cluster) for more details.
        """
        return pulumi.get(self, "kubernetes_version")

    @property
    @pulumi.getter(name="linuxProfile")
    def linux_profile(self) -> Optional['outputs.ContainerServiceLinuxProfileResponse']:
        """
        The profile for Linux VMs in the Managed Cluster.
        """
        return pulumi.get(self, "linux_profile")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxAgentPools")
    def max_agent_pools(self) -> int:
        """
        The max number of agent pools for the managed cluster.
        """
        return pulumi.get(self, "max_agent_pools")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkProfile")
    def network_profile(self) -> Optional['outputs.ContainerServiceNetworkProfileResponse']:
        """
        The network configuration profile.
        """
        return pulumi.get(self, "network_profile")

    @property
    @pulumi.getter(name="nodeResourceGroup")
    def node_resource_group(self) -> Optional[str]:
        """
        The name of the resource group containing agent pool nodes.
        """
        return pulumi.get(self, "node_resource_group")

    @property
    @pulumi.getter(name="podIdentityProfile")
    def pod_identity_profile(self) -> Optional['outputs.ManagedClusterPodIdentityProfileResponse']:
        """
        See [use AAD pod identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity) for more details on AAD pod identity integration.
        """
        return pulumi.get(self, "pod_identity_profile")

    @property
    @pulumi.getter(name="powerState")
    def power_state(self) -> 'outputs.PowerStateResponse':
        """
        The Power State of the cluster.
        """
        return pulumi.get(self, "power_state")

    @property
    @pulumi.getter(name="privateFQDN")
    def private_fqdn(self) -> str:
        """
        The FQDN of private cluster.
        """
        return pulumi.get(self, "private_fqdn")

    @property
    @pulumi.getter(name="privateLinkResources")
    def private_link_resources(self) -> Optional[Sequence['outputs.PrivateLinkResourceResponse']]:
        """
        Private link resources associated with the cluster.
        """
        return pulumi.get(self, "private_link_resources")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The current provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="servicePrincipalProfile")
    def service_principal_profile(self) -> Optional['outputs.ManagedClusterServicePrincipalProfileResponse']:
        """
        Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        """
        return pulumi.get(self, "service_principal_profile")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ManagedClusterSKUResponse']:
        """
        The managed cluster SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="windowsProfile")
    def windows_profile(self) -> Optional['outputs.ManagedClusterWindowsProfileResponse']:
        """
        The profile for Windows VMs in the Managed Cluster.
        """
        return pulumi.get(self, "windows_profile")


class AwaitableGetManagedClusterResult(GetManagedClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedClusterResult(
            aad_profile=self.aad_profile,
            addon_profiles=self.addon_profiles,
            agent_pool_profiles=self.agent_pool_profiles,
            api_server_access_profile=self.api_server_access_profile,
            auto_scaler_profile=self.auto_scaler_profile,
            auto_upgrade_profile=self.auto_upgrade_profile,
            azure_portal_fqdn=self.azure_portal_fqdn,
            disable_local_accounts=self.disable_local_accounts,
            disk_encryption_set_id=self.disk_encryption_set_id,
            dns_prefix=self.dns_prefix,
            enable_pod_security_policy=self.enable_pod_security_policy,
            enable_rbac=self.enable_rbac,
            extended_location=self.extended_location,
            fqdn=self.fqdn,
            fqdn_subdomain=self.fqdn_subdomain,
            http_proxy_config=self.http_proxy_config,
            id=self.id,
            identity=self.identity,
            identity_profile=self.identity_profile,
            kubernetes_version=self.kubernetes_version,
            linux_profile=self.linux_profile,
            location=self.location,
            max_agent_pools=self.max_agent_pools,
            name=self.name,
            network_profile=self.network_profile,
            node_resource_group=self.node_resource_group,
            pod_identity_profile=self.pod_identity_profile,
            power_state=self.power_state,
            private_fqdn=self.private_fqdn,
            private_link_resources=self.private_link_resources,
            provisioning_state=self.provisioning_state,
            service_principal_profile=self.service_principal_profile,
            sku=self.sku,
            tags=self.tags,
            type=self.type,
            windows_profile=self.windows_profile)


def get_managed_cluster(resource_group_name: Optional[str] = None,
                        resource_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagedClusterResult:
    """
    Managed cluster.


    :param str resource_group_name: The name of the resource group.
    :param str resource_name: The name of the managed cluster resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:containerservice/v20210501:getManagedCluster', __args__, opts=opts, typ=GetManagedClusterResult).value

    return AwaitableGetManagedClusterResult(
        aad_profile=pulumi.get(__ret__, 'aad_profile'),
        addon_profiles=pulumi.get(__ret__, 'addon_profiles'),
        agent_pool_profiles=pulumi.get(__ret__, 'agent_pool_profiles'),
        api_server_access_profile=pulumi.get(__ret__, 'api_server_access_profile'),
        auto_scaler_profile=pulumi.get(__ret__, 'auto_scaler_profile'),
        auto_upgrade_profile=pulumi.get(__ret__, 'auto_upgrade_profile'),
        azure_portal_fqdn=pulumi.get(__ret__, 'azure_portal_fqdn'),
        disable_local_accounts=pulumi.get(__ret__, 'disable_local_accounts'),
        disk_encryption_set_id=pulumi.get(__ret__, 'disk_encryption_set_id'),
        dns_prefix=pulumi.get(__ret__, 'dns_prefix'),
        enable_pod_security_policy=pulumi.get(__ret__, 'enable_pod_security_policy'),
        enable_rbac=pulumi.get(__ret__, 'enable_rbac'),
        extended_location=pulumi.get(__ret__, 'extended_location'),
        fqdn=pulumi.get(__ret__, 'fqdn'),
        fqdn_subdomain=pulumi.get(__ret__, 'fqdn_subdomain'),
        http_proxy_config=pulumi.get(__ret__, 'http_proxy_config'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        identity_profile=pulumi.get(__ret__, 'identity_profile'),
        kubernetes_version=pulumi.get(__ret__, 'kubernetes_version'),
        linux_profile=pulumi.get(__ret__, 'linux_profile'),
        location=pulumi.get(__ret__, 'location'),
        max_agent_pools=pulumi.get(__ret__, 'max_agent_pools'),
        name=pulumi.get(__ret__, 'name'),
        network_profile=pulumi.get(__ret__, 'network_profile'),
        node_resource_group=pulumi.get(__ret__, 'node_resource_group'),
        pod_identity_profile=pulumi.get(__ret__, 'pod_identity_profile'),
        power_state=pulumi.get(__ret__, 'power_state'),
        private_fqdn=pulumi.get(__ret__, 'private_fqdn'),
        private_link_resources=pulumi.get(__ret__, 'private_link_resources'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        service_principal_profile=pulumi.get(__ret__, 'service_principal_profile'),
        sku=pulumi.get(__ret__, 'sku'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        windows_profile=pulumi.get(__ret__, 'windows_profile'))
def get_managed_cluster_output(resource_group_name: Optional[pulumi.Input[str]] = None,
                               resource_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetManagedClusterResult]:
    """
    Managed cluster.


    :param str resource_group_name: The name of the resource group.
    :param str resource_name: The name of the managed cluster resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:containerservice/v20210501:getManagedCluster', __args__, opts=opts, typ=GetManagedClusterResult)
    return __ret__.apply(lambda __response__: GetManagedClusterResult(
        aad_profile=pulumi.get(__response__, 'aad_profile'),
        addon_profiles=pulumi.get(__response__, 'addon_profiles'),
        agent_pool_profiles=pulumi.get(__response__, 'agent_pool_profiles'),
        api_server_access_profile=pulumi.get(__response__, 'api_server_access_profile'),
        auto_scaler_profile=pulumi.get(__response__, 'auto_scaler_profile'),
        auto_upgrade_profile=pulumi.get(__response__, 'auto_upgrade_profile'),
        azure_portal_fqdn=pulumi.get(__response__, 'azure_portal_fqdn'),
        disable_local_accounts=pulumi.get(__response__, 'disable_local_accounts'),
        disk_encryption_set_id=pulumi.get(__response__, 'disk_encryption_set_id'),
        dns_prefix=pulumi.get(__response__, 'dns_prefix'),
        enable_pod_security_policy=pulumi.get(__response__, 'enable_pod_security_policy'),
        enable_rbac=pulumi.get(__response__, 'enable_rbac'),
        extended_location=pulumi.get(__response__, 'extended_location'),
        fqdn=pulumi.get(__response__, 'fqdn'),
        fqdn_subdomain=pulumi.get(__response__, 'fqdn_subdomain'),
        http_proxy_config=pulumi.get(__response__, 'http_proxy_config'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        identity_profile=pulumi.get(__response__, 'identity_profile'),
        kubernetes_version=pulumi.get(__response__, 'kubernetes_version'),
        linux_profile=pulumi.get(__response__, 'linux_profile'),
        location=pulumi.get(__response__, 'location'),
        max_agent_pools=pulumi.get(__response__, 'max_agent_pools'),
        name=pulumi.get(__response__, 'name'),
        network_profile=pulumi.get(__response__, 'network_profile'),
        node_resource_group=pulumi.get(__response__, 'node_resource_group'),
        pod_identity_profile=pulumi.get(__response__, 'pod_identity_profile'),
        power_state=pulumi.get(__response__, 'power_state'),
        private_fqdn=pulumi.get(__response__, 'private_fqdn'),
        private_link_resources=pulumi.get(__response__, 'private_link_resources'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        service_principal_profile=pulumi.get(__response__, 'service_principal_profile'),
        sku=pulumi.get(__response__, 'sku'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        windows_profile=pulumi.get(__response__, 'windows_profile')))
