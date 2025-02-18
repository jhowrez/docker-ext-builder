package exporter

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/moby/moby/pkg/jsonmessage"
	"github.com/sirupsen/logrus"
)

func buildImage(cli *client.Client, dockerfile string) string {
	ctx := context.Background()
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFile := dockerfile
	dockerFileReader, err := os.Open(dockerFile)
	if err != nil {
		log.Fatal(err, " :unable to open Dockerfile")
	}
	readDockerFile, err := io.ReadAll(dockerFileReader)
	if err != nil {
		log.Fatal(err, " :unable to read dockerfile")
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context:    dockerFileTarReader,
			Dockerfile: dockerFile,
			Remove:     true})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()

	outID, err := parseDockerDaemonJsonMessages(imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}

	return outID
}

func parseDockerDaemonJsonMessages(r io.Reader) (string, error) {
	var result string
	decoder := json.NewDecoder(r)
	for {
		var jsonMessage jsonmessage.JSONMessage
		if err := decoder.Decode(&jsonMessage); err != nil {
			if err == io.EOF {
				break
			}
			return result, err
		}
		if err := jsonMessage.Error; err != nil {
			return result, err
		}
		if jsonMessage.Aux != nil {
			var r types.BuildResult
			if err := json.Unmarshal(*jsonMessage.Aux, &r); err != nil {
				logrus.Warnf("Failed to unmarshal aux message. Cause: %s", err)
			} else {
				result = r.ID
			}
		}
	}
	return result, nil
}
