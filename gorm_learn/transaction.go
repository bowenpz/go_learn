package main

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var ErrRollback = errors.New("rollback transaction")

type ConcurrentTrx struct {
	Rollback   bool
	ExecSeq    []int
	TrxSQLList [][]func(tx *gorm.DB) error
}

func (c *ConcurrentTrx) AddSQL(trxIndex int, sql func(tx *gorm.DB) error) {
	if len(c.TrxSQLList) == 0 {
		c.TrxSQLList = make([][]func(tx *gorm.DB) error, 10)
	}
	c.ExecSeq = append(c.ExecSeq, trxIndex)
	c.TrxSQLList[trxIndex] = append(c.TrxSQLList[trxIndex], sql)
}

func (c *ConcurrentTrx) Execute() {
	chanList := make([]chan int, 0)
	for i := 0; i < len(c.TrxSQLList); i++ {
		chanList = append(chanList, make(chan int, 1))
	}

	wg := sync.WaitGroup{}
	doneChan := make(chan int, 1)
	for trxIndex, sqlList := range c.TrxSQLList {
		if len(sqlList) != 0 {
			wg.Add(1)

			sqlList := sqlList
			trxIndex := trxIndex
			go func() {
				err := db.Transaction(func(tx *gorm.DB) error {
					for sqlIndex, sqlFunc := range sqlList {
						<-chanList[trxIndex]

						err := sqlFunc(tx)
						if err != nil {
							return err
						}
						fmt.Printf("Success to exec SQL, trx index: %d, sql index: %d\n", trxIndex, sqlIndex+1)
						doneChan <- 1
					}

					if c.Rollback {
						return ErrRollback
					}
					return nil
				})
				if err != nil && err != ErrRollback {
					if IsDeadlock(err) {
						fmt.Printf("\n\nA deadlock has occurred!")
						PrintLatestDeadlock()
					}
					panic(fmt.Errorf("failed to exec transaction, trx index: %d, err:\n%s", trxIndex, err))
				}
				wg.Done()
			}()
		}
	}

	for _, seq := range c.ExecSeq {
		chanList[seq] <- 1
		select {
		case <-doneChan:
		case <-time.After(time.Second):
			fmt.Printf("\nwait sql time out, jump it.\n")
		}
	}
	wg.Wait()
}

func IsDeadlock(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1213
}
