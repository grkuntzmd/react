// Code generated by reactGen. DO NOT EDIT.

package react

// InsProps defines the properties for the <Ins> element.
type InsProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	Cite                    string
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	DateTime string
	ID       string
	Key      string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (i *InsProps) assign(_v *_InsProps) {

	_v.AriaExpanded = i.AriaExpanded

	_v.AriaHasPopup = i.AriaHasPopup

	_v.AriaLabelledBy = i.AriaLabelledBy

	_v.Cite = i.Cite

	_v.ClassName = i.ClassName

	_v.DangerouslySetInnerHTML = i.DangerouslySetInnerHTML

	if i.DataSet != nil {
		for dk, dv := range i.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	_v.DateTime = i.DateTime

	if i.ID != "" {
		_v.ID = i.ID
	}

	if i.Key != "" {
		_v.Key = i.Key
	}

	if i.OnChange != nil {
		_v.o.Set("onChange", i.OnChange.OnChange)
	}

	if i.OnClick != nil {
		_v.o.Set("onClick", i.OnClick.OnClick)
	}

	if i.Ref != nil {
		_v.o.Set("ref", i.Ref.Ref)
	}

	_v.Role = i.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = i.Style.hack()

}
