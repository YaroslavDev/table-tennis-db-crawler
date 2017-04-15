package rubber

import (
	"upper.io/db.v3/mysql"
)

type RubberRepository interface {
	GetRubbers() ([]Rubber, error)
	SaveRubber(rubber *Rubber) error
}

type rubberRepositoryImpl struct {
	connectionUrl *mysql.ConnectionURL
}

func NewRubberRepository(connectionUrl *mysql.ConnectionURL) RubberRepository {
	return &rubberRepositoryImpl{connectionUrl: connectionUrl}
}

func (r rubberRepositoryImpl) GetRubbers() ([]Rubber, error) {
	session, err := mysql.Open(r.connectionUrl)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var rubbers []Rubber
	err = session.Collection("rubber").Find().OrderBy("name ASC").All(&rubbers)
	if err != nil {
		return nil, err
	}
	return rubbers, nil
}

func (r rubberRepositoryImpl) SaveRubber(rubber *Rubber) error {
	session, err := mysql.Open(r.connectionUrl)
	if err != nil {
		return err
	}
	defer session.Close()

	_, err = session.Exec(
		`INSERT INTO rubber
		(name, speed, spin, control, tackiness, weight, sponge_hardness, gears, throw_angle, consistency, durability)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		name = VALUES(name),
		speed = VALUES(speed),
		spin = VALUES(spin),
		control = VALUES(control),
		tackiness = VALUES(tackiness),
		weight = VALUES(weight),
		sponge_hardness = VALUES(sponge_hardness),
		gears = VALUES(gears),
		throw_angle = VALUES(throw_angle),
		consistency = VALUES(consistency),
		durability = VALUES(durability)
		`, rubber.Name, rubber.Speed, rubber.Spin, rubber.Control, rubber.Tackiness,
		rubber.Weight, rubber.SpongeHardness, rubber.Gears, rubber.ThrowAngle, rubber.Consistency, rubber.Durability)
	return err
}
