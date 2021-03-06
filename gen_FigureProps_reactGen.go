// Code generated by reactGen. DO NOT EDIT.

package react

// FigureProps defines the properties for the <Figure> element.
type FigureProps struct {
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

func (f *FigureProps) assign(_v *_FigureProps) {

	_v.AriaExpanded = f.AriaExpanded

	_v.AriaHasPopup = f.AriaHasPopup

	_v.AriaLabelledBy = f.AriaLabelledBy

	_v.ClassName = f.ClassName

	_v.DangerouslySetInnerHTML = f.DangerouslySetInnerHTML

	if f.DataSet != nil {
		for dk, dv := range f.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if f.ID != "" {
		_v.ID = f.ID
	}

	if f.Key != "" {
		_v.Key = f.Key
	}

	if f.OnChange != nil {
		_v.o.Set("onChange", f.OnChange.OnChange)
	}

	if f.OnClick != nil {
		_v.o.Set("onClick", f.OnClick.OnClick)
	}

	if f.Ref != nil {
		_v.o.Set("ref", f.Ref.Ref)
	}

	_v.Role = f.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = f.Style.hack()

}
