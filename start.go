package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"entdemo/ent"
	"entdemo/ent/car"
	"entdemo/ent/group"
	"entdemo/ent/lineuser"
	"entdemo/ent/user"
	"entdemo/linebot/route"

	// "entdemo/linebot/route"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
)

func main() {
	r := gin.Default()

	client, err := ent.Open("postgres", "host=localhost port=6789 user=teerapat dbname=teerapat password=admin1234 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()
	// ctx := context.Background()
	
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r.POST("/callback", route.HandlerReply)
	r.Run(":7777")
}

// ------------------ FUNCTION --------------------

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("tss").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
 
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("tss")).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)

	return u, nil
}

// ----------------------- Example entity relation -------------------------------

func CreateCarsToUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create Car
	tesla, err := client.Car.Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create Car
	ford, err := client.Car.Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// Create User then add two cars to User
	u, err := client.User.Create().
		SetAge(25).
		SetName("arm").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil

}

func QueryCarsFromUser(ctx context.Context, u *ent.User) error {
	cars, err := u.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed query user cars: %w", err)
	}
	log.Println("returend cars:", cars)

	ford, err := u.QueryCars().
		Where(car.Model("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println(ford)
	return nil
}

// Query User(Owner) from Car
func QueryUserFromCar(ctx context.Context, u *ent.User) error {
	cars, err := u.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed queruing user cars: %w", err)
	}

	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying owner car: %w", err)
		}
		log.Printf("car %q owner %q", c.Model, owner.Name)
	}

	return nil
}

// ---------------- Create Graph --------------------------------
func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// Then, create the cars, and attach them to the users created above.
	err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		// Attach this car to Ariel.
		SetOwner(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()).
		// Attach this car to Ariel.
		SetOwner(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		// Attach this car to Neta.
		SetOwner(neta).
		Exec(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers().                // (User(Name=Ariel, Age=30),)
		QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
	a8m := client.User.
		Query().
		Where(
			user.HasCars(),
			user.Name("Ariel"),
		).
		OnlyX(ctx)
	cars, err := a8m.QueryGroups().
		QueryUsers().
		QueryCars().
		Where(
			car.Not(
				car.Model("Mazda"),
			),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}

	log.Println("cars returned:", cars)

	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}

func CreateBuyPayLater(ctx context.Context, client *ent.Client)  {
	lineuser := client.LineUser.Query().
		Where(lineuser.DisplyaName("Dew")).
		OnlyX(ctx)

	// Create CreditLater
	credit_later := client.CreditLater.
		Create().
		SetTransactionRef("14").
		// SetDate(time.Now())
		SetBranch("สมุทรปราการ").
		SetAmount(5000).
		SetInstallment(0).
		SaveX(ctx)
	
	lineuser = lineuser.Update().
		SetCreditlaters(credit_later).
		SaveX(ctx)
	
	log.Println("Result ------------ ", lineuser)
	
}
