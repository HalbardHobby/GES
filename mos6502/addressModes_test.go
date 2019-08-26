package mos6502

import "testing"

func TestCPU_a(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.a(); gotValue != tt.wantValue {
				t.Errorf("CPU.a() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_abs(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.abs(); gotValue != tt.wantValue {
				t.Errorf("CPU.abs() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_absX(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.absX(); gotValue != tt.wantValue {
				t.Errorf("CPU.absX() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_absY(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.absY(); gotValue != tt.wantValue {
				t.Errorf("CPU.absY() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_imm(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.imm(); gotValue != tt.wantValue {
				t.Errorf("CPU.imm() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_impl(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.impl(); gotValue != tt.wantValue {
				t.Errorf("CPU.impl() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_ind(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.ind(); gotValue != tt.wantValue {
				t.Errorf("CPU.ind() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_xInd(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.xInd(); gotValue != tt.wantValue {
				t.Errorf("CPU.xInd() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_indY(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.indY(); gotValue != tt.wantValue {
				t.Errorf("CPU.indY() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_rel(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.rel(); gotValue != tt.wantValue {
				t.Errorf("CPU.rel() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_zpg(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.zpg(); gotValue != tt.wantValue {
				t.Errorf("CPU.zpg() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_zpgX(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.zpgX(); gotValue != tt.wantValue {
				t.Errorf("CPU.zpgX() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestCPU_zpgY(t *testing.T) {
	tests := []struct {
		name      string
		cpu       *CPU
		wantValue uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.cpu.zpgY(); gotValue != tt.wantValue {
				t.Errorf("CPU.zpgY() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
