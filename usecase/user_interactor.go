package usecase

import "clean_arch_golang/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) (err error) {
	_, err = i.UserRepository.Store(u)
	return err
}

func (i *UserInteractor) Users() (users domain.Users, err error) {
	users, err = i.UserRepository.FindAll()
	return
}

func (i *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = i.UserRepository.FindById(id)
	return
}
