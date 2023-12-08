package graphqlHandlers

import (
	"Knoxiaes/fairesults/graph/model"
	"context"
)


func CreateResult(ctx context.Context, input model.NewResult) (*model.Result,error){
    return &model.Result{},nil
}
