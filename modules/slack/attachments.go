// Package slack contains struct for build payload of slack
package slack

// Field struct
type Field struct {
	title string `json:"title,omitempty"`
	value string `json:"value,omitempty"`
	short bool   `json:"short,omitempty"`
}

// Action struct
type Action struct {
	Name    string `json:"Name,omitempty"`
	Text    string `json:"Text,omitempty"`
	Type    string `json:"Type,omitempty"`
	Value   string `json:"Value,omitempty"`
	Style   string
	Confirm struct {
		Title       string `json:"title,omitempty"`
		Text        string `json:"text,omitempty"`
		OkText      string `json:"ok_text,omitempty"`
		DismissText string `json:"dismiss_text,omitempty"`
	}
}

// Attachment struct
type Attachment struct {
	Fallback   string    `json:"fallback,omitempty"`
	Color      string    `json:"color,omitempty"`
	Pretext    string    `json:"pretext,omitempty"`
	AuthorName string    `json:"author_name,omitempty"`
	AuthorLink string    `json:"author_link,omitempty"`
	AuthorIcon string    `json:"author_icon,omitempty"`
	Title      string    `json:"title,omitempty"`
	TitleLink  string    `json:"title_link,omitempty"`
	Text       string    `json:"text,omitempty"`
	Fields     []*Field  `json:"fields,omitempty"`
	Actions    []*Action `json:"actions,omitempty"`
	ImageURL   string    `json:"image_url,omitempty"`
	ThumbURL   string    `json:"thumb_url,omitempty"`
	Footer     string    `json:"footer,omitempty"`
	FooterIcon string    `json:"footer_icon,omitempty"`
	Ts         int32     `json:"ts,omitempty"`
}
