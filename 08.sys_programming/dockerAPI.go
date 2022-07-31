package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func listConteners() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return (err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return (err)
	}

	for _, container := range containers {
		fmt.Println("Images:", container.Image, "withID: ", container.ID)
	}
	return nil
}

type Container struct {
	ID         string `json:"id"`
	Names      []string
	Image      string
	ImageID    string
	Command    string
	Created    int64
	Ports      []Port
	SizeRw     int64 `json:",omitempty"`
	SizeRootFs int64 `json:",omitempty"`
	Labels     map[string]string
	State      string
	Hostconfig struct {
		NetworkMode string `json:",omitempty"`
	}
	NetworkSettings *SummaryNetworkSettings
	Mounts          []MountPoint
}

func listImages() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return (err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return (err)
	}

	for _, image := range images {
		fmt.Printf("Images %s with size %d\n", image.RepoTags, image.Size)
	}
	return nil
}

func main() {
	fmt.Println("The available images are : ")
	err := listImages()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The runnig Containers are : ")
	err = listConteners()
	if err != nil {
		fmt.Println(err)
	}

}
