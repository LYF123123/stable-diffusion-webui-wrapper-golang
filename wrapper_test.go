package wrapper

import "testing"

func TestText2Imgapi(t *testing.T) {
	c:=NewWrapperClient()
	c.SetAPIUrl("http://192.168.1.99:7860")
	result,err:=c.Text2Imgapi("maltese puppy")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(result)
}
