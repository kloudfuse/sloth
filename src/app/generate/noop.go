package generate

import (
	"context"

	"github.com/kloudfuse/sloth/src/alert"
	"github.com/kloudfuse/sloth/src/info"
	"github.com/kloudfuse/sloth/src/prometheus"
	"github.com/prometheus/prometheus/model/rulefmt"
)

type noopSLIRecordingRulesGenerator bool

const NoopSLIRecordingRulesGenerator = noopSLIRecordingRulesGenerator(false)

func (noopSLIRecordingRulesGenerator) GenerateSLIRecordingRules(ctx context.Context, slo prometheus.SLO, alerts alert.MWMBAlertGroup) ([]rulefmt.Rule, error) {
	return nil, nil
}

type noopMetadataRecordingRulesGenerator bool

const NoopMetadataRecordingRulesGenerator = noopMetadataRecordingRulesGenerator(false)

func (noopMetadataRecordingRulesGenerator) GenerateMetadataRecordingRules(ctx context.Context, info info.Info, slo prometheus.SLO, alerts alert.MWMBAlertGroup) ([]rulefmt.Rule, error) {
	return nil, nil
}

type noopSLOAlertRulesGenerator bool

const NoopSLOAlertRulesGenerator = noopSLOAlertRulesGenerator(false)

func (noopSLOAlertRulesGenerator) GenerateSLOAlertRules(ctx context.Context, slo prometheus.SLO, alerts alert.MWMBAlertGroup) ([]rulefmt.Rule, error) {
	return nil, nil
}
