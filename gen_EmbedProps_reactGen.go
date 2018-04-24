// Code generated by reactGen. DO NOT EDIT.

package react

// EmbedProps defines the properties for the <Embed> element.
type EmbedProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	Height string
	ID     string
	Key    string

	OnChange
	OnClick

	Ref
	Role  string
	Src   string
	Style *CSS
	Type  string
	Width string
}

func (e *EmbedProps) assign(_v *_EmbedProps) {

	_v.AriaExpanded = e.AriaExpanded

	_v.AriaHasPopup = e.AriaHasPopup

	_v.AriaLabelledBy = e.AriaLabelledBy

	_v.ClassName = e.ClassName

	_v.DangerouslySetInnerHTML = e.DangerouslySetInnerHTML

	if e.DataSet != nil {
		for dk, dv := range e.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	_v.Height = e.Height

	if e.ID != "" {
		_v.ID = e.ID
	}

	if e.Key != "" {
		_v.Key = e.Key
	}

	if e.OnChange != nil {
		_v.o.Set("onChange", e.OnChange.OnChange)
	}

	if e.OnClick != nil {
		_v.o.Set("onClick", e.OnClick.OnClick)
	}

	if e.Ref != nil {
		_v.o.Set("ref", e.Ref.Ref)
	}

	_v.Role = e.Role

	_v.Src = e.Src

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = e.Style.hack()

	_v.Type = e.Type

	_v.Width = e.Width

}
