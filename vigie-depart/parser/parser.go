// Package parser découpe une ligne du journal en champs et vérifie qu'elle est valide.
package parser

import "errors"

// ErrLigneInvalide signale une ligne qui ne respecte pas le format attendu.
var ErrLigneInvalide = errors.New("ligne invalide")

// NbChamps est le nombre de champs d'une ligne du journal.
const NbChamps = 5

// Evenement représente une tentative de connexion.
type Evenement struct {
	Date, Heure, Resultat, Utilisateur, IP string
}

// Analyser découpe une ligne et renvoie un Evenement,
// ou une erreur qui enveloppe ErrLigneInvalide.
func Analyser(ligne string) (Evenement, error) {
	// Exercice 2 : à compléter.
	return Evenement{}, nil
}

// Journal analyse toutes les lignes du texte.
// Les lignes vides sont sautées sans avertissement ;
// les lignes invalides sont écartées et signalées dans avertissements.
func Journal(texte string) (evenements []Evenement, avertissements []string) {
	// Exercice 3 : à compléter.
	return nil, nil
}
