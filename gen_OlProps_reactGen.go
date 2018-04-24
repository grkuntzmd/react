// Code generated by reactGen. DO NOT EDIT.

package react

// OlProps defines the properties for the <Ol> element.
type OlProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	ClassName               string
	Compact                 string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Reversed bool
	Role     string
	Start    string
	Style    *CSS
	Type     string
}

func (o *OlProps) assign(_v *_OlProps) {

	_v.AriaExpanded = o.AriaExpanded

	_v.AriaHasPopup = o.AriaHasPopup

	_v.AriaLabelledBy = o.AriaLabelledBy

	_v.ClassName = o.ClassName

	_v.Compact = o.Compact

	_v.DangerouslySetInnerHTML = o.DangerouslySetInnerHTML

	if o.DataSet != nil {
		for dk, dv := range o.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if o.ID != "" {
		_v.ID = o.ID
	}

	if o.Key != "" {
		_v.Key = o.Key
	}

	if o.OnChange != nil {
		_v.o.Set("onChange", o.OnChange.OnChange)
	}

	if o.OnClick != nil {
		_v.o.Set("onClick", o.OnClick.OnClick)
	}

	if o.Ref != nil {
		_v.o.Set("ref", o.Ref.Ref)
	}

	_v.Reversed = o.Reversed

	_v.Role = o.Role

	_v.Start = o.Start

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = o.Style.hack()

	_v.Type = o.Type

}
