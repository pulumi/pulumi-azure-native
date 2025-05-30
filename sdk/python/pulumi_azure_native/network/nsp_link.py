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

__all__ = ['NspLinkArgs', 'NspLink']

@pulumi.input_type
class NspLinkArgs:
    def __init__(__self__, *,
                 network_security_perimeter_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 auto_approved_remote_perimeter_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 link_name: Optional[pulumi.Input[builtins.str]] = None,
                 local_inbound_profiles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 remote_inbound_profiles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a NspLink resource.
        :param pulumi.Input[builtins.str] network_security_perimeter_name: The name of the network security perimeter.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input[builtins.str] auto_approved_remote_perimeter_resource_id: Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
        :param pulumi.Input[builtins.str] description: A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
        :param pulumi.Input[builtins.str] link_name: The name of the NSP link.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] local_inbound_profiles: Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] remote_inbound_profiles: Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode.
        """
        pulumi.set(__self__, "network_security_perimeter_name", network_security_perimeter_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if auto_approved_remote_perimeter_resource_id is not None:
            pulumi.set(__self__, "auto_approved_remote_perimeter_resource_id", auto_approved_remote_perimeter_resource_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if link_name is not None:
            pulumi.set(__self__, "link_name", link_name)
        if local_inbound_profiles is not None:
            pulumi.set(__self__, "local_inbound_profiles", local_inbound_profiles)
        if remote_inbound_profiles is not None:
            pulumi.set(__self__, "remote_inbound_profiles", remote_inbound_profiles)

    @property
    @pulumi.getter(name="networkSecurityPerimeterName")
    def network_security_perimeter_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the network security perimeter.
        """
        return pulumi.get(self, "network_security_perimeter_name")

    @network_security_perimeter_name.setter
    def network_security_perimeter_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "network_security_perimeter_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="autoApprovedRemotePerimeterResourceId")
    def auto_approved_remote_perimeter_resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
        """
        return pulumi.get(self, "auto_approved_remote_perimeter_resource_id")

    @auto_approved_remote_perimeter_resource_id.setter
    def auto_approved_remote_perimeter_resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auto_approved_remote_perimeter_resource_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="linkName")
    def link_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the NSP link.
        """
        return pulumi.get(self, "link_name")

    @link_name.setter
    def link_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "link_name", value)

    @property
    @pulumi.getter(name="localInboundProfiles")
    def local_inbound_profiles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles.
        """
        return pulumi.get(self, "local_inbound_profiles")

    @local_inbound_profiles.setter
    def local_inbound_profiles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "local_inbound_profiles", value)

    @property
    @pulumi.getter(name="remoteInboundProfiles")
    def remote_inbound_profiles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode.
        """
        return pulumi.get(self, "remote_inbound_profiles")

    @remote_inbound_profiles.setter
    def remote_inbound_profiles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "remote_inbound_profiles", value)


@pulumi.type_token("azure-native:network:NspLink")
class NspLink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_approved_remote_perimeter_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 link_name: Optional[pulumi.Input[builtins.str]] = None,
                 local_inbound_profiles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 network_security_perimeter_name: Optional[pulumi.Input[builtins.str]] = None,
                 remote_inbound_profiles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The network security perimeter link resource

        Uses Azure REST API version 2023-08-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-02-01-preview.

        Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] auto_approved_remote_perimeter_resource_id: Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
        :param pulumi.Input[builtins.str] description: A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
        :param pulumi.Input[builtins.str] link_name: The name of the NSP link.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] local_inbound_profiles: Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles.
        :param pulumi.Input[builtins.str] network_security_perimeter_name: The name of the network security perimeter.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] remote_inbound_profiles: Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NspLinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The network security perimeter link resource

        Uses Azure REST API version 2023-08-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-02-01-preview.

        Other available API versions: 2021-02-01-preview, 2023-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param NspLinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NspLinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_approved_remote_perimeter_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 link_name: Optional[pulumi.Input[builtins.str]] = None,
                 local_inbound_profiles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 network_security_perimeter_name: Optional[pulumi.Input[builtins.str]] = None,
                 remote_inbound_profiles: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NspLinkArgs.__new__(NspLinkArgs)

            __props__.__dict__["auto_approved_remote_perimeter_resource_id"] = auto_approved_remote_perimeter_resource_id
            __props__.__dict__["description"] = description
            __props__.__dict__["link_name"] = link_name
            __props__.__dict__["local_inbound_profiles"] = local_inbound_profiles
            if network_security_perimeter_name is None and not opts.urn:
                raise TypeError("Missing required property 'network_security_perimeter_name'")
            __props__.__dict__["network_security_perimeter_name"] = network_security_perimeter_name
            __props__.__dict__["remote_inbound_profiles"] = remote_inbound_profiles
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["local_outbound_profiles"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["remote_outbound_profiles"] = None
            __props__.__dict__["remote_perimeter_guid"] = None
            __props__.__dict__["remote_perimeter_location"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:network/v20210201preview:NspLink"), pulumi.Alias(type_="azure-native:network/v20230701preview:NspLink"), pulumi.Alias(type_="azure-native:network/v20230801preview:NspLink"), pulumi.Alias(type_="azure-native:network/v20240601preview:NspLink"), pulumi.Alias(type_="azure-native:network/v20240701:NspLink")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NspLink, __self__).__init__(
            'azure-native:network:NspLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NspLink':
        """
        Get an existing NspLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NspLinkArgs.__new__(NspLinkArgs)

        __props__.__dict__["auto_approved_remote_perimeter_resource_id"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["local_inbound_profiles"] = None
        __props__.__dict__["local_outbound_profiles"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["remote_inbound_profiles"] = None
        __props__.__dict__["remote_outbound_profiles"] = None
        __props__.__dict__["remote_perimeter_guid"] = None
        __props__.__dict__["remote_perimeter_location"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["type"] = None
        return NspLink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoApprovedRemotePerimeterResourceId")
    def auto_approved_remote_perimeter_resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
        """
        return pulumi.get(self, "auto_approved_remote_perimeter_resource_id")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[builtins.str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="localInboundProfiles")
    def local_inbound_profiles(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles.
        """
        return pulumi.get(self, "local_inbound_profiles")

    @property
    @pulumi.getter(name="localOutboundProfiles")
    def local_outbound_profiles(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Local Outbound profile names from which Outbound is allowed. In current version, it is readonly property and it's value is set to ['*'] to allow outbound from all profiles. In later version, user will be able to modify it.
        """
        return pulumi.get(self, "local_outbound_profiles")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state of the NSP Link resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="remoteInboundProfiles")
    def remote_inbound_profiles(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode.
        """
        return pulumi.get(self, "remote_inbound_profiles")

    @property
    @pulumi.getter(name="remoteOutboundProfiles")
    def remote_outbound_profiles(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Remote Outbound profile names from which Outbound is allowed. In current version, it is readonly property and it's value is set to ['*'] to allow outbound from all profiles. In later version, user will be able to modify it.
        """
        return pulumi.get(self, "remote_outbound_profiles")

    @property
    @pulumi.getter(name="remotePerimeterGuid")
    def remote_perimeter_guid(self) -> pulumi.Output[builtins.str]:
        """
        Remote NSP Guid with which the link gets created.
        """
        return pulumi.get(self, "remote_perimeter_guid")

    @property
    @pulumi.getter(name="remotePerimeterLocation")
    def remote_perimeter_location(self) -> pulumi.Output[builtins.str]:
        """
        Remote NSP location with which the link gets created.
        """
        return pulumi.get(self, "remote_perimeter_location")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        The NSP link state.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

