package i3line

type Block struct {
	Name                string `json:"name"`
	Instance            string `json:"instance"`
	FullText            string `json:"full_text"`
	ShortText           string `json:"short_text"`
	Color               string `json:"color"`
	Background          string `json:"background"`
	Border              string `json:"border"`
	MinWidth            int    `json:"min_width"`
	Align               string `json:"align"`
	Urgent              bool   `json:"urgent"`
	Separator           bool   `json:"separator"`
	SeparatorBlockWidth int    `json:"separator_block_width"`
	Markup              string `json:"markup"`
}

func NewDefaultBlock(name, instance, text string) Block {
	return Block{
		name,
		instance,
		text,
		text,
		"#ffffff",
		"#000000",
		"#000000",
		25,
		"left",
		false,
		true,
		9,
		"none"}
}

func NewErrorBlock(name, instance, text string) Block {
	return Block{
		name,
		instance,
		text,
		text,
		"#ff0000",
		"#000000",
		"#000000",
		25,
		"left",
		false,
		true,
		9,
		"none"}
}
