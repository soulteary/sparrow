package datatypes

type Models struct {
	Models     []ModelListItem  `json:"models"`
	Categories []ModelsCategory `json:"categories,omitempty"`
}

type ModelsCategory struct {
	Category             string `json:"category"`
	HumanCategoryName    string `json:"human_category_name"`
	SubscriptionLevel    string `json:"subscription_level"`
	DefaultModel         string `json:"default_model"`
	BrowsingModel        any    `json:"browsing_model"`
	CodeInterpreterModel any    `json:"code_interpreter_model"`
	PluginsModel         any    `json:"plugins_model"`
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
	MaxTokens:   8191,
	Title:       "Default (GPT-3.5)",
	Description: "Our fastest model, great for most everyday tasks.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{3, 5},
		Speed:       []int{5, 5},
		Conciseness: []int{2, 5},
	},
}

var MODEL_TEXT_DAVINCI_002_RENDER_SHA_CATEGORY = ModelsCategory{
	BrowsingModel:        nil,
	Category:             "gpt_3.5",
	CodeInterpreterModel: nil,
	DefaultModel:         "text-davinci-002-render-sha",
	HumanCategoryName:    "GPT-3.5",
	PluginsModel:         nil,
	SubscriptionLevel:    "free",
}

// discard 23.05.14
// var MODEL_TEXT_DAVINCI_002_RENDER_PAID = ModelListItem{
// 	Slug:        "text-davinci-002-render-paid",
// 	MaxTokens:   4097,
// 	Title:       "Legacy (GPT-3.5)",
// 	Description: "The previous ChatGPT Plus model",
// 	Tags:        []string{},
// 	QualitativeProperties: ModelListQualitativeProperties{
// 		Reasoning:   []int{3, 5},
// 		Speed:       []int{2, 5},
// 		Conciseness: []int{1, 5},
// 	},
// }

var MODEL_GPT4 = ModelListItem{
	Slug:        "gpt-4",
	MaxTokens:   4095,
	Title:       "GPT-4",
	Description: "Our most capable model, great for tasks that require creativity and advanced reasoning.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{4, 5},
	},
}

var MODEL_GPT4_CATEGORY = ModelsCategory{
	BrowsingModel:        "gpt-4-browsing",
	Category:             "gpt_4",
	CodeInterpreterModel: "gpt-4-code-interpreter",
	DefaultModel:         "gpt-4",
	HumanCategoryName:    "GPT-4",
	PluginsModel:         "gpt-4-plugins",
	SubscriptionLevel:    "plus",
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

var MODEL_FLAGSTUDIO = ModelListItem{
	Slug:        "flag-studio",
	MaxTokens:   1000,
	Title:       "FlagStudio",
	Description: "FlagStudio is a text-to-image platform developed by BAAI's z-lab and FlagAI team.\n\nIt supports 18-language text-to-image generation including Chinese and English, and aims to provide advanced AI art creation experience.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{4, 5},
		Speed:       []int{4, 5},
		Conciseness: []int{3, 5},
	},
}
