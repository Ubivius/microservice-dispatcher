package ubiviuscontroller

import (
	"context"
	"errors"

	v1 "agones.dev/agones/pkg/apis/agones/v1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getAgonesConfig() (*versioned.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err, "Could not create in cluster config")
		return nil, err
	}

	// Access to standard Kubernetes resources through the Kubernetes Clientset
	// We don't actually need this for this example, but it's just here for
	// illustrative purposes
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err, "Could not create the kubernetes clientset")
		return nil, err
	}

	// Access to the Agones resources through the Agones Clientset
	// Note that we reuse the same config as we used for the Kubernetes Clientset
	agonesClient, err := versioned.NewForConfig(config)
	if err != nil {
		log.Error(err, "Could not create the agones api clientset")
		return nil, err
	}

	return agonesClient, nil
}

func GetGameserverIP(id string) (string, error) {

	agonesClient, _ := getAgonesConfig()

	gameServer, err := agonesClient.AgonesV1().GameServers("default").Get(context.TODO(), "ubivius-game-"+id, metav1.GetOptions{})
	if err != nil {
		log.Error(err, "Error getting game server")
		return "", err
	}

	gameServerIP := &gameServer.Status.Address
	log.Info("Game server name is: %s", gameServer.ObjectMeta.Name)
	log.Info("Game server ip address is: %s", *gameServerIP)
	return *gameServerIP, nil
}

func GetReadyGameserver() (v1.GameServer, error) {
	agonesClient, _ := getAgonesConfig()

	gameServerList, err := agonesClient.AgonesV1().GameServers("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error(err, "Error getting game server list")
		return v1.GameServer{}, err
	}

	for _, gameServer := range gameServerList.Items {
		if gameServer.Status.State == "Ready" {
			return gameServer, nil
		}
	}
	return v1.GameServer{}, errors.New("No game server available")
}
