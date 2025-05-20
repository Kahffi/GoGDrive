package main

import (
	"context"

	"github.com/Kahffi/GoGDrive/config"
	"github.com/Kahffi/GoGDrive/gogdrive"
)

func main() {
	ctx := context.Background()

	config := &config.DriveClientConfig{
		CredentialsPath: "./credentials.json",
	}

	dc, err := gogdrive.NewDriveClient(ctx, config)
	if err != nil {
		panic(err)
	}

	dc.List()
}
