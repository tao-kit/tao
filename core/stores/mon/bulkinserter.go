package mon

import (
	"context"
	"time"

	"github.com/sllt/tao/core/executors"
	"github.com/sllt/tao/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	flushInterval = time.Second
	maxBulkRows   = 1000
)

type (
	// ResultHandler is a handler that used to handle results.
	ResultHandler func(*mongo.InsertManyResult, error)

	// A BulkInserter is used to insert bulk of mongo records.
	BulkInserter struct {
		executor *executors.PeriodicalExecutor
		inserter *dbInserter
	}
)

// Deprecated. Use NewBatchInserter instead.
func NewBulkInserter(coll *mongo.Collection, interval ...time.Duration) *BulkInserter {
	return newBulkInserter(coll, interval...)
}

// NewBatchInserter returns a BulkInserter.
func NewBatchInserter(coll Collection, interval ...time.Duration) (*BulkInserter, error) {
	cloneColl, err := coll.Clone()
	if err != nil {
		return nil, err
	}

	return newBulkInserter(cloneColl, interval...), nil
}

// newBulkInserter returns a BulkInserter.
func newBulkInserter(coll *mongo.Collection, interval ...time.Duration) *BulkInserter {
	inserter := &dbInserter{
		collection: coll,
	}

	duration := flushInterval
	if len(interval) > 0 {
		duration = interval[0]
	}

	return &BulkInserter{
		executor: executors.NewPeriodicalExecutor(duration, inserter),
		inserter: inserter,
	}
}

// Flush flushes the inserter, writes all pending records.
func (bi *BulkInserter) Flush() {
	bi.executor.Flush()
}

// Insert inserts doc.
func (bi *BulkInserter) Insert(doc interface{}) {
	bi.executor.Add(doc)
}

// SetResultHandler sets the result handler.
func (bi *BulkInserter) SetResultHandler(handler ResultHandler) {
	bi.executor.Sync(func() {
		bi.inserter.resultHandler = handler
	})
}

type dbInserter struct {
	collection    *mongo.Collection
	documents     []interface{}
	resultHandler ResultHandler
}

func (in *dbInserter) AddTask(doc interface{}) bool {
	in.documents = append(in.documents, doc)
	return len(in.documents) >= maxBulkRows
}

func (in *dbInserter) Execute(objs interface{}) {
	docs := objs.([]interface{})
	if len(docs) == 0 {
		return
	}

	result, err := in.collection.InsertMany(context.Background(), docs)
	if in.resultHandler != nil {
		in.resultHandler(result, err)
	} else if err != nil {
		logx.Error(err)
	}
}

func (in *dbInserter) RemoveAll() interface{} {
	documents := in.documents
	in.documents = nil
	return documents
}
