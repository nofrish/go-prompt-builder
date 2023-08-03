package go_prompt_builder

import (
	"fmt"
	"testing"
)

func TestPromptTemplate(t *testing.T) {

	exampleSelector := NewExampleSelector()
	exampleSelector.AddOneExample("example1", "this is the first example")
	exampleSelector.AddOneExample("example2", "this is the second example")
	exampleSelector.AddOneExample("example3", "this is the third example")

	template := NewPromptTemplate("Hello {name}, welcome to {place}! \ncheck out these examples: \n{examples}")
	template.WithExampleSelector(exampleSelector)

	args := map[string]interface{}{
		"name":  "John",
		"place": "the jungle",
	}

	prompt, err := template.Format(args)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(prompt)
}
