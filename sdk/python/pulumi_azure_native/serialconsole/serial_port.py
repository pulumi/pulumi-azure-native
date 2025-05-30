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
from ._enums import *

__all__ = ['SerialPortArgs', 'SerialPort']

@pulumi.input_type
class SerialPortArgs:
    def __init__(__self__, *,
                 parent_resource: pulumi.Input[builtins.str],
                 parent_resource_type: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 resource_provider_namespace: pulumi.Input[builtins.str],
                 serial_port: Optional[pulumi.Input[builtins.str]] = None,
                 state: Optional[pulumi.Input['SerialPortState']] = None):
        """
        The set of arguments for constructing a SerialPort resource.
        :param pulumi.Input[builtins.str] parent_resource: The resource name, or subordinate path, for the parent of the serial port. For example: the name of the virtual machine.
        :param pulumi.Input[builtins.str] parent_resource_type: The resource type of the parent resource.  For example: 'virtualMachines' or 'virtualMachineScaleSets'
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input[builtins.str] resource_provider_namespace: The namespace of the resource provider.
        :param pulumi.Input[builtins.str] serial_port: The name of the serial port to create.
        :param pulumi.Input['SerialPortState'] state: Specifies whether the port is enabled for a serial console connection.
        """
        pulumi.set(__self__, "parent_resource", parent_resource)
        pulumi.set(__self__, "parent_resource_type", parent_resource_type)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "resource_provider_namespace", resource_provider_namespace)
        if serial_port is not None:
            pulumi.set(__self__, "serial_port", serial_port)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="parentResource")
    def parent_resource(self) -> pulumi.Input[builtins.str]:
        """
        The resource name, or subordinate path, for the parent of the serial port. For example: the name of the virtual machine.
        """
        return pulumi.get(self, "parent_resource")

    @parent_resource.setter
    def parent_resource(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "parent_resource", value)

    @property
    @pulumi.getter(name="parentResourceType")
    def parent_resource_type(self) -> pulumi.Input[builtins.str]:
        """
        The resource type of the parent resource.  For example: 'virtualMachines' or 'virtualMachineScaleSets'
        """
        return pulumi.get(self, "parent_resource_type")

    @parent_resource_type.setter
    def parent_resource_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "parent_resource_type", value)

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
    @pulumi.getter(name="resourceProviderNamespace")
    def resource_provider_namespace(self) -> pulumi.Input[builtins.str]:
        """
        The namespace of the resource provider.
        """
        return pulumi.get(self, "resource_provider_namespace")

    @resource_provider_namespace.setter
    def resource_provider_namespace(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_provider_namespace", value)

    @property
    @pulumi.getter(name="serialPort")
    def serial_port(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the serial port to create.
        """
        return pulumi.get(self, "serial_port")

    @serial_port.setter
    def serial_port(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "serial_port", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input['SerialPortState']]:
        """
        Specifies whether the port is enabled for a serial console connection.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input['SerialPortState']]):
        pulumi.set(self, "state", value)


@pulumi.type_token("azure-native:serialconsole:SerialPort")
class SerialPort(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent_resource: Optional[pulumi.Input[builtins.str]] = None,
                 parent_resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_provider_namespace: Optional[pulumi.Input[builtins.str]] = None,
                 serial_port: Optional[pulumi.Input[builtins.str]] = None,
                 state: Optional[pulumi.Input['SerialPortState']] = None,
                 __props__=None):
        """
        Represents the serial port of the parent resource.

        Uses Azure REST API version 2018-05-01. In version 2.x of the Azure Native provider, it used API version 2018-05-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] parent_resource: The resource name, or subordinate path, for the parent of the serial port. For example: the name of the virtual machine.
        :param pulumi.Input[builtins.str] parent_resource_type: The resource type of the parent resource.  For example: 'virtualMachines' or 'virtualMachineScaleSets'
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input[builtins.str] resource_provider_namespace: The namespace of the resource provider.
        :param pulumi.Input[builtins.str] serial_port: The name of the serial port to create.
        :param pulumi.Input['SerialPortState'] state: Specifies whether the port is enabled for a serial console connection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SerialPortArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents the serial port of the parent resource.

        Uses Azure REST API version 2018-05-01. In version 2.x of the Azure Native provider, it used API version 2018-05-01.

        :param str resource_name: The name of the resource.
        :param SerialPortArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SerialPortArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent_resource: Optional[pulumi.Input[builtins.str]] = None,
                 parent_resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_provider_namespace: Optional[pulumi.Input[builtins.str]] = None,
                 serial_port: Optional[pulumi.Input[builtins.str]] = None,
                 state: Optional[pulumi.Input['SerialPortState']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SerialPortArgs.__new__(SerialPortArgs)

            if parent_resource is None and not opts.urn:
                raise TypeError("Missing required property 'parent_resource'")
            __props__.__dict__["parent_resource"] = parent_resource
            if parent_resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'parent_resource_type'")
            __props__.__dict__["parent_resource_type"] = parent_resource_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if resource_provider_namespace is None and not opts.urn:
                raise TypeError("Missing required property 'resource_provider_namespace'")
            __props__.__dict__["resource_provider_namespace"] = resource_provider_namespace
            __props__.__dict__["serial_port"] = serial_port
            __props__.__dict__["state"] = state
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:serialconsole/v20180501:SerialPort")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SerialPort, __self__).__init__(
            'azure-native:serialconsole:SerialPort',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SerialPort':
        """
        Get an existing SerialPort resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SerialPortArgs.__new__(SerialPortArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        return SerialPort(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies whether the port is enabled for a serial console connection.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

