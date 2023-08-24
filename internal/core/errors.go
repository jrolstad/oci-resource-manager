package core

import "github.com/jrolstad/oci-resource-manager/internal/logging"

func ThrowIfError(err error) {
	if err != nil {
		logging.LogPanic(err)
	}
}
