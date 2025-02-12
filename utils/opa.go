package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type OPAConnection struct {
	URL string
}

func (conn *OPAConnection) Evaluate(path, input string) (string, error) {
	// Evaluate the input using the OPA connection
	// POST /v1/data/httpapi/authz

	url := fmt.Sprintf("%s%s", conn.URL, path)

	resp, err := http.Post(url, "application/json", strings.NewReader(input))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Read the response body
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (conn *OPAConnection) User(email string) (map[string]any, error) {
	users, err := conn.Users()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user["email"] == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

type TypeValue struct {
	Type  string
	Value string
}

type Package struct {
	Path []TypeValue `json:"path"`
}
type Ast struct {
	Package Package `json:"package"`
}

type Policy struct {
	ID  string `json:"id"`
	Raw string `json:"raw"`
	Ast Ast    `json:"ast"`
}

func (p Policy) Path() string {
	var path []string
	for _, tv := range p.Ast.Package.Path {
		path = append(path, tv.Value)
	}

	return fmt.Sprintf("/v1/%s", strings.Join(path, "/"))
}

func (conn *OPAConnection) Policies() ([]Policy, error) {
	// Get the list of policies from the OPA connection
	// GET /v1/data/policies

	resp, err := http.Get(fmt.Sprintf("%s/v1/policies", conn.URL))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	resultB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := map[string]any{}
	if err := json.Unmarshal(resultB, &result); err != nil {
		return nil, err
	}

	// Extract the policies from the result field
	policiesB, err := json.Marshal(result["result"])
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into a slice of policies
	var policies []Policy
	if err := json.Unmarshal(policiesB, &policies); err != nil {
		return nil, err
	}

	return policies, nil
}

func (conn *OPAConnection) Users() ([]map[string]any, error) {
	// Get the list of users from the OPA connection
	// GET /v1/data/users

	resp, err := http.Get(fmt.Sprintf("%s/v1/data/users", conn.URL))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	usersB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// usersB looks like {"decision_id":"4de5779c-7cab-4e51-b626-0b53f7ac1395","result":{"admin":{"groups":["admin"],"maxAmount":1000},"marcel":{"groups":["admin","user"],"maxAmount":23},"user":{"groups":["user"],"maxAmount":15}}}
	// We need to extract the users from the result field

	// Unmarshal the response body into a map
	var result map[string]any
	if err := json.Unmarshal(usersB, &result); err != nil {
		return nil, err
	}

	// Extract the users from the result field
	usersB, err = json.Marshal(result["result"])
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into a slice of users
	var users []map[string]any
	if err := json.Unmarshal(usersB, &users); err != nil {
		return nil, err
	}

	return users, nil
}
