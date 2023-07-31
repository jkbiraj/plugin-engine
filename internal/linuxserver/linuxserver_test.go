package linuxserver

import (
	"reflect"
	"testing"
)

func Test_service_GetLinuxServerData(t *testing.T) {
	type fields struct {
		config           Config
		isInitialRequest bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    MockLinuxServerData
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
			got, err := s.GetLinuxServerData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLinuxServerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLinuxServerData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
