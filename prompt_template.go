package go_prompt_builder

import (
	"errors"
	"strings"
)

const (
	KeyWordExamples     = "Examples"
	KeyWordOutputFormat = "OutputFormat"
)

type PromptTemplate struct {
	template        string
	exampleSelector *ExampleSelector
}

func NewPromptTemplate(template string) *PromptTemplate {
	return &PromptTemplate{template: template}
}

func FromString(template string) *PromptTemplate {
	return NewPromptTemplate(template)
}

func (t *PromptTemplate) WithExampleSelector(selector *ExampleSelector) *PromptTemplate {
	t.exampleSelector = selector
	return t
}

func (t *PromptTemplate) Format(args map[string]interface{}) (string, error) {

	/*
	 * TODO give more specific error messages
	 */

	sb := strings.Builder{}

	for left := t.template; len(left) > 0; {
		idx := strings.Index(left, "{")
		if idx == -1 {
			sb.WriteString(left)
			break
		}
		sb.WriteString(left[:idx])

		left = left[idx+1:]
		idx = strings.Index(left, "}")
		if idx == -1 {
			return "", errors.New("invalid template, missing right curly brace")
		}
		key := left[:idx]

		// TODO I should find a more elegant way to do this
		if key == KeyWordExamples && t.exampleSelector != nil {
			selected, err := t.exampleSelector.selectFunc(t.exampleSelector.examples)
			if err != nil {
				return "", err
			}
			sb.WriteString(strings.Join(selected, "\n"))
		} else if val, ok := args[key]; ok {
			sb.WriteString(val.(string))
		} else {
			return "", errors.New("mismatched template and key value pairs")
		}

		left = left[idx+1:]
	}

	return sb.String(), nil
}
