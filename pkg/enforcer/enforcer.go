package enforcer

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/vitorhugoze/enforcer/internal/rules"
)

var ruleEnforcer RuleEnforcer

type RuleEnforcer struct {
	RuleFunctions map[string]func(fieldName, fieldVal string, params ...string)
}

func GetEnforcer() *RuleEnforcer {
	return &ruleEnforcer
}

func SetFaultHandler(faultHandler func(reason string)) {
	rules.FaultHandler = faultHandler
}

func (r *RuleEnforcer) ValidateRules(val any) {
	reflectVal := reflect.ValueOf(val)
	for i := range reflectVal.NumField() {
		fieldVal := reflectVal.Field(i)
		if fieldVal.Kind() == reflect.Struct {
			r.ValidateRules(fieldVal.Interface())
		} else {
			structField := reflectVal.Type().Field(i)
			applyRules(structField.Tag.Get("rules"), structField.Name, fieldVal.String())
		}
	}
}

func applyRules(rawRules string, fieldName, fieldVal string) {
	if len(rawRules) > 0 {
		rules := strings.Split(rawRules, ",")
		for _, r := range rules {
			completeRule := strings.Split(strings.TrimSpace(r), "=")

			ruleName := completeRule[0]
			if ruleEnforcer.RuleFunctions[ruleName] != nil {
				ruleEnforcer.RuleFunctions[ruleName](fieldName, fieldVal, completeRule...)
			} else {
				panic(fmt.Sprintf("%v is not a valid rule", ruleName))
			}
		}
	}
}

func init() {
	ruleEnforcer = RuleEnforcer{}
	ruleEnforcer.RuleFunctions = make(map[string]func(fieldName string, fieldVal string, params ...string))
	ruleEnforcer.RuleFunctions["required"] = rules.RequiredRule
	ruleEnforcer.RuleFunctions["min"] = rules.MinRule
	ruleEnforcer.RuleFunctions["max"] = rules.MaxRule
	ruleEnforcer.RuleFunctions["email"] = rules.EmailRule
	ruleEnforcer.RuleFunctions["password"] = rules.PasswordRule
}
