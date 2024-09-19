package supaclient

import (
	"context"

	"github.com/meloneg/mln_data_pool/supabase"
)

// su "github.com/darwishdev/supabase-go"

type SupabaseServiceInterface interface {
	SignUp(c context.Context, req supabase.UserCredentials) (user *supabase.AuthenticatedDetails, err error)
	BucketCreate(c context.Context, req supabase.BucketOption) (bucket *supabase.Bucket, err error)
	BucketUpload(req BucketUploadRequest) (supabase.FileResponse, error)
	SingIn(c context.Context, req supabase.UserCredentials) (user *supabase.AuthenticatedDetails, err error)
}

type SupabaseService struct {
	Client *supabase.Client
}

func NewSupabaseService(supabaseUrl string, supabaseKey string) (*SupabaseService, error) {
	supa := supabase.CreateClient(supabaseUrl, supabaseKey)

	return &SupabaseService{
		Client: supa,
	}, nil
}
