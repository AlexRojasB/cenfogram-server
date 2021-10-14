package user_service

import (
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"

	userRepository "github.com/AlexRojasB/go-mongoAtlas-connection.git/repositories/user.repository"
)

func Create(user m.User) error {
	err := userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func Read(loginUser m.User) (m.User, error) {
	user, err := userRepository.Read(loginUser)
	if err != nil {
		return user, err
	}
	return user, nil
}

func Update(user m.User, userId string) error {
	err := userRepository.Update(user, userId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	err := userRepository.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
