package usecase

import "clean_arch_golang/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	id, err := i.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = i.UserRepository.FindById(id)
	return user, err
}

func (i *UserInteractor) Users() (users domain.Users, err error) {
	users, err = i.UserRepository.FindAll()
	return
}

func (i *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = i.UserRepository.FindById(id)
	return
}
