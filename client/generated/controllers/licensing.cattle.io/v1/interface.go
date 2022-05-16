/*
Copyright 2022.

All Rights Reserved
*/
// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/ebauman/klicense/api/v1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	Request() RequestController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) Request() RequestController {
	return NewRequestController(schema.GroupVersionKind{Group: "licensing.cattle.io", Version: "v1", Kind: "Request"}, "requests", true, c.controllerFactory)
}
