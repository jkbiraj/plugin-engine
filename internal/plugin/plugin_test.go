package plugin

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"pluggin-engine/internal/firewall"
	"pluggin-engine/internal/firewall/mocks"
	"reflect"
	"testing"
)

func Test_service_FetchFirewallData(t *testing.T) {
	type fields struct {
		config                Config
		mockFirewallServiceFn func(mockFirewallService *mocks.MockService)
	}
	tests := []struct {
		name    string
		fields  fields
		want    FirewallDataResponse
		wantErr bool
	}{
		{
			name: "GetFirewallData returns error",
			fields: fields{
				mockFirewallServiceFn: func(mockFirewallService *mocks.MockService) {
					mockFirewallService.EXPECT().GetFirewallData().DoAndReturn(func() (firewall.MockFirewallData, error) {
						return firewall.MockFirewallData{}, fmt.Errorf("test error")
					})
				},
			},
			wantErr: true,
		},
		{
			name: "GetFirewallData returns firewall data",
			fields: fields{
				mockFirewallServiceFn: func(mockFirewallService *mocks.MockService) {
					mockFirewallService.EXPECT().GetFirewallData().DoAndReturn(func() (firewall.MockFirewallData, error) {
						return firewall.MockFirewallData{
							CPUUtilization: 1.2345,
						}, nil
					})
				},
			},
			wantErr: false,
			want: FirewallDataResponse{
				CPUUtilization: 1.2345,
				Interfaces:     map[string]InterfaceStats{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockFirewallService := mocks.NewMockService(mockCtrl)

			if tt.fields.mockFirewallServiceFn != nil {
				tt.fields.mockFirewallServiceFn(mockFirewallService)
			}
			s := &service{
				config:          tt.fields.config,
				fireWallService: mockFirewallService,
			}
			got, err := s.FetchFirewallData()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchFirewallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchFirewallData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
