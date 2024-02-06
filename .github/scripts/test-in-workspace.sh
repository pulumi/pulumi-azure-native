#!/bin/bash
set -euxo pipefail

test "${1}" = "" && echo "Usage: $0 <TestSet> (see pulumi/examples 'make specific_test_set')" && exit 1
TestSet=${1}

go work init
go work use -r sdk/pulumi-azure-native-sdk
go work use p-examples/misc/test

cd examples
make specific_test_set TestSet=$TestSet
