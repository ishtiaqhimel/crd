# Custom Resource Definition
**Note:** CustomResourceDefinition is the successor of the deprecated ThirdPartyResource.

This particular example demonstrates how to generate a client for CustomResources using `k8s.io/code-generator`. The clientset can be generated using the `./hack/update-codegen.sh` script.

The `update-codegen` script will automatically generate the following files and directories:
- `pkg/apis/cr/v1/zz_generated.deepcopy.go`
- `pkg/client/`


## Goals
- Understand CRs
- Define a CR
- Create a controller for it

