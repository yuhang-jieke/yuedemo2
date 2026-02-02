package pkg

import "testing"

func TestUpload(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test", args{address: "C:\\Users\\ZhuanZ\\OneDrive\\图片\\9444703483_1852610764.400x400.jpg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Upload(tt.args.address)
		})
	}
}
