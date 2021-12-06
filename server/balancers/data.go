package balancers

import (
	"database/sql"
	"fmt"
)

type Balancer struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	UsedMachines []int64 `json:"usedMachines"`
	TotalMachinesCount int64 `json:"totalMachinesCount"`
}

type Machine struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Worked bool `json:"worked"`
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
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Balancer, 0)
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
