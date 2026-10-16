// Package fileio lit et écrit les fichiers texte.
package fileio

// Lire renvoie le contenu du fichier sous forme de texte.
// Fichier absent : erreur "fichier introuvable : <chemin>".
// Autre erreur : "lecture de <chemin> : <erreur d'origine>" (avec %w).
func Lire(chemin string) (string, error) {
	// Exercice 1 : à compléter.
	return "", nil
}

// Ecrire écrit le texte dans le fichier (créé ou remplacé).
// En cas d'erreur : "écriture de <chemin> : <erreur d'origine>" (avec %w).
func Ecrire(chemin string, texte string) error {
	// Exercice 1 : à compléter.
	return nil
}
