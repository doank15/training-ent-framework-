package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"training-ent/ent"
	"training-ent/ent/car"
	"training-ent/ent/group"
	"training-ent/ent/user"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_postgresql_db password=go_postgresql_pass sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := entc.Generate("./schema", &gen.Config{}); err != nil {
		log.Fatalf("failed generating ent client: %v", err)
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().SetAge(30).SetName("Daniel").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Query().Where(user.Name("Daniel")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.Create().SetModel("Tesla").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	ford, err := client.Car.Create().SetModel("Ford").SetRegisteredAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	daniel, err := client.User.Create().SetName("Daniel").SetAge(20).AddCars(tesla, ford).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", daniel)
	return daniel, nil
}

func QueryCars(ctx context.Context, u *ent.User) error {
	cars, err := u.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying cars: %w", err)
	}
	log.Println("returned cars: ", cars)

	ford, err := u.QueryCars().Where(car.Model("Ford")).Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying cars: %w", err)
	}
	log.Println("returned car: ", ford)
	return nil
}

func QueryCarUsers(ctx context.Context, u *ent.User) error {
	cars, err := u.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying cars: %w", err)
	}
	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("fa")
		}
		log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetName("Daniel").
		SetAge(30).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	neta, err := client.User.
		Create().
		SetName("Neta").
		SetAge(20).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		// set owner for this car
		SetOwner(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}

	err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		SetOwner(neta).
		Exec(ctx)
	if err != nil {
		return err
	}

	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(neta, a8m).
		Exec(ctx)
	if err != nil {
		return err
	}

	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func QueryGitHub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")).
		QueryUsers().
		QueryCars().
		All(ctx)
	if err != nil {
		return err
	}
	log.Println("returned cars: ", cars)
	return nil
}

func QueryDanielCars(ctx context.Context, client *ent.Client) error {
	daniel := client.User.Query().Where(
		user.HasCars(),
		user.Name("Daniel"),
	).OnlyX(ctx)

	cars, err := daniel.QueryGroups().QueryUsers().QueryCars().Where(car.Not(car.Model("Mazda"))).All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned: ", cars)
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.Query().Where(group.HasUsers()).All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	return nil
}
