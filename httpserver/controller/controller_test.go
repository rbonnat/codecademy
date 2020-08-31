package controller

import (
    "testing"
)

func TestHandleInsertPic(t *testing.T) {
    tests := [] struct {
        name string
        endpoint string
        bucketName string
        file string
        ID string
    }{
        {
            "Success",
            "",
            "",
            "",
                "",
        },
    }

    for _, test := range tests {
        fileStore := &MockFileStore{}
        fileStore.On("Get", test.endpoint, test.bucketName).Return(test.file)

        dbStore := &MockDBStore{}
        dbStore.On("Get", test.endpoint).Return(test.file)

    }
}