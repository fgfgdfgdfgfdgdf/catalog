package gift

import (
	"testing"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func Test_validateQuery(t *testing.T) {
	type args struct {
		q *entity.GiftQuery
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateQuery(tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("validateQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
