package feed_service

import (
	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"

	feedRepository "github.com/AlexRojasB/go-mongoAtlas-connection.git/repositories/feed.repository"
)

func Create(feed m.Feed) (interface{}, error) {
	id, err := feedRepository.Create(feed)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func Read(skipAmount int) (m.Feeds, error) {
	feeds, err := feedRepository.Read(skipAmount)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

func ReadFromOwner(ownerId string) (m.Feeds, error) {
	feeds, err := feedRepository.ReadAllFromOwner(ownerId)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

func Update(feed m.Feed) error {
	err := feedRepository.UpdateCommentsOrLikes(feed)
	if err != nil {
		return err
	}
	return nil
}

func Delete(feedId string) error {
	err := feedRepository.Delete(feedId)
	if err != nil {
		return err
	}
	return nil
}
