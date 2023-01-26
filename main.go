package main

import (
	"context"
	"flag"
	myv1 "github.com/ishtiaqhimel/crd/pkg/apis/crd.com/v1"
	sbclientset "github.com/ishtiaqhimel/crd/pkg/client/clientset/versioned"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	log.Println("Configuring KubeConfig...")

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	crdClient, err := crdclientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	customCRD := v1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "sakiibbhais.crd.com",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "crd.com",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Type: "object",

							Properties: map[string]v1.JSONSchemaProps{
								"spec": {
									Type: "object",
									Properties: map[string]v1.JSONSchemaProps{
										"name": {
											Type: "string",
										},
										"replicas": {
											Type: "integer",
										},
										"container": {
											Type: "object",
											Properties: map[string]v1.JSONSchemaProps{
												"image": {
													Type: "string",
												},
												"port": {
													Type: "integer",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			Scope: "Namespaced",
			Names: v1.CustomResourceDefinitionNames{
				Kind:     "SakiibBhai",
				Plural:   "sakiibbhais",
				Singular: "sakiibbhai",
				ShortNames: []string{
					"sb",
				},
				Categories: []string{
					"all",
				},
			},
		},
	}
	ctx := context.TODO()
	// deleting existing CR
	_ = crdClient.ApiextensionsV1().CustomResourceDefinitions().Delete(ctx, customCRD.Name, metav1.DeleteOptions{})
	// creating new one
	_, err = crdClient.ApiextensionsV1().CustomResourceDefinitions().Create(ctx, &customCRD, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	log.Println("CRD is Created!")

	log.Println("Press ctrl+c to create a SakiibBhai")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	log.Println("Creating SakiibBhai...")

	client, err := sbclientset.NewForConfig(config)
	sbObj := &myv1.SakiibBhai{
		ObjectMeta: metav1.ObjectMeta{
			Name: "sakiibbhai",
		},
		Spec: myv1.SakiibBhaiSpec{
			Name:     "CustomSakiibBhai",
			Replicas: intptr(2),
			Container: myv1.ContainerSpec{
				Image: "ishtiaq99/go-api-server",
				Port:  3000,
			},
		},
	}
	_, err = client.CrdV1().SakiibBhais("default").Create(ctx, sbObj, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	log.Println("SakiibBhai Created!!")

	log.Println("Press ctrl+c to clean up...")
	signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	err = client.CrdV1().SakiibBhais("default").Delete(ctx, "sakiibbhai", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}

	err = crdClient.ApiextensionsV1().CustomResourceDefinitions().Delete(ctx, "sakiibbhais.crd.com", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	log.Println("Cleaned Up!!!")
}

func intptr(i int32) *int32 {
	return &i
}
