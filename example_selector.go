package go_prompt_builder

type SelectFunc func(examples map[string]string) (selected []string, err error)

type ExampleSelector struct {
	examples   map[string]string
	selectFunc SelectFunc
}

func NewExampleSelector() *ExampleSelector {
	return &ExampleSelector{
		examples:   make(map[string]string),
		selectFunc: DefaultSelectFunction,
	}
}

// AddOneExample adds one example to the ExampleSelector.
func (s *ExampleSelector) AddOneExample(name, value string) {
	s.examples[name] = value
}

// AddExamples adds multiple examples to the ExampleSelector.
func (s *ExampleSelector) AddExamples(examples map[string]string) {
	for k, v := range examples {
		s.examples[k] = v
	}
}

// DefaultSelectFunction is the default implementation of SelectFunc.
// It will select all examples.
func DefaultSelectFunction(examples map[string]string) (selected []string, err error) {
	for k := range examples {
		selected = append(selected, k)
	}
	return selected, nil
}
