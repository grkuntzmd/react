// Code generated by reactGen. DO NOT EDIT.

package react

// BodyProps defines the properties for the <Body> element.
type BodyProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID             string
	Key            string
	OnAfterPrint   string
	OnBeforePrint  string
	OnBeforeUnload string
	OnBlur         string

	OnChange
	OnClick

	OnError          string
	OnFocus          string
	OnHashChange     string
	OnLanguageChange string
	OnLoad           string
	OnMessage        string
	OnOffline        string
	OnOnline         string
	OnPopState       string
	OnRedo           string
	OnResize         string
	OnStorage        string
	OnUndo           string
	OnUnload         string
	Ref
	Role  string
	Style *CSS
}

func (b *BodyProps) assign(_v *_BodyProps) {

	_v.AriaExpanded = b.AriaExpanded

	_v.AriaHasPopup = b.AriaHasPopup

	_v.AriaLabelledBy = b.AriaLabelledBy

	_v.ClassName = b.ClassName

	_v.DangerouslySetInnerHTML = b.DangerouslySetInnerHTML

	if b.DataSet != nil {
		for dk, dv := range b.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if b.ID != "" {
		_v.ID = b.ID
	}

	if b.Key != "" {
		_v.Key = b.Key
	}

	_v.OnAfterPrint = b.OnAfterPrint

	_v.OnBeforePrint = b.OnBeforePrint

	_v.OnBeforeUnload = b.OnBeforeUnload

	_v.OnBlur = b.OnBlur

	if b.OnChange != nil {
		_v.o.Set("onChange", b.OnChange.OnChange)
	}

	if b.OnClick != nil {
		_v.o.Set("onClick", b.OnClick.OnClick)
	}

	_v.OnError = b.OnError

	_v.OnFocus = b.OnFocus

	_v.OnHashChange = b.OnHashChange

	_v.OnLanguageChange = b.OnLanguageChange

	_v.OnLoad = b.OnLoad

	_v.OnMessage = b.OnMessage

	_v.OnOffline = b.OnOffline

	_v.OnOnline = b.OnOnline

	_v.OnPopState = b.OnPopState

	_v.OnRedo = b.OnRedo

	_v.OnResize = b.OnResize

	_v.OnStorage = b.OnStorage

	_v.OnUndo = b.OnUndo

	_v.OnUnload = b.OnUnload

	if b.Ref != nil {
		_v.o.Set("ref", b.Ref.Ref)
	}

	_v.Role = b.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = b.Style.hack()

}