// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/arm2pulumi"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	goprovider "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
	"reflect"
)

func (k *azureNativeProvider) construct(ctx *pulumi.Context, typ, name string, inputs goprovider.ConstructInputs,
	options pulumi.ResourceOption) (*goprovider.ConstructResult, error) {
	switch typ {
	case "azure-native:extensions:ArmTemplate":
		return k.constructArmTemplate(ctx, name, inputs, options)
	default:
		return nil, errors.Errorf("unknown resource type %s", typ)
	}
}

// constructArmTemplate is an implementation of Construct for the example ArmTemplate component.
// It demonstrates converting the raw ConstructInputs to the component's args struct, creating
// the component, and returning its URN and state (outputs).
func (k *azureNativeProvider) constructArmTemplate(ctx *pulumi.Context, name string, inputs goprovider.ConstructInputs,
	options pulumi.ResourceOption) (*goprovider.ConstructResult, error) {
	args := &ArmTemplateArgs{}
	if err := inputs.CopyTo(args); err != nil {
		return nil, errors.Wrap(err, "setting args")
	}

	uncompressed, err := gzip.NewReader(bytes.NewReader(k.schemaBytes))
	if err != nil {
		return nil, errors.Wrap(err, "expand compressed schema")
	}

	var pkgSpec schema.PackageSpec
	if err = json.NewDecoder(uncompressed).Decode(&pkgSpec); err != nil {
		return nil, fmt.Errorf("deserializing schema: %w", err)
	}

	component := &ArmTemplate{}
	err = ctx.RegisterComponentResource("azure-native:extensions:ArmTemplate", name, component, options)
	if err != nil {
		return nil, err
	}

	component.Result = pulumi.All(args.ResourceGroupName, args.Content).ApplyT(
		func(args []interface{}) (string, error) {
			resourceGroupName := args[0].(string)
			content:= args[1].(string)
			res, err := arm2pulumi.DecodeTemplate(&pkgSpec, k.resourceMap, content)
			if err != nil {
				return "", err
			}

			for _, v := range res {
				tok := v["token"].(string)
				name := v["name"].(string)
				args := v["args"].(map[string]interface{})
				if resourceGroupName != "" {
					args["resourceGroupName"] = resourceGroupName
				}

				var res Foo
				err := ctx.RegisterResource(tok, name, UntypedArgs(args), &res, pulumi.Parent(component))
				if err != nil {
					return "", err
				}
			}

			return fmt.Sprintf("%+v", res), nil
		},
	).(pulumi.StringOutput)

	// Return the component resource's URN and state. `NewConstructResult` automatically sets the
	// ConstructResult's state based on resource struct fields tagged with `pulumi:` tags with a value
	// that is convertible to `pulumi.Input`.
	return goprovider.NewConstructResult(component)
}

type ArmTemplateArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Content pulumi.StringInput `pulumi:"content"`
}

type ArmTemplate struct {
	pulumi.ResourceState
	Result pulumi.StringOutput `pulumi:"result"`
}

type Foo struct {
	pulumi.CustomResourceState
}

type UntypedArgs map[string]interface{}

func (UntypedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]interface{})(nil)).Elem()
}
