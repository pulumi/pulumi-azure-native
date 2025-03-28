# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .fabric_capacity import *
from .get_fabric_capacity import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.fabric.v20231101 as __v20231101
    v20231101 = __v20231101
    import pulumi_azure_native.fabric.v20250115preview as __v20250115preview
    v20250115preview = __v20250115preview
else:
    v20231101 = _utilities.lazy_import('pulumi_azure_native.fabric.v20231101')
    v20250115preview = _utilities.lazy_import('pulumi_azure_native.fabric.v20250115preview')

