// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"sync"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetPairCreatedEventHash() string {
	return "0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9"
}

type PairCreatedEvent struct {
	EventToken0 string
	EventToken1 string
	EventPair   string
	EventArg3   decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:77abbbb1-f51e-4d7c-8e7a-85e9e84712e4,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:77abbbb1-f51e-4d7c-8e7a-85e9e84712e4,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:77abbbb1-f51e-4d7c-8e7a-85e9e84712e4,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetCreatePairMethodHash() string {
	return "c9c65396"
}

type CreatePairMethod struct {
	MethodTokenA string
	MethodTokenB string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:e2151d2a-d2a0-4a6c-886c-c2c798a133c7,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:e2151d2a-d2a0-4a6c-886c-c2c798a133c7,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSetFeeToMethodHash() string {
	return "f46901ed"
}

type SetFeeToMethod struct {
	MethodFeeTo string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:22f6c275-5672-4a27-82f1-9fc5184d5b9b,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:22f6c275-5672-4a27-82f1-9fc5184d5b9b,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSetFeeToSetterMethodHash() string {
	return "a2e74af6"
}

type SetFeeToSetterMethod struct {
	MethodFeeToSetter string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:c75fe3fe-fc9b-4228-bf45-9fc40f2886dd,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:c75fe3fe-fc9b-4228-bf45-9fc40f2886dd,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models
type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

var tokenCache = sync.Map{}

// Config
type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}
type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
}
