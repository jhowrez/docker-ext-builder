package exporter

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/client"
)

func exportContent(cli *client.Client, contID string, exportPath string, outputFilename string) error {
	ctx := context.Background()
	r, _, err := cli.CopyFromContainer(ctx, contID, exportPath)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, r)
	if err != nil {
		return err
	}

	return nil
}
