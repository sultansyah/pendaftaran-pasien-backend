package register

import "context"

type RegisterService interface {
	GetAll(ctx context.Context)
}
