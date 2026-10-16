package parser

import (
	"errors"
	"testing"
)

// Exercice 2 : Analyser sur des lignes valides.
func TestEx2AnalyserLignesValides(t *testing.T) {
	cas := []struct {
		ligne   string
		attendu Evenement
	}{
		{"2026-10-14 08:14:09 FAIL admin 10.0.0.5",
			Evenement{"2026-10-14", "08:14:09", "FAIL", "admin", "10.0.0.5"}},
		{"2026-10-14 08:00:03 ok alice 192.168.1.20",
			Evenement{"2026-10-14", "08:00:03", "OK", "alice", "192.168.1.20"}},
		{"2026-10-14   08:00:07\tFail   bob  192.168.1.33\r",
			Evenement{"2026-10-14", "08:00:07", "FAIL", "bob", "192.168.1.33"}},
	}
	for _, c := range cas {
		obtenu, err := Analyser(c.ligne)
		if err != nil {
			t.Errorf("Analyser(%q) a renvoyé l'erreur %q alors que la ligne est valide", c.ligne, err)
			continue
		}
		if obtenu != c.attendu {
			t.Errorf("Analyser(%q)\n  obtenu  : %+v\n  attendu : %+v", c.ligne, obtenu, c.attendu)
		}
	}
}

// Exercice 2 : Analyser refuse les lignes invalides avec une erreur qui enveloppe ErrLigneInvalide.
func TestEx2AnalyserLignesInvalides(t *testing.T) {
	cas := []struct {
		ligne   string
		message string
	}{
		{"2026-10-14 08:00:03 FAIL admin", "ligne invalide : 4 champs au lieu de 5"},
		{"2026-10-14 08:00:03 FAIL admin 10.0.0.5 ssh", "ligne invalide : 6 champs au lieu de 5"},
		{"", "ligne invalide : 0 champs au lieu de 5"},
		{"2026-10-14 08:00:03 DENIED admin 10.0.0.5", `ligne invalide : résultat inconnu "DENIED"`},
	}
	for _, c := range cas {
		_, err := Analyser(c.ligne)
		if err == nil {
			t.Errorf("Analyser(%q) devrait renvoyer une erreur, elle a renvoyé nil", c.ligne)
			continue
		}
		if !errors.Is(err, ErrLigneInvalide) {
			t.Errorf("Analyser(%q) : l'erreur doit envelopper ErrLigneInvalide (avec %%w)", c.ligne)
		}
		if err.Error() != c.message {
			t.Errorf("Analyser(%q) : message d'erreur\n  obtenu  : %q\n  attendu : %q", c.ligne, err.Error(), c.message)
		}
	}
}

// Exercice 3 : Journal sur un texte complet.
func TestEx3Journal(t *testing.T) {
	cas := []struct {
		nom            string
		texte          string
		nbEvenements   int
		avertissements []string
	}{
		{"texte vide", "", 0, nil},
		{"lignes vides et ligne invalide",
			"2026-10-14 08:00:03 OK alice 192.168.1.20\n\n2026-10-14 08:00:05 FAIL admin\n2026-10-14 08:00:09 FAIL admin 10.0.0.5\n",
			2, []string{"ligne 3 ignorée : ligne invalide : 4 champs au lieu de 5"}},
		{"fichier Windows",
			"2026-10-14 08:00:03 OK alice 192.168.1.20\r\n\r\n2026-10-14 08:00:09 FAIL admin 10.0.0.5\r\n",
			2, nil},
		{"résultat inconnu en dernière ligne",
			"2026-10-14 08:00:03 OK alice 192.168.1.20\n2026-10-14 08:00:09 TIMEOUT eve 192.168.1.60",
			1, []string{`ligne 2 ignorée : ligne invalide : résultat inconnu "TIMEOUT"`}},
	}
	for _, c := range cas {
		evenements, avertissements := Journal(c.texte)
		if len(evenements) != c.nbEvenements {
			t.Errorf("%s : %d événements au lieu de %d", c.nom, len(evenements), c.nbEvenements)
		}
		if len(avertissements) != len(c.avertissements) {
			t.Errorf("%s : %d avertissements au lieu de %d\n  obtenus : %q", c.nom, len(avertissements), len(c.avertissements), avertissements)
			continue
		}
		for i := range avertissements {
			if avertissements[i] != c.avertissements[i] {
				t.Errorf("%s : avertissement\n  obtenu  : %q\n  attendu : %q", c.nom, avertissements[i], c.avertissements[i])
			}
		}
	}
}

// Exercice 3 : Journal garde les événements dans l'ordre du fichier.
func TestEx3JournalContenu(t *testing.T) {
	evenements, _ := Journal("2026-10-14 08:00:03 ok alice 192.168.1.20\n2026-10-14 08:00:09 FAIL admin 10.0.0.5\n")
	attendus := []Evenement{
		{"2026-10-14", "08:00:03", "OK", "alice", "192.168.1.20"},
		{"2026-10-14", "08:00:09", "FAIL", "admin", "10.0.0.5"},
	}
	if len(evenements) != len(attendus) {
		t.Fatalf("%d événements au lieu de %d", len(evenements), len(attendus))
	}
	for i := range attendus {
		if evenements[i] != attendus[i] {
			t.Errorf("événement %d\n  obtenu  : %+v\n  attendu : %+v", i+1, evenements[i], attendus[i])
		}
	}
}
