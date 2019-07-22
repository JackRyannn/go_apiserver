package service

//service这一层主要做一些复杂逻辑的处理，是业务层，介于接入层和数据存储层（ORM）之间
import (
	"apiserver/model"
	"apiserver/util"
	"fmt"
	"sync"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}

func ListProduct(name string, offset, limit int) ([]*model.ProductInfo, uint64, error) {
	infos := make([]*model.ProductInfo, 0)
	products, count, err := model.ListProduct(name, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, product := range products {
		ids = append(ids, product.Id)
	}

	wg := sync.WaitGroup{}
	productList := model.ProductList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.ProductInfo, len(products)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range products {
		wg.Add(1)
		go func(u *model.ProductModel) {
			defer wg.Done()

			productList.Lock.Lock()
			defer productList.Lock.Unlock()
			productList.IdMap[u.Id] = &model.ProductInfo{
				Id:          u.Id,
				Name:        u.Name,
				Author:      u.Author,
				Intro:       u.Intro,
				Label:       u.Label,
				Create_Time: u.CreatedAt.Format("2006-01-02 15:04:05"),
				Update_Time: u.UpdatedAt.Format("2006-01-02 15:04:05"),
				State:       u.State,
				Pic:         u.Pic,
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, productList.IdMap[id])
	}

	return infos, count, nil
}

func ListTag(name string, offset, limit int) ([]*model.TagInfo, uint64, error) {
	infos := make([]*model.TagInfo, 0)
	tags, count, err := model.ListTag(name, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, tag := range tags {
		ids = append(ids, tag.Id)
	}

	wg := sync.WaitGroup{}
	tagList := model.TagList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.TagInfo, len(tags)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range tags {
		wg.Add(1)
		go func(u *model.TagModel) {
			defer wg.Done()

			tagList.Lock.Lock()
			defer tagList.Lock.Unlock()
			tagList.IdMap[u.Id] = &model.TagInfo{
				Id:          u.Id,
				Name:        u.Name,
				Source:      u.Source,
				Category:    u.Category,
				Property:    u.Property,
				Create_Time: u.CreatedAt.Format("2006-01-02 15:04:05"),
				Update_Time: u.UpdatedAt.Format("2006-01-02 15:04:05"),
				State:       u.State,
				Operator:    u.Operator,
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, tagList.IdMap[id])
	}

	return infos, count, nil
}
