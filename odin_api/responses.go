package odin_api

import "time"

type OdinUser struct {
	Principal          string      `json:"principal"`
	Username           string      `json:"username"`
	Bio                interface{} `json:"bio"`
	Image              interface{} `json:"image"`
	Referrer           interface{} `json:"referrer"`
	Admin              bool        `json:"admin"`
	RefCode            string      `json:"ref_code"`
	Profit             interface{} `json:"profit"`
	TotalAssetValue    interface{} `json:"total_asset_value"`
	ReferralEarnings   int         `json:"referral_earnings"`
	ReferralCount      int         `json:"referral_count"`
	AccessAllowed      bool        `json:"access_allowed"`
	BetaAccessCodes    string      `json:"beta_access_codes"`
	BtcDepositAddress  string      `json:"btc_deposit_address"`
	BtcWalletAddress   string      `json:"btc_wallet_address"`
	BlifeID            string      `json:"blife_id"`
	CreatedAt          time.Time   `json:"created_at"`
	RuneDepositAddress string      `json:"rune_deposit_address"`
}

type OdinUserBalance struct {
	Data  []BalanceDetail `json:"data"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
	Count int             `json:"count"`
}

type BalanceDetail struct {
	ID           string      `json:"id"`
	Ticker       string      `json:"ticker"`
	Rune         string      `json:"rune"`
	Name         string      `json:"name"`
	Balance      int         `json:"balance"`
	Image        interface{} `json:"image"`
	Divisibility int         `json:"divisibility"`
	Decimals     int         `json:"decimals"`
	RuneID       string      `json:"rune_id"`
	Trading      bool        `json:"trading"`
	Deposits     bool        `json:"deposits"`
	Withdrawals  bool        `json:"withdrawals"`
}

type OdinFunTokens struct {
	Data  []TokenDetail `json:"data"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Count int           `json:"count"`
}

type TokenDetail struct {
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	Image              string      `json:"image"`
	Creator            string      `json:"creator"`
	CreatedTime        time.Time   `json:"created_time"`
	Volume             int         `json:"volume"`
	Bonded             bool        `json:"bonded"`
	IcrcLedger         string      `json:"icrc_ledger"`
	Price              int         `json:"price"`
	Marketcap          int64       `json:"marketcap"`
	Rune               string      `json:"rune"`
	Featured           bool        `json:"featured"`
	HolderCount        int         `json:"holder_count"`
	HolderTop          int         `json:"holder_top"`
	HolderDev          int         `json:"holder_dev"`
	CommentCount       int         `json:"comment_count"`
	Sold               int64       `json:"sold"`
	Twitter            string      `json:"twitter"`
	Website            string      `json:"website"`
	Telegram           string      `json:"telegram"`
	LastCommentTime    interface{} `json:"last_comment_time"`
	SellCount          int         `json:"sell_count"`
	BuyCount           int         `json:"buy_count"`
	Ticker             string      `json:"ticker"`
	BtcLiquidity       int         `json:"btc_liquidity"`
	TokenLiquidity     int         `json:"token_liquidity"`
	UserBtcLiquidity   int         `json:"user_btc_liquidity"`
	UserTokenLiquidity int         `json:"user_token_liquidity"`
	UserLpTokens       int         `json:"user_lp_tokens"`
	TotalSupply        int64       `json:"total_supply"`
	SwapFees           int         `json:"swap_fees"`
	SwapFees24         int         `json:"swap_fees_24"`
	SwapVolume         int         `json:"swap_volume"`
	SwapVolume24       int         `json:"swap_volume_24"`
	Threshold          int64       `json:"threshold"`
	TxnCount           int         `json:"txn_count"`
	Divisibility       int         `json:"divisibility"`
	Decimals           int         `json:"decimals"`
	Withdrawals        bool        `json:"withdrawals"`
	Deposits           bool        `json:"deposits"`
	Trading            bool        `json:"trading"`
	External           bool        `json:"external"`
	Price5M            int         `json:"price_5m"`
	Price1H            int         `json:"price_1h"`
	Price6H            int         `json:"price_6h"`
	Price1D            int         `json:"price_1d"`
	RuneID             string      `json:"rune_id"`
	LastActionTime     time.Time   `json:"last_action_time"`
	TwitterVerified    bool        `json:"twitter_verified"`
}

type Holders struct {
	Data []struct {
		User         string `json:"user"`
		Token        string `json:"token"`
		Balance      int64  `json:"balance"`
		UserUsername string `json:"user_username"`
		UserImage    string `json:"user_image"`
	} `json:"data"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Count int `json:"count"`
}

type TokenTraders struct {
	Data []struct {
		ID           string      `json:"id"`
		User         string      `json:"user"`
		Token        string      `json:"token"`
		Time         time.Time   `json:"time"`
		Buy          bool        `json:"buy"`
		AmountBtc    int         `json:"amount_btc"`
		AmountToken  int64       `json:"amount_token"`
		Price        int         `json:"price"`
		Bonded       bool        `json:"bonded"`
		UserUsername string      `json:"user_username"`
		UserImage    interface{} `json:"user_image"`
		Decimals     int         `json:"decimals"`
		Divisibility int         `json:"divisibility"`
	} `json:"data"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Count int `json:"count"`
}

type BTCInfo struct {
	ID       int       `json:"id"`
	Symbol   string    `json:"symbol"`
	Datetime time.Time `json:"datetime"`
	Amount   float64   `json:"amount"`
}
