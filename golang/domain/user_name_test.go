package domain

import "testing"

func TestNewUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    Username
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				username: "test_user1",
			},
			want:    Username("test_user1"),
			wantErr: false,
		},
		{
			name: "OK16",
			args: args{
				username: "ok16ok16ok16ok16",
			},
			want:    Username("ok16ok16ok16ok16"),
			wantErr: false,
		},
		{
			name: "Err16+1",
			args: args{
				username: "err17err17err17er",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "ErrContain-",
			args: args{
				username: "contain-name",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
