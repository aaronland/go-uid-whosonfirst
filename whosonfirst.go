package whosonfirst

import (
	"context"
	"fmt"
	"github.com/aaronland/go-artisanal-integers"
	_ "github.com/aaronland/go-brooklynintegers-api"
	"github.com/aaronland/go-uid"
)

const WHOSONFIRST_SCHEME string = "whosonfirst"

func init() {
	ctx := context.Background()
	uid.RegisterProvider(ctx, WHOSONFIRST_SCHEME, NewWhosOnFirstProvider)
}

type WhosOnFirstProvider struct {
	uid.Provider
	client artisanalinteger.Client
}

type WhosOnFirstUID struct {
	uid.UID
	id int64
}

func NewWhosOnFirstProvider(ctx context.Context, uri string) (uid.Provider, error) {

	client, err := artisanalinteger.NewClient(ctx, "brooklynintegers://")

	if err != nil {
		return nil, fmt.Errorf("Failed to create artisanal integer client, %w", err)
	}

	pr := &WhosOnFirstProvider{
		client: client,
	}

	return pr, nil
}

func (pr *WhosOnFirstProvider) UID(ctx context.Context, args ...interface{}) (uid.UID, error) {
	return NewWhosOnFirstUID(ctx, pr.client)
}

func NewWhosOnFirstUID(ctx context.Context, args ...interface{}) (uid.UID, error) {

	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid arguments")
	}

	cl, ok := args[0].(artisanalinteger.Client)

	if !ok {
		return nil, fmt.Errorf("Invalid client")
	}

	i, err := cl.NextInt()

	if err != nil {
		return nil, fmt.Errorf("Failed to create new integerm %w", err)
	}

	u := &WhosOnFirstUID{
		id: i,
	}

	return u, nil
}

func (u *WhosOnFirstUID) Value() any {
	return u.id
}

func (u *WhosOnFirstUID) String() string {
	return fmt.Sprintf("%v", u.Value())
}
