package _switch

import (
	"reflect"
	"testing"
)

func Test_service_GetSwitchData(t *testing.T) {
	type fields struct {
		config           Config
		isInitialRequest bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    MockSwitchData
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
			got, err := s.GetSwitchData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSwitchData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSwitchData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
