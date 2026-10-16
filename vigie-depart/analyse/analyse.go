// Package analyse calcule les statistiques du journal.
// Ses fonctions sont pures : elles ne lisent ni n'écrivent aucun fichier et n'affichent rien.
package analyse

import "vigie/parser"

// Stats regroupe les comptages du journal.
type Stats struct {
	Evenements  int            // lignes valides
	Ignorees    int            // lignes invalides
	Reussites   int            // résultat OK
	Echecs      int            // résultat FAIL
	EchecsParIP map[string]int // adresse IP -> nombre d'échecs
}

// Compte associe un nom (adresse IP ou utilisateur) à un nombre d'échecs.
type Compte struct {
	Nom    string
	Nombre int
}

// Calculer compte les réussites, les échecs et les échecs par adresse IP.
func Calculer(evenements []parser.Evenement, ignorees int) Stats {
	// Exercice 4 : à compléter.
	return Stats{}
}

// Classer transforme la map en tranche triée par nombre décroissant,
// puis par nom (ordre alphabétique) en cas d'égalité.
func Classer(nombres map[string]int) []Compte {
	// Exercice 5 : à compléter.
	return nil
}
