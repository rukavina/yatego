package yatego

// Base is base component, can just answer call, the rest you have to programm it
type Base struct {
	componentCommon
	componentCallback
	componentYate
	componentMedia
}

// pseudo 2nd step constructor
func (b *Base) init() {
	b.Listen(MsgCallExecute, func(call *Call, msg *Message) *CallbackResult {
		b.logger.Infof("Component [%s] going to answer the call from [%s] to [%s]", b.name, call.Caller, call.Called)
		b.Answer(call, msg)
		return NewCallbackResult(ResEnter, "")
	})
	b.initListeners()
}

// hook method
func (b *Base) initListeners() {

}

// NewBaseComponent generates new base component
func NewBaseComponent(name string, engine *Engine, logger Logger, config map[string]interface{}) *Base {
	common := componentCommon{
		name:   name,
		logger: logger,
		config: config,
	}
	cb := componentCallback{
		callbacks: make(map[string]Callback),
	}
	yate := componentYate{
		engine:            engine,
		messagesToInstall: make(map[string]InstallDef),
		messagesToWatch:   make([]string, 0),
	}
	media := componentMedia{
		yate,
	}
	com := &Base{
		common,
		cb,
		yate,
		media,
	}
	com.init()
	return com
}
