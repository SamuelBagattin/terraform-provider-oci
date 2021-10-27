// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TODO:  Code Partitioning - Keep only one DefinedTagsToMap() function and remove the other one after all services are partitioned
func DefinedTagsToMap(definedTags map[string]map[string]interface{}) map[string]interface{} {
	var tags = make(map[string]interface{})
	if len(definedTags) > 0 {
		for namespace, keys := range definedTags {
			for key, value := range keys {
				tags[namespace+"."+key] = value
			}
		}
	}
	return tags
}

func definedTagsToMap(definedTags map[string]map[string]interface{}) map[string]interface{} {
	var tags = make(map[string]interface{})
	if len(definedTags) > 0 {
		for namespace, keys := range definedTags {
			for key, value := range keys {
				tags[namespace+"."+key] = value
			}
		}
	}
	return tags
}

// TODO:  Code Partitioning - Keep only one MapToDefinedTags() function and remove the other one after all services are partitioned
func MapToDefinedTags(rawMap map[string]interface{}) (map[string]map[string]interface{}, error) {
	definedTags := make(map[string]map[string]interface{})
	if len(rawMap) > 0 {
		for key, value := range rawMap {
			var keyComponents = strings.Split(key, ".")
			if len(keyComponents) != 2 {
				return nil, fmt.Errorf("invalid key structure found %s", key)
			}
			var namespace = keyComponents[0]
			if _, ok := definedTags[namespace]; !ok {
				definedTags[namespace] = make(map[string]interface{})
			}
			definedTags[namespace][keyComponents[1]] = value
		}
	}
	return definedTags, nil
}

func mapToDefinedTags(rawMap map[string]interface{}) (map[string]map[string]interface{}, error) {
	definedTags := make(map[string]map[string]interface{})
	if len(rawMap) > 0 {
		for key, value := range rawMap {
			var keyComponents = strings.Split(key, ".")
			if len(keyComponents) != 2 {
				return nil, fmt.Errorf("invalid key structure found %s", key)
			}
			var namespace = keyComponents[0]
			if _, ok := definedTags[namespace]; !ok {
				definedTags[namespace] = make(map[string]interface{})
			}
			definedTags[namespace][keyComponents[1]] = value
		}
	}
	return definedTags, nil
}

// TODO:  Code Partitioning - Keep only one DefinedTagsDiffSuppressFunction() function and remove the other one after all services are partitioned
func DefinedTagsDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old != "" && new != "" {
		return false
	}

	// Find the specific defined_tag key name (mainly if a resource supports tagging at multiple levels)
	// For example: "create_vnic_details.0.defined_tags.mynamespace.mykey" => "create_vnic_details.0.defined_tags"
	keyParts := strings.Split(key, ".")
	definedTagKeyParts := []string{}
	for _, keyPart := range keyParts {
		definedTagKeyParts = append(definedTagKeyParts, keyPart)
		if strings.EqualFold(keyPart, "defined_tags") {
			break
		}
	}

	//Old value comes from refreshed state, while new value comes from config
	oldRaw, newRaw := d.GetChange(strings.Join(definedTagKeyParts, "."))
	if newRaw == nil || oldRaw == nil {
		return false
	}

	newValue, newValueOk := newRaw.(map[string]interface{})
	oldValue, oldValueOk := oldRaw.(map[string]interface{})
	if !newValueOk || !oldValueOk {
		return false
	}

	lowerCaseNewValueMap := toLowerCaseKeyMap(newValue)
	lowerCaseOldValueMap := toLowerCaseKeyMap(oldValue)

	if reflect.DeepEqual(lowerCaseOldValueMap, lowerCaseNewValueMap) {
		return true
	}
	return false
}

func definedTagsDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old != "" && new != "" {
		return false
	}

	// Find the specific defined_tag key name (mainly if a resource supports tagging at multiple levels)
	// For example: "create_vnic_details.0.defined_tags.mynamespace.mykey" => "create_vnic_details.0.defined_tags"
	keyParts := strings.Split(key, ".")
	definedTagKeyParts := []string{}
	for _, keyPart := range keyParts {
		definedTagKeyParts = append(definedTagKeyParts, keyPart)
		if strings.EqualFold(keyPart, "defined_tags") {
			break
		}
	}

	//Old value comes from refreshed state, while new value comes from config
	oldRaw, newRaw := d.GetChange(strings.Join(definedTagKeyParts, "."))
	if newRaw == nil || oldRaw == nil {
		return false
	}

	newValue, newValueOk := newRaw.(map[string]interface{})
	oldValue, oldValueOk := oldRaw.(map[string]interface{})
	if !newValueOk || !oldValueOk {
		return false
	}

	lowerCaseNewValueMap := toLowerCaseKeyMap(newValue)
	lowerCaseOldValueMap := toLowerCaseKeyMap(oldValue)

	if reflect.DeepEqual(lowerCaseOldValueMap, lowerCaseNewValueMap) {
		return true
	}
	return false
}

func toLowerCaseKeyMap(original map[string]interface{}) map[string]interface{} {
	lowercaseKeyMap := make(map[string]interface{}, len(original))
	for key, value := range original {
		lowercaseKeyMap[strings.ToLower(key)] = value
	}
	return lowercaseKeyMap
}

func systemTagsToMap(systemTags map[string]map[string]interface{}) map[string]interface{} {
	return definedTagsToMap(systemTags)
}

func mapToSystemTags(rawMap map[string]interface{}) (map[string]map[string]interface{}, error) {
	systemTags := make(map[string]map[string]interface{})
	if len(rawMap) > 0 {
		for key, value := range rawMap {
			var keyComponents = strings.Split(key, ".")
			if len(keyComponents) != 2 {
				return nil, fmt.Errorf("invalid key structure found %s", key)
			}
			var namespace = keyComponents[0]
			if _, ok := systemTags[namespace]; !ok {
				systemTags[namespace] = make(map[string]interface{})
			}
			systemTags[namespace][keyComponents[1]] = value
		}
	}
	return systemTags, nil
}
