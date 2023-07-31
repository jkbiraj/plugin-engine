package loadbalancer

import (
	"reflect"
	"testing"
)

func Test_service_GetLoadBalancerData(t *testing.T) {
	type fields struct {
		config           Config
		isInitialRequest bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    MockLoadBalancerData
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
			got, err := s.GetLoadBalancerData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoadBalancerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoadBalancerData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
