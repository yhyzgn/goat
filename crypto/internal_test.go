package crypto

import "testing"

func TestBase64Encoder_Encode(t *testing.T) {
	be := new(Base64Encoder)
	src := []byte("我试试a")
	bys := be.Encode(src)
	t.Log(string(bys))
	t.Log(string(new(Base64Decoder).Decode(bys)))
	t.Log("======================================")

	bue := new(Base64URLEncoder)
	bys = bue.Encode(src)
	t.Log(string(bys))
	t.Log(string(new(Base64URLDecoder).Decode(bys)))
	t.Log("======================================")

	he := new(HexEncoder)
	bys = he.Encode(src)
	t.Log(string(bys))
	t.Log(string(new(HexDecoder).Decode(bys)))
}
