// Code generated by reactGen. DO NOT EDIT.

package react

// DlProps defines the properties for the <Dl> element.
type DlProps struct {
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

func (d *DlProps) assign(_v *_DlProps) {

	_v.AriaExpanded = d.AriaExpanded

	_v.AriaHasPopup = d.AriaHasPopup

	_v.AriaLabelledBy = d.AriaLabelledBy

	_v.ClassName = d.ClassName

	_v.DangerouslySetInnerHTML = d.DangerouslySetInnerHTML

	if d.DataSet != nil {
		for dk, dv := range d.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if d.ID != "" {
		_v.ID = d.ID
	}

	if d.Key != "" {
		_v.Key = d.Key
	}

	if d.OnChange != nil {
		_v.o.Set("onChange", d.OnChange.OnChange)
	}

	if d.OnClick != nil {
		_v.o.Set("onClick", d.OnClick.OnClick)
	}

	if d.Ref != nil {
		_v.o.Set("ref", d.Ref.Ref)
	}

	_v.Role = d.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = d.Style.hack()

}
