// Code generated by reactGen. DO NOT EDIT.

package react

// RPProps defines the properties for the <RP> element.
type RPProps struct {
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

func (r *RPProps) assign(_v *_RPProps) {

	_v.AriaExpanded = r.AriaExpanded

	_v.AriaHasPopup = r.AriaHasPopup

	_v.AriaLabelledBy = r.AriaLabelledBy

	_v.ClassName = r.ClassName

	_v.DangerouslySetInnerHTML = r.DangerouslySetInnerHTML

	if r.DataSet != nil {
		for dk, dv := range r.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if r.ID != "" {
		_v.ID = r.ID
	}

	if r.Key != "" {
		_v.Key = r.Key
	}

	if r.OnChange != nil {
		_v.o.Set("onChange", r.OnChange.OnChange)
	}

	if r.OnClick != nil {
		_v.o.Set("onClick", r.OnClick.OnClick)
	}

	if r.Ref != nil {
		_v.o.Set("ref", r.Ref.Ref)
	}

	_v.Role = r.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = r.Style.hack()

}