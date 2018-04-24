// Code generated by reactGen. DO NOT EDIT.

package react

// LinkProps defines the properties for the <Link> element.
type LinkProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	As                      string
	ClassName               string
	CrossOrigin             string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	Disabled  bool
	Href      string
	HrefLang  string
	ID        string
	Integrity string
	Key       string
	Media     string
	Methods   string

	OnChange
	OnClick

	Prefetch string
	Ref
	ReferrerPolicy string
	Rel            string
	Role           string
	Sizes          string
	Style          *CSS
	Target         string
	Type           string
}

func (l *LinkProps) assign(_v *_LinkProps) {

	_v.AriaExpanded = l.AriaExpanded

	_v.AriaHasPopup = l.AriaHasPopup

	_v.AriaLabelledBy = l.AriaLabelledBy

	_v.As = l.As

	_v.ClassName = l.ClassName

	_v.CrossOrigin = l.CrossOrigin

	_v.DangerouslySetInnerHTML = l.DangerouslySetInnerHTML

	if l.DataSet != nil {
		for dk, dv := range l.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	_v.Disabled = l.Disabled

	_v.Href = l.Href

	_v.HrefLang = l.HrefLang

	if l.ID != "" {
		_v.ID = l.ID
	}

	_v.Integrity = l.Integrity

	if l.Key != "" {
		_v.Key = l.Key
	}

	_v.Media = l.Media

	_v.Methods = l.Methods

	if l.OnChange != nil {
		_v.o.Set("onChange", l.OnChange.OnChange)
	}

	if l.OnClick != nil {
		_v.o.Set("onClick", l.OnClick.OnClick)
	}

	_v.Prefetch = l.Prefetch

	if l.Ref != nil {
		_v.o.Set("ref", l.Ref.Ref)
	}

	_v.ReferrerPolicy = l.ReferrerPolicy

	_v.Rel = l.Rel

	_v.Role = l.Role

	_v.Sizes = l.Sizes

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = l.Style.hack()

	_v.Target = l.Target

	_v.Type = l.Type

}
