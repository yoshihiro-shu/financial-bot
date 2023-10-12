package notification_test

import (
	"testing"

	"github.com/yoshihiro-shu/financial-bot/internal/notification"
)

const accessToken = "inORb1lJ64ajtqaGh12s3UoaZgnmv1eWMZL2o4x4ucN"

func TestSendMsg(t *testing.T) {
	client := notification.NewNotification(accessToken)

	tests := []struct {
		name string
		Args string
		Want error
	}{
		{
			name: "test",
			Args: "test",
			Want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := client.SendMsg(test.Args)
			if err != test.Want {
				t.Errorf("got %v, want %v", err, test.Want)
			}
		})
	}
}

func FuzzSendMsg(f *testing.F) {
	client := notification.NewNotification(accessToken)
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, msg string) {
		err := client.SendMsg(msg)
		if err != nil {
			t.Error(err)
		}
	})
}