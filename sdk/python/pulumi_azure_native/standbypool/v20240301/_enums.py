# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'RefillPolicy',
    'VirtualMachineState',
]


class RefillPolicy(str, Enum):
    """
    Specifies refill policy of the pool.
    """
    ALWAYS = "always"
    """
    A refill policy that standby pool is automatically refilled to maintain maxReadyCapacity.
    """


class VirtualMachineState(str, Enum):
    """
    Specifies the desired state of virtual machines in the pool.
    """
    RUNNING = "Running"
    """
    The virtual machine is up and running.
    """
    DEALLOCATED = "Deallocated"
    """
    The virtual machine has released the lease on the underlying hardware and is powered off.
    """
