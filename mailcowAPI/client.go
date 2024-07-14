package mailcowapi

type MailcowAPIClient struct {
	Scheme string
	Host   string
	Token  string
}

func NewClient(scheme string, host string, token string) *MailcowAPIClient {
	return &MailcowAPIClient{
		Scheme: scheme,
		Host:   host,
		Token:  token,
	}
}

func (client *MailcowAPIClient) Get(endpoint string, writeTo any) error {
	return nil
}
