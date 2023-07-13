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
	BrowsingModel        string `json:"browsing_model"`
	CodeInterpreterModel string `json:"code_interpreter_model"`
	PluginsModel         string `json:"plugins_model"`
}

type ModelProductFeaturesAttachments struct {
	Type string `json:"type"`
}

type ModelProductFeatures struct {
	Attachments ModelProductFeaturesAttachments `json:"attachments,omitempty"`
}

type ModelListItem struct {
	Slug                  string                         `json:"slug"`
	MaxTokens             int                            `json:"max_tokens"`
	Title                 string                         `json:"title"`
	Description           string                         `json:"description"`
	Tags                  []string                       `json:"tags"`                   // "beta", "confidential", "alpha"
	QualitativeProperties ModelListQualitativeProperties `json:"qualitative_properties"` // removed 0713, keep a version for now
	Capabilities          struct{}                       `json:"capabilities"`           // added 0713
	ProductFeatures       ModelProductFeatures           `json:"product_features"`       // added 0713
	EnabledTools          []string                       `json:"enabled_tools,omitempty"`
}

type ModelListQualitativeProperties struct {
	Reasoning   []int `json:"reasoning,omitempty"`
	Speed       []int `json:"speed,omitempty"`
	Conciseness []int `json:"conciseness,omitempty"`
}

// Models represents the list of models available to the user

// discard 23.07.13
// var MODEL_TEXT_DAVINCI_002_PLUGINS = ModelListItem{
// 	Slug:                  "text-davinci-002-plugins",
// 	MaxTokens:             8195,
// 	Title:                 "Plugins",
// 	Description:           "An experimental model that knows when and how to use plugins",
// 	Tags:                  []string{"alpha"},
// 	QualitativeProperties: ModelListQualitativeProperties{},
// 	EnabledTools:          []string{"tools3"},
// }

var MODEL_TEXT_DAVINCI_002_RENDER_SHA = ModelListItem{
	Slug:        "text-davinci-002-render-sha",
	MaxTokens:   8191,
	Title:       "Default (GPT-3.5)",
	Description: "Our fastest model, great for most everyday tasks.",
	Tags:        []string{"gpt3.5"},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{3, 5},
		Speed:       []int{5, 5},
		Conciseness: []int{2, 5},
	},
}

var MODEL_TEXT_DAVINCI_002_RENDER_SHA_MOBILE = ModelListItem{
	Slug:        "text-davinci-002-render-sha-mobile",
	MaxTokens:   8191,
	Title:       "Default (GPT-3.5) (Mobile)",
	Description: "Our fastest model, great for most everyday tasks.",
	Tags:        []string{"mobile", "gpt3.5"},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{3, 5},
		Speed:       []int{5, 5},
		Conciseness: []int{2, 5},
	},
}

var MODEL_TEXT_DAVINCI_002_RENDER_SHA_CATEGORY = ModelsCategory{
	Category:             "gpt_3.5",
	HumanCategoryName:    "GPT-3.5",
	SubscriptionLevel:    "free",
	DefaultModel:         "text-davinci-002-render-sha",
	BrowsingModel:        "text-davinci-002-render-sha-browsing",
	CodeInterpreterModel: "text-davinci-002-render-sha-code-interpreter",
	PluginsModel:         "text-davinci-002-render-sha-plugins",
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
	Tags:        []string{"gpt4"},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{4, 5},
	},
}

var MODEL_GPT4_BROWSING = ModelListItem{
	Slug:                  "gpt-4-browsing",
	MaxTokens:             4095,
	Title:                 "Web Browsing",
	Description:           "An experimental model that knows when and how to browse the internet",
	Tags:                  []string{"gpt4", "beta"},
	QualitativeProperties: ModelListQualitativeProperties{},
	EnabledTools:          []string{"tools"},
}

var MODEL_GPT4_CODE = ModelListItem{
	Slug:        "gpt-4-code-interpreter",
	MaxTokens:   8192,
	Title:       "Code Interpreter",
	Description: "An experimental model that can solve tasks by generating Python code and executing it in a Jupyter notebook.\nYou can upload any kind of file, and ask model to analyse it, or produce a new file which you can download.",
	Tags:        []string{"gpt4", "beta"},
	ProductFeatures: ModelProductFeatures{
		Attachments: ModelProductFeaturesAttachments{
			Type: "code_interpreter",
		},
	},
	EnabledTools: []string{"tools2"},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{4, 5},
	},
}

var MODEL_GPT4_PLUGIN = ModelListItem{
	Slug:                  "gpt-4-plugins",
	MaxTokens:             8192,
	Title:                 "Plugins",
	Description:           "An experimental model that knows when and how to use plugins",
	Tags:                  []string{"gpt4", "beta"},
	QualitativeProperties: ModelListQualitativeProperties{},
	EnabledTools:          []string{"tools3"},
}

var MODEL_GPT4_MOBILE = ModelListItem{
	Slug:        "gpt-4-mobile",
	MaxTokens:   4095,
	Title:       "GPT-4 (Mobile, V2)",
	Description: "Our most capable model, great for tasks that require creativity and advanced reasoning.",
	Tags:        []string{"gpt4", "mobile"},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{4, 5},
	},
}

var MODEL_GPT4_CATEGORY = ModelsCategory{
	Category:             "gpt_4",
	HumanCategoryName:    "GPT-4",
	SubscriptionLevel:    "plus",
	DefaultModel:         "gpt-4",
	BrowsingModel:        "gpt-4-browsing",
	CodeInterpreterModel: "gpt-4-code-interpreter",
	PluginsModel:         "gpt-4-plugins",
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

var MODEL_OTHER = ModelListItem{
	Slug:        "other",
	MaxTokens:   4095,
	Title:       "Other",
	Description: "Our most capable model, great for tasks that require creativity and advanced reasoning.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{2, 5},
		Conciseness: []int{4, 5},
	},
}

var MODEL_OTHER_CATEGORY = ModelsCategory{
	BrowsingModel:        "other",
	Category:             "other",
	CodeInterpreterModel: "other-code-interpreter",
	DefaultModel:         "other",
	HumanCategoryName:    "other",
	PluginsModel:         "other-plugins",
	SubscriptionLevel:    "free",
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

var MODEL_CLAUDE = ModelListItem{
	Slug:        "claude",
	MaxTokens:   8192,
	Title:       "Claude",
	Description: "Claude is a next-generation AI assistant based on Anthropic's research into training helpful, honest, and harmless AI systems.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{3, 5},
		Conciseness: []int{4, 5},
	},
}

var MODEL_GITHUB_TOP = ModelListItem{
	Slug:        "github-top",
	MaxTokens:   1024,
	Title:       "GitHub Top",
	Description: "GitHub Top is Another demo API example.",
	Tags:        []string{},
	QualitativeProperties: ModelListQualitativeProperties{
		Reasoning:   []int{5, 5},
		Speed:       []int{3, 5},
		Conciseness: []int{4, 5},
	},
}
