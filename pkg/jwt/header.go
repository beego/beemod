package jwt

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}
