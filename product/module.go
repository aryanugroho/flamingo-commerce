package product

import (
	"flamingo.me/flamingo-commerce/product/interfaces/controller"
	"flamingo.me/flamingo-commerce/product/interfaces/templatefunctions"
	"flamingo.me/flamingo/framework/config"
	"flamingo.me/flamingo/framework/dingo"
	"flamingo.me/flamingo/framework/router"
	"flamingo.me/flamingo/framework/template"
)

// Module registers our profiler
type Module struct{}

// Configure the product URL
func (m *Module) Configure(injector *dingo.Injector) {
	injector.BindMap(new(template.CtxFunc), "getProduct").To(templatefunctions.GetProduct{})
	injector.BindMap(new(template.Func), "getProductUrl").To(templatefunctions.GetProductUrl{})

	router.Bind(injector, new(routes))
}

// DefaultConfig for this module
func (m *Module) DefaultConfig() config.Map {
	return config.Map{
		"core.product.view.template": "product/product",
		"templating": config.Map{
			"product": config.Map{
				"attributeRenderer": config.Map{},
			},
		},
		"pagination.defaultPageSize": float64(36),
	}
}

type routes struct {
	controller *controller.View
}

func (r *routes) Inject(controller *controller.View) {
	r.controller = controller
}

func (r *routes) Routes(registry *router.Registry) {
	registry.HandleGet("product.view", r.controller.Get)
	registry.Route("/product/:marketplacecode/:name.html", `product.view(marketplacecode, name, backurl?="")`).Normalize("marketplacecode", "name")
	registry.Route("/product/:marketplacecode/:variantcode/:name.html", `product.view(marketplacecode, variantcode, name, backurl?="")`).Normalize("marketplacecode", "name", "variantcode")
}
