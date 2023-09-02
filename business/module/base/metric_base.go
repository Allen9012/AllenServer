package base

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type MetricsBase struct {
	Name string
}

func (m *MetricsBase) GetName() string {
	return m.Name
}

func (m *MetricsBase) SetName(str string) {
	m.Name = str
}
