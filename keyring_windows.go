package keyring

import (
	"syscall"

	"github.com/danieljoos/wincred"
)

type windowsKeychain struct{}

// Get gets a secret from the keyring given a service name and account.
func (k windowsKeychain) Get(service, account string) (string, error) {
	cred, err := wincred.GetGenericCredential(k.credName(service, account))
	if err != nil {
		if err == syscall.ERROR_NOT_FOUND {
			return "", ErrNotFound
		}
		return "", err
	}

	return string(cred.CredentialBlob), nil
}

// Set stores stores account and pass in the keyring under the defined service
// name.
func (k windowsKeychain) Set(service, account, password string) error {
	cred := wincred.NewGenericCredential(k.credName(service, account))
	cred.UserName = account
	cred.CredentialBlob = []byte(password)
	return cred.Write()
}

// Delete deletes a secret, identified by service & account, from the keyring.
func (k windowsKeychain) Delete(service, account string) error {
	cred, err := wincred.GetGenericCredential(k.credName(service, account))
	if err != nil {
		if err == syscall.ERROR_NOT_FOUND {
			return ErrNotFound
		}
		return err
	}

	return cred.Delete()
}

// credName combines service and account to a single string.
func (k windowsKeychain) credName(service, account string) string {
	return service + "/" + account
}

func init() {
	provider = windowsKeychain{}
}
