package wallet

import (
	"context"
	"path/filepath"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/k0yote/backend-wallet/config"
	"github.com/k0yote/backend-wallet/util"
	"google.golang.org/api/option"
)

func entropy(config config.Config, len int32) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var (
		path string
		err  error
	)

	if path, err = util.GetRootPath(); err != nil {
		return nil, err
	}

	client, err := kms.NewKeyManagementClient(ctx, option.WithCredentialsFile(filepath.Join(path, config.GoogleCredentialPath)))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.GenerateRandomBytesRequest{
		Location:        config.GoogleKmsLocation,
		LengthBytes:     len,
		ProtectionLevel: kmspb.ProtectionLevel_HSM,
	}

	result, err := client.GenerateRandomBytes(ctx, req)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
