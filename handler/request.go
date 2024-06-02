package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"` //colocou o * pois por default sempre é false, mas preciso saber se a pessoa preencheu, assim aponto para o endereco de memoria para saber se tem algo la
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("Company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("Location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("Link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("Salary", "numerico")
	}
	if r.Remote == nil {
		return errParamIsRequired("Salary", "Boleano")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"` //colocou o * pois por default sempre é false, mas preciso saber se a pessoa preencheu, assim aponto para o endereco de memoria para saber se tem algo la
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *UpdateOpeningRequest) Validate() error {

	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Salary > 0 || r.Remote != nil {
		logger.Errorf("Company: %v", r.Company)
		return nil

	}
	return fmt.Errorf("at least one fild must be provided")
}
