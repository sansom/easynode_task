package service

import "github.com/uduncloud/easynode_task/config"

type Process interface {
	Start()
}

type CreateTask interface {
	GetLastBlockNumber(v *config.BlockConfig) error
}
