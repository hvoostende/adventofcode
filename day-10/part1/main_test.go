package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_bot_getLowestMC(t *testing.T) {
	type fields struct {
		botID      string
		microChips []int
		low        string
		high       string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test1",
			fields{botID: "bot1", microChips: []int{10, 20}, low: "", high: ""},
			10},
		{"test2",
			fields{botID: "bot1", microChips: []int{20, 10}, low: "", high: ""},
			10},
		{"test3",
			fields{botID: "bot1", microChips: []int{20, 20}, low: "", high: ""},
			20},
		{"test4",
			fields{botID: "bot1", microChips: []int{20, -10}, low: "foo", high: "bar"},
			-10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bot{
				botID:      tt.fields.botID,
				microChips: tt.fields.microChips,
				low:        tt.fields.low,
				high:       tt.fields.high,
			}
			if got := b.getLowestMC(); got != tt.want {
				t.Errorf("bot.getLowestMC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bot_getHighestMC(t *testing.T) {
	type fields struct {
		botID      string
		microChips []int
		low        string
		high       string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test1",
			fields{botID: "bot1", microChips: []int{10, 20}, low: "", high: ""},
			20},
		{"test2",
			fields{botID: "bot1", microChips: []int{20, 10}, low: "", high: ""},
			20},
		{"test3",
			fields{botID: "bot1", microChips: []int{20, 20}, low: "", high: ""},
			20},
		{"test4",
			fields{botID: "bot1", microChips: []int{-20, -10}, low: "foo", high: "bar"},
			-10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bot{
				botID:      tt.fields.botID,
				microChips: tt.fields.microChips,
				low:        tt.fields.low,
				high:       tt.fields.high,
			}
			if got := b.getHighestMC(); got != tt.want {
				t.Errorf("bot.getHighestMC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		s   string
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{"test01",
			args{"test string", nil}},
		{"test02",
			args{"test string", errors.New("testerror")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.s, tt.args.err)
		})
	}
}

func Test_assignInputToBot(t *testing.T) {
	type args struct {
		parsedInstructions []string
		botArmy            map[string]bot
	}
	tests := []struct {
		name string
		args args
		want map[string]bot
	}{
		{"Test01-Happy_New_Bot",
			args{[]string{"value", "5", "goes", "to", "bot", "2"}, make(map[string]bot)},
			map[string]bot{
				"bot2": bot{"bot2", []int{5}, "", ""}}},
		{"Test02-Happy_Excisting_Bot ",
			args{[]string{"value", "5", "goes", "to", "bot", "2"}, map[string]bot{
				"bot2": bot{"bot2", []int{2}, "bot8", "bot34"}}},
			map[string]bot{
				"bot2": bot{"bot2", []int{2, 5}, "bot8", "bot34"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignInputToBot(tt.args.parsedInstructions, tt.args.botArmy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assignInputToBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCommandToBot(t *testing.T) {
	type args struct {
		parsedInstructions []string
		botArmy            map[string]bot
	}
	tests := []struct {
		name string
		args args
		want map[string]bot
	}{
		{"Test01-Happy_New_Bot",
			args{[]string{"bot", "2", "gives", "low", "to", "bot", "1", "and", "high", "to", "bot", "0"}, make(map[string]bot)},
			map[string]bot{
				"bot2": bot{botID: "bot2", low: "bot1", high: "bot0"}}},
		{"Test02-Happy_Excisting_Bot ",
			args{[]string{"bot", "2", "gives", "low", "to", "bot", "1", "and", "high", "to", "bot", "0"}, map[string]bot{
				"bot2": bot{"bot2", []int{2}, "", ""}}},
			map[string]bot{
				"bot2": bot{botID: "bot2", microChips: []int{2}, low: "bot1", high: "bot0"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignCommandToBot(tt.args.parsedInstructions, tt.args.botArmy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assignCommandToBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recruitBot(t *testing.T) {
	type args struct {
		instruction string
		botArmy     map[string]bot
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test01",
			args{"value 5 goes to bot 2", make(map[string]bot)},
			false},
		{"test02",
			args{"bot 2 gives low to bot 1 and high to bot 0", make(map[string]bot)},
			false},
		{"test03",
			args{"wrong 5 goes to bot 2", make(map[string]bot)},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := recruitBot(tt.args.instruction, tt.args.botArmy); (err != nil) != tt.wantErr {
				t.Errorf("recruitBot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_buildBotArmy(t *testing.T) {
	type args struct {
		location string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]bot
		wantErr bool
	}{
		{"test1",
			args{"../test.txt"},
			map[string]bot{
				"bot0": bot{botID: "bot0", low: "output2", high: "output0"},
				"bot1": bot{botID: "bot1", microChips: []int{3}, low: "output1", high: "bot0"},
				"bot2": bot{botID: "bot2", microChips: []int{5, 2}, low: "bot1", high: "bot0"},
			},
			false,
		},
		{"test2",
			args{"nofile.txt"},
			make(map[string]bot),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildBotArmy(tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildBotArmy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildBotArmy() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_giveMicroChips(t *testing.T) {
// 	type args struct {
// 		b       bot
// 		botArmy map[string]bot
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want map[string]bot
// 	}{
// 		{"test01",
// 			args{bot{}, map[string]bot{
// 				"bot0": bot{botID: "bot0", low: "output2", high: "output0"},
// 				"bot1": bot{botID: "bot1", microChips: []int{3}, low: "output1", high: "bot0"},
// 				"bot2": bot{botID: "bot2", microChips: []int{5, 2}, low: "bot1", high: "bot0"},
// 			}},
// 			map[string]bot{
// 				"bot0": bot{botID: "bot0", microChips: []int{5}, low: "output2", high: "output0"},
// 				"bot1": bot{botID: "bot1", microChips: []int{3, 2}, low: "output1", high: "bot0"},
// 				"bot2": bot{botID: "bot2", microChips: []int{}, low: "bot1", high: "bot0"},
// 			}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := giveMicroChips(tt.args.b, tt.args.botArmy); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("giveMicroChips() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }
