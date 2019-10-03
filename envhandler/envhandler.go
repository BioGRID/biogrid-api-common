// Copyright 2019 BioGRID Project. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

// Based on https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f

package envhandler

import (
    "os"
    "strconv"
    "strings"
)

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
		return value
    }

    return defaultVal
}

// Simple helper function to read an environment or return a default value
func GetEnvAsString(key string, defaultVal string) string {
    return getEnv(key,defaultVal)
}

// Simple helper function to read an environment variable into integer or return a default value
func GetEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
		return value
    }

    return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func GetEnvAsBool(name string, defaultVal bool) bool {
    valStr := getEnv(name, "")
    if val, err := strconv.ParseBool(valStr); err == nil {
		return val
    }

    return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
    valStr := getEnv(name, "")

    if valStr == "" {
		return defaultVal
    }

    val := strings.Split(valStr, sep)

    return val
}