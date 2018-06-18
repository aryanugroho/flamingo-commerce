package templatefunctions

import (
	"flamingo.me/flamingo-commerce/product/application"
	"flamingo.me/flamingo-commerce/product/domain"
)

type (
	// GetProductUrl is exported as a template function
	GetProductUrl struct {
		UrlService *application.UrlService `inject:""`
	}
)

// Name alias for use in template
func (tf GetProductUrl) Name() string {
	return "getProductUrl"
}

// Func returns the JSON object
func (tf GetProductUrl) Func() interface{} {
	return func(p domain.BasicProduct) string {
		if p == nil {
			return ""
		}
		url, err := tf.UrlService.Get(p, "")
		if err != nil {
			return ""
		}
		return url
	}
}
