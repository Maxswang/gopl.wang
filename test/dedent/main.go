package main

import (
	"fmt"

	"github.com/lithammer/dedent"
)

func main() {
	// s := `
	// 	Lorem ipsum dolor sit amet,
	// 	consectetur adipiscing elit.

	// 	Curabitur justo tellus, facilisis nec efficitur dictum,
	// 	fermentum vitae ligula. Sed eu convallis sapien.`

	s := `
		┌──────────────────────────────────────────────────────────┐
		│ KUBEADM                                                  │
		│ Easily bootstrap a secure Kubernetes cluster             │
		│                                                          │
		│ Please give us feedback at:                              │
		│ https://github.com/kubernetes/kubeadm/issues             │
		└──────────────────────────────────────────────────────────┘
	
	Example usage

		Create a two-machine cluster with one control-plane node
		(which controls the cluster), and one worker node
		(where your workloads, like Pods and Deployments run)
		┌──────────────────────────────────────────────────────────┐
		│ On the first machine:                                    │
		├──────────────────────────────────────────────────────────┤
		│ control-plane# kubeadm init                              │
		└──────────────────────────────────────────────────────────┘
		
		┌──────────────────────────────────────────────────────────┐
		│ On the second machine:                                   │
		├──────────────────────────────────────────────────────────┤
		│ worker# kubeadm join <arguments-returned-from-init>      │
		└──────────────────────────────────────────────────────────┘
		
		You can then repeat the second step on as many other machines as you like.

		`
	fmt.Println(dedent.Dedent(s))
	// fmt.Println("-------------------------------")
	// fmt.Println(s)
}
