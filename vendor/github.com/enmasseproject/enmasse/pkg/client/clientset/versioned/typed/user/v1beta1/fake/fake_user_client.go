/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/enmasseproject/enmasse/pkg/client/clientset/versioned/typed/user/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeUserV1beta1 struct {
	*testing.Fake
}

func (c *FakeUserV1beta1) MessagingUsers(namespace string) v1beta1.MessagingUserInterface {
	return &FakeMessagingUsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeUserV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}