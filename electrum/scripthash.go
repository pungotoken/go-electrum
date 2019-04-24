package electrum

// GetBalanceResp represents the response to GetBalance().
type GetBalanceResp struct {
	Result GetBalanceResult `json:"result"`
}

// GetBalanceResult represents the content of the result field in the response to GetBalance().
type GetBalanceResult struct {
	Confirmed   float64 `json:"confirmed"`
	Unconfirmed float64 `json:"unconfirmed"`
}

// GetBalance returns the confirmed and unconfirmed balance for a scripthash.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#blockchain-scripthash-get-balance
func (s *Server) GetBalance(scripthash string) (GetBalanceResult, error) {
	var resp GetBalanceResp

	err := s.request("blockchain.scripthash.get_balance", []interface{}{scripthash}, &resp)
	if err != nil {
		return GetBalanceResult{}, err
	}

	return resp.Result, err
}

// GetMempoolResp represents the response to GetHistory() and GetMempool().
type GetMempoolResp struct {
	Result []*GetMempoolResult `json:"result"`
}

// GetMempoolResult represents the content of the result field in the response
// to GetHistory() and GetMempool().
type GetMempoolResult struct {
	Hash   string `json:"tx_hash"`
	Height int32  `json:"height"`
	Fee    uint32 `json:"fee,omitempty"`
}

// GetHistory returns the confirmed and unconfirmed history for a scripthash.
func (s *Server) GetHistory(scripthash string) ([]*GetMempoolResult, error) {
	var resp GetMempoolResp

	err := s.request("blockchain.scripthash.get_history", []interface{}{scripthash}, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, err
}

// GetMempool returns the unconfirmed transacations of a scripthash.
func (s *Server) GetMempool(scripthash string) ([]*GetMempoolResult, error) {
	var resp GetMempoolResp

	err := s.request("blockchain.scripthash.get_mempool", []interface{}{scripthash}, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, err
}

// ListUnspentResp represents the response to ListUnspent()
type ListUnspentResp struct {
	Result []*ListUnspentResult `json:"result"`
}

// ListUnspentResult represents the content of the result field in the response to ListUnspent()
type ListUnspentResult struct {
	Height   uint32 `json:"height"`
	Position uint32 `json:"tx_pos"`
	Hash     string `json:"tx_hash"`
	Value    uint64 `json:"value"`
}

// ListUnspent returns an ordered list of UTXOs for a scripthash.
func (s *Server) ListUnspent(scripthash string) ([]*ListUnspentResult, error) {
	var resp ListUnspentResp

	err := s.request("blockchain.scripthash.listunspent", []interface{}{scripthash}, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, err
}

// ListUnspentAddress is a deprecated function for listing unspent transactions for a given address.
func (s *Server) ListUnspentAddress(address string) ([]*ListUnspentResult, error) {
	var resp ListUnspentResp

	err := s.request("blockchain.address.listunspent", []interface{}{address}, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, err
}
