package api

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// Part 9
// The apiparser_parser.go file was split into multiple files because it
// was too large and caused a possible memory overflow during Taoctl installation.

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
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
	ID() antlr.TerminalNode

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) GetLp() antlr.Token { return s.lp }

func (s *BodyContext) GetRp() antlr.Token { return s.rp }

func (s *BodyContext) SetLp(v antlr.Token) { s.lp = v }

func (s *BodyContext) SetRp(v antlr.Token) { s.rp = v }

func (s *BodyContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Body() (localctx IBodyContext) {
	this := p
	_ = this

	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ApiParserParserRULE_body)
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
		p.SetState(301)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*BodyContext).lp = _m
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(302)
			p.Match(ApiParserParserID)
		}

	}
	{
		p.SetState(305)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*BodyContext).rp = _m
	}

	return localctx
}

// IReplybodyContext is an interface to support dynamic dispatch.
type IReplybodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnToken returns the returnToken token.
	GetReturnToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetReturnToken sets the returnToken token.
	SetReturnToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext

	// IsReplybodyContext differentiates from other interfaces.
	IsReplybodyContext()
}

type ReplybodyContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	returnToken antlr.Token
	lp          antlr.Token
	rp          antlr.Token
}

func NewEmptyReplybodyContext() *ReplybodyContext {
	var p = new(ReplybodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_replybody
	return p
}

func (*ReplybodyContext) IsReplybodyContext() {}

func NewReplybodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplybodyContext {
	var p = new(ReplybodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_replybody

	return p
}

func (s *ReplybodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplybodyContext) GetReturnToken() antlr.Token { return s.returnToken }

func (s *ReplybodyContext) GetLp() antlr.Token { return s.lp }

func (s *ReplybodyContext) GetRp() antlr.Token { return s.rp }

func (s *ReplybodyContext) SetReturnToken(v antlr.Token) { s.returnToken = v }

func (s *ReplybodyContext) SetLp(v antlr.Token) { s.lp = v }

func (s *ReplybodyContext) SetRp(v antlr.Token) { s.rp = v }

func (s *ReplybodyContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ReplybodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplybodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplybodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterReplybody(s)
	}
}

func (s *ReplybodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitReplybody(s)
	}
}

func (s *ReplybodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitReplybody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Replybody() (localctx IReplybodyContext) {
	this := p
	_ = this

	localctx = NewReplybodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ApiParserParserRULE_replybody)
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
		p.SetState(307)

		var _m = p.Match(ApiParserParserT__9)

		localctx.(*ReplybodyContext).returnToken = _m
	}
	{
		p.SetState(308)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*ReplybodyContext).lp = _m
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16843200) != 0 {
		{
			p.SetState(309)
			p.DataType()
		}

	}
	{
		p.SetState(312)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*ReplybodyContext).rp = _m
	}

	return localctx
}

// IKvLitContext is an interface to support dynamic dispatch.
type IKvLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	LINE_VALUE() antlr.TerminalNode

	// IsKvLitContext differentiates from other interfaces.
	IsKvLitContext()
}

type KvLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
}

func NewEmptyKvLitContext() *KvLitContext {
	var p = new(KvLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_kvLit
	return p
}

func (*KvLitContext) IsKvLitContext() {}

func NewKvLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvLitContext {
	var p = new(KvLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_kvLit

	return p
}

func (s *KvLitContext) GetParser() antlr.Parser { return s.parser }

func (s *KvLitContext) GetKey() antlr.Token { return s.key }

func (s *KvLitContext) GetValue() antlr.Token { return s.value }

func (s *KvLitContext) SetKey(v antlr.Token) { s.key = v }

func (s *KvLitContext) SetValue(v antlr.Token) { s.value = v }

func (s *KvLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *KvLitContext) LINE_VALUE() antlr.TerminalNode {
	return s.GetToken(ApiParserParserLINE_VALUE, 0)
}

func (s *KvLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterKvLit(s)
	}
}

func (s *KvLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitKvLit(s)
	}
}

func (s *KvLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitKvLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) KvLit() (localctx IKvLitContext) {
	this := p
	_ = this

	localctx = NewKvLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ApiParserParserRULE_kvLit)

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
		p.SetState(314)

		var _m = p.Match(ApiParserParserID)

		localctx.(*KvLitContext).key = _m
	}
	checkKeyValue(p)
	{
		p.SetState(316)

		var _m = p.Match(ApiParserParserLINE_VALUE)

		localctx.(*KvLitContext).value = _m
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *ServiceNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceName() (localctx IServiceNameContext) {
	this := p
	_ = this

	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ApiParserParserRULE_serviceName)
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
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(318)
			p.Match(ApiParserParserID)
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ApiParserParserT__10 {
			{
				p.SetState(319)
				p.Match(ApiParserParserT__10)
			}

		}

		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPathItem() []IPathItemContext
	PathItem(i int) IPathItemContext

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_path

	return p
}
