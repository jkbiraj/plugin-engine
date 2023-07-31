package router

import (
	"reflect"
	"testing"
)

func Test_service_GetRouterData(t *testing.T) {
	type fields struct {
		config           Config
		isInitialRequest bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    MockRouterData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				config:           tt.fields.config,
				isInitialRequest: tt.fields.isInitialRequest,
			}
			got, err := s.GetRouterData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRouterData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRouterData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
