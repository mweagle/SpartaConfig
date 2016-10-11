package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
	environments "github.com/mweagle/SpartaConfig/environments"
)

// Standard AWS λ function
func helloWorld(event *json.RawMessage,
	context *sparta.LambdaContext,
	w http.ResponseWriter,
	logger *logrus.Logger) {

	fmt.Fprintf(w, "Hello %s", environments.Name)
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
	parseErrors := sparta.ParseOptions(nil)
	if nil != parseErrors {
		os.Exit(2)
	}

	lambdaFn := sparta.NewLambda(sparta.IAMRoleDefinition{},
		helloWorld,
		nil)
	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)

	hooks := &sparta.WorkflowHooks{
		Context:          map[string]interface{}{},
		ServiceDecorator: environments.ServiceDecoratorHook(sparta.OptionsGlobal.BuildTags),
	}

	err := sparta.MainEx(fmt.Sprintf("SpartaHelloWorld-%s", sparta.OptionsGlobal.BuildTags),
		fmt.Sprintf("Test HelloWorld resource command"),
		lambdaFunctions,
		nil,
		nil,
		hooks)
	if err != nil {
		os.Exit(1)
	}
}
