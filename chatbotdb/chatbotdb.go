package chatbotdb

import (
	"context"
	"path"

	"github.com/graphql-go/graphql"
	"go.uber.org/zap"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
)

// NewChatBotDB - new chatbot db
func NewChatBotDB(dbpath string, httpAddr string, engine string) (*ankadb.AnkaDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   "chatbotdb",
		Engine: engine,
		PathDB: path.Join(dbpath, "chatbotdb"),
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, newDBLogic())
	if ankaDB == nil {
		jarvisbase.Error("NewChatBotDB", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("NewChatBotDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	return ankaDB, err
}

// chatbotDB -
type chatbotDB struct {
	schema graphql.Schema
}

// newDBLogic -
func newDBLogic() ankadb.DBLogic {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    typeQuery,
			Mutation: typeMutation,
			// Types:    curTypes,
		},
	)

	return &chatbotDB{
		schema: schema,
	}
}

// OnQuery -
func (logic *chatbotDB) OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:         logic.schema,
		RequestString:  request,
		VariableValues: values,
		Context:        ctx,
	})
	// if len(result.Errors) > 0 {
	// 	fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	// }

	return result, nil
}

// OnQueryStream -
func (logic *chatbotDB) OnQueryStream(ctx context.Context, request string, values map[string]interface{}, funcOnQueryStream ankadb.FuncOnQueryStream) error {
	return nil
}
