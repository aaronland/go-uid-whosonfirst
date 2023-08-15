package whosonfirst

import (
	_ "github.com/aaronland/go-brooklynintegers-api"
	_ "github.com/aaronland/go-uid-artisanal"	
)

import (
	"context"

	"github.com/aaronland/go-uid"
)

const WHOSONFIRST_SCHEME string = "whosonfirst"

func init() {
	ctx := context.Background()
	uid.RegisterProvider(ctx, WHOSONFIRST_SCHEME, NewWhosOnFirstProvider)
}

func NewWhosOnFirstProvider(ctx context.Context, uri string) (uid.Provider, error) {
	return uid.NewProvider(ctx, "brooklynintegers://")
}
