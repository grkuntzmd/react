// Code generated by reactGen. DO NOT EDIT.

package react

// AbbrProps defines the properties for the <Abbr> element.
type AbbrProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (a *AbbrProps) assign(_v *_AbbrProps) {

	_v.AriaExpanded = a.AriaExpanded

	_v.AriaHasPopup = a.AriaHasPopup

	_v.AriaLabelledBy = a.AriaLabelledBy

	_v.ClassName = a.ClassName

	_v.DangerouslySetInnerHTML = a.DangerouslySetInnerHTML

	if a.DataSet != nil {
		for dk, dv := range a.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if a.ID != "" {
		_v.ID = a.ID
	}

	if a.Key != "" {
		_v.Key = a.Key
	}

	if a.OnChange != nil {
		_v.o.Set("onChange", a.OnChange.OnChange)
	}

	if a.OnClick != nil {
		_v.o.Set("onClick", a.OnClick.OnClick)
	}

	if a.Ref != nil {
		_v.o.Set("ref", a.Ref.Ref)
	}

	_v.Role = a.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = a.Style.hack()

}