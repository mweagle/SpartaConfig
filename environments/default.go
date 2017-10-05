// +build !staging,!production

package environments

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws/session"
	gocf "github.com/mweagle/go-cloudformation"
	sparta "github.com/mweagle/Sparta"
)

// Name is the default configuration
const Name = ""

// ServiceDecoratorHook returns a service decorator hook with the supplied
// tags
func ServiceDecoratorHook(buildTags string) sparta.ServiceDecoratorHook {
	return func(context map[string]interface{},
		serviceName string,
		template *gocf.Template,
		S3Bucket string,
		buildID string,
		awsSession *session.Session,
		noop bool,
		logger *logrus.Logger) error {
		if len(buildTags) <= 0 {
			return fmt.Errorf("Please provide a --tags value for environment target")
		}
		return nil
	}
}
