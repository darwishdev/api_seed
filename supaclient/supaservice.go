package supaclient

import (
	"context"
	"io"

	"github.com/meloneg/mln_data_pool/supabase"
	// supa "github.com/darwishdev/supabase-go"
)

func (s *SupabaseService) SingIn(c context.Context, req supabase.UserCredentials) (user *supabase.AuthenticatedDetails, err error) {
	user, err = s.Client.Auth.SignIn(c, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *SupabaseService) SignUp(c context.Context, req supabase.UserCredentials) (user *supabase.AuthenticatedDetails, err error) {
	user, err = s.Client.Auth.SignUp(c, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SupabaseService) BucketCreate(c context.Context, req supabase.BucketOption) (bucket *supabase.Bucket, err error) {
	bucket, err = s.Client.Storage.CreateBucket(c, req)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

type BucketUploadRequest struct {
	BucketName string
	Path       string
	Reader     io.Reader
	FileType   string
}

func (s *SupabaseService) BucketUpload(req BucketUploadRequest) (supabase.FileResponse, error) {
	return s.Client.Storage.From(req.BucketName).Upload(req.Path, req.Reader, req.FileType), nil

}
