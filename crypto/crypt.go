package crypto

// Crypt ...
type Crypt struct {
	mode    Mode
	padding Padding
	encoder Encoder
	decoder Decoder
	Encryptor
	Decryptor
}

// ModeECB ...
func (c *Crypt) ModeECB() {
	c.mode = ECB
}

// ModeCBC ...
func (c *Crypt) ModeCBC() {
	c.mode = CBC
}

// ModeCTR ...
func (c *Crypt) ModeCTR() {
	c.mode = CTR
}

// ModeOFB ...
func (c *Crypt) ModeOFB() {
	c.mode = OFB
}

// ModeCFB ...
func (c *Crypt) ModeCFB() {
	c.mode = CFB
}

// PaddingNo ...
func (c *Crypt) PaddingNo() {
	c.padding = No
}

// PaddingZero ...
func (c *Crypt) PaddingZero() {
	c.padding = Zero
}

// PaddingPKCS1 ...
func (c *Crypt) PaddingPKCS1() {
	c.padding = PKCS1
}

// PaddingPKCS5 ...
func (c *Crypt) PaddingPKCS5() {
	c.padding = PKCS5
}

// PaddingPKCS7 ...
func (c *Crypt) PaddingPKCS7() {
	c.padding = PKCS7
}

// PaddingISO10126 ...
func (c *Crypt) PaddingISO10126() {
	c.padding = ISO10126
}

// PaddingOAEP ...
func (c *Crypt) PaddingOAEP() {
	c.padding = OAEP
}

// PaddingSSL3 ...
func (c *Crypt) PaddingSSL3() {
	c.padding = SSL3
}

// Encoder ...
func (c *Crypt) Encoder(encoder Encoder) {
	c.encoder = encoder
}

// Decoder ...
func (c *Crypt) Decoder(decoder Decoder) {
	c.decoder = decoder
}
