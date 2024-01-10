package entity

// Dimensions is the representation of a width and height.
type Dimensions struct {
	Width  float64
	Height float64
}

// AppendMap appends the dimensions to a map.
func (d *Dimensions) AppendMap(label string, m map[string]interface{}) map[string]interface{} {
	if d.Width != 0 {
		m[label+"_dimension_width"] = d.Width
	}

	if d.Height != 0 {
		m[label+"_dimension_height"] = d.Height
	}

	return m
}
