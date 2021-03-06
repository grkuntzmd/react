// Code generated by reactGen. DO NOT EDIT.

package react

// SectionProps defines the properties for the <Section> element.
type SectionProps struct {
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

func (s *SectionProps) assign(_v *_SectionProps) {

	_v.AriaExpanded = s.AriaExpanded

	_v.AriaHasPopup = s.AriaHasPopup

	_v.AriaLabelledBy = s.AriaLabelledBy

	_v.ClassName = s.ClassName

	_v.DangerouslySetInnerHTML = s.DangerouslySetInnerHTML

	if s.DataSet != nil {
		for dk, dv := range s.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if s.ID != "" {
		_v.ID = s.ID
	}

	if s.Key != "" {
		_v.Key = s.Key
	}

	if s.OnChange != nil {
		_v.o.Set("onChange", s.OnChange.OnChange)
	}

	if s.OnClick != nil {
		_v.o.Set("onClick", s.OnClick.OnClick)
	}

	if s.Ref != nil {
		_v.o.Set("ref", s.Ref.Ref)
	}

	_v.Role = s.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = s.Style.hack()

}
