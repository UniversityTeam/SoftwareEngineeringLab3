package balancers

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type Balancer struct {
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	TotalMachinesCount int64  `json:"totalMachinesCount"`
	UsedMachines       string `json:"usedMachines"`
}

type Machine struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Worked bool   `json:"worked"`
}

type CheckResult struct {
	Exists bool `json:"exists"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListBalancers() ([]*Balancer, error) {
	rows, err := s.Db.Query("SELECT * FROM balancers")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Balancer
	for rows.Next() {
		var c Balancer
		if err := rows.Scan(&c.Id, &c.Name, &c.TotalMachinesCount, &c.UsedMachines); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Balancer, 0)
	}

	for i := 0; i < len(res); i++ {
		var machines []int
		var workedMachines []int
		json.Unmarshal([]byte(res[i].UsedMachines), &machines)

		for j := 0; j < len(machines); j++ {
			check, _ := s.Db.Query("select exists(select 1 from \"machines\" where id=$1 AND worked=true)", machines[j])
			var checkRes []*CheckResult
			for check.Next() {
				var c CheckResult
				if err := check.Scan(&c.Exists); err != nil {
					return nil, err
				}
				checkRes = append(checkRes, &c)
			}

			if checkRes[0].Exists {
				workedMachines = append(workedMachines, machines[j])
			}
		}

		correctData, err := json.Marshal(workedMachines)
		if err != nil {
			return nil, err
		}
		res[i].UsedMachines = string(correctData)
		res[i].TotalMachinesCount = int64(len(machines))

	}

	return res, nil
}

func (s *Store) UpdateMachine(id int64, worked bool) error {
	if id < 0 {
		return fmt.Errorf("Machine id is not provided")
	}
	_, err := s.Db.Exec("UPDATE machines set \"worked\" = $1 WHERE (id) = $2", worked, id)
	return err
}
