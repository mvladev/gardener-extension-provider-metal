package shoot

import (
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	"github.com/gardener/gardener/extensions/pkg/webhook/shoot"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	// DefaultAddOptions are the default AddOptions for AddToManager.
	DefaultAddOptions = AddOptions{}
)

// AddOptions are options to apply when adding the Packet shoot webhook to the manager.
type AddOptions struct{}

var logger = log.Log.WithName("metal-shoot-webhook")

// AddToManagerWithOptions creates a webhook with the given options and adds it to the manager.
func AddToManagerWithOptions(mgr manager.Manager, opts AddOptions) (*extensionswebhook.Webhook, error) {
	logger.Info("Adding webhook to manager")
	return shoot.New(mgr, shoot.Args{
		Types:   []runtime.Object{},
		Mutator: NewMutator(),
	})
}

// AddToManager creates a webhook with the default options and adds it to the manager.
func AddToManager(mgr manager.Manager) (*extensionswebhook.Webhook, error) {
	return AddToManagerWithOptions(mgr, DefaultAddOptions)
}
