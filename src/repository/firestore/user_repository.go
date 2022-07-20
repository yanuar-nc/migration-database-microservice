package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"github.com/yanuar-nc/migration-database-microservice/src/domain"
)

// Repository struct
type Repository struct {
	client *firestore.Client
}

// NewRepository function
func NewRepository(client *firestore.Client) *Repository {
	return &Repository{client: client}
}

func (l *Repository) Save(ctx context.Context, data *domain.User) error {
	_, _, err := l.client.Collection("test").Add(ctx, &domain.UserDetail{
		Form: domain.UserDetailForm{
			Personal: *data,
		},
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *Repository) Update(ctx context.Context, data *domain.User) error {
	return nil
}

func (l *Repository) FetchAll(ctx context.Context, filter domain.Filter) ([]domain.User, error) {

	iter, err := l.client.Collection("test").
		Where("partner.data.name", "==", "alhamsya10000").
		Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for _, s := range iter {
		var user domain.UserDetail
		err = toStruct(s.Data(), &user)
		if err != nil {
			return nil, err
		}
		user.Form.Personal.ID = user.ID
		user.Form.Personal.CreatedAt = s.CreateTime
		users = append(users, user.Form.Personal)
	}

	return users, nil
}

func (l *Repository) MigrationUpdate(ctx context.Context, data *domain.Migration) error {
	return nil
}

func (l *Repository) MigrationGet(ctx context.Context) (*domain.Migration, error) {
	return &domain.Migration{}, nil
}

func (l *Repository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return &domain.User{}, nil
}

func toStruct(v map[string]interface{}, r interface{}) (err error) {
	ms, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "map",
		Result:  &r,
	})
	if err != nil {
		return err
	}
	err = ms.Decode(v)
	if err != nil {
		return err
	}
	return
}
