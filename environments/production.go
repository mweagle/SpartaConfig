// +build production

package environments

import (
	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws/session"
	gocf "github.com/mweagle/go-cloudformation"
	sparta "github.com/mweagle/Sparta"
)

// Name is the production configuration
const Name = "production"

func ServiceDecoratorHook(buildTags string) sparta.ServiceDecoratorHook {
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
			Value:       Name,
		}
		return nil
	}
}
