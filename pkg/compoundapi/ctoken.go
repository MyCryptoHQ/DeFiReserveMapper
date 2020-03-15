package compoundapi

import (
	"github.com/postables/go-compound/client"
)

func FetchCompoundCTokens(compoundClient *client.Client) (CTokens, error) {
	cTokensApiReturn, err := compoundClient.GetCTokens()
	if err != nil {
		return nil, err
	}
	var cTokens CTokens
	cTokens = cTokensApiReturn.CToken
	return cTokens, nil
}