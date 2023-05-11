package datatypes

type Models struct {
	Models []ModelListItem `json:"models"`
}

type ModelListItem struct {
	Slug                  string                         `json:"slug"`
	MaxTokens             int                            `json:"max_tokens"`
	Title                 string                         `json:"title"`
	Description           string                         `json:"description"`
	Tags                  []string                       `json:"tags"`
	QualitativeProperties ModelListQualitativeProperties `json:"qualitative_properties"`
	EnabledTools          []string                       `json:"enabled_tools,omitempty"`
}

type ModelListQualitativeProperties struct {
	Reasoning   []int `json:"reasoning,omitempty"`
	Speed       []int `json:"speed,omitempty"`
	Conciseness []int `json:"conciseness,omitempty"`
}

// Models represents the list of models available to the user
var MODEL_TEXT_DAVINCI_002_PLUGINS = ModelListItem{
	Slug:                  "text-davinci-002-plugins",
	MaxTokens:             8195,
	Title:                 "Plugins",
	Description:           "An experimental model that knows when and how to use plugins",
	Tags:                  []string{"alpha"},
	QualitativeProperties: ModelListQualitativeProperties{},
	EnabledTools:          []string{"tools3"},
}

var MODEL_TEXT_DAVINCI_002_RENDER_SHA = ModelListItem{
	Slug:        "text-davinci-002-render-sha",
	MaxTokens:   4097,
	Title:       "Default (GPT-3.5)",
	Description: "Optimized for speed, currently available to Plus users",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{3, 5},
		Speed:       []int{5, 5},
		Conciseness: []int{2, 5},
	},
}

var MODEL_TEXT_DAVINCI_002_RENDER_PAID = ModelListItem{
	Slug:        "text-davinci-002-render-paid",
	MaxTokens:   4097,
	Title:       "Legacy (GPT-3.5)",
	Description: "The previous ChatGPT Plus model",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{3, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{1, 5},
	},
}

var MODEL_GPT4 = ModelListItem{
	Slug:        "gpt-4",
	MaxTokens:   4095,
	Title:       "GPT-4",
	Description: "Our most advanced model, available to Plus subscribers.\n\nGPT-4 excels at tasks that require advanced reasoning, complex instruction understanding, and more creativity.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{4, 5},
	},
}

var MODEL_OPENAI_API_3_5 = ModelListItem{
	Slug:        "official-api",
	MaxTokens:   4097,
	Title:       "OpenAI API (GPT-3.5)",
	Description: "The version based on the official interface package has more stable service capabilities.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{3, 5},
		Speed:       []int{5, 5},
		Conciseness: []int{2, 5},
	},
}

var MODEL_NO_MODELS = ModelListItem{
	Slug:        "no-models",
	MaxTokens:   1,
	Title:       "No models",
	Description: "No models available.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{0, 5},
		Speed:       []int{0, 5},
		Conciseness: []int{0, 5},
	},
}

var MODEL_MIDJOURNEY = ModelListItem{
	Slug:        "mid-journey",
	MaxTokens:   1000,
	Title:       "Mid Journey",
	Description: "The drawing model with the best effect at present.\n\nAn artificial intelligence program developed by the research laboratory of the same name, which can generate images based on text, will enter the public testing stage on July 12, 2022. Users can use Discord's robot instructions Operational.\n\nThe research lab is led by David Holz, founder of Leap Motion.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{4, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{3, 5},
	},
}
