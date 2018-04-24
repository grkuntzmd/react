// Code generated by reactGen. DO NOT EDIT.

package react

// MetaProps defines the properties for the <Meta> element.
type MetaProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	CharSet                 string
	ClassName               string
	Content                 string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	HTTPEquiv string
	ID        string
	Key       string
	Name      string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (m *MetaProps) assign(_v *_MetaProps) {

	_v.AriaExpanded = m.AriaExpanded

	_v.AriaHasPopup = m.AriaHasPopup

	_v.AriaLabelledBy = m.AriaLabelledBy

	_v.CharSet = m.CharSet

	_v.ClassName = m.ClassName

	_v.Content = m.Content

	_v.DangerouslySetInnerHTML = m.DangerouslySetInnerHTML

	if m.DataSet != nil {
		for dk, dv := range m.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	_v.HTTPEquiv = m.HTTPEquiv

	if m.ID != "" {
		_v.ID = m.ID
	}

	if m.Key != "" {
		_v.Key = m.Key
	}

	_v.Name = m.Name

	if m.OnChange != nil {
		_v.o.Set("onChange", m.OnChange.OnChange)
	}

	if m.OnClick != nil {
		_v.o.Set("onClick", m.OnClick.OnClick)
	}

	if m.Ref != nil {
		_v.o.Set("ref", m.Ref.Ref)
	}

	_v.Role = m.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = m.Style.hack()

}
