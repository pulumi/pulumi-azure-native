# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = ['LoadBalancerArgs', 'LoadBalancer']

@pulumi.input_type
class LoadBalancerArgs:
    def __init__(__self__, *,
                 addresses: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 advertise_mode: pulumi.Input[Union[builtins.str, 'AdvertiseMode']],
                 resource_uri: pulumi.Input[builtins.str],
                 bgp_peers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 load_balancer_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_selector: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a LoadBalancer resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] addresses: IP Range
        :param pulumi.Input[Union[builtins.str, 'AdvertiseMode']] advertise_mode: Advertise Mode
        :param pulumi.Input[builtins.str] resource_uri: The fully qualified Azure Resource manager identifier of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] bgp_peers: The list of BGP peers it should advertise to. Null or empty means to advertise to all peers.
        :param pulumi.Input[builtins.str] load_balancer_name: The name of the LoadBalancer
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] service_selector: A dynamic label mapping to select related services. For instance, if you want to create a load balancer only for services with label "a=b", then please specify {"a": "b"} in the field.
        """
        pulumi.set(__self__, "addresses", addresses)
        pulumi.set(__self__, "advertise_mode", advertise_mode)
        pulumi.set(__self__, "resource_uri", resource_uri)
        if bgp_peers is not None:
            pulumi.set(__self__, "bgp_peers", bgp_peers)
        if load_balancer_name is not None:
            pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        if service_selector is not None:
            pulumi.set(__self__, "service_selector", service_selector)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        IP Range
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="advertiseMode")
    def advertise_mode(self) -> pulumi.Input[Union[builtins.str, 'AdvertiseMode']]:
        """
        Advertise Mode
        """
        return pulumi.get(self, "advertise_mode")

    @advertise_mode.setter
    def advertise_mode(self, value: pulumi.Input[Union[builtins.str, 'AdvertiseMode']]):
        pulumi.set(self, "advertise_mode", value)

    @property
    @pulumi.getter(name="resourceUri")
    def resource_uri(self) -> pulumi.Input[builtins.str]:
        """
        The fully qualified Azure Resource manager identifier of the resource.
        """
        return pulumi.get(self, "resource_uri")

    @resource_uri.setter
    def resource_uri(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_uri", value)

    @property
    @pulumi.getter(name="bgpPeers")
    def bgp_peers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The list of BGP peers it should advertise to. Null or empty means to advertise to all peers.
        """
        return pulumi.get(self, "bgp_peers")

    @bgp_peers.setter
    def bgp_peers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "bgp_peers", value)

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the LoadBalancer
        """
        return pulumi.get(self, "load_balancer_name")

    @load_balancer_name.setter
    def load_balancer_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "load_balancer_name", value)

    @property
    @pulumi.getter(name="serviceSelector")
    def service_selector(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A dynamic label mapping to select related services. For instance, if you want to create a load balancer only for services with label "a=b", then please specify {"a": "b"} in the field.
        """
        return pulumi.get(self, "service_selector")

    @service_selector.setter
    def service_selector(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "service_selector", value)


@pulumi.type_token("azure-native:kubernetesruntime:LoadBalancer")
class LoadBalancer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 advertise_mode: Optional[pulumi.Input[Union[builtins.str, 'AdvertiseMode']]] = None,
                 bgp_peers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 load_balancer_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_uri: Optional[pulumi.Input[builtins.str]] = None,
                 service_selector: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        A LoadBalancer resource for an Arc connected cluster (Microsoft.Kubernetes/connectedClusters)

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2024-03-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] addresses: IP Range
        :param pulumi.Input[Union[builtins.str, 'AdvertiseMode']] advertise_mode: Advertise Mode
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] bgp_peers: The list of BGP peers it should advertise to. Null or empty means to advertise to all peers.
        :param pulumi.Input[builtins.str] load_balancer_name: The name of the LoadBalancer
        :param pulumi.Input[builtins.str] resource_uri: The fully qualified Azure Resource manager identifier of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] service_selector: A dynamic label mapping to select related services. For instance, if you want to create a load balancer only for services with label "a=b", then please specify {"a": "b"} in the field.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoadBalancerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A LoadBalancer resource for an Arc connected cluster (Microsoft.Kubernetes/connectedClusters)

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2024-03-01.

        :param str resource_name: The name of the resource.
        :param LoadBalancerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoadBalancerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 advertise_mode: Optional[pulumi.Input[Union[builtins.str, 'AdvertiseMode']]] = None,
                 bgp_peers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 load_balancer_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_uri: Optional[pulumi.Input[builtins.str]] = None,
                 service_selector: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoadBalancerArgs.__new__(LoadBalancerArgs)

            if addresses is None and not opts.urn:
                raise TypeError("Missing required property 'addresses'")
            __props__.__dict__["addresses"] = addresses
            if advertise_mode is None and not opts.urn:
                raise TypeError("Missing required property 'advertise_mode'")
            __props__.__dict__["advertise_mode"] = advertise_mode
            __props__.__dict__["bgp_peers"] = bgp_peers
            __props__.__dict__["load_balancer_name"] = load_balancer_name
            if resource_uri is None and not opts.urn:
                raise TypeError("Missing required property 'resource_uri'")
            __props__.__dict__["resource_uri"] = resource_uri
            __props__.__dict__["service_selector"] = service_selector
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:kubernetesruntime/v20231001preview:LoadBalancer"), pulumi.Alias(type_="azure-native:kubernetesruntime/v20240301:LoadBalancer")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(LoadBalancer, __self__).__init__(
            'azure-native:kubernetesruntime:LoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LoadBalancer':
        """
        Get an existing LoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LoadBalancerArgs.__new__(LoadBalancerArgs)

        __props__.__dict__["addresses"] = None
        __props__.__dict__["advertise_mode"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["bgp_peers"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["service_selector"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return LoadBalancer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        IP Range
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="advertiseMode")
    def advertise_mode(self) -> pulumi.Output[builtins.str]:
        """
        Advertise Mode
        """
        return pulumi.get(self, "advertise_mode")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="bgpPeers")
    def bgp_peers(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The list of BGP peers it should advertise to. Null or empty means to advertise to all peers.
        """
        return pulumi.get(self, "bgp_peers")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Resource provision state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceSelector")
    def service_selector(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        A dynamic label mapping to select related services. For instance, if you want to create a load balancer only for services with label "a=b", then please specify {"a": "b"} in the field.
        """
        return pulumi.get(self, "service_selector")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

