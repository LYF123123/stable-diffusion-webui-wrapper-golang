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

func TestGetSdModels(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetSdModels()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetPromptStyles(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetPromptStyles()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetRealesrganModels(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetRealesrganModels()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetFaceRestorers(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetFaceRestorers()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetEmbedding(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetEmbeddings()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetHypernetworks(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetHypernetworks()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetUpscalers(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetUpscalers()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetCmdFlags(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetCmdFlags()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetCurrentUser(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetCurrentUser()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestLoginCheck(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.LoginCheck()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetToken(t *testing.T) {
	c := NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result, err := c.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
