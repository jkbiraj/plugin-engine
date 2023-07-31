package windowsserver

import (
	"reflect"
	"testing"
)

func Test_service_GetWindowsServerData(t *testing.T) {
	type fields struct {
		config           Config
		isInitialRequest bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    MockWindowsServerData
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
			got, err := s.GetWindowsServerData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWindowsServerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWindowsServerData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
