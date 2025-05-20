package gogdrive

import (
	"context"

	"github.com/Kahffi/GoGDrive/config"
	"github.com/Kahffi/GoGDrive/internal/auth"
	"google.golang.org/api/drive/v3"
)

type driveClient struct {
	cfg *config.DriveClientConfig
	ds  *drive.Service
}

func NewDriveClient(ctx context.Context, cfg *config.DriveClientConfig) (*driveClient, error) {

	ds, err := auth.Authenticate(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &driveClient{
		cfg: cfg,
		ds:  ds,
	}, nil
}
