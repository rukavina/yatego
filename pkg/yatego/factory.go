package yatego

import "github.com/rukavina/minidic"

// Factory returns factory object
type Factory struct {
	container minidic.Container
}

// NewFactory factory constructor
func NewFactory() *Factory {
	return &Factory{
		container: dic(),
	}
}

// Container returns DIC container
func (f *Factory) Container() minidic.Container {
	if f.container == nil {
		f.container = dic()
	}
	return f.container
}

// Controller get controller service instance
func (f *Factory) Controller(cl CallflowLoader) *Controller {
	c := f.Container().Get("controller").(*Controller)
	if cl != nil {
		c.callflowLoader = cl
	}
	return c
}

// BaseComponent generates base component
func (f *Factory) BaseComponent() Component {
	fac := BaseComponentFactory(f.Container())
	return fac("", "start", map[string]interface{}{})
}
