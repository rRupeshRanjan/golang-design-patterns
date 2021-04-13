package builder

type PersonalComputer struct {
	cpuCores   int
	cpuFlavour string
	ram        int
	hdd        string
	os         string
}

type PCBuilder interface {
	setCpuCores() PCBuilder
	setCpuFlavour() PCBuilder
	setRam() PCBuilder
	setHdd() PCBuilder
	setOs() PCBuilder
	GetPc() PersonalComputer
}

type HomeEditionPcBuilder struct {
	pc PersonalComputer
}

func (b *HomeEditionPcBuilder) setCpuCores() PCBuilder {
	b.pc.cpuCores = 4
	return b
}

func (b *HomeEditionPcBuilder) setCpuFlavour() PCBuilder {
	b.pc.cpuFlavour = "i3"
	return b
}

func (b *HomeEditionPcBuilder) setRam() PCBuilder {
	b.pc.ram = 4
	return b
}

func (b *HomeEditionPcBuilder) setHdd() PCBuilder {
	b.pc.hdd = "500GB"
	return b
}

func (b *HomeEditionPcBuilder) setOs() PCBuilder {
	b.pc.os = "windows 10"
	return b
}

func (b *HomeEditionPcBuilder) GetPc() PersonalComputer {
	return b.pc
}

type ProfessionalEditionPcBuilder struct {
	pc PersonalComputer
}

func (b *ProfessionalEditionPcBuilder) setCpuCores() PCBuilder {
	b.pc.cpuCores = 16
	return b
}

func (b *ProfessionalEditionPcBuilder) setCpuFlavour() PCBuilder {
	b.pc.cpuFlavour = "i9"
	return b
}

func (b *ProfessionalEditionPcBuilder) setRam() PCBuilder {
	b.pc.ram = 16
	return b
}

func (b *ProfessionalEditionPcBuilder) setHdd() PCBuilder {
	b.pc.hdd = "2TB"
	return b
}

func (b *ProfessionalEditionPcBuilder) setOs() PCBuilder {
	b.pc.os = "linux garuda"
	return b
}

func (b *ProfessionalEditionPcBuilder) GetPc() PersonalComputer {
	return b.pc
}

type Manufacturer struct {
	b PCBuilder
}

func (m *Manufacturer) SetBuilder(builder PCBuilder) {
	m.b = builder
}

func (m *Manufacturer) ConstructPC() PersonalComputer {
	return m.b.setCpuCores().setCpuFlavour().setRam().setOs().setHdd().GetPc()
}
