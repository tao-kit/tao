package api

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// Part 7
// The apiparser_parser.go file was split into multiple files because it
// was too large and caused a possible memory overflow during Taoctl installation.

// IServiceSpecContext is an interface to support dynamic dispatch.
type IServiceSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ServiceApi() IServiceApiContext
	AtServer() IAtServerContext

	// IsServiceSpecContext differentiates from other interfaces.
	IsServiceSpecContext()
}

type ServiceSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceSpecContext() *ServiceSpecContext {
	var p = new(ServiceSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceSpec
	return p
}

func (*ServiceSpecContext) IsServiceSpecContext() {}

func NewServiceSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceSpecContext {
	var p = new(ServiceSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceSpec

	return p
}

func (s *ServiceSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceSpecContext) ServiceApi() IServiceApiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceApiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceApiContext)
}

func (s *ServiceSpecContext) AtServer() IAtServerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtServerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtServerContext)
}

func (s *ServiceSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterServiceSpec(s)
	}
}

func (s *ServiceSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitServiceSpec(s)
	}
}

func (s *ServiceSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceSpec() (localctx IServiceSpecContext) {
	this := p
	_ = this

	localctx = NewServiceSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ApiParserParserRULE_serviceSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserATSERVER {
		{
			p.SetState(239)
			p.AtServer()
		}

	}
	{
		p.SetState(242)
		p.ServiceApi()
	}

	return localctx
}

// IAtServerContext is an interface to support dynamic dispatch.
type IAtServerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ATSERVER() antlr.TerminalNode
	AllKvLit() []IKvLitContext
	KvLit(i int) IKvLitContext

	// IsAtServerContext differentiates from other interfaces.
	IsAtServerContext()
}

type AtServerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtServerContext() *AtServerContext {
	var p = new(AtServerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_atServer
	return p
}

func (*AtServerContext) IsAtServerContext() {}

func NewAtServerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtServerContext {
	var p = new(AtServerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atServer

	return p
}

func (s *AtServerContext) GetParser() antlr.Parser { return s.parser }

func (s *AtServerContext) GetLp() antlr.Token { return s.lp }

func (s *AtServerContext) GetRp() antlr.Token { return s.rp }

func (s *AtServerContext) SetLp(v antlr.Token) { s.lp = v }

func (s *AtServerContext) SetRp(v antlr.Token) { s.rp = v }

func (s *AtServerContext) ATSERVER() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATSERVER, 0)
}

func (s *AtServerContext) AllKvLit() []IKvLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvLitContext); ok {
			len++
		}
	}

	tst := make([]IKvLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvLitContext); ok {
			tst[i] = t.(IKvLitContext)
			i++
		}
	}

	return tst
}

func (s *AtServerContext) KvLit(i int) IKvLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvLitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *AtServerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtServerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtServerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterAtServer(s)
	}
}

func (s *AtServerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitAtServer(s)
	}
}

func (s *AtServerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtServer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtServer() (localctx IAtServerContext) {
	this := p
	_ = this

	localctx = NewAtServerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ApiParserParserRULE_atServer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(ApiParserParserATSERVER)
	}
	{
		p.SetState(245)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*AtServerContext).lp = _m
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(246)
			p.KvLit()
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(251)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*AtServerContext).rp = _m
	}

	return localctx
}

// IServiceApiContext is an interface to support dynamic dispatch.
type IServiceApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetServiceToken returns the serviceToken token.
	GetServiceToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetServiceToken sets the serviceToken token.
	SetServiceToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// Getter signatures
	ServiceName() IServiceNameContext
	ID() antlr.TerminalNode
	AllServiceRoute() []IServiceRouteContext
	ServiceRoute(i int) IServiceRouteContext

	// IsServiceApiContext differentiates from other interfaces.
	IsServiceApiContext()
}

type ServiceApiContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	serviceToken antlr.Token
	lbrace       antlr.Token
	rbrace       antlr.Token
}

func NewEmptyServiceApiContext() *ServiceApiContext {
	var p = new(ServiceApiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceApi
	return p
}

func (*ServiceApiContext) IsServiceApiContext() {}

func NewServiceApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceApiContext {
	var p = new(ServiceApiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceApi

	return p
}

func (s *ServiceApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceApiContext) GetServiceToken() antlr.Token { return s.serviceToken }

func (s *ServiceApiContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *ServiceApiContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *ServiceApiContext) SetServiceToken(v antlr.Token) { s.serviceToken = v }

func (s *ServiceApiContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *ServiceApiContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *ServiceApiContext) ServiceName() IServiceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceApiContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ServiceApiContext) AllServiceRoute() []IServiceRouteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceRouteContext); ok {
			len++
		}
	}

	tst := make([]IServiceRouteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceRouteContext); ok {
			tst[i] = t.(IServiceRouteContext)
			i++
		}
	}

	return tst
}

func (s *ServiceApiContext) ServiceRoute(i int) IServiceRouteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceRouteContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceRouteContext)
}

func (s *ServiceApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceApiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterServiceApi(s)
	}
}

func (s *ServiceApiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitServiceApi(s)
	}
}

func (s *ServiceApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceApi() (localctx IServiceApiContext) {
	this := p
	_ = this

	localctx = NewServiceApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ApiParserParserRULE_serviceApi)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	match(p, "service")
	{
		p.SetState(254)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ServiceApiContext).serviceToken = _m
	}
	{
		p.SetState(255)
		p.ServiceName()
	}
	{
		p.SetState(256)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*ServiceApiContext).lbrace = _m
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&180224) != 0 {
		{
			p.SetState(257)
			p.ServiceRoute()
		}

		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(263)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*ServiceApiContext).rbrace = _m
	}

	return localctx
}

// IServiceRouteContext is an interface to support dynamic dispatch.
type IServiceRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Route() IRouteContext
	AtServer() IAtServerContext
	AtHandler() IAtHandlerContext
	AtDoc() IAtDocContext

	// IsServiceRouteContext differentiates from other interfaces.
	IsServiceRouteContext()
}

type ServiceRouteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceRouteContext() *ServiceRouteContext {
	var p = new(ServiceRouteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceRoute
	return p
}

func (*ServiceRouteContext) IsServiceRouteContext() {}

func NewServiceRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceRouteContext {
	var p = new(ServiceRouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceRoute

	return p
}

func (s *ServiceRouteContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceRouteContext) Route() IRouteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRouteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRouteContext)
}

func (s *ServiceRouteContext) AtServer() IAtServerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtServerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtServerContext)
}
