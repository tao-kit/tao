package api

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// Part 4
// The apiparser_parser.go file was split into multiple files because it
// was too large and caused a possible memory overflow during Taoctl installation.

func (s *TypeStructContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeStructContext) GetStructName() antlr.Token { return s.structName }

func (s *TypeStructContext) GetStructToken() antlr.Token { return s.structToken }

func (s *TypeStructContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *TypeStructContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *TypeStructContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *TypeStructContext) SetStructToken(v antlr.Token) { s.structToken = v }

func (s *TypeStructContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *TypeStructContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *TypeStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *TypeStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *TypeStructContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeStructContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *TypeStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeStruct(s)
	}
}

func (s *TypeStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeStruct(s)
	}
}

func (s *TypeStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeStruct() (localctx ITypeStructContext) {
	this := p
	_ = this

	localctx = NewTypeStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ApiParserParserRULE_typeStruct)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	checkKeyword(p)
	{
		p.SetState(156)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeStructContext).structName = _m
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(157)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeStructContext).structToken = _m
		}

	}
	{
		p.SetState(160)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeStructContext).lbrace = _m
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(161)
				p.Field()
			}

		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	{
		p.SetState(167)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*TypeStructContext).rbrace = _m
	}

	return localctx
}

// ITypeAliasContext is an interface to support dynamic dispatch.
type ITypeAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// GetAssign returns the assign token.
	GetAssign() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// SetAssign sets the assign token.
	SetAssign(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext
	ID() antlr.TerminalNode

	// IsTypeAliasContext differentiates from other interfaces.
	IsTypeAliasContext()
}

type TypeAliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	assign antlr.Token
}

func NewEmptyTypeAliasContext() *TypeAliasContext {
	var p = new(TypeAliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeAlias
	return p
}

func (*TypeAliasContext) IsTypeAliasContext() {}

func NewTypeAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAliasContext {
	var p = new(TypeAliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeAlias

	return p
}

func (s *TypeAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAliasContext) GetAlias() antlr.Token { return s.alias }

func (s *TypeAliasContext) GetAssign() antlr.Token { return s.assign }

func (s *TypeAliasContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *TypeAliasContext) SetAssign(v antlr.Token) { s.assign = v }

func (s *TypeAliasContext) DataType() IDataTypeContext {
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

func (s *TypeAliasContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeAlias(s)
	}
}

func (s *TypeAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeAlias(s)
	}
}

func (s *TypeAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeAlias() (localctx ITypeAliasContext) {
	this := p
	_ = this

	localctx = NewTypeAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ApiParserParserRULE_typeAlias)
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
	checkKeyword(p)
	{
		p.SetState(170)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeAliasContext).alias = _m
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(171)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeAliasContext).assign = _m
		}

	}
	{
		p.SetState(174)
		p.DataType()
	}

	return localctx
}

// ITypeBlockStructContext is an interface to support dynamic dispatch.
type ITypeBlockStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// GetStructToken returns the structToken token.
	GetStructToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// SetStructToken sets the structToken token.
	SetStructToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsTypeBlockStructContext differentiates from other interfaces.
	IsTypeBlockStructContext()
}

type TypeBlockStructContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	rbrace      antlr.Token
}

func NewEmptyTypeBlockStructContext() *TypeBlockStructContext {
	var p = new(TypeBlockStructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockStruct
	return p
}

func (*TypeBlockStructContext) IsTypeBlockStructContext() {}

func NewTypeBlockStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockStructContext {
	var p = new(TypeBlockStructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockStruct

	return p
}

func (s *TypeBlockStructContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockStructContext) GetStructName() antlr.Token { return s.structName }

func (s *TypeBlockStructContext) GetStructToken() antlr.Token { return s.structToken }

func (s *TypeBlockStructContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *TypeBlockStructContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *TypeBlockStructContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *TypeBlockStructContext) SetStructToken(v antlr.Token) { s.structToken = v }

func (s *TypeBlockStructContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *TypeBlockStructContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *TypeBlockStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *TypeBlockStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *TypeBlockStructContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeBlockStructContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *TypeBlockStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeBlockStruct(s)
	}
}

func (s *TypeBlockStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeBlockStruct(s)
	}
}

func (s *TypeBlockStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlockStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlockStruct() (localctx ITypeBlockStructContext) {
	this := p
	_ = this

	localctx = NewTypeBlockStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ApiParserParserRULE_typeBlockStruct)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	checkKeyword(p)
	{
		p.SetState(177)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockStructContext).structName = _m
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(178)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeBlockStructContext).structToken = _m
		}

	}
	{
		p.SetState(181)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeBlockStructContext).lbrace = _m
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(182)
				p.Field()
			}

		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	{
		p.SetState(188)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*TypeBlockStructContext).rbrace = _m
	}

	return localctx
}

// ITypeBlockAliasContext is an interface to support dynamic dispatch.
type ITypeBlockAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// GetAssign returns the assign token.
	GetAssign() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// SetAssign sets the assign token.
	SetAssign(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext
	ID() antlr.TerminalNode

	// IsTypeBlockAliasContext differentiates from other interfaces.
	IsTypeBlockAliasContext()
}

type TypeBlockAliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	assign antlr.Token
}

func NewEmptyTypeBlockAliasContext() *TypeBlockAliasContext {
	var p = new(TypeBlockAliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockAlias
	return p
}

func (*TypeBlockAliasContext) IsTypeBlockAliasContext() {}

func NewTypeBlockAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockAliasContext {
	var p = new(TypeBlockAliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockAlias

	return p
}
