package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Error struct {
	Error string `json:"error"`
}

const (
	keyContentType       = "Content-Type"
	valueApplicationJson = "application/json"

	linuxServerVirtualDevice   = "linuxServer"
	windowsServerVirtualDevice = "windowsServer"
	routerVirtualDevice        = "router"
	switchVirtualDevice        = "switch"
	firewallVirtualDevice      = "firewall"
	loadBalancerVirtualDevice  = "loadBalancer"
)

// FetchVirtualDeviceData gives the data corresponding to virtual device for valid request
func (h *handler) FetchVirtualDeviceData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(keyContentType, valueApplicationJson)
	if !r.URL.Query().Has(h.queryParamVirtualDevice) {
		logrus.Errorf("Empty value for virtualDevice")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Error{
				Error: "Empty value for virtualDevice"},
		)
		return
	}
	virtualDevice := r.URL.Query().Get(h.queryParamVirtualDevice)
	defer finishRequest(virtualDevice, time.Now())
	switch virtualDevice {
	case linuxServerVirtualDevice:
		logrus.Debugf("Virtual device is: %s", virtualDevice)
		linuxServerData, err := h.PluginService.FetchLinuxServerData()
		if err != nil {
			logrus.Errorf("Internal error occurred for: %s" + virtualDevice)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				Error{
					Error: fmt.Sprintf("Internal server error: %v", err)},
			)
			return
		}
		logrus.Infof("Request processed success for: %s", virtualDevice)
		handleResponse(w, linuxServerData)
		return

	case windowsServerVirtualDevice:
		logrus.Debugf("Virtual device is: %s", virtualDevice)
		windowsServerData, err := h.PluginService.FetchWindowsServerData()
		if err != nil {
			logrus.Errorf("Internal error occurred for: %s" + virtualDevice)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				Error{
					Error: fmt.Sprintf("Internal server error: %v", err)},
			)
			return
		}
		logrus.Infof("Request processed success for: %s", virtualDevice)
		handleResponse(w, windowsServerData)
		return

	case routerVirtualDevice:
		logrus.Debugf("Virtual device is: %s", virtualDevice)
		routerData, err := h.PluginService.FetchRouterData()
		if err != nil {
			logrus.Errorf("Internal error occurred for: %s" + virtualDevice)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				Error{
					Error: fmt.Sprintf("Internal server error: %v", err)},
			)
			return
		}
		logrus.Infof("Request processed success for: %s", virtualDevice)
		handleResponse(w, routerData)
		return

	case switchVirtualDevice:
		logrus.Debugf("Virtual device is: %s", virtualDevice)
		switchData, err := h.PluginService.FetchSwitchData()
		if err != nil {
			logrus.Errorf("Internal error occurred for: %s" + virtualDevice)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				Error{
					Error: fmt.Sprintf("Internal server error: %v", err)},
			)
			return
		}
		logrus.Infof("Request processed success for: %s", virtualDevice)
		handleResponse(w, switchData)
		return

	case firewallVirtualDevice:
		logrus.Debugf("Virtual device is: %s", virtualDevice)
		fireWallData, err := h.PluginService.FetchFirewallData()
		if err != nil {
			logrus.Errorf("Internal error occurred for: %s" + virtualDevice)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				Error{
					Error: fmt.Sprintf("Internal server error: %v", err)},
			)
			return
		}
		logrus.Infof("Request processed success for: %s", virtualDevice)
		handleResponse(w, fireWallData)
		return

	case loadBalancerVirtualDevice:
		logrus.Debugf("Virtual device is: %s", virtualDevice)
		loadBalancerData, err := h.PluginService.FetchLoadBalancerData()
		if err != nil {
			logrus.Errorf("Internal error occurred for: %s" + virtualDevice)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				Error{
					Error: fmt.Sprintf("Internal server error: %v", err)},
			)
			return
		}
		logrus.Infof("Request processed success for: %s", virtualDevice)
		handleResponse(w, loadBalancerData)
		return

	default:
		logrus.Debugf("Virtual device is: %s", virtualDevice)
		logrus.Errorf("Request process failure for :%s", virtualDevice)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			Error{
				Error: "Invalid value for virtualDevice"},
		)
		return
	}

}

// handleResponse handles the response
func handleResponse(w http.ResponseWriter, virtualDeviceData interface{}) {
	responseJSON, err := json.Marshal(virtualDeviceData)
	if err != nil {
		logrus.Errorf("Internal Server Error")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the JSON response to the client
	w.Write(responseJSON)
}

// finishRequest finish the request by adding additional logger
func finishRequest(virtualDevice string, start time.Time) {
	logrus.Infof("Finishing http request for: %s, elapsedTimeMS: %v", virtualDevice, start)
}
