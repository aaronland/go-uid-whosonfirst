package whosonfirst

import (
	"context"
	"fmt"
	"testing"

	"github.com/aaronland/go-uid"
)

func TestWhosonfirstProvider(t *testing.T) {

	ctx := context.Background()

	uri := fmt.Sprintf("%s://", WHOSONFIRST_SCHEME)

	pr, err := uid.NewProvider(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create provider for %s, %v", uri, err)
	}

	id, err := pr.UID(ctx)

	if err != nil {
		t.Fatalf("Failed to create whosonfirst UID, %v", err)
	}

	if id.String() == "0" {
		t.Fatalf("Invalid ID for whosonfirst provider")
	}

	_, ok := uid.AsInt64(id)

	if !ok {
		t.Fatalf("Expected value to be int64")
	}
}
