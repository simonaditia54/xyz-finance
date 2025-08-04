package transaction

// import (
// 	"errors"

// 	"xyz-finance/internal/model"
// )

// type Repository interface {
// 	GetLimit(consumerID string, tenor int) (*model.Limit, error)
// 	UpdateLimit(limit *model.Limit) error
// 	SaveTransaction(tx model.Transaction) error
// 	FindAll() []model.Transaction
// }

// type memoryRepo struct {
// 	limits       map[string]*model.Limit      // key: consumerID_tenor
// 	transactions map[string]model.Transaction // key: transaction ID
// }

// func NewInMemoryRepo() Repository {
// 	return &memoryRepo{
// 		limits:       make(map[string]*model.Limit),
// 		transactions: make(map[string]model.Transaction),
// 	}
// }

// func (r *memoryRepo) GetLimit(consumerID string, tenor int) (*model.Limit, error) {
// 	key := consumerID + "_" + string(rune(tenor))
// 	limit, ok := r.limits[key]
// 	if !ok {
// 		return nil, errors.New("limit tidak ditemukan")
// 	}
// 	return limit, nil
// }

// func (r *memoryRepo) UpdateLimit(limit *model.Limit) error {
// 	key := limit.ConsumerID + "_" + string(rune(limit.TenorMonth))
// 	r.limits[key] = limit
// 	return nil
// }

// func (r *memoryRepo) SaveTransaction(tx model.Transaction) error {
// 	r.transactions[tx.ID] = tx
// 	return nil
// }

// func (r *memoryRepo) FindAll() []model.Transaction {
// 	var list []model.Transaction
// 	for _, tx := range r.transactions {
// 		list = append(list, tx)
// 	}
// 	return list
// }
