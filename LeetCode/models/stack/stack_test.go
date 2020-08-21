package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple testing what different between Fatal and Error
func TestNew(t *testing.T) {
	stringStack := StringStack{}
	assert.NotNil(t, stringStack)

	stringStack.Push("+")
	stringStack.Push("/")
	stringStack.Push("*")
	res := []string{}
	for len(stringStack) > 0 {
		x, y := stringStack.Pop()
		if y == true {
			fmt.Printf("stringStack.Pop:%v\n", x)
			res = append(res, x.(string))
		}
	}

	expected := []string{"*", "/", "+"}
	assert.Equal(t, expected, res)
	// assert.Nil(t, c)
}

// // Testing
// func TestCar_SetName(t *testing.T) {
// 	type fields struct {
// 		Name  string
// 		Price float32
// 	}
// 	type args struct {
// 		name string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   string
// 	}{
// 		{
// 			name: "input empty name",
// 			fields: fields{
// 				Name:  "foo",
// 				Price: 100,
// 			},
// 			args: args{
// 				name: "",
// 			},
// 			want: "foo",
// 		},
// 		{
// 			name: "input name",
// 			fields: fields{
// 				Name:  "foo",
// 				Price: 100,
// 			},
// 			args: args{
// 				name: "bar",
// 			},
// 			want: "bar",
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			c := &Car{
// 				Name:  tt.fields.Name,
// 				Price: tt.fields.Price,
// 			}
// 			if got := c.SetName(tt.args.name); got != tt.want {
// 				t.Errorf("Car.SetName() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
