// +build js

// Code generated by elementGen. DO NOT EDIT.

package react_test

import (
	"testing"

	"honnef.co/go/js/dom"

	"myitcv.io/react"
	"myitcv.io/react/testutils"
)

func TestAbbrElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Abbr(&react.AbbrProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <abbr> element")
	}
}

func TestAcronymElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Acronym(&react.AcronymProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <acronym> element")
	}
}

func TestAddressElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Address(&react.AddressProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <address> element")
	}
}

func TestAppletElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Applet(&react.AppletProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLAppletElement); !ok {
		t.Fatal("Failed to find <applet> element")
	}
}

func TestAreaElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Area(&react.AreaProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLAreaElement); !ok {
		t.Fatal("Failed to find <area> element")
	}
}

func TestArticleElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Article(&react.ArticleProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <article> element")
	}
}

func TestAsideElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Aside(&react.AsideProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <aside> element")
	}
}

func TestAudioElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Audio(&react.AudioProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLAudioElement); !ok {
		t.Fatal("Failed to find <audio> element")
	}
}

func TestBElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.B(&react.BProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <b> element")
	}
}

func TestBaseElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Base(&react.BaseProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLBaseElement); !ok {
		t.Fatal("Failed to find <base> element")
	}
}

func TestBaseFontElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.BaseFont(&react.BaseFontProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <basefont> element")
	}
}

func TestBdiElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Bdi(&react.BdiProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <bdi> element")
	}
}

func TestBdoElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Bdo(&react.BdoProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <bdo> element")
	}
}

func TestBlockQuoteElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.BlockQuote(&react.BlockQuoteProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <blockquote> element")
	}
}

func TestBodyElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Body(&react.BodyProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLBodyElement); !ok {
		t.Fatal("Failed to find <body> element")
	}
}

func TestCanvasElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Canvas(&react.CanvasProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLCanvasElement); !ok {
		t.Fatal("Failed to find <canvas> element")
	}
}

func TestCaptionElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Caption(&react.CaptionProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <caption> element")
	}
}

func TestCiteElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Cite(&react.CiteProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <cite> element")
	}
}

func TestColElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Col(&react.ColProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <col> element")
	}
}

func TestColgroupElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Colgroup(&react.ColgroupProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <colgroup> element")
	}
}

func TestDDElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.DD(&react.DDProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <dd> element")
	}
}

func TestDataElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Data(&react.DataProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLDataElement); !ok {
		t.Fatal("Failed to find <data> element")
	}
}

func TestDataListElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.DataList(&react.DataListProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLDataListElement); !ok {
		t.Fatal("Failed to find <datalist> element")
	}
}

func TestDelElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Del(&react.DelProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <del> element")
	}
}

func TestDetailsElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Details(&react.DetailsProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <details> element")
	}
}

func TestDfnElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Dfn(&react.DfnProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <dfn> element")
	}
}

func TestDialogElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Dialog(&react.DialogProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <dialog> element")
	}
}

func TestDlElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Dl(&react.DlProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <dl> element")
	}
}

func TestDtElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Dt(&react.DtProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <dt> element")
	}
}

func TestEmElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Em(&react.EmProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <em> element")
	}
}

func TestEmbedElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Embed(&react.EmbedProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLEmbedElement); !ok {
		t.Fatal("Failed to find <embed> element")
	}
}

func TestFieldSetElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.FieldSet(&react.FieldSetProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLFieldSetElement); !ok {
		t.Fatal("Failed to find <fieldset> element")
	}
}

func TestFigCaptionElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.FigCaption(&react.FigCaptionProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <figcaption> element")
	}
}

func TestFigureElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Figure(&react.FigureProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <figure> element")
	}
}

func TestH2Elem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.H2(&react.H2Props{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <h2> element")
	}
}

func TestH5Elem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.H5(&react.H5Props{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <h5> element")
	}
}

func TestH6Elem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.H6(&react.H6Props{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <h6> element")
	}
}

func TestHGroupElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.HGroup(&react.HGroupProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <hgroup> element")
	}
}

func TestHTMLElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.HTML(&react.HTMLProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <html> element")
	}
}

func TestHeadElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Head(&react.HeadProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLHeadElement); !ok {
		t.Fatal("Failed to find <head> element")
	}
}

func TestHeaderElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Header(&react.HeaderProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <header> element")
	}
}

func TestInsElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Ins(&react.InsProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <ins> element")
	}
}

func TestKbdElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Kbd(&react.KbdProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <kbd> element")
	}
}

func TestLegendElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Legend(&react.LegendProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLLegendElement); !ok {
		t.Fatal("Failed to find <legend> element")
	}
}

func TestLinkElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Link(&react.LinkProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLLinkElement); !ok {
		t.Fatal("Failed to find <link> element")
	}
}

func TestMainElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Main(&react.MainProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <main> element")
	}
}

func TestMapElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Map(&react.MapProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLMapElement); !ok {
		t.Fatal("Failed to find <map> element")
	}
}

func TestMarkElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Mark(&react.MarkProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <mark> element")
	}
}

func TestMenuElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Menu(&react.MenuProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLMenuElement); !ok {
		t.Fatal("Failed to find <menu> element")
	}
}

func TestMetaElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Meta(&react.MetaProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLMetaElement); !ok {
		t.Fatal("Failed to find <meta> element")
	}
}

func TestMeterElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Meter(&react.MeterProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLMeterElement); !ok {
		t.Fatal("Failed to find <meter> element")
	}
}

func TestNoScriptElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.NoScript(&react.NoScriptProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <noscript> element")
	}
}

func TestObjectElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Object(&react.ObjectProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLObjectElement); !ok {
		t.Fatal("Failed to find <object> element")
	}
}

func TestOlElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Ol(&react.OlProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <ol> element")
	}
}

func TestOptGroupElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.OptGroup(&react.OptGroupProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLOptGroupElement); !ok {
		t.Fatal("Failed to find <optgroup> element")
	}
}

func TestOutputElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Output(&react.OutputProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLOutputElement); !ok {
		t.Fatal("Failed to find <output> element")
	}
}

func TestParamElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Param(&react.ParamProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLParamElement); !ok {
		t.Fatal("Failed to find <param> element")
	}
}

func TestPictureElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Picture(&react.PictureProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <picture> element")
	}
}

func TestProgressElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Progress(&react.ProgressProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLProgressElement); !ok {
		t.Fatal("Failed to find <progress> element")
	}
}

func TestQElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Q(&react.QProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <q> element")
	}
}

func TestRPElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.RP(&react.RPProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <rp> element")
	}
}

func TestRTElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.RT(&react.RTProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <rt> element")
	}
}

func TestRTCElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.RTC(&react.RTCProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <rtc> element")
	}
}

func TestRubyElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Ruby(&react.RubyProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <ruby> element")
	}
}

func TestSampElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Samp(&react.SampProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <samp> element")
	}
}

func TestScriptElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Script(&react.ScriptProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLScriptElement); !ok {
		t.Fatal("Failed to find <script> element")
	}
}

func TestSectionElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Section(&react.SectionProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <section> element")
	}
}

func TestSlotElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Slot(&react.SlotProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <slot> element")
	}
}

func TestSmallElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Small(&react.SmallProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <small> element")
	}
}

func TestSourceElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Source(&react.SourceProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLSourceElement); !ok {
		t.Fatal("Failed to find <source> element")
	}
}

func TestStrikeElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Strike(&react.StrikeProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <strike> element")
	}
}

func TestStrongElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Strong(&react.StrongProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <strong> element")
	}
}

func TestStyleElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Style(&react.StyleProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLStyleElement); !ok {
		t.Fatal("Failed to find <style> element")
	}
}

func TestSubElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Sub(&react.SubProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <sub> element")
	}
}

func TestTbodyElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Tbody(&react.TbodyProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <tbody> element")
	}
}

func TestTdElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Td(&react.TdProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <td> element")
	}
}

func TestTemplateElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Template(&react.TemplateProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <template> element")
	}
}

func TestTfootElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Tfoot(&react.TfootProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <tfoot> element")
	}
}

func TestThElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Th(&react.ThProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <th> element")
	}
}

func TestTheadElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Thead(&react.TheadProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <thead> element")
	}
}

func TestTimeElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Time(&react.TimeProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLTimeElement); !ok {
		t.Fatal("Failed to find <time> element")
	}
}

func TestTitleElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Title(&react.TitleProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLTitleElement); !ok {
		t.Fatal("Failed to find <title> element")
	}
}

func TestTrElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Tr(&react.TrProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <tr> element")
	}
}

func TestTrackElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Track(&react.TrackProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLTrackElement); !ok {
		t.Fatal("Failed to find <track> element")
	}
}

func TestUElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.U(&react.UProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <u> element")
	}
}

func TestVarElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Var(&react.VarProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <var> element")
	}
}

func TestVideoElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Video(&react.VideoProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLVideoElement); !ok {
		t.Fatal("Failed to find <video> element")
	}
}

func TestWbrElem(t *testing.T) {
	class := "test"

	x := testutils.Wrapper(react.Wbr(&react.WbrProps{ClassName: class}))
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.BasicHTMLElement); !ok {
		t.Fatal("Failed to find <wbr> element")
	}
}