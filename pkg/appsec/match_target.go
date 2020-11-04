package appsec

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// MatchTarget represents a collection of MatchTarget
//
// See: MatchTarget.GetMatchTarget()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type (
	// MatchTarget  contains operations available on MatchTarget  resource
	// See: // appsec v1
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getmatchtarget
	MatchTarget interface {
		GetMatchTargets(ctx context.Context, params GetMatchTargetsRequest) (*GetMatchTargetsResponse, error)
		GetMatchTarget(ctx context.Context, params GetMatchTargetRequest) (*GetMatchTargetResponse, error)
		CreateMatchTarget(ctx context.Context, params CreateMatchTargetRequest) (*CreateMatchTargetResponse, error)
		UpdateMatchTarget(ctx context.Context, params UpdateMatchTargetRequest) (*UpdateMatchTargetResponse, error)
		RemoveMatchTarget(ctx context.Context, params RemoveMatchTargetRequest) (*RemoveMatchTargetResponse, error)
	}

	// GetMatchTargetRequest is the argument for GetProperties
	GetMatchTargetRequest struct {
		ConfigID      int `json:"configId"`
		ConfigVersion int `json:"configVersion"`
		TargetID      int `json:"targetId"`
	}

	// GetMatchTargetsRequest is the argument for GetProperties
	GetMatchTargetsRequest struct {
		ConfigID      int `json:"configId"`
		ConfigVersion int `json:"configVersion"`
	}

	// UpdateMatchTargetRequest is the argument for GetProperties
	UpdateMatchTargetRequest struct {
		Type                      string `json:"type"`
		ConfigID                  int    `json:"configId"`
		ConfigVersion             int    `json:"configVersion"`
		DefaultFile               string `json:"defaultFile"`
		EffectiveSecurityControls struct {
			ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
			ApplyBotmanControls           bool `json:"applyBotmanControls"`
			ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
			ApplyRateControls             bool `json:"applyRateControls"`
			ApplyReputationControls       bool `json:"applyReputationControls"`
			ApplySlowPostControls         bool `json:"applySlowPostControls"`
		} `json:"effectiveSecurityControls"`
		Hostnames                    []string `json:"hostnames"`
		IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
		FilePaths                    []string `json:"filePaths"`
		FileExtensions               []string `json:"fileExtensions"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"sequence"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	// RemoveMatchTargetRequest is the argument for GetProperties
	RemoveMatchTargetRequest struct {
		ConfigID      int `json:"configId"`
		ConfigVersion int `json:"configVersion"`
		TargetID      int `json:"targetId"`
	}

	// CreateMatchTargetRequest is the argument for GetProperties
	CreateMatchTargetRequest struct {
		Type                      string `json:"type"`
		ConfigID                  int    `json:"configId"`
		ConfigVersion             int    `json:"configVersion"`
		DefaultFile               string `json:"defaultFile"`
		EffectiveSecurityControls struct {
			ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
			ApplyBotmanControls           bool `json:"applyBotmanControls"`
			ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
			ApplyRateControls             bool `json:"applyRateControls"`
			ApplyReputationControls       bool `json:"applyReputationControls"`
			ApplySlowPostControls         bool `json:"applySlowPostControls"`
		} `json:"effectiveSecurityControls"`
		FileExtensions               []string `json:"fileExtensions"`
		FilePaths                    []string `json:"filePaths"`
		Hostnames                    []string `json:"hostnames"`
		IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"sequence"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	// CreateMatchTargetResponse is the argument for GetProperties
	CreateMatchTargetResponse struct {
		MType                     string `json:"type"`
		ConfigID                  int    `json:"configId"`
		ConfigVersion             int    `json:"configVersion"`
		DefaultFile               string `json:"defaultFile"`
		EffectiveSecurityControls struct {
			ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
			ApplyBotmanControls           bool `json:"applyBotmanControls"`
			ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
			ApplyRateControls             bool `json:"applyRateControls"`
			ApplyReputationControls       bool `json:"applyReputationControls"`
			ApplySlowPostControls         bool `json:"applySlowPostControls"`
		} `json:"effectiveSecurityControls"`
		Hostnames                    []string `json:"hostnames"`
		IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
		FilePaths                    []string `json:"filePaths"`
		FileExtensions               []string `json:"fileExtensions"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"sequence"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	//GetMatchTargetResponse ...
	GetMatchTargetResponse struct {
		Type                      string `json:"type"`
		ConfigID                  int    `json:"configId"`
		ConfigVersion             int    `json:"configVersion"`
		DefaultFile               string `json:"defaultFile"`
		EffectiveSecurityControls struct {
			ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
			ApplyBotmanControls           bool `json:"applyBotmanControls"`
			ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
			ApplyRateControls             bool `json:"applyRateControls"`
			ApplyReputationControls       bool `json:"applyReputationControls"`
			ApplySlowPostControls         bool `json:"applySlowPostControls"`
		} `json:"effectiveSecurityControls"`
		Hostnames                    []string `json:"hostnames"`
		IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
		FilePaths                    []string `json:"filePaths"`
		FileExtensions               []string `json:"fileExtensions"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"sequence"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	// GetMatchTargetResponse ...
	GetMatchTargetsResponse struct {
		MatchTargets struct {
			APITargets []struct {
				ConfigID                  int    `json:"configId"`
				ConfigVersion             int    `json:"configVersion"`
				Sequence                  int    `json:"sequence"`
				TargetID                  int    `json:"targetId"`
				Type                      string `json:"type"`
				EffectiveSecurityControls struct {
					ApplyAPIConstraints           bool `json:"applyApiConstraints"`
					ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
					ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
					ApplyRateControls             bool `json:"applyRateControls"`
					ApplyReputationControls       bool `json:"applyReputationControls"`
					ApplySlowPostControls         bool `json:"applySlowPostControls"`
				} `json:"effectiveSecurityControls"`
				SecurityPolicy struct {
					PolicyID string `json:"policyId"`
				} `json:"securityPolicy"`
				Apis []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"apis"`
				BypassNetworkLists []struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"bypassNetworkLists"`
			} `json:"apiTargets"`
			WebsiteTargets []struct {
				ConfigID                     int           `json:"configId"`
				ConfigVersion                int           `json:"configVersion"`
				DefaultFile                  string        `json:"defaultFile"`
				IsNegativeFileExtensionMatch bool          `json:"isNegativeFileExtensionMatch"`
				IsNegativePathMatch          bool          `json:"isNegativePathMatch"`
				Sequence                     int           `json:"sequence"`
				TargetID                     int           `json:"targetId"`
				Type                         string        `json:"type"`
				FileExtensions               []string      `json:"fileExtensions"`
				FilePaths                    []string      `json:"filePaths"`
				Hostnames                    []interface{} `json:"hostnames"`
				EffectiveSecurityControls    struct {
					ApplyAPIConstraints           bool `json:"applyApiConstraints"`
					ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
					ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
					ApplyRateControls             bool `json:"applyRateControls"`
					ApplyReputationControls       bool `json:"applyReputationControls"`
					ApplySlowPostControls         bool `json:"applySlowPostControls"`
				} `json:"effectiveSecurityControls"`
				SecurityPolicy struct {
					PolicyID string `json:"policyId"`
				} `json:"securityPolicy"`
				BypassNetworkLists []struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"bypassNetworkLists"`
			} `json:"websiteTargets"`
		} `json:"matchTargets"`
	}

	// UpdateMatchTargetResponse ...
	UpdateMatchTargetResponse struct {
		Type                      string `json:"type"`
		ConfigID                  int    `json:"configId"`
		ConfigVersion             int    `json:"configVersion"`
		DefaultFile               string `json:"defaultFile"`
		EffectiveSecurityControls struct {
			ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
			ApplyBotmanControls           bool `json:"applyBotmanControls"`
			ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
			ApplyRateControls             bool `json:"applyRateControls"`
			ApplyReputationControls       bool `json:"applyReputationControls"`
			ApplySlowPostControls         bool `json:"applySlowPostControls"`
		} `json:"effectiveSecurityControls"`
		Hostnames                    []string `json:"hostnames"`
		IsNegativeFileExtensionMatch bool     `json:"isNegativeFileExtensionMatch"`
		IsNegativePathMatch          bool     `json:"isNegativePathMatch"`
		FilePaths                    []string `json:"filePaths"`
		FileExtensions               []string `json:"fileExtensions"`
		SecurityPolicy               struct {
			PolicyID string `json:"policyId"`
		} `json:"securityPolicy"`
		Sequence           int `json:"sequence"`
		TargetID           int `json:"targetId"`
		BypassNetworkLists []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"bypassNetworkLists"`
	}

	// RemoveMatchTargetResponse ...
	RemoveMatchTargetResponse struct {
		MatchTargets struct {
			APITargets []struct {
				ConfigID                  int    `json:"configId"`
				ConfigVersion             int    `json:"configVersion"`
				Sequence                  int    `json:"sequence"`
				TargetID                  int    `json:"targetId"`
				Type                      string `json:"type"`
				EffectiveSecurityControls struct {
					ApplyAPIConstraints           bool `json:"applyApiConstraints"`
					ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
					ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
					ApplyRateControls             bool `json:"applyRateControls"`
					ApplyReputationControls       bool `json:"applyReputationControls"`
					ApplySlowPostControls         bool `json:"applySlowPostControls"`
				} `json:"effectiveSecurityControls"`
				SecurityPolicy struct {
					PolicyID string `json:"policyId"`
				} `json:"securityPolicy"`
				Apis []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"apis"`
				BypassNetworkLists []struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"bypassNetworkLists"`
			} `json:"apiTargets"`
			WebsiteTargets []struct {
				ConfigID                     int           `json:"configId"`
				ConfigVersion                int           `json:"configVersion"`
				DefaultFile                  string        `json:"defaultFile"`
				IsNegativeFileExtensionMatch bool          `json:"isNegativeFileExtensionMatch"`
				IsNegativePathMatch          bool          `json:"isNegativePathMatch"`
				Sequence                     int           `json:"sequence"`
				TargetID                     int           `json:"targetId"`
				Type                         string        `json:"type"`
				FileExtensions               []string      `json:"fileExtensions"`
				FilePaths                    []string      `json:"filePaths"`
				Hostnames                    []interface{} `json:"hostnames"`
				EffectiveSecurityControls    struct {
					ApplyAPIConstraints           bool `json:"applyApiConstraints"`
					ApplyApplicationLayerControls bool `json:"applyApplicationLayerControls"`
					ApplyNetworkLayerControls     bool `json:"applyNetworkLayerControls"`
					ApplyRateControls             bool `json:"applyRateControls"`
					ApplyReputationControls       bool `json:"applyReputationControls"`
					ApplySlowPostControls         bool `json:"applySlowPostControls"`
				} `json:"effectiveSecurityControls"`
				SecurityPolicy struct {
					PolicyID string `json:"policyId"`
				} `json:"securityPolicy"`
				BypassNetworkLists []struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"bypassNetworkLists"`
			} `json:"websiteTargets"`
		} `json:"matchTargets"`
	}

	// BypassNetworkList ...
	BypassNetworkList struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}

	Hostnames struct {
		Hostnames string `json:"hostnames"`
	}
)

// Validate validates GetMatchTargetRequest
func (v GetMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
		"TargetID":      validation.Validate(v.TargetID, validation.Required),
	}.Filter()
}

// Validate validates GetMatchTargetsRequest
func (v GetMatchTargetsRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
	}.Filter()
}

// Validate validates CreateMatchTargetRequest
func (v CreateMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
	}.Filter()
}

// Validate validates UpdateMatchTargetRequest
func (v UpdateMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
		"TargetID":      validation.Validate(v.TargetID, validation.Required),
	}.Filter()
}

// Validate validates RemoveMatchTargetRequest
func (v RemoveMatchTargetRequest) Validate() error {
	return validation.Errors{
		"ConfigID":      validation.Validate(v.ConfigID, validation.Required),
		"ConfigVersion": validation.Validate(v.ConfigVersion, validation.Required),
		"TargetID":      validation.Validate(v.TargetID, validation.Required),
	}.Filter()
}

func (p *appsec) GetMatchTarget(ctx context.Context, params GetMatchTargetRequest) (*GetMatchTargetResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetMatchTarget")

	var rval GetMatchTargetResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets/%d?includeChildObjectName=true",
		params.ConfigID,
		params.ConfigVersion,
		params.TargetID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getmatchtarget request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getproperties request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

func (p *appsec) GetMatchTargets(ctx context.Context, params GetMatchTargetsRequest) (*GetMatchTargetsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("GetMatchTargets")

	var rval GetMatchTargetsResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets",
		params.ConfigID,
		params.ConfigVersion,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getmatchtargets request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getmatchtargets request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

// Update will update a MatchTarget.
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#putmatchtarget

func (p *appsec) UpdateMatchTarget(ctx context.Context, params UpdateMatchTargetRequest) (*UpdateMatchTargetResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("UpdateMatchTarget")

	putURL := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets/%d",
		params.ConfigID,
		params.ConfigVersion,
		params.TargetID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, putURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create create MatchTargetrequest: %w", err)
	}

	var rval UpdateMatchTargetResponse
	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("create MatchTarget request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, p.Error(resp)
	}

	return &rval, nil
}

// Create will create a new matchtarget.
//
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#postmatchtarget
func (p *appsec) CreateMatchTarget(ctx context.Context, params CreateMatchTargetRequest) (*CreateMatchTargetResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	logger := p.Log(ctx)
	logger.Debug("CreateMatchTarget")

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets",
		params.ConfigID,
		params.ConfigVersion,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create create matchtarget request: %w", err)
	}

	var rval CreateMatchTargetResponse

	resp, err := p.Exec(req, &rval, params)
	if err != nil {
		return nil, fmt.Errorf("create matchtargetrequest failed: %w", err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}

// Delete will delete a MatchTarget
//
//
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html#deletematchtarget

func (p *appsec) RemoveMatchTarget(ctx context.Context, params RemoveMatchTargetRequest) (*RemoveMatchTargetResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrStructValidation, err.Error())
	}

	var rval RemoveMatchTargetResponse

	logger := p.Log(ctx)
	logger.Debug("RemoveMatchTarget")

	uri, err := url.Parse(fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/match-targets/%d",
		params.ConfigID,
		params.ConfigVersion,
		params.TargetID,
	),
	)
	if err != nil {
		return nil, fmt.Errorf("failed parse url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create delmatchtarget request: %w", err)
	}

	resp, errd := p.Exec(req, nil)
	if errd != nil {
		logger.Debug("No JSON on DELETE")
		return nil, fmt.Errorf("delmatchtarget request failed: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil
}
