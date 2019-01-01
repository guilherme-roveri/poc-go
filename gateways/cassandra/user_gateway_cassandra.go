package cassandra

import (
	"fmt"
	"log"

	"github.com/gbroveri/users/domains"
	"github.com/gbroveri/users/gateways"
	"github.com/gocql/gocql"
	uuid "github.com/nu7hatch/gouuid"
)

type userGatewayCassandra struct {
	session *gocql.Session
}

//NewUserGatewayCassandra returns new cassandra gw for user
func NewUserGatewayCassandra(s *gocql.Session) gateways.UserGateway {
	return &userGatewayCassandra{session: s}
}

func (g *userGatewayCassandra) Save(u domains.User) (domains.User, error) {
	var user = u
	if u2, err := uuid.NewV4(); err != nil {
		fmt.Println("Error while inserting User")
		fmt.Println(err)
		return domains.User{}, err
	} else {
		user.ID = u2.String()
	}
	if err := g.session.Query("INSERT INTO users(id, username, email, password, identifiers) VALUES(?, ?, ?, ?, ?)",
		user.ID, user.Username, user.Email, user.Password, user.Identifiers).Exec(); err != nil {
		fmt.Println("Error while inserting User")
		fmt.Println(err)
		return domains.User{}, err
	}
	return user, nil
}

func (g *userGatewayCassandra) Get(id string) (domains.User, error) {
	var user = domains.User{}
	if err := g.session.Query(`SELECT id, email, username, password, identifiers FROM users WHERE id = ? LIMIT 1`,
		id).Consistency(gocql.One).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Identifiers); err != nil {
		log.Fatal(err)
		return user, err
	} else {
		return user, nil
	}
}
