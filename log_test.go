/**
 * Package zap_log
 * @author zhangjie
 * @email zhangjie@rockontrol.com
 * @date 2021/6/23
 * @description
 */
package zap_log

import (
	"reflect"
	"testing"
)

func TestLogInit(t *testing.T) {
	logger, err := LogInit("test.log", false, "dev")
	if err != nil {
		t.Fatal("init log fail")
	}
	if logger == nil {
		t.Fatal("init log fail")
	}
	if logger.logger == nil {
		t.Fatal("init log fail")
	}

}

func TestLog_Info(t *testing.T) {
	logger, err := LogInit("test.log", true, "dev")
	if err != nil {
		t.Fatal("init log fail")
	}
	if logger == nil {
		t.Fatal("init log fail")
	}
	if logger.logger == nil {
		t.Fatal("init log fail")
	}
	logger.Info(map[string]interface{}{"foo": "bar"})
}

func TestEnv_Enabled(t *testing.T) {
	type args struct {
		env Env
	}
	tests := []struct {
		name string
		e    Env
		args args
		want bool
	}{
		{
			name: "dev",
			e:    Env__Dev,
			args: args{Env__Dev},
			want: true,
		},
		{
			name: "dev",
			e:    Env__Dev,
			args: args{Env__Test},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Enabled(tt.args.env); got != tt.want {
				t.Errorf("Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnv_Get(t *testing.T) {
	tests := []struct {
		name string
		e    Env
		want interface{}
	}{
		{
			name: "dev",
			e:    Env__Dev,
			want: Env__Dev,
		},
		{
			name: "test",
			e:    Env__Test,
			want: Env__Test,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnv_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		e       Env
		want    []byte
		wantErr bool
	}{
		{
			name:    "dev",
			e:       Env__Dev,
			want:    []byte(Env__Dev.String()),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnv_Set(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		e       Env
		args    args
		wantErr bool
	}{
		{
			name:    "dev",
			e:       Env__Dev,
			args:    args{s: "dev"},
			wantErr: false,
		},
		{
			name:    "null",
			e:       Env__Dev,
			args:    args{s: ""},
			wantErr: false,
		},
		{
			name:    "Dev",
			e:       Env__Dev,
			args:    args{s: "Dev"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Set(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnv_String(t *testing.T) {
	tests := []struct {
		name string
		e    Env
		want string
	}{
		{
			name: "dev",
			e:    Env__Dev,
			want: "dev",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnv_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		e       Env
		args    args
		wantErr bool
	}{
		{
			name:    "dev",
			e:       Env__Dev,
			args:    args{text: []byte("dev")},
			wantErr: false,
		},
		{
			name:    "dev",
			e:       Env__Dev,
			args:    args{text: []byte("de1")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnv_unmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		e    Env
		args args
		want bool
	}{
		{
			name: "dev",
			e:    Env__Dev,
			args: args{text: []byte("dev")},
			want: true,
		},
		{
			name: "dev",
			e:    Env__Dev,
			args: args{text: []byte("de1")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.unmarshalText(tt.args.text); got != tt.want {
				t.Errorf("unmarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}
