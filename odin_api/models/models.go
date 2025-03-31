package models

// AuthRequest 身份验证请求结构
type AuthRequest struct {
	PublicKey string `json:"publickey"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	Referrer  string `json:"referrer"`
}

// AuthToken 身份验证响应结构
type AuthToken struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token       string `json:"token"`
		PrincipalID string `json:"principal_id"`
	} `json:"data"`
}

// CommentRequest 发表评论请求结构
type CommentRequest struct {
	Message string `json:"message"`
}

// BTCInfo 比特币价格信息结构
type BTCInfo struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Price float64 `json:"price"`
	} `json:"data"`
}

// OdinFunTokens Odin.fun令牌列表结构
type OdinFunTokens struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []TokenData `json:"data"`
}

// TokenData 令牌数据结构
type TokenData struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Symbol              string  `json:"symbol"`
	Description         string  `json:"description"`
	LogoURL             string  `json:"logo_url"`
	Price               float64 `json:"price"`
	MarketCap           float64 `json:"marketcap"`
	Supply              float64 `json:"supply"`
	TradingVolume       float64 `json:"trading_volume"`
	TradingVolume24h    float64 `json:"trading_volume_24h"`
	HolderCount         int     `json:"holder_count"`
	LastActionPrice     float64 `json:"last_action_price"`
	LastActionTime      string  `json:"last_action_time"`
	LastActionTimestamp int64   `json:"last_action_timestamp"`
	Website             string  `json:"website"`
	Twitter             string  `json:"twitter"`
	Telegram            string  `json:"telegram"`
	Discord             string  `json:"discord"`
	Medium              string  `json:"medium"`
	CreatorID           string  `json:"creator_id"`
	CreatorName         string  `json:"creator_name"`
	Verified            bool    `json:"verified"`
}

// Holders 持有者列表结构
type Holders struct {
	Status  bool         `json:"status"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []HolderData `json:"data"`
}

// HolderData 持有者数据结构
type HolderData struct {
	PrincipalID    string  `json:"principal_id"`
	Username       string  `json:"username"`
	Amount         float64 `json:"amount"`
	AvatarURL      string  `json:"avatar_url"`
	ProfileBgColor string  `json:"profile_bg_color"`
}

// OdinUser 用户信息结构
type OdinUser struct {
	Status  bool     `json:"status"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}

// UserData 用户数据结构
type UserData struct {
	PrincipalID    string `json:"principal_id"`
	Username       string `json:"username"`
	AvatarURL      string `json:"avatar_url"`
	ProfileBgColor string `json:"profile_bg_color"`
	Followers      int    `json:"followers"`
	Following      int    `json:"following"`
	Verified       bool   `json:"verified"`
	Bio            string `json:"bio"`
	Website        string `json:"website"`
	Twitter        string `json:"twitter"`
	Telegram       string `json:"telegram"`
	Discord        string `json:"discord"`
}

// UserBalances 用户余额列表结构
type UserBalances struct {
	Status  bool          `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []UserBalance `json:"data"`
}

// UserBalance 用户余额结构
type UserBalance struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Symbol    string  `json:"symbol"`
	LogoURL   string  `json:"logo_url"`
	Amount    float64 `json:"amount"`
	Value     float64 `json:"value"`
	Price     float64 `json:"price"`
	CreatorID string  `json:"creator_id"`
}

// TokenTrades 令牌交易列表结构
type TokenTrades struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []TradeData `json:"data"`
}

// TradeData 交易数据结构
type TradeData struct {
	ID          string  `json:"id"`
	TokenID     string  `json:"token_id"`
	TokenSymbol string  `json:"token_symbol"`
	TokenName   string  `json:"token_name"`
	TokenLogo   string  `json:"token_logo"`
	Amount      float64 `json:"amount"`
	Price       float64 `json:"price"`
	Value       float64 `json:"value"`
	Type        string  `json:"type"`
	BuyerID     string  `json:"buyer_id"`
	BuyerName   string  `json:"buyer_name"`
	SellerID    string  `json:"seller_id"`
	SellerName  string  `json:"seller_name"`
	Time        string  `json:"time"`
	Timestamp   int64   `json:"timestamp"`
}

// TokenTarget 用于获取令牌交易的目标结构
type TokenTarget struct {
	Id                  string `json:"id"`
	LastActionTimestamp int64  `json:"last_action_timestamp"`
}
