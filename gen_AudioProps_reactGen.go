// Code generated by reactGen. DO NOT EDIT.

package react

// AudioProps defines the properties for the <Audio> element.
type AudioProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	AutoPlay                string
	Buffered                string
	ClassName               string
	Controls                string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID                     string
	Key                    string
	Loop                   string
	MozCurrentSampleOffset string
	Muted                  string

	OnChange
	OnClick

	Played  string
	Preload string
	Ref
	Role   string
	Src    string
	Style  *CSS
	Volume string
}

func (a *AudioProps) assign(_v *_AudioProps) {

	_v.AriaExpanded = a.AriaExpanded

	_v.AriaHasPopup = a.AriaHasPopup

	_v.AriaLabelledBy = a.AriaLabelledBy

	_v.AutoPlay = a.AutoPlay

	_v.Buffered = a.Buffered

	_v.ClassName = a.ClassName

	_v.Controls = a.Controls

	_v.DangerouslySetInnerHTML = a.DangerouslySetInnerHTML

	if a.DataSet != nil {
		for dk, dv := range a.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if a.ID != "" {
		_v.ID = a.ID
	}

	if a.Key != "" {
		_v.Key = a.Key
	}

	_v.Loop = a.Loop

	_v.MozCurrentSampleOffset = a.MozCurrentSampleOffset

	_v.Muted = a.Muted

	if a.OnChange != nil {
		_v.o.Set("onChange", a.OnChange.OnChange)
	}

	if a.OnClick != nil {
		_v.o.Set("onClick", a.OnClick.OnClick)
	}

	_v.Played = a.Played

	_v.Preload = a.Preload

	if a.Ref != nil {
		_v.o.Set("ref", a.Ref.Ref)
	}

	_v.Role = a.Role

	_v.Src = a.Src

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = a.Style.hack()

	_v.Volume = a.Volume

}
