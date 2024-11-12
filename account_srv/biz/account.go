package biz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"testProject/microservice/account_srv/internal"
	"testProject/microservice/account_srv/model"
	"testProject/microservice/account_srv/proto/pb"
	"testProject/microservice/custom_error"
)

type AccountServer struct {
	pb.AccountServiceServer
}

// Paginate 翻页功能，防止请求过多
func Paginate(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo == 0 {
			pageNo = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNo - 1) * pageSize
		// Offset 偏移量指定在开始返回记录之前要跳过的记录数
		return db.Offset(offset).Limit(pageSize)
	}
}

func Model2Pb(account *model.Account) *pb.AccountRes {
	accountRes := &pb.AccountRes{
		Id:       int32(account.ID),
		Mobile:   account.Mobile,
		Password: account.Password,
		Nickname: account.NickName,
		Gender:   account.Gender,
		Role:     int32(account.Role),
	}
	return accountRes
}

// FindById 根据Id查找数据
func FindById(ctx context.Context, dest interface{}, id int32) *gorm.DB {
	return internal.DB.First(&dest, id)
}

// FindByMobile 根据手机号查找数据
func FindByMobile(ctx context.Context, dest interface{}, mobile string) *gorm.DB {
	return internal.DB.First(&dest, mobile)
}

// GetAccountList 获取用户列表，所有用户信息
func (a *AccountServer) GetAccountList(ctx context.Context, req *pb.PagingRequest) (*pb.AccountListRes, error) {
	var accountList []model.Account
	//result := internal.DB.Find(&accountList)
	//做一个可以分页的版本：
	//作用域将当前数据库连接传递给参数 `func(DB) DB`，可用于动态添加条件
	result := internal.DB.Scopes(Paginate(int(req.PageNo), int(req.PageSize))).Find(&accountList)
	//这拿到的是model的，我们需要返回pb，所以需要一个方法来转换
	if result.Error != nil {
		return nil, result.Error
	}
	accountListRes := &pb.AccountListRes{}
	accountListRes.Total = int32(result.RowsAffected)
	for _, account := range accountList {
		accountRes := Model2Pb(&account)
		accountListRes.AccountList = append(accountListRes.AccountList, accountRes)
	}
	return accountListRes, nil
}

// GetAccountByMobile 根据手机号获取用户信息
func (a *AccountServer) GetAccountByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.AccountRes, error) {
	//用来接收数据
	var account model.Account
	result := internal.DB.Where("mobile = ?", req.Mobile).First(&account)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}
	accountRes := Model2Pb(&account)
	return accountRes, nil
}

// GetAccountById 根据Id获取用户信息
func (a *AccountServer) GetAccountById(ctx context.Context, req *pb.IdRequest) (*pb.AccountRes, error) {
	var account model.Account
	result := FindById(ctx, &account, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}
	accountRes := Model2Pb(&account)
	return accountRes, nil
}

// AddAccount 添加用户信息
func (a *AccountServer) AddAccount(ctx context.Context, req *pb.AddAccountRequest) (*pb.AccountRes, error) {
	var account model.Account
	result := FindByMobile(ctx, &account, req.Mobile)
	if result.RowsAffected != 0 {
		return nil, errors.New(custom_error.AccountIsExist)
	}
	salt, pwd := GetMd5(req.Password)
	account.Mobile = req.Mobile
	account.Password = pwd
	account.NickName = req.NickName
	account.Salt = salt
	account.Gender = req.Gender
	result = internal.DB.Create(&account)
	if result.Error != nil {
		return nil, errors.New(custom_error.InternalError)
	}
	accountRes := Model2Pb(&account)
	return accountRes, nil
}

// UpdateAccount 修改用户信息
func (a *AccountServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountRes, error) {
	var account model.Account
	result := FindById(ctx, &account, int32(req.Id))
	if result.RowsAffected == 0 {
		return nil, errors.New(custom_error.AccountNotFound)
	}
	account.Mobile = req.Mobile
	account.Gender = req.Gender
	account.NickName = req.NickName
	salt, pwd := GetMd5(req.Password)
	account.Salt = salt
	account.Password = pwd
	account.Role = int(req.Role)

	result = internal.DB.Save(&account)
	if result.Error != nil {
		return &pb.UpdateAccountRes{
			Result: false,
		}, errors.New(custom_error.InternalError)
	}
	return &pb.UpdateAccountRes{
		Result: true,
	}, nil
}

// CheckPassword 验证密码是否正确
func (a *AccountServer) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordRes, error) {
	var account model.Account
	result := FindById(ctx, &account, req.AccountId)
	if result.Error != nil {
		return nil, errors.New(custom_error.InternalError)
	}
	if account.Salt == "" {
		return nil, errors.New(custom_error.SaltError)
	}
	verify := Verify(req.Password, account.Salt, account.Password)
	if !verify {
		return &pb.CheckPasswordRes{
			Result: false,
		}, errors.New(custom_error.AccountPasswordError)
	}
	return &pb.CheckPasswordRes{
		Result: true,
	}, nil
}

func (a *AccountServer) mustEmbedUnimplementedAccountServiceServer() {}
