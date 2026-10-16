package analyse

import (
	"testing"

	"vigie/parser"
)

// Exercice 4 : Calculer compte les réussites, les échecs et les échecs par adresse IP.
func TestEx4Calculer(t *testing.T) {
	evenements := []parser.Evenement{
		{Date: "2026-10-14", Heure: "08:00:03", Resultat: "OK", Utilisateur: "alice", IP: "192.168.1.20"},
		{Date: "2026-10-14", Heure: "08:00:05", Resultat: "FAIL", Utilisateur: "admin", IP: "10.0.0.5"},
		{Date: "2026-10-14", Heure: "08:00:07", Resultat: "FAIL", Utilisateur: "root", IP: "10.0.0.5"},
		{Date: "2026-10-14", Heure: "08:00:09", Resultat: "FAIL", Utilisateur: "bob", IP: "192.168.1.33"},
		{Date: "2026-10-14", Heure: "08:00:11", Resultat: "OK", Utilisateur: "bob", IP: "192.168.1.33"},
	}
	s := Calculer(evenements, 2)

	verifier := func(nom string, obtenu, attendu int) {
		if obtenu != attendu {
			t.Errorf("Calculer : %s = %d au lieu de %d", nom, obtenu, attendu)
		}
	}
	verifier("Evenements", s.Evenements, 5)
	verifier("Ignorees", s.Ignorees, 2)
	verifier("Reussites", s.Reussites, 2)
	verifier("Echecs", s.Echecs, 3)

	attendu := map[string]int{"10.0.0.5": 2, "192.168.1.33": 1}
	if len(s.EchecsParIP) != len(attendu) {
		t.Errorf("Calculer : EchecsParIP contient %d adresses au lieu de %d : %v", len(s.EchecsParIP), len(attendu), s.EchecsParIP)
	}
	for ip, n := range attendu {
		if s.EchecsParIP[ip] != n {
			t.Errorf("Calculer : EchecsParIP[%q] = %d au lieu de %d", ip, s.EchecsParIP[ip], n)
		}
	}
}

// Exercice 4 : aucun événement, tout vaut zéro.
func TestEx4CalculerVide(t *testing.T) {
	s := Calculer(nil, 0)
	if s.Evenements != 0 || s.Ignorees != 0 || s.Reussites != 0 || s.Echecs != 0 || len(s.EchecsParIP) != 0 {
		t.Errorf("Calculer(nil, 0) devrait tout mettre à zéro, obtenu : %+v", s)
	}
}

// Exercice 5 : Classer trie par nombre décroissant, puis par ordre alphabétique en cas d'égalité.
func TestEx5Classer(t *testing.T) {
	cas := []struct {
		nom     string
		entree  map[string]int
		attendu []Compte
	}{
		{"map vide", map[string]int{}, nil},
		{"tri décroissant", map[string]int{"a": 1, "b": 3, "c": 2},
			[]Compte{{"b", 3}, {"c", 2}, {"a", 1}}},
		{"égalités", map[string]int{"192.168.1.33": 2, "10.0.0.5": 4, "10.0.0.12": 2, "172.16.0.9": 4},
			[]Compte{{"10.0.0.5", 4}, {"172.16.0.9", 4}, {"10.0.0.12", 2}, {"192.168.1.33", 2}}},
	}
	for _, c := range cas {
		// On recommence 20 fois : l'ordre de parcours d'une map change d'une fois à l'autre.
		for essai := 0; essai < 20; essai++ {
			obtenu := Classer(c.entree)
			if len(obtenu) != len(c.attendu) {
				t.Fatalf("%s : %d éléments au lieu de %d : %v", c.nom, len(obtenu), len(c.attendu), obtenu)
			}
			for i := range obtenu {
				if obtenu[i] != c.attendu[i] {
					t.Fatalf("%s (essai %d) :\n  obtenu  : %v\n  attendu : %v", c.nom, essai+1, obtenu, c.attendu)
				}
			}
		}
	}
}
