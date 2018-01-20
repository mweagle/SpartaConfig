package environments

import (
	"github.com/aws/aws-sdk-go/aws/session"
	sparta "github.com/mweagle/Sparta"
	"github.com/mweagle/SpartaConfig/environments/targets"
	gocf "github.com/mweagle/go-cloudformation"
	"github.com/sirupsen/logrus"
)

// ServiceDecoratorHook returns a service decorator hook with the supplied
// tags
func ServiceDecoratorHook() sparta.ServiceDecoratorHookFunc {
	return func(context map[string]interface{},
		serviceName string,
		template *gocf.Template,
		S3Bucket string,
		buildID string,
		awsSession *session.Session,
		noop bool,
		logger *logrus.Logger) error {
		template.Outputs["Environment"] = &gocf.Output{
			Description: "Sparta Config target environment",
			Value:       targets.Name,
		}
		return nil
	}
}
