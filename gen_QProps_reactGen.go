// Code generated by reactGen. DO NOT EDIT.

package react

// QProps defines the properties for the <Q> element.
type QProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	Cite                    string
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

func (q *QProps) assign(_v *_QProps) {

	_v.AriaExpanded = q.AriaExpanded

	_v.AriaHasPopup = q.AriaHasPopup

	_v.AriaLabelledBy = q.AriaLabelledBy

	_v.Cite = q.Cite

	_v.ClassName = q.ClassName

	_v.DangerouslySetInnerHTML = q.DangerouslySetInnerHTML

	if q.DataSet != nil {
		for dk, dv := range q.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if q.ID != "" {
		_v.ID = q.ID
	}

	if q.Key != "" {
		_v.Key = q.Key
	}

	if q.OnChange != nil {
		_v.o.Set("onChange", q.OnChange.OnChange)
	}

	if q.OnClick != nil {
		_v.o.Set("onClick", q.OnClick.OnClick)
	}

	if q.Ref != nil {
		_v.o.Set("ref", q.Ref.Ref)
	}

	_v.Role = q.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = q.Style.hack()

}
