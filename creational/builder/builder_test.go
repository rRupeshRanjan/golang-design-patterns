package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturer := Manufacturer{}

	homeEditionPcBuilder := HomeEditionPcBuilder{}
	manufacturer.SetBuilder(&homeEditionPcBuilder)
	manufacturer.ConstructPC()

	if homeEditionPcBuilder.pc.os != "windows 10" {
		t.Errorf("Expected os: %s, instead got: %s", "windows 10", homeEditionPcBuilder.pc.os)
	} else if homeEditionPcBuilder.pc.hdd != "500GB" {
		t.Errorf("Expected hdd: %s, instead got: %s", "500GB", homeEditionPcBuilder.pc.hdd)
	} else if homeEditionPcBuilder.pc.ram != 4 {
		t.Errorf("Expected ram: %d, instead got: %d", 4, homeEditionPcBuilder.pc.ram)
	} else if homeEditionPcBuilder.pc.cpuFlavour != "i3" {
		t.Errorf("Expected cpu flavour: %s, instead got: %s", "i3", homeEditionPcBuilder.pc.cpuFlavour)
	} else if homeEditionPcBuilder.pc.cpuCores != 4 {
		t.Errorf("Expected cpu cores: %d, instead got: %d", 4, homeEditionPcBuilder.pc.cpuCores)
	}

	professionalEditionPcBuilder := ProfessionalEditionPcBuilder{}
	manufacturer.SetBuilder(&professionalEditionPcBuilder)
	manufacturer.ConstructPC()

	if professionalEditionPcBuilder.pc.os != "linux garuda" {
		t.Errorf("Expected os: %s, instead got: %s", "linux garuda", professionalEditionPcBuilder.pc.os)
	} else if professionalEditionPcBuilder.pc.hdd != "2TB" {
		t.Errorf("Expected hdd: %s, instead got: %s", "2TB", professionalEditionPcBuilder.pc.hdd)
	} else if professionalEditionPcBuilder.pc.ram != 16 {
		t.Errorf("Expected ram: %d, instead got: %d", 16, professionalEditionPcBuilder.pc.ram)
	} else if professionalEditionPcBuilder.pc.cpuFlavour != "i9" {
		t.Errorf("Expected cpu flavour: %s, instead got: %s", "i9", professionalEditionPcBuilder.pc.cpuFlavour)
	} else if professionalEditionPcBuilder.pc.cpuCores != 16 {
		t.Errorf("Expected cpu cores: %d, instead got: %d", 16, professionalEditionPcBuilder.pc.cpuCores)
	}
}
