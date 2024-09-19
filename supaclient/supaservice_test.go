package supaclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/meloneg/mln_data_pool/common/convertor"
	"github.com/meloneg/mln_data_pool/common/random"
	"github.com/meloneg/mln_data_pool/supabase"
)

type SigneUpTest struct {
	name      string
	req       *supabase.UserCredentials
	expectErr bool
}

func getValidSignupRequest(field string, value interface{}) *supabase.UserCredentials {
	validRole := &supabase.UserCredentials{
		Email:    random.RandomEmail(),
		Password: random.RandomString(6),
	}
	if field != "" {
		err := convertor.SetField(validRole, field, value)
		if err != nil {
			log.Fatal(err)
		}
	}
	return validRole
}

func TestSignUp(t *testing.T) {
	// Define a slice of test cases
	testcases := []SigneUpTest{
		{
			name:      "ValidUser",
			req:       getValidSignupRequest("", ""),
			expectErr: false,
		},
	}
	fmt.Println("user  ")

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the RoleCreate function with the role data from the current test case
			user, err := service.SignUp(context.Background(), *tc.req)

			fmt.Println("user  " + user.User.ID)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}

			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			//delete all  roles created during test
			// testQueries.RoleDelete(context.Background(), createdRole.RoleID)

		})
	}
}

type BucketCreateTest struct {
	name      string
	req       *supabase.BucketOption
	expectErr bool
}

func getValidBucketCreateRequest(field string, value interface{}) *supabase.BucketOption {
	valid := &supabase.BucketOption{
		Id:     "images2",
		Name:   "images2",
		Public: true,
	}
	if field != "" {
		err := convertor.SetField(valid, field, value)
		if err != nil {
			log.Fatal(err)
		}
	}
	return valid
}

func TestBucketCreate(t *testing.T) {
	// Define a slice of test cases
	testcases := []BucketCreateTest{
		{
			name:      "ValidBucket",
			req:       getValidBucketCreateRequest("", ""),
			expectErr: false,
		},
	}
	fmt.Println("buket  ")

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the RoleCreate function with the role data from the current test case
			bucket, err := service.BucketCreate(context.Background(), *tc.req)

			fmt.Println("buket  " + bucket.Name)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}

			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			//delete all  roles created during test
			// testQueries.RoleDelete(context.Background(), createdRole.RoleID)

		})
	}
}

func TestBucketUpload(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		// Define the path to the image file
		imageFilePath := "./noimg.webp"

		// Open and read the image file
		imageFile, err := os.Open(imageFilePath)
		if err != nil {
			fmt.Printf("Error opening image file: %v\n", err)
			return
		}
		defer imageFile.Close()

		// Read the image file and convert it to an io.Reader
		imageData := &bytes.Buffer{}
		_, err = io.Copy(imageData, imageFile)
		if err != nil {
			fmt.Printf("Error reading image file: %v\n", err)
			return
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		valid := &BucketUploadRequest{
			BucketName: "images2",
			Path:       "./initial/noimg.webp",
			Reader:     imageData,
			FileType:   "image/webp",
		}

		// defer file.Close()
		// Call the RoleCreate function with the role data from the current test case
		f, err := service.BucketUpload(*valid)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		fmt.Println("file  " + f.Key)
		// If the current test case expects an error and no error occurred, fail the test

	})

}
