// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitoring

import (
	"context"
	"encoding/json"
	"fmt"

	"emperror.dev/errors"
	"github.com/antihax/optional"
	"github.com/banzaicloud/banzai-cli/.gen/pipeline"
	"github.com/banzaicloud/banzai-cli/internal/cli"
)

type ActivateManager struct {
	baseManager
}

func (ActivateManager) BuildRequestInteractively(banzaiCLI cli.Cli) (*pipeline.ActivateClusterFeatureRequest, error) {

	grafana, err := askGrafana(banzaiCLI, grafanaAndPrometheusSpec{})
	if err != nil {
		return nil, errors.WrapIf(err, "error during getting Grafana options")
	}

	prometheus, err := askPrometheus(banzaiCLI, grafanaAndPrometheusSpec{})
	if err != nil {
		return nil, errors.WrapIf(err, "error during getting Prometheus options")
	}

	alertmanager, err := askAlertManager(alertmanagerSpec{})
	if err != nil {
		return nil, errors.WrapIf(err, "error during getting Alertmanager options")
	}

	return &pipeline.ActivateClusterFeatureRequest{
		Spec: map[string]interface{}{
			"grafana":      grafana,
			"prometheus":   prometheus,
			"alertmanager": alertmanager,
		},
	}, nil
}

func (ActivateManager) ValidateRequest(req interface{}) error {
	var request pipeline.ActivateClusterFeatureRequest
	if err := json.Unmarshal([]byte(req.(string)), &request); err != nil {
		return errors.WrapIf(err, "request is not valid JSON")
	}

	return validateSpec(request.Spec)
}

func NewActivateManager() *ActivateManager {
	return &ActivateManager{}
}

func askIngress(compType string, defaults baseSpec) (*baseSpec, error) {
	var isEnabled bool
	var isIngressEnabled bool

	var domain string
	var path string

	if err := doQuestions([]questionMaker{
		questionConfirm{
			questionBase: questionBase{
				message: fmt.Sprintf("Do you want to enable %s?", compType),
			},
			defaultValue: defaults.Enabled,
			output:       &isEnabled,
		},
	}); err != nil {
		return nil, errors.WrapIf(err, fmt.Sprintf("error during getting %s enabled", compType))
	}

	if isEnabled {
		if err := doQuestions([]questionMaker{
			questionConfirm{
				questionBase: questionBase{
					message: fmt.Sprintf("Do you want to enable %s Ingress?", compType),
				},
				defaultValue: defaults.Public.Enabled,
				output:       &isIngressEnabled,
			},
		}); err != nil {
			return nil, errors.WrapIf(err, fmt.Sprintf("error during getting %s ingress enabled", compType))
		}

		if isIngressEnabled {

			if err := doQuestions([]questionMaker{
				questionInput{
					questionBase: questionBase{
						message: fmt.Sprintf("Please provide %s Ingress domain:", compType),
						help:    "Leave empty to use cluster's IP",
					},
					defaultValue: defaults.Public.Domain,
					output:       &domain,
				},
				questionInput{
					questionBase: questionBase{
						message: fmt.Sprintf("Please provide %s Ingress path:", compType),
					},
					defaultValue: defaults.Public.Path,
					output:       &path,
				},
			}); err != nil {
				return nil, errors.WrapIf(err, "error during asking ingress fields")
			}

		}
	}
	return &baseSpec{
		Enabled: isEnabled,
		Public: ingressSpec{
			Enabled: isIngressEnabled,
			Domain:  domain,
			Path:    path,
		},
	}, nil
}

func askGrafana(banzaiCLI cli.Cli, defaults grafanaAndPrometheusSpec) (*grafanaAndPrometheusSpec, error) {
	grafanaBase, err := askIngress("Grafana", defaults.baseSpec)
	if err != nil {
		return nil, errors.WrapIf(err, "error during getting Grafana options")
	}

	var secretId string
	if grafanaBase.Enabled {
		secretId, err = askSecret(banzaiCLI, passwordSecretType, defaults.SecretId)
		if err != nil {
			return nil, errors.WrapIf(err, "error during getting Grafana secret")
		}
	}

	return &grafanaAndPrometheusSpec{
		baseSpec: *grafanaBase,
		SecretId: secretId,
	}, nil
}

func askPrometheus(banzaiCLI cli.Cli, defaults grafanaAndPrometheusSpec) (*grafanaAndPrometheusSpec, error) {
	prometheusBase, err := askIngress("Prometheus", defaults.baseSpec)
	if err != nil {
		return nil, errors.WrapIf(err, "error during getting Prometheus options")
	}

	var secretId string
	if prometheusBase.Enabled {
		secretId, err = askSecret(banzaiCLI, passwordSecretType, defaults.SecretId)
		if err != nil {
			return nil, errors.WrapIf(err, "error during getting Grafana secret")
		}
	}

	return &grafanaAndPrometheusSpec{
		baseSpec: *prometheusBase,
		SecretId: secretId,
	}, nil
}

func askAlertManager(defaults alertmanagerSpec) (*alertmanagerSpec, error) {
	alertmanagerBase, err := askIngress("Alertmanager", defaults.baseSpec)
	if err != nil {
		return nil, errors.WrapIf(err, "error during getting Alertmanager options")
	}

	var response = &alertmanagerSpec{baseSpec: *alertmanagerBase}

	if alertmanagerBase.Enabled {
		// Slack
		slackProperties, err := askSlackProperties()
		if err != nil {
			return nil, errors.WrapIf(err, "failed to get Slack options")
		}

		// Pagerduty
		pagerdutyProperties, err := askPagerdutyProperties()
		if err != nil {
			return nil, errors.WrapIf(err, "failed to get Pagerduty options")
		}

		// Email
		emailProperties, err := askEmailProperties()
		if err != nil {
			return nil, errors.WrapIf(err, "failed to get Email options")
		}

		response.Provider = providerSpec{
			Slack:     *slackProperties,
			Pagerduty: *pagerdutyProperties,
			Email:     *emailProperties,
		}
	}

	return response, nil
}

func askSlackProperties() (*slackPropertiesSpec, error) {
	var slackSpec slackPropertiesSpec
	var needSlack bool
	if err := doQuestions([]questionMaker{
		questionConfirm{
			questionBase: questionBase{
				message: "Do you want to enable Slack provider?",
			},
			defaultValue: false,
			output:       &needSlack,
		},
	}); err != nil {
		return nil, errors.WrapIf(err, "error during getting Slack enabled")
	}

	if needSlack {
		var url string
		var channel string
		var sendResolved bool
		if err := doQuestions([]questionMaker{
			questionInput{
				questionBase: questionBase{
					message: "Provide Slack API url:",
				},
				output: &url,
			},
			questionInput{
				questionBase: questionBase{
					message: "Provide Slack channel:",
				},
				output: &channel,
			},
			questionConfirm{
				questionBase: questionBase{
					message: "Send resolved notifications as well",
				},
				defaultValue: false,
				output:       &sendResolved,
			},
		}); err != nil {
			return nil, errors.WrapIf(err, "error during getting Slack options")
		}

		return &slackPropertiesSpec{
			Enabled:      needSlack,
			ApiUrl:       url,
			Channel:      channel,
			SendResolved: sendResolved,
		}, nil
	}

	return &slackSpec, nil
}

func askPagerdutyProperties() (*pagerdutyPropertiesSpec, error) {
	var pagerdutySpec pagerdutyPropertiesSpec
	var needPagerduty bool
	if err := doQuestions([]questionMaker{
		questionConfirm{
			questionBase: questionBase{
				message: "Do you want to enable Pagerduty provider?",
			},
			defaultValue: false,
			output:       &needPagerduty,
		},
	}); err != nil {
		return nil, errors.WrapIf(err, "error during getting Pagerduty enabled")
	}

	if needPagerduty {
		var routingKey string
		var serviceKey string
		var url string
		var sendResolved bool
		if err := doQuestions([]questionMaker{
			questionInput{
				questionBase: questionBase{
					message: "Provide Pagerduty url:",
				},
				output: &url,
			},
			questionInput{
				questionBase: questionBase{
					message: "Provide routing key:",
				},
				output: &routingKey,
			},
			questionInput{
				questionBase: questionBase{
					message: "Provide service key:",
				},
				output: &serviceKey,
			},
			questionConfirm{
				questionBase: questionBase{
					message: "Send resolved notifications as well",
				},
				defaultValue: false,
				output:       &sendResolved,
			},
		}); err != nil {
			return nil, errors.WrapIf(err, "error during getting Slack options")
		}

		return &pagerdutyPropertiesSpec{
			Enabled:      needPagerduty,
			RoutingKey:   routingKey,
			ServiceKey:   serviceKey,
			Url:          url,
			SendResolved: sendResolved,
		}, nil
	}

	return &pagerdutySpec, nil
}

func askEmailProperties() (*emailPropertiesSpec, error) {
	var emailSpec emailPropertiesSpec
	var needEmail bool
	if err := doQuestions([]questionMaker{
		questionConfirm{
			questionBase: questionBase{
				message: "Do you want to enable Email provider?",
			},
			defaultValue: false,
			output:       &needEmail,
		},
	}); err != nil {
		return nil, errors.WrapIf(err, "error during getting Email enabled")
	}

	if needEmail {
		var to string
		var from string
		var sendResolved bool
		if err := doQuestions([]questionMaker{
			questionInput{
				questionBase: questionBase{
					message: "Provide destination of the alert message:",
				},
				output: &to,
			},
			questionInput{
				questionBase: questionBase{
					message: "Provide sender of the alert message:",
				},
				output: &from,
			},
			questionConfirm{
				questionBase: questionBase{
					message: "Send resolved notifications as well",
				},
				defaultValue: false,
				output:       &sendResolved,
			},
		}); err != nil {
			return nil, errors.WrapIf(err, "error during getting Slack options")
		}

		return &emailPropertiesSpec{
			Enabled:      needEmail,
			To:           to,
			From:         from,
			SendResolved: sendResolved,
		}, nil
	}

	return &emailSpec, nil
}

func askSecret(banzaiCLI cli.Cli, secretType, defaultValue string) (string, error) {

	orgID := banzaiCLI.Context().OrganizationID()
	secrets, _, err := banzaiCLI.Client().SecretsApi.GetSecrets(
		context.Background(),
		orgID,
		&pipeline.GetSecretsOpts{
			Type_: optional.NewString(secretType),
		},
	)
	if err != nil {
		return "", errors.WrapIf(err, "failed to get Vault secret(s)")
	}

	if len(secrets) == 0 {
		// TODO (colin): add option to create new secret
		return "", nil
	}

	const skip = "skip"

	var secretName string
	var defaultSecretName = skip
	secretOptions := make([]string, len(secrets)+1)
	secretIds := make(map[string]string, len(secrets))
	secretOptions[0] = skip
	for i, s := range secrets {
		secretOptions[i+1] = s.Name
		secretIds[s.Name] = s.Id
		if s.Id == defaultValue {
			defaultSecretName = s.Name
		}
	}

	if err := doQuestions([]questionMaker{questionSelect{
		questionInput: questionInput{
			questionBase: questionBase{
				message: "Provider secret:",
			},
			defaultValue: defaultSecretName,
			output:       &secretName,
		},
		options: secretOptions,
	}}); err != nil {
		return "", errors.WrapIf(err, "error during getting secret")
	}

	if secretName == skip {
		return "", nil
	}

	return secretIds[secretName], nil
}