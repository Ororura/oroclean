package formatutil

import "testing"

func TestBytes(t *testing.T) {
	tests := []struct {
		name string
		size int64
		want string
	}{
		{
			name: "bytes",
			size: 512,
			want: "512 B",
		},
		{
			name: "kibibytes",
			size: 1024,
			want: "1.00 KiB",
		},
		{
			name: "mebibytes",
			size: 1024 * 1024,
			want: "1.00 MiB",
		},
		{
			name: "gibibytes",
			size: 1024 * 1024 * 1024,
			want: "1.00 GiB",
		},
		{
			name: "tebibytes",
			size: 1024 * 1024 * 1024 * 1024,
			want: "1.00 TiB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bytes(tt.size)

			if got != tt.want {
				t.Fatalf(
					"Bytes(%d) = %q, want %q",
					tt.size,
					got,
					tt.want,
				)
			}
		})
	}
}
