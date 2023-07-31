package routes

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"pluggin-engine/internal/plugin"
	"pluggin-engine/internal/plugin/mocks"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func Test_handler_FetchVirtualDeviceData(t *testing.T) {
	type fields struct {
		Config Config
	}
	type args struct {
		virtualDevice      string
		mockPluginFunction func(mockPlugin *mocks.MockService)
	}
	tests := []struct {
		name       string
		wantStatus int
		fields     fields
		args       args
	}{
		{
			name:       "Invalid test",
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "linux server - FetchLinuxServerData returns error",
			args: args{
				virtualDevice: "linuxServer",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchLinuxServerData().DoAndReturn(func() (plugin.LinuxServerDataResponse, error) {
						return plugin.LinuxServerDataResponse{}, fmt.Errorf("test error")
					})
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "linux server - FetchLinuxServerData returns data",
			args: args{
				virtualDevice: "linuxServer",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchLinuxServerData().DoAndReturn(func() (plugin.LinuxServerDataResponse, error) {
						return plugin.LinuxServerDataResponse{}, nil
					})
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "windows server - FetchWindowsServerData returns error",
			args: args{
				virtualDevice: "windowsServer",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchWindowsServerData().DoAndReturn(func() (plugin.WindowsServerDataResponse, error) {
						return plugin.WindowsServerDataResponse{}, fmt.Errorf("test error")
					})
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "windows server - FetchWindowsServerData returns data",
			args: args{
				virtualDevice: "windowsServer",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchWindowsServerData().DoAndReturn(func() (plugin.WindowsServerDataResponse, error) {
						return plugin.WindowsServerDataResponse{}, nil
					})
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "router - FetchRouterData returns error",
			args: args{
				virtualDevice: "router",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchRouterData().DoAndReturn(func() (plugin.WindowsServerDataResponse, error) {
						return plugin.WindowsServerDataResponse{}, fmt.Errorf("test error")
					})
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "router - FetchRouterData returns data",
			args: args{
				virtualDevice: "router",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchRouterData().DoAndReturn(func() (plugin.RouterDataResponse, error) {
						return plugin.RouterDataResponse{}, nil
					})
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "switch - FetchSwitchData returns error",
			args: args{
				virtualDevice: "switch",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchSwitchData().DoAndReturn(func() (plugin.SwitchDataResponse, error) {
						return plugin.SwitchDataResponse{}, fmt.Errorf("test error")
					})
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "switch - FetchSwitchData returns data",
			args: args{
				virtualDevice: "switch",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchSwitchData().DoAndReturn(func() (plugin.SwitchDataResponse, error) {
						return plugin.SwitchDataResponse{}, nil
					})
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "firewall - FetchSwitchData returns error",
			args: args{
				virtualDevice: "firewall",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchFirewallData().DoAndReturn(func() (plugin.FirewallDataResponse, error) {
						return plugin.FirewallDataResponse{}, fmt.Errorf("test error")
					})
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "firewall - FetchSwitchData returns data",
			args: args{
				virtualDevice: "firewall",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchFirewallData().DoAndReturn(func() (plugin.FirewallDataResponse, error) {
						return plugin.FirewallDataResponse{}, nil
					})
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "loadBalancer - FetchLoadBalancerData returns error",
			args: args{
				virtualDevice: "loadBalancer",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchLoadBalancerData().DoAndReturn(func() (plugin.LoadBalancerDataResponse, error) {
						return plugin.LoadBalancerDataResponse{}, fmt.Errorf("test error")
					})
				},
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "loadBalancer - FetchLoadBalancerData returns data",
			args: args{
				virtualDevice: "loadBalancer",
				mockPluginFunction: func(mockPlugin *mocks.MockService) {
					mockPlugin.EXPECT().FetchLoadBalancerData().DoAndReturn(func() (plugin.LoadBalancerDataResponse, error) {
						return plugin.LoadBalancerDataResponse{}, nil
					})
				},
			},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockPlugin := mocks.NewMockService(mockCtrl)
			server := newTestServer(t, mockPlugin)
			defer server.Close()

			if tt.args.mockPluginFunction != nil {
				tt.args.mockPluginFunction(mockPlugin)
			}

			url := server.URL + fmt.Sprintf("/fetchData?virtualDevice=%s", tt.args.virtualDevice)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			client := server.Client()
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			defer resp.Body.Close()
			_, err = io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if tt.wantStatus != resp.StatusCode {
				t.Fatalf("Status ErrorCode = %v, want %v", resp.StatusCode, tt.wantStatus)
			}
		})
	}
}

func newTestServer(t *testing.T, pluginService plugin.Service) *httptest.Server {
	router := mux.NewRouter() // router for advanced routing capabilities
	config := Config{
		Port:                    ":8080",
		queryParamVirtualDevice: queryParamVirtualDevice,
		PluginService:           pluginService,
	}
	handler := newHandler(config)
	router.HandleFunc("/fetchData", handler.FetchVirtualDeviceData).Methods("GET") // Define routes

	return httptest.NewTLSServer(router)
}
