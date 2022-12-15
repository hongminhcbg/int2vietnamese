package int2vietnamese

import "testing"

func TestInt2VietNamese(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		// TODO: Add test cases.
		{
			3, "ba",
		},
		{
			17, "mười bảy",
		},
		{
			98, "chín mươi tám",
		},
		{
			123, "một trăm hai mươi ba",
		},
		{
			207, "hai trăm lẻ bảy",
		},
		{
			999, "chín trăm chín mươi chín",
		},
		{
			1024, "một nghìn không trăm hai mươi tư",
		},
		{
			2048, "hai nghìn không trăm bốn mươi tám",
		},
		{
			10008, "mười nghìn không trăm lẻ tám",
		},
		{
			10050, "mười nghìn không trăm năm mươi",
		},
		{
			99081, "chín mươi chín nghìn không trăm tám mươi mốt",
		},
		{
			1000000, "một triệu",
		},
		{
			1000003, "một triệu không trăm lẻ ba",
		},
		{
			1004953, "một triệu không trăm lẻ bốn nghìn chín trăm năm mươi ba",
		},
		{
			100000000, "một trăm triệu",
		},
		{
			1000000000, "một tỷ",
		},
		{
			1003000400, "một tỷ không trăm lẻ ba triệu bốn trăm",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Int2VietNamese(tt.in); got != tt.want {
				t.Errorf("Int2VietNamese() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
