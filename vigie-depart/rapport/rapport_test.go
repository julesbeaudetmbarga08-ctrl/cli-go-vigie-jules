package rapport

import (
	"strings"
	"testing"

	"vigie/analyse"
)

// comparer signale la première ligne qui diffère entre deux textes.
func comparer(t *testing.T, nom, obtenu, attendu string) {
	t.Helper()
	if obtenu == attendu {
		return
	}
	lo := strings.Split(obtenu, "\n")
	la := strings.Split(attendu, "\n")
	for i := 0; i < len(lo) || i < len(la); i++ {
		var a, b string
		if i < len(lo) {
			a = lo[i]
		} else {
			a = "(ligne absente)"
		}
		if i < len(la) {
			b = la[i]
		} else {
			b = "(ligne absente)"
		}
		if a != b {
			t.Errorf("%s : première différence à la ligne %d\n  obtenu  : %q\n  attendu : %q", nom, i+1, a, b)
			return
		}
	}
}

// Exercice 6 : rapport du journal d'exemple, seuil 3.
func TestEx6FormaterJournal(t *testing.T) {
	stats := analyse.Stats{
		Evenements: 17, Ignorees: 2, Reussites: 6, Echecs: 11,
		EchecsParIP: map[string]int{"192.168.1.41": 1, "10.0.0.5": 6, "172.16.0.9": 3, "192.168.1.33": 1},
	}
	attendu := `RAPPORT VIGIE
Événements analysés : 17
Lignes ignorées : 2
Connexions réussies : 6
Connexions échouées : 11

Échecs par adresse IP :
- 10.0.0.5 : 6
- 172.16.0.9 : 3
- 192.168.1.33 : 1
- 192.168.1.41 : 1

Alertes (au moins 3 échecs) :
- 10.0.0.5 : 6
- 172.16.0.9 : 3
`
	comparer(t, "journal d'exemple", Formater(stats, 3), attendu)
}

// Exercice 6 : aucun échec, aucune alerte.
func TestEx6FormaterSansEchec(t *testing.T) {
	stats := analyse.Stats{Evenements: 3, Reussites: 3, EchecsParIP: map[string]int{}}
	attendu := `RAPPORT VIGIE
Événements analysés : 3
Lignes ignorées : 0
Connexions réussies : 3
Connexions échouées : 0

Échecs par adresse IP :
Aucun échec.

Alertes (au moins 3 échecs) :
Aucune alerte.
`
	comparer(t, "sans échec", Formater(stats, 3), attendu)
}

// Exercice 6 : le seuil est inclus (au moins N échecs) et s'affiche dans le titre.
func TestEx6FormaterSeuil(t *testing.T) {
	stats := analyse.Stats{
		Evenements: 4, Echecs: 4,
		EchecsParIP: map[string]int{"10.0.0.5": 2, "172.16.0.9": 1, "10.0.0.12": 1},
	}
	attendu := `RAPPORT VIGIE
Événements analysés : 4
Lignes ignorées : 0
Connexions réussies : 0
Connexions échouées : 4

Échecs par adresse IP :
- 10.0.0.5 : 2
- 10.0.0.12 : 1
- 172.16.0.9 : 1

Alertes (au moins 2 échecs) :
- 10.0.0.5 : 2
`
	comparer(t, "seuil 2", Formater(stats, 2), attendu)
}
