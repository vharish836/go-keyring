package keyring

import "fmt"

// provider set in the init function by the relevant os file e.g.:
// keyring_linux.go
var provider Keyring = fallbackServiceProvider{}

var (
	// ErrNotFound is the expected error if the secret isn't found in the
	// keyring.
	ErrNotFound = fmt.Errorf("secret not found in keyring")
)

// Keyring provides a simple set/get interface for a keyring service.
type Keyring interface {
	// Set password in keyring for user.
	Set(service, account, password string) error
	// Get password from keyring given service and user name.
	Get(service, account string) (string, error)
	// Delete secret from keyring.
	Delete(service, account string) error
}

// Set password in keyring for user.
func Set(service, account, password string) error {
	return provider.Set(service, account, password)
}

// Get password from keyring given service and user name.
func Get(service, account string) (string, error) {
	return provider.Get(service, account)
}

// Delete secret from keyring.
func Delete(service, account string) error {
	return provider.Delete(service, account)
}
