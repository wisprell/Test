package mysql

import "time"

func Select(id string) (*CounterModel, error) {
	db := GetMysql()
	var err error
	var model CounterModel
	err = db.Debug().Table(TableNameCounterModel).
		Where("id = ?", id).Scan(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func CreateLockTable(name string) (*CounterModel, error) {
	db := GetMysql()
	var err error
	tx := db.Begin()
	err = tx.Debug().Exec("lock tables counter_models write").Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	model := CounterModel{
		Name:      name,
		Count:     0,
		CreatedAt: time.Now(),
	}
	err = tx.Debug().Table(TableNameCounterModel).
		Create(&model).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Debug().Exec("unlock tables").Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &model, nil
}

func Create(name string) (*CounterModel, error) {
	db := GetMysql()
	var err error
	model := CounterModel{
		Name:      name,
		Count:     0,
		CreatedAt: time.Now(),
	}
	err = db.Debug().Table(TableNameCounterModel).
		Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func Update(id string) (*CounterModel, error) {
	db := GetMysql()
	var err error
	var model CounterModel
	err = db.Debug().Table(TableNameCounterModel).
		Where("id = ?", id).Scan(&model).Error
	if err != nil {
		return nil, err
	}
	model.Count += 1
	model.UpdatedAt = time.Now()
	err = db.Debug().Table(TableNameCounterModel).
		Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func Delete(id string) error {
	db := GetMysql()
	var err error
	err = db.Debug().Table(TableNameCounterModel).
		Where("id = ?", id).Delete(&CounterModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

