package reusable

type AssetType string

const AssetTypeToken AssetType = "token"
const AssetTypeCoin AssetType = "coin"

type TransactionType string

const TransactionTypeBuy TransactionType = "BUY"
const TransactionTypeCoin TransactionType = "SELL"

type AdStatus string

const AdStatusOffline AdStatus = "Offline"
const AdStatusOnline AdStatus = "Online"

type OrderStatus string

const (
	OrderStatusPending         OrderStatus = "Pending"
	OrderStatusAppealed        OrderStatus = "Appealed"
	OrderStatusCanceled        OrderStatus = "Canceled"
	OrderStatusPaymentReceived OrderStatus = "PaymentReceived"
	OrderStatusPaymentSent     OrderStatus = "PaymentSent"
	OrderStatusCompleted       OrderStatus = "Completed"
)
