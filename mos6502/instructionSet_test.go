package mos6502

import "testing"

func TestCPU_adc(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.adc(tt.args.operand)
		})
	}
}

func TestCPU_and(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.and(tt.args.operand)
		})
	}
}

func TestCPU_asl(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.asl(tt.args.operand)
		})
	}
}

func TestCPU_bcc(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bcc(tt.args.operand)
		})
	}
}

func TestCPU_bcs(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bcs(tt.args.operand)
		})
	}
}

func TestCPU_beq(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.beq(tt.args.operand)
		})
	}
}

func TestCPU_bit(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bit(tt.args.operand)
		})
	}
}

func TestCPU_bmi(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bmi(tt.args.operand)
		})
	}
}

func TestCPU_bne(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bne(tt.args.operand)
		})
	}
}

func TestCPU_bpl(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bpl(tt.args.operand)
		})
	}
}

func TestCPU_brk(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.brk(tt.args.operand)
		})
	}
}

func TestCPU_bvc(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bvc(tt.args.operand)
		})
	}
}

func TestCPU_bvs(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.bvs(tt.args.operand)
		})
	}
}

func TestCPU_clc(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.clc(tt.args.operand)
		})
	}
}

func TestCPU_cld(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cld(tt.args.operand)
		})
	}
}

func TestCPU_cli(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cli(tt.args.operand)
		})
	}
}

func TestCPU_clv(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.clv(tt.args.operand)
		})
	}
}

func TestCPU_cmp(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cmp(tt.args.operand)
		})
	}
}

func TestCPU_cpx(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cpx(tt.args.operand)
		})
	}
}

func TestCPU_cpy(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.cpy(tt.args.operand)
		})
	}
}

func TestCPU_dec(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.dec(tt.args.operand)
		})
	}
}

func TestCPU_dex(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.dex(tt.args.operand)
		})
	}
}

func TestCPU_dey(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.dey(tt.args.operand)
		})
	}
}

func TestCPU_eor(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.eor(tt.args.operand)
		})
	}
}

func TestCPU_inc(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.inc(tt.args.operand)
		})
	}
}

func TestCPU_inx(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.inx(tt.args.operand)
		})
	}
}

func TestCPU_iny(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.iny(tt.args.operand)
		})
	}
}

func TestCPU_jmp(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.jmp(tt.args.operand)
		})
	}
}

func TestCPU_jsr(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.jsr(tt.args.operand)
		})
	}
}

func TestCPU_lda(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.lda(tt.args.operand)
		})
	}
}

func TestCPU_ldx(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ldx(tt.args.operand)
		})
	}
}

func TestCPU_ldy(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ldy(tt.args.operand)
		})
	}
}

func TestCPU_lsr(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.lsr(tt.args.operand)
		})
	}
}

func TestCPU_nop(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.nop(tt.args.operand)
		})
	}
}

func TestCPU_ora(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ora(tt.args.operand)
		})
	}
}

func TestCPU_pha(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.pha(tt.args.operand)
		})
	}
}

func TestCPU_php(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.php(tt.args.operand)
		})
	}
}

func TestCPU_pla(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.pla(tt.args.operand)
		})
	}
}

func TestCPU_plp(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.plp(tt.args.operand)
		})
	}
}

func TestCPU_rol(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.rol(tt.args.operand)
		})
	}
}

func TestCPU_ror(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.ror(tt.args.operand)
		})
	}
}

func TestCPU_rti(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.rti(tt.args.operand)
		})
	}
}

func TestCPU_rts(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.rts(tt.args.operand)
		})
	}
}

func TestCPU_sbc(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sbc(tt.args.operand)
		})
	}
}

func TestCPU_sec(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sec(tt.args.operand)
		})
	}
}

func TestCPU_sed(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sed(tt.args.operand)
		})
	}
}

func TestCPU_sei(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sei(tt.args.operand)
		})
	}
}

func TestCPU_sta(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sta(tt.args.operand)
		})
	}
}

func TestCPU_stx(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.stx(tt.args.operand)
		})
	}
}

func TestCPU_sty(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.sty(tt.args.operand)
		})
	}
}

func TestCPU_tax(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tax(tt.args.operand)
		})
	}
}

func TestCPU_tay(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tay(tt.args.operand)
		})
	}
}

func TestCPU_tsx(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tsx(tt.args.operand)
		})
	}
}

func TestCPU_txa(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.txa(tt.args.operand)
		})
	}
}

func TestCPU_txs(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.txs(tt.args.operand)
		})
	}
}

func TestCPU_tya(t *testing.T) {
	type args struct {
		operand addressMode
	}
	tests := []struct {
		name string
		cpu  *cpu
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cpu.tya(tt.args.operand)
		})
	}
}
