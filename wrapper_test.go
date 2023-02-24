package wrapper

import "testing"

func TestText2Img(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.Text2Img("maltese puppy")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetMemory(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetMemory()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetSdModels(t *testing.T){
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetSdModels()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetPromptStyles(t *testing.T){
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetPromptStyles()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
