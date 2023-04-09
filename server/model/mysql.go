package model

import "gorm.io/gorm"

type MySql struct {
	*gorm.DB
}

func (m *MySql) Create(value interface{}) *gorm.DB {
	return m.DB.Create(value)
}

func (m *MySql) Find(out interface{}, where ...interface{}) *gorm.DB {
	return m.DB.Find(out, where...)
}

func (m *MySql) Where(query interface{}, args ...interface{}) *gorm.DB {
	return m.DB.Where(query, args...)
}

func (m *MySql) First(out interface{}, where ...interface{}) *gorm.DB {
	return m.DB.First(out, where...)
}
