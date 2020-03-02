package main

import (
	"reflect"
	"testing"
)

//func Test_getBinaryTree(t *testing.T) {
//	tests := []struct {
//		name string
//		want binaryTree
//	}{}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := getBinaryTree(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("getBinaryTree() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

// func Test_binaryTree_addNode(t *testing.T) {
// 	type fields struct {
// 		root *node
// 	}
// 	type args struct {
// 		data int64
// 	}
// 	tests := []struct {
// 		name           string
// 		initial, final binaryTree
// 		want           bool
// 		args           args
// 	}{
// 		{
// 			name:    "Test inserting tree root must add the node",
// 			initial: binaryTree{}, final: binaryTree{root: &node{data: 1}},
// 			want: true,
// 			args: args{data: 1},
// 		},
// 		{
// 			name:    "Test inserting tree root but wrong value expected",
// 			initial: binaryTree{}, final: binaryTree{root: &node{data: 1}},
// 			want: false,
// 			args: args{data: 2},
// 		},
// 		{
// 			name:    "Test inserting left node",
// 			initial: binaryTree{root: &node{data: 50}},
// 			final:   binaryTree{root: &node{data: 50, left: &node{data: 10}}},
// 			//initial: binaryTree{root: &node{data: 50}}, final: binaryTree{root: &node{data: 10}},
// 			want: true,
// 			args: args{data: 10},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.initial.addNode(tt.args.data)
//
// 			if tt.want != reflect.DeepEqual(tt.final, tt.initial) {
// 				t.Errorf("%v", spew.Sprintf("%v -> %v", tt.initial, tt.final))
// 			}
// 		})
// 	}
// }

// RECURSIVE TESTS

func Test_node_addNode(t *testing.T) {

	type args struct {
		newNode *node
	}

	tests := []struct {
		name    string
		initial *node
		final   *node
		args    args
		want    bool
	}{
		{
			name:    "Add Node to left",
			want:    true,
			initial: &node{data: 50},
			final:   &node{data: 50, left: &node{data: 25}},
			args: args{
				newNode: &node{data: 25},
			},
		},
		{
			name:    "Add Node to left but it should be to right",
			want:    false,
			initial: &node{data: 50},
			final:   &node{data: 50, left: &node{data: 75}},
			args: args{
				newNode: &node{data: 75},
			},
		},
		{
			name:    "Add Node to right",
			want:    true,
			initial: &node{data: 50},
			final:   &node{data: 50, right: &node{data: 75}},
			args: args{
				newNode: &node{data: 75},
			},
		},
		{
			name:    "Add Node to right but it should be to left",
			want:    false,
			initial: &node{data: 50},
			final:   &node{data: 50, right: &node{data: 10}},
			args: args{
				newNode: &node{data: 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.addNode(tt.args.newNode)
			if tt.want != reflect.DeepEqual(tt.final, tt.initial) {
				t.Errorf("addNode() = %v, want %v", tt.final, tt.initial)
			}
		})
	}
}

func Test_node_addNode_multiples(t *testing.T) {

	type args struct {
		newNodes []*node
	}

	tests := []struct {
		name    string
		initial *node
		final   *node
		args    args
		want    bool
	}{
		{
			name:    "Add Node Multiple Left right Left",
			want:    true,
			initial: &node{data: 50},
			final: &node{
				data: 50,
				left: &node{
					data: 25,
					right: &node{
						data: 35,
					},
				},
				right: &node{
					data: 75,
				},
			},
			args: args{
				newNodes: []*node{
					&node{data: 25},
					&node{data: 75},
					&node{data: 35},
				},
			},
		},

		{
			name:    "Add Node Multiple wrong Left right Left",
			want:    false,
			initial: &node{data: 50},
			final: &node{
				data: 50,
				left: &node{
					data: 25,
					right: &node{
						data: 35,
					},
				},
				right: &node{
					data: 75,
				},
			},
			args: args{
				newNodes: []*node{
					&node{data: 25},
					&node{data: 15},
					&node{data: 35},
				},
			},
		},

		{
			name:    "Add Node Multiple right Left right Left",
			want:    true,
			initial: &node{data: 50},
			final: &node{
				data: 50,
				left: &node{
					data: 25,
					right: &node{
						data: 33,
					},
				},
				right: &node{
					data: 75,
					left: &node{
						data: 55,
					},
				},
			},
			args: args{
				newNodes: []*node{
					&node{data: 75},
					&node{data: 25},
					&node{data: 55},
					&node{data: 33},
				},
			},
		},

		{
			name:    "Add Node Multiple wrong right Left right Left",
			want:    false,
			initial: &node{data: 50},
			final: &node{
				data: 50,
				left: &node{
					data: 25,
					right: &node{
						data: 55,
					},
				},
				right: &node{
					data: 75,
					left: &node{
						data: 33,
					},
				},
			},
			args: args{
				newNodes: []*node{
					&node{data: 75},
					&node{data: 25},
					&node{data: 55},
					&node{data: 33},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, r := range tt.args.newNodes {
				tt.initial.addNode(r)
			}

			if tt.want != reflect.DeepEqual(tt.final, tt.initial) {
				t.Errorf("addNode() = %v, want %v", tt.final, tt.initial)
			}
		})
	}
}

func ExampleMain() {
	main()
	// output:
	// 0
}

func Test_main(t *testing.T) {

	defer recoverMe(t)

	root := &node{}
	root.addNode(nil)

}

func recoverMe(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("I was expecting a panic")
	}
}
