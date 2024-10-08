// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package cli

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/ultravioletrs/cocos/agent"
	"github.com/ultravioletrs/cocos/internal"
	"google.golang.org/grpc/metadata"
)

var decompressDataset bool

func (cli *CLI) NewDatasetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "data",
		Short:   "Upload a dataset",
		Example: "data <dataset_path> <private_key_file_path>",
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			datasetPath := args[0]

			log.Println("Uploading dataset:", datasetPath)

			f, err := os.Stat(datasetPath)
			if err != nil {
				log.Fatalf("Error reading dataset file: %v", err)
			}

			var dataset []byte

			if f.IsDir() {
				dataset, err = internal.ZipDirectoryToMemory(datasetPath)
				if err != nil {
					log.Fatalf("Error zipping dataset directory: %v", err)
				}
			} else {
				dataset, err = os.ReadFile(datasetPath)
				if err != nil {
					log.Fatalf("Error reading dataset file: %v", err)
				}
			}

			dataReq := agent.Dataset{
				Dataset:  dataset,
				Filename: path.Base(datasetPath),
			}

			privKeyFile, err := os.ReadFile(args[1])
			if err != nil {
				log.Fatalf("Error reading private key file: %v", err)
			}

			pemBlock, _ := pem.Decode(privKeyFile)

			privKey := decodeKey(pemBlock)

			ctx := metadata.NewOutgoingContext(cmd.Context(), metadata.New(make(map[string]string)))
			if err := cli.agentSDK.Data(addDatasetMetadata(ctx), dataReq, privKey); err != nil {
				log.Fatalf("Error uploading dataset: %v", err)
			}

			log.Println("Successfully uploaded dataset")
		},
	}

	cmd.Flags().BoolVarP(&decompressDataset, "decompress", "d", false, "Decompress the dataset on agent")
	return cmd
}

func decodeKey(b *pem.Block) interface{} {
	switch b.Type {
	case rsaKeyType:
		privKey, err := x509.ParsePKCS8PrivateKey(b.Bytes)
		if err != nil {
			privKey, err = x509.ParsePKCS1PrivateKey(b.Bytes)
			if err != nil {
				log.Fatalf("Error parsing private key: %v", err)
			}
		}
		return privKey
	case ecdsaKeyType:
		privKey, err := x509.ParseECPrivateKey(b.Bytes)
		if err != nil {
			log.Fatalf("Error parsing private key: %v", err)
		}
		return privKey
	default:
		log.Fatalf("Error decoding key")
		return nil
	}
}

func addDatasetMetadata(ctx context.Context) context.Context {
	return agent.DecompressToContext(ctx, decompressDataset)
}
