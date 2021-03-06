// Code generated by reactGen. DO NOT EDIT.

package react

// IProps are the props for a <i> component
type IProps struct {
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
	Src   string
	Style *CSS
}

func (i *IProps) assign(_v *_IProps) {

	_v.AriaExpanded = i.AriaExpanded

	_v.AriaHasPopup = i.AriaHasPopup

	_v.AriaLabelledBy = i.AriaLabelledBy

	_v.ClassName = i.ClassName

	_v.DangerouslySetInnerHTML = i.DangerouslySetInnerHTML

	if i.DataSet != nil {
		for dk, dv := range i.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

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

	_v.Src = i.Src

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = i.Style.hack()

}
