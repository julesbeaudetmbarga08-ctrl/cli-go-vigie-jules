package fileio

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Exercice 1 : Ecrire puis Lire doivent rendre exactement le même texte.
func TestEx1EcrireEtLire(t *testing.T) {
	chemin := filepath.Join(t.TempDir(), "essai.txt")
	texte := "Événement accentué\nseconde ligne\n"

	if err := Ecrire(chemin, texte); err != nil {
		t.Fatalf("Ecrire(%q) a renvoyé l'erreur : %v", chemin, err)
	}
	lu, err := Lire(chemin)
	if err != nil {
		t.Fatalf("Lire(%q) a renvoyé l'erreur : %v", chemin, err)
	}
	if lu != texte {
		t.Errorf("Lire après Ecrire a renvoyé %q au lieu de %q", lu, texte)
	}
}

// Exercice 1 : Ecrire remplace un fichier qui existe déjà.
func TestEx1EcrireRemplace(t *testing.T) {
	chemin := filepath.Join(t.TempDir(), "essai.txt")
	if err := os.WriteFile(chemin, []byte("ancien contenu bien plus long"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := Ecrire(chemin, "nouveau"); err != nil {
		t.Fatalf("Ecrire a renvoyé l'erreur : %v", err)
	}
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		t.Fatal(err)
	}
	if string(contenu) != "nouveau" {
		t.Errorf("après Ecrire, le fichier contient %q au lieu de %q", contenu, "nouveau")
	}
}

// Exercice 1 : un fichier absent donne un message clair qui le nomme.
func TestEx1LireFichierAbsent(t *testing.T) {
	chemin := filepath.Join(t.TempDir(), "absent.txt")
	_, err := Lire(chemin)
	if err == nil {
		t.Fatalf("Lire(%q) sur un fichier absent devrait renvoyer une erreur, elle a renvoyé nil", chemin)
	}
	attendu := "fichier introuvable : " + chemin
	if err.Error() != attendu {
		t.Errorf("message d'erreur :\n  obtenu  : %q\n  attendu : %q", err.Error(), attendu)
	}
}

// Exercice 1 : une écriture impossible renvoie une erreur avec son contexte.
func TestEx1EcrireDossierAbsent(t *testing.T) {
	chemin := filepath.Join(t.TempDir(), "dossier-absent", "rapport.txt")
	err := Ecrire(chemin, "texte")
	if err == nil {
		t.Fatalf("Ecrire(%q) dans un dossier absent devrait renvoyer une erreur, elle a renvoyé nil", chemin)
	}
	debut := "écriture de " + chemin + " : "
	if !strings.HasPrefix(err.Error(), debut) {
		t.Errorf("le message d'erreur devrait commencer par %q (contexte ajouté avec %%w)\n  obtenu : %q", debut, err.Error())
	}
}
