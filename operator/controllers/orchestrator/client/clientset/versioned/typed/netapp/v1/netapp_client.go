// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/operator/controllers/orchestrator/apis/netapp/v1"
	"github.com/netapp/trident/operator/controllers/orchestrator/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type TridentV1Interface interface {
	RESTClient() rest.Interface
	TridentOrchestratorsGetter
}

// TridentV1Client is used to interact with features provided by the trident.netapp.io group.
type TridentV1Client struct {
	restClient rest.Interface
}

func (c *TridentV1Client) TridentOrchestrators() TridentOrchestratorInterface {
	return newTridentOrchestrators(c)
}

// NewForConfig creates a new TridentV1Client for the given config.
func NewForConfig(c *rest.Config) (*TridentV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TridentV1Client{client}, nil
}

// NewForConfigOrDie creates a new TridentV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TridentV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TridentV1Client for the given RESTClient.
func New(c rest.Interface) *TridentV1Client {
	return &TridentV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TridentV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
