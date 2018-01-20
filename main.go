package main

import (
	"context"
	"fmt"
	"os"

	sparta "github.com/mweagle/Sparta"
	environments "github.com/mweagle/SpartaConfig/environments"
	"github.com/mweagle/SpartaConfig/environments/targets"
)

// Standard AWS λ function
func helloWorld(ctx context.Context) (string, error) {
	return fmt.Sprintf("Hello %s", targets.Name), nil
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
	lambdaFn := sparta.HandleAWSLambda("spartaConfig",
		helloWorld,
		sparta.IAMRoleDefinition{})
	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)

	workflowHooks := &sparta.WorkflowHooks{
		ServiceDecorators: []sparta.ServiceDecoratorHookHandler{
			environments.ServiceDecoratorHook(),
		},
	}

	err := sparta.MainEx("SpartaConfig",
		fmt.Sprintf("Test SpartaConfig environments"),
		lambdaFunctions,
		nil,
		nil,
		workflowHooks,
		false)
	if err != nil {
		os.Exit(1)
	}
}
