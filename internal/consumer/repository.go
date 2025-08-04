package consumer

// import (
// 	"errors"
// 	"xyz-finance/internal/model"
// )

// type Repository interface {
// 	Save(c model.Consumer) error
// 	FindByID(id string) (*model.Consumer, error)
// 	FindByNIK(nik string) (*model.Consumer, error)
// }

// type memoryRepo struct {
// 	storage map[string]model.Consumer
// }

// func NewInMemoryRepo() Repository {
// 	return &memoryRepo{
// 		storage: make(map[string]model.Consumer),
// 	}
// }

// func (r *memoryRepo) Save(c model.Consumer) error {
// 	r.storage[c.ID] = c
// 	return nil
// }

// func (r *memoryRepo) FindByID(id string) (*model.Consumer, error) {
// 	consumer, ok := r.storage[id]
// 	if !ok {
// 		return nil, errors.New("consumer not found")
// 	}
// 	return &consumer, nil
// }

//	func (r *memoryRepo) FindByNIK(nik string) (*model.Consumer, error) {
//		for _, c := range r.storage {
//			if c.NIK == nik {
//				return &c, nil
//			}
//		}
//		return nil, errors.New("not found")
//	}
