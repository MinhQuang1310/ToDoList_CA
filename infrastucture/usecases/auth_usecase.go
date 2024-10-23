// // infrastucture/usecases/auth_usecase.go
package usecases

// import (
// 	"context"

// 	models "cleanAchitech/entities"
// 	"cleanAchitech/infrastucture/repository"
// )

// type AuthUseCase struct {
// 	userRepo repository.UserRepository
// }

// func NewAuthUseCase(userRepo repository.UserRepository) *AuthUseCase {
// 	return &AuthUseCase{userRepo: userRepo}
// }

// func (a *AuthUseCase) Register(ctx context.Context, user *models.User) error {
// 	// Thêm logic xử lý đăng ký tại đây
// 	return a.userRepo.Create(ctx, user)
// }

// func (a *AuthUseCase) Login(ctx context.Context, email, password string) (*models.User, error) {
// 	// Thêm logic xử lý đăng nhập tại đây
// 	user, err := a.userRepo.GetByEmail(ctx, email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Kiểm tra mật khẩu và trả về user nếu hợp lệ
// 	return user, nil
// }
