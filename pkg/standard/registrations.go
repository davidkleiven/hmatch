package standard

import (
	"fmt"
	"strings"

	"github.com/davidkleiven/hmatch/pkg"
)

type Registration struct {
	Name         string
	Registration pkg.DrawbarSetting
}

type Registrations []Registration

func (r Registrations) ToString() string {
	horizontalLine := strings.Repeat("-", 52)
	result := fmt.Sprintf("%s\n| %30s | %15s |\n%s\n", horizontalLine, "Name", "Positition", horizontalLine)
	for _, v := range r {
		result += fmt.Sprintf("| %30s | %15s |\n", v.Name, pkg.FormatDrawbarPositions(v.Registration[:]))
	}
	result += fmt.Sprintf("%s\n", horizontalLine)
	return result
}

func AllRegistations() Registrations {
	return []Registration{
		{
			Name:         "Tibia 16'",
			Registration: pkg.DrawbarSetting{7, 2, 0, 0, 2, 0, 0, 0, 0},
		},
		{
			Name:         "Bourdon 16'",
			Registration: pkg.DrawbarSetting{5, 4, 3, 1, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Diapason 16'",
			Registration: pkg.DrawbarSetting{6, 4, 3, 3, 2, 2, 0, 0, 0},
		},
		{
			Name:         "Solo Strings 16'",
			Registration: pkg.DrawbarSetting{2, 5, 4, 4, 2, 1, 0, 0, 0},
		},
		{
			Name:         "Contra Viol 16'",
			Registration: pkg.DrawbarSetting{2, 4, 3, 2, 1, 0, 0, 0, 0},
		},
		{
			Name:         "Contra Celeste 16'",
			Registration: pkg.DrawbarSetting{2, 3, 4, 3, 2, 1, 0, 0, 0},
		},
		{
			Name:         "Vox Humana 16'",
			Registration: pkg.DrawbarSetting{1, 4, 3, 1, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Oboe Horn 16'",
			Registration: pkg.DrawbarSetting{4, 7, 6, 3, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Saxophone 16'",
			Registration: pkg.DrawbarSetting{2, 7, 3, 2, 1, 0, 0, 0, 0},
		},
		{
			Name:         "Clarinet 16'",
			Registration: pkg.DrawbarSetting{3, 5, 4, 3, 2, 0, 0, 0, 0},
		},
		{
			Name:         "English Horn 16'",
			Registration: pkg.DrawbarSetting{2, 5, 4, 4, 2, 1, 0, 0, 0},
		},
		{
			Name:         "Ophicleide 16'",
			Registration: pkg.DrawbarSetting{4, 7, 6, 0, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Tibia 8'",
			Registration: pkg.DrawbarSetting{0, 8, 2, 4, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Concert Flute 8'",
			Registration: pkg.DrawbarSetting{0, 6, 4, 2, 1, 0, 0, 0, 0},
		},
		{
			Name:         "Diapason 8'",
			Registration: pkg.DrawbarSetting{0, 5, 6, 4, 2, 1, 1, 0, 0},
		},
		{
			Name:         "Solo Strings 8'",
			Registration: pkg.DrawbarSetting{0, 2, 3, 6, 6, 5, 4, 2, 0},
		},
		{
			Name:         "Viol d’Orchestre 8'",
			Registration: pkg.DrawbarSetting{0, 2, 4, 4, 4, 3, 2, 0, 0},
		},
		{
			Name:         "Viole Celeste 8'",
			Registration: pkg.DrawbarSetting{0, 2, 3, 2, 3, 2, 1, 0, 0},
		},
		{
			Name:         "Vox Humana 8'",
			Registration: pkg.DrawbarSetting{0, 3, 4, 0, 0, 3, 3, 2, 0},
		},
		{
			Name:         "Oboe Horn 8'",
			Registration: pkg.DrawbarSetting{0, 4, 7, 6, 3, 0, 0, 0, 0},
		},
		{
			Name:         "Saxophone 8'",
			Registration: pkg.DrawbarSetting{0, 2, 4, 7, 8, 5, 0, 0, 0},
		},
		{
			Name:         "Clarinet 8'",
			Registration: pkg.DrawbarSetting{0, 8, 3, 8, 2, 7, 0, 0, 0},
		},
		{
			Name:         "English Horn 8'",
			Registration: pkg.DrawbarSetting{0, 3, 5, 7, 7, 5, 4, 0, 0},
		},
		{
			Name:         "Tuba 8'",
			Registration: pkg.DrawbarSetting{0, 5, 6, 8, 0, 4, 0, 0, 0},
		},
		{
			Name:         "Flute 4'",
			Registration: pkg.DrawbarSetting{0, 0, 8, 0, 3, 0, 3, 0, 0},
		},
		{
			Name:         "Piccolo 4'",
			Registration: pkg.DrawbarSetting{0, 0, 6, 0, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Octave 4'",
			Registration: pkg.DrawbarSetting{0, 0, 5, 4, 5, 3, 2, 1, 0},
		},
		{
			Name:         "Solo Strings 4'",
			Registration: pkg.DrawbarSetting{0, 0, 4, 3, 6, 5, 5, 5, 0},
		},
		{
			Name:         "Viol 4'",
			Registration: pkg.DrawbarSetting{0, 0, 3, 4, 4, 2, 3, 2, 0},
		},
		{
			Name:         "Octave Celeste 4'",
			Registration: pkg.DrawbarSetting{0, 0, 3, 2, 4, 2, 2, 0, 0},
		},
		{
			Name:         "Vox Humana 4'",
			Registration: pkg.DrawbarSetting{0, 0, 4, 3, 3, 0, 4, 2, 0},
		},
		{
			Name:         "Oboe Horn 4'",
			Registration: pkg.DrawbarSetting{0, 0, 6, 0, 6, 3, 1, 0, 0},
		},
		{
			Name:         "Clarion 4'",
			Registration: pkg.DrawbarSetting{0, 0, 5, 1, 5, 2, 3, 0, 0},
		},
		{
			Name:         "Tibia 2'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 6, 0, 0, 1, 0, 0},
		},
		{
			Name:         "Piccolo 2'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 5, 1, 1, 1, 0, 0},
		},
		{
			Name:         "Twelfth",
			Registration: pkg.DrawbarSetting{0, 0, 6, 0, 2, 0, 0, 2, 0},
		},
		{
			Name:         "Gedeckt 16'",
			Registration: pkg.DrawbarSetting{3, 2, 2, 0, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Geigen Diapason 8'",
			Registration: pkg.DrawbarSetting{0, 4, 7, 4, 3, 2, 2, 1, 1},
		},
		{
			Name:         "Stopped Diapason 8'",
			Registration: pkg.DrawbarSetting{0, 5, 1, 4, 1, 1, 0, 0, 0},
		},
		{
			Name:         "Aeoline 8'",
			Registration: pkg.DrawbarSetting{0, 2, 3, 1, 1, 0, 0, 0, 0},
		},
		{
			Name:         "Salicional 8'",
			Registration: pkg.DrawbarSetting{0, 2, 4, 3, 3, 2, 2, 0, 0},
		},
		{
			Name:         "Viol Da Gamba 8'",
			Registration: pkg.DrawbarSetting{0, 2, 4, 2, 3, 2, 2, 1, 0},
		},
		{
			Name:         "Voix Celeste",
			Registration: pkg.DrawbarSetting{0, 2, 4, 3, 4, 4, 3, 2, 0},
		},
		{
			Name:         "Octave Geigen 4'",
			Registration: pkg.DrawbarSetting{0, 0, 4, 1, 4, 2, 3, 1, 0},
		},
		{
			Name:         "Traverse Flute 4'",
			Registration: pkg.DrawbarSetting{0, 0, 6, 0, 1, 0, 1, 0, 0},
		},
		{
			Name:         "Fifteenth 2'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 2, 2, 2, 3, 0, 0},
		},
		{
			Name:         "Cornet V",
			Registration: pkg.DrawbarSetting{0, 4, 4, 4, 6, 5, 3, 3, 0},
		},
		{
			Name:         "Oboe 8'",
			Registration: pkg.DrawbarSetting{0, 4, 7, 6, 4, 2, 1, 0, 0},
		},
		{
			Name:         "Trumpet 8'",
			Registration: pkg.DrawbarSetting{0, 6, 8, 6, 7, 5, 3, 1, 0},
		},
		{
			Name:         "Clarion 4'",
			Registration: pkg.DrawbarSetting{0, 0, 5, 4, 5, 4, 5, 2, 0},
		},
		{
			Name:         "Double Trumpet 16'",
			Registration: pkg.DrawbarSetting{3, 6, 6, 4, 2, 0, 0, 0, 0},
		},
		{
			Name:         "Vox Humana 8'",
			Registration: pkg.DrawbarSetting{0, 3, 4, 5, 3, 5, 4, 2, 0},
		},
		{
			Name:         "French Horn 8'",
			Registration: pkg.DrawbarSetting{0, 7, 5, 3, 1, 0, 0, 0, 0},
		},
		{
			Name:         "English Horn 8'",
			Registration: pkg.DrawbarSetting{0, 3, 7, 4, 4, 3, 2, 0, 0},
		},
		{
			Name:         "Grosse Flute 8'",
			Registration: pkg.DrawbarSetting{0, 8, 7, 2, 3, 1, 0, 0, 0},
		},
		{
			Name:         "Tromba 8'",
			Registration: pkg.DrawbarSetting{0, 6, 8, 6, 8, 6, 4, 2, 0},
		},
		{
			Name:         "Bassoon 8'",
			Registration: pkg.DrawbarSetting{0, 8, 7, 5, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Heckelphone 8'",
			Registration: pkg.DrawbarSetting{0, 6, 7, 7, 6, 4, 0, 0, 0},
		},
		{
			Name:         "1st Open Diapason 8'",
			Registration: pkg.DrawbarSetting{0, 6, 6, 3, 4, 2, 2, 1, 0},
		},
		{
			Name:         "2nd Open Diapason 8'",
			Registration: pkg.DrawbarSetting{0, 6, 4, 3, 3, 2, 2, 1, 0},
		},
		{
			Name:         "Hohl Flute 8'",
			Registration: pkg.DrawbarSetting{0, 5, 3, 1, 1, 0, 0, 0, 0},
		},
		{
			Name:         "Gemshorn 8'",
			Registration: pkg.DrawbarSetting{0, 3, 5, 1, 1, 1, 0, 0, 0},
		},
		{
			Name:         "Octave 4'",
			Registration: pkg.DrawbarSetting{0, 0, 5, 2, 5, 3, 4, 2, 0},
		},
		{
			Name:         "Flute Triangulaire 4'",
			Registration: pkg.DrawbarSetting{0, 0, 5, 0, 3, 1, 0, 0, 0},
		},
		{
			Name:         "Super Octave 2'",
			Registration: pkg.DrawbarSetting{0, 0, 6, 2, 3, 4, 0, 0, 0},
		},
		{
			Name:         "Mixture IV",
			Registration: pkg.DrawbarSetting{0, 0, 5, 6, 4, 1, 2, 3, 0},
		},
		{
			Name:         "Tromba 8'",
			Registration: pkg.DrawbarSetting{0, 1, 6, 5, 5, 4, 3, 0, 0},
		},
		{
			Name:         "Dulciana 8'",
			Registration: pkg.DrawbarSetting{0, 3, 3, 2, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Unda Maris 8'",
			Registration: pkg.DrawbarSetting{0, 3, 4, 2, 1, 0, 0, 0, 0},
		},
		{
			Name:         "Melodia 8'",
			Registration: pkg.DrawbarSetting{0, 5, 2, 0, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Concert Flute 8'",
			Registration: pkg.DrawbarSetting{0, 4, 5, 1, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Diapason 8'",
			Registration: pkg.DrawbarSetting{0, 3, 4, 1, 2, 2, 0, 0, 0},
		},
		{
			Name:         "Orchestral Flute 8'",
			Registration: pkg.DrawbarSetting{0, 0, 8, 0, 5, 0, 0, 0, 0},
		},
		{
			Name:         "Flute d’Amour 4'",
			Registration: pkg.DrawbarSetting{0, 0, 5, 0, 1, 0, 0, 0, 0},
		},
		{
			Name:         "Principal 4'",
			Registration: pkg.DrawbarSetting{0, 0, 4, 1, 5, 1, 1, 2, 0},
		},
		{
			Name:         "Flageolet 2'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 4, 1, 1, 2, 0, 0},
		},
		{
			Name:         "Nazard 2-2/3'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 4, 0, 3, 0, 0, 0},
		},
		{
			Name:         "Tierce 1-3/5'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 0, 0, 5, 0, 0, 0},
		},
		{
			Name:         "Clarinet 8'",
			Registration: pkg.DrawbarSetting{0, 7, 4, 6, 2, 4, 2, 0, 0},
		},
		{
			Name:         "Gedeckt 8'",
			Registration: pkg.DrawbarSetting{0, 5, 1, 4, 1, 1, 0, 0, 0},
		},
		{
			Name:         "Salicional 8'",
			Registration: pkg.DrawbarSetting{0, 3, 4, 3, 3, 1, 1, 0, 0},
		},
		{
			Name:         "Vox Celeste 8'",
			Registration: pkg.DrawbarSetting{0, 2, 3, 2, 2, 1, 1, 0, 0},
		},
		{
			Name:         "Principal 4'",
			Registration: pkg.DrawbarSetting{0, 0, 5, 1, 5, 0, 3, 1, 0},
		},
		{
			Name:         "Harmonic Flute 4'",
			Registration: pkg.DrawbarSetting{0, 0, 8, 0, 4, 0, 1, 1, 0},
		},
		{
			Name:         "Piccolo 2'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 6, 1, 3, 2, 0, 0},
		},
		{
			Name:         "Siffloete 1'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 0, 0, 0, 0, 5, 0},
		},
		{
			Name:         "Mixture 3 ranks",
			Registration: pkg.DrawbarSetting{0, 0, 0, 8, 7, 0, 6, 4, 0},
		},
		{
			Name:         "Contra Fagotte 16'",
			Registration: pkg.DrawbarSetting{1, 7, 5, 3, 2, 1, 0, 0, 0},
		},
		{
			Name:         "Trumpet 8'",
			Registration: pkg.DrawbarSetting{0, 6, 7, 8, 6, 5, 3, 0, 0},
		},
		{
			Name:         "Quintadena 16'",
			Registration: pkg.DrawbarSetting{2, 3, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Principal 8'",
			Registration: pkg.DrawbarSetting{0, 5, 7, 5, 4, 2, 1, 0, 0},
		},
		{
			Name:         "Hohl Flute 8'",
			Registration: pkg.DrawbarSetting{0, 6, 3, 2, 0, 0, 0, 0, 0},
		},
		{
			Name:         "Octave 4'",
			Registration: pkg.DrawbarSetting{0, 0, 6, 2, 6, 1, 2, 1, 0},
		},
		{
			Name:         "Octave 2'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 7, 0, 0, 3, 0, 0},
		},
		{
			Name:         "Mixture 4 ranks",
			Registration: pkg.DrawbarSetting{0, 0, 0, 6, 4, 0, 6, 4, 0},
		},
		{
			Name:         "Gedeckt 8'",
			Registration: pkg.DrawbarSetting{0, 5, 0, 3, 0, 1, 0, 0, 0},
		},
		{
			Name:         "Flute d’Amour 4'",
			Registration: pkg.DrawbarSetting{0, 0, 6, 0, 3, 0, 2, 0, 0},
		},
		{
			Name:         "Principal 2'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 6, 0, 0, 2, 0, 0},
		},
		{
			Name:         "Quint 1 1/3'",
			Registration: pkg.DrawbarSetting{0, 0, 0, 0, 0, 0, 6, 0, 0},
		},
		{
			Name:         "Clarinet 8'",
			Registration: pkg.DrawbarSetting{0, 4, 2, 6, 2, 4, 2, 1, 0},
		},
	}
}

func FindRegistration(pattern string) Registrations {
	pattern = strings.ToLower(pattern)
	registrations := []Registration{}
	for _, reg := range AllRegistations() {
		if strings.Contains(strings.ToLower(reg.Name), pattern) {
			registrations = append(registrations, reg)
		}
	}
	return registrations
}
