package jwt

type Payload struct {
	Issuer     string `json:"iss,omitempty"`
	Subject    string `json:"sub,omitempty"`
	Audience   string `json:"aud,omitempty"`
	Expiration int64  `json:"exp,omitempty"`
	NotBefore  int64  `json:"nbf,omitempty"`
	IssuedAt   int64  `json:"iat,omitempty"`
	JwtId      string `json:"jti,omitempty"`
}
