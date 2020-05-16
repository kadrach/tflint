// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsWafXSSMatchSetInvalidNameRule checks the pattern is valid
type AwsWafXSSMatchSetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWafXSSMatchSetInvalidNameRule returns new rule with default attributes
func NewAwsWafXSSMatchSetInvalidNameRule() *AwsWafXSSMatchSetInvalidNameRule {
	return &AwsWafXSSMatchSetInvalidNameRule{
		resourceType:  "aws_waf_xss_match_set",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsWafXSSMatchSetInvalidNameRule) Name() string {
	return "aws_waf_xss_match_set_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafXSSMatchSetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafXSSMatchSetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafXSSMatchSetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafXSSMatchSetInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
