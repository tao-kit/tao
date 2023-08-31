package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Part 10
// The apiparser_parser.go file was split into multiple files because it
// was too large and caused a possible memory overflow during Taoctl installation.

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllPathItem() []IPathItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathItemContext); ok {
			len++
		}
	}

	tst := make([]IPathItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathItemContext); ok {
			tst[i] = t.(IPathItemContext)
			i++
		}
	}

	return tst
}

func (s *PathContext) PathItem(i int) IPathItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathItemContext); ok {
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

	return t.(IPathItemContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Path() (localctx IPathContext) {
	this := p
	_ = this

	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ApiParserParserRULE_path)
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

	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ApiParserParserT__11 || _la == ApiParserParserT__12 {
			p.SetState(341)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case ApiParserParserT__11:
				{
					p.SetState(326)
					p.Match(ApiParserParserT__11)
				}

				{
					p.SetState(327)
					p.PathItem()
				}
				p.SetState(332)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == ApiParserParserT__10 {
					{
						p.SetState(328)
						p.Match(ApiParserParserT__10)
					}
					{
						p.SetState(329)
						p.PathItem()
					}

					p.SetState(334)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}

			case ApiParserParserT__12:
				{
					p.SetState(335)
					p.Match(ApiParserParserT__12)
				}

				{
					p.SetState(336)
					p.PathItem()
				}
				p.SetState(339)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == ApiParserParserT__10 {
					{
						p.SetState(337)
						p.Match(ApiParserParserT__10)
					}
					{
						p.SetState(338)
						p.PathItem()
					}

				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(343)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.Match(ApiParserParserT__11)
		}

	}

	return localctx
}

// IPathItemContext is an interface to support dynamic dispatch.
type IPathItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllLetterOrDigit() []antlr.TerminalNode
	LetterOrDigit(i int) antlr.TerminalNode

	// IsPathItemContext differentiates from other interfaces.
	IsPathItemContext()
}

type PathItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathItemContext() *PathItemContext {
	var p = new(PathItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ApiParserParserRULE_pathItem
	return p
}

func (*PathItemContext) IsPathItemContext() {}

func NewPathItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathItemContext {
	var p = new(PathItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_pathItem

	return p
}

func (s *PathItemContext) GetParser() antlr.Parser { return s.parser }

func (s *PathItemContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *PathItemContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *PathItemContext) AllLetterOrDigit() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserLetterOrDigit)
}

func (s *PathItemContext) LetterOrDigit(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserLetterOrDigit, i)
}

func (s *PathItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterPathItem(s)
	}
}

func (s *PathItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitPathItem(s)
	}
}

func (s *PathItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPathItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) PathItem() (localctx IPathItemContext) {
	this := p
	_ = this

	localctx = NewPathItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ApiParserParserRULE_pathItem)
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
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID || _la == ApiParserParserLetterOrDigit {
		{
			p.SetState(348)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ApiParserParserID || _la == ApiParserParserLetterOrDigit) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *ApiParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *FieldContext = nil
		if localctx != nil {
			t = localctx.(*FieldContext)
		}
		return p.Field_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ApiParserParser) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return isNormal(p)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
