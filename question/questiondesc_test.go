package question

import "testing"

func TestLeetCodeDesc_WriteDesc(t *testing.T) {
	type fields struct {
		desc string
		code string
	}
	type args struct {
		q    string
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"q1",
			args{"perfect-squares", "./tmp"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := NewLeetCode(tt.args.q)
			if err != nil {
				t.Errorf("NewLeetCode fail err=%v, wantErr %v", err, tt.wantErr)
			}
			if err := l.WriteDesc(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("WriteDesc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLeetCodeDesc_WriteCode(t *testing.T) {
	type fields struct {
		desc string
		code string
	}
	type args struct {
		q    string
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"q1",
			args{"perfect-squares", "./tmp"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := NewLeetCode(tt.args.q)
			if err != nil {
				t.Errorf("NewLeetCode fail err=%v, wantErr %v", err, tt.wantErr)
			}
			if err := l.WriteCode(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("WriteDesc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
