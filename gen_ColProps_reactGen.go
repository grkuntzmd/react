// Code generated by reactGen. DO NOT EDIT.

package react

// ColProps defines the properties for the <Col> element.
type ColProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	BGColor                 string
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Span  string
	Style *CSS
}

func (c *ColProps) assign(_v *_ColProps) {

	_v.AriaExpanded = c.AriaExpanded

	_v.AriaHasPopup = c.AriaHasPopup

	_v.AriaLabelledBy = c.AriaLabelledBy

	_v.BGColor = c.BGColor

	_v.ClassName = c.ClassName

	_v.DangerouslySetInnerHTML = c.DangerouslySetInnerHTML

	if c.DataSet != nil {
		for dk, dv := range c.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if c.ID != "" {
		_v.ID = c.ID
	}

	if c.Key != "" {
		_v.Key = c.Key
	}

	if c.OnChange != nil {
		_v.o.Set("onChange", c.OnChange.OnChange)
	}

	if c.OnClick != nil {
		_v.o.Set("onClick", c.OnClick.OnClick)
	}

	if c.Ref != nil {
		_v.o.Set("ref", c.Ref.Ref)
	}

	_v.Role = c.Role

	_v.Span = c.Span

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = c.Style.hack()

}
