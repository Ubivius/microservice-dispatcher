package controller

import (
	"context"
	"fmt"

	agonesv1 "agones.dev/agones/pkg/apis/agones/v1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func CreateGameserver() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err, "Could not create in cluster config")
	}

	// Access to standard Kubernetes resources through the Kubernetes Clientset
	// We don't actually need this for this example, but it's just here for
	// illustrative purposes
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err, "Could not create the kubernetes clientset")
	}

	// Access to the Agones resources through the Agones Clientset
	// Note that we reuse the same config as we used for the Kubernetes Clientset
	agonesClient, err := versioned.NewForConfig(config)
	if err != nil {
		log.Error(err, "Could not create the agones api clientset")
	}

	// Create a GameServer
	gs := &agonesv1.GameServer{ObjectMeta: metav1.ObjectMeta{GenerateName: "simple-game-server", Namespace: "default"},
		Spec: agonesv1.GameServerSpec{
			Container: "simple-game-server",
			Ports: []agonesv1.GameServerPort{{
				ContainerPort: 7654,
				HostPort:      7654,
				Name:          "gameport",
				PortPolicy:    agonesv1.Static,
				Protocol:      corev1.ProtocolUDP,
			}},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{Name: "simple-game-server", Image: "gcr.io/agones-images/simple-game-server:0.3"}},
				},
			},
		},
	}
	newGS, err := agonesClient.AgonesV1().GameServers("default").Create(context.TODO(), gs, metav1.CreateOptions{})
	if err != nil {
		log.Error(err, "Error creating game server")
	}

	fmt.Printf("New game servers' name is: %s", newGS.ObjectMeta.Name)
}
