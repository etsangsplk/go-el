package el

// Patch contains a group path and value
type Patch map[Expression]interface{}

// Patcher use to patch in memory struct with path
type Patcher struct{}

// PatchIt do patch work
func (p *Patcher) PatchIt(target interface{}, patch Patch) error {

	for path, value := range patch {

		targetValue, err := path.Execute(target)
		if err != nil {
			return err
		}

		err = targetValue.SetValue(value)
		if err != nil {
			return err
		}

	}

	return nil
}
