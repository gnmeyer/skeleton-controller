package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	workqueue "k8s.io/client-go/util/workqueue"
)

type controller struct {
	clientset kubernetes.Interface
	depLister appslisters.DeploymentLister
	//informer maintains a cache
	depCacheSyncd cache.InformerSynced
	queue         workqueue.RateLimitingInterface
}

func newController(clientset kubernetes.Interface, depInformer appsinformers.DeploymentInformer) *controller {
	c := &controller{
		clientset:     clientset,
		depLister:     depInformer.Lister(),
		depCacheSyncd: depInformer.Informer().HasSynced,
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ekpose"),
	}

	depInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    handleAdd,
			DeleteFunc: handleDel,
		},
	)
	return c

}

func (c *controller) run(ch <-chan struct{}) {
	fmt.Println("starting controller")
	if !cache.WaitForCacheSync(ch, c.depCacheSyncd) {
		fmt.Printf("error waiting for cache to sync\n")
	}

	//run worker fun while ch is open
	go wait.Until(c.worker, 1*time.Second, ch)

	<-ch //block until ch is closed

}

func (c *controller) worker() {

}

func handleAdd(obj interface{}) {
	fmt.Println("add was called")
}

func handleDel(obj interface{}) {
	fmt.Println("delete was called")

}
