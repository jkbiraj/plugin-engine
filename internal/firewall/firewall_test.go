package firewall

import (
	"reflect"
	"testing"
)

func Test_service_GetFirewallData(t *testing.T) {
	type fields struct {
		config           Config
		isInitialRequest bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    MockFirewallData
		wantErr bool
	}{
		{
			name: "firewall data unmarshal error",
			fields: fields{
				config: Config{
					MockFirewallData: "unmarshal failure data",
				},
				isInitialRequest: true,
			},
			want:    MockFirewallData{},
			wantErr: true,
		},
		{
			name: "get firewall data success",
			fields: fields{
				config: Config{
					MockFirewallData: "{\"cpu_utilization\": 10.8}",
				},
				isInitialRequest: true,
			},
			want: MockFirewallData{
				CPUUtilization: 10.8,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				config:           tt.fields.config,
				isInitialRequest: tt.fields.isInitialRequest,
			}
			got, err := s.GetFirewallData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirewallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirewallData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
