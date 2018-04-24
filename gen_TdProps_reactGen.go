// Code generated by reactGen. DO NOT EDIT.

package react

// TdProps defines the properties for the <Td> element.
type TdProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	BGColor                 string
	ClassName               string
	ColSpan                 string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	Headers string
	ID      string
	Key     string

	OnChange
	OnClick

	Ref
	Role    string
	RowSpan string
	Style   *CSS
}

func (t *TdProps) assign(_v *_TdProps) {

	_v.AriaExpanded = t.AriaExpanded

	_v.AriaHasPopup = t.AriaHasPopup

	_v.AriaLabelledBy = t.AriaLabelledBy

	_v.BGColor = t.BGColor

	_v.ClassName = t.ClassName

	_v.ColSpan = t.ColSpan

	_v.DangerouslySetInnerHTML = t.DangerouslySetInnerHTML

	if t.DataSet != nil {
		for dk, dv := range t.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	_v.Headers = t.Headers

	if t.ID != "" {
		_v.ID = t.ID
	}

	if t.Key != "" {
		_v.Key = t.Key
	}

	if t.OnChange != nil {
		_v.o.Set("onChange", t.OnChange.OnChange)
	}

	if t.OnClick != nil {
		_v.o.Set("onClick", t.OnClick.OnClick)
	}

	if t.Ref != nil {
		_v.o.Set("ref", t.Ref.Ref)
	}

	_v.Role = t.Role

	_v.RowSpan = t.RowSpan

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = t.Style.hack()

}
